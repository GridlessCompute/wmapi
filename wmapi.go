// Package whatsminer provides a low-level API interface for Whatsminer ASICs.
// This package handles the core communication protocol, encryption, and token management.
package wmapi

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"maps"
	"net"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/GehirnInc/crypt/md5_crypt"
	"github.com/andreburgaud/crypt2go/ecb"
)

// WhatsminerAccessToken represents a reusable token to access and/or control a single Whatsminer ASIC.
type WhatsminerAccessToken struct {
	Created        time.Time
	IPAddress      string
	Port           int
	AdminPassword  string
	Cipher         cipher.Block
	Sign           string
	OnRefreshError func(err error)

	mu   sync.Mutex
	stop chan bool
}

// NewWhatsminerAccessToken creates a new instance of WhatsminerAccessToken.
func NewWhatsminerAccessToken(ipAddress string, port int, adminPassword string) (*WhatsminerAccessToken, error) {

	token := &WhatsminerAccessToken{
		Created:   time.Now(),
		IPAddress: ipAddress,
		Port:      port,
		stop:      make(chan bool),
	}

	if adminPassword != "" {
		if err := token.EnableWriteAccess(adminPassword); err != nil {
			return nil, fmt.Errorf("error while trying to enable write access: %w", err)
		}
	}

	return token, nil
}

// Close stops the background token refresh goroutine.
func (t *WhatsminerAccessToken) Close() {
	select {
	case <-t.stop:
		// Channel is already closed.
		return
	default:
		close(t.stop)
	}
}

func (t *WhatsminerAccessToken) getTokenInfo() (map[string]any, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(t.IPAddress, fmt.Sprintf("%d", t.Port)), 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to miner: %w", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte(`{"cmd": "get_token"}`)); err != nil {
		return nil, fmt.Errorf("failed to send get_token command: %w", err)
	}

	response, err := io.ReadAll(conn)
	if err != nil {
		return nil, fmt.Errorf("failed to read token response: %w", err)
	}

	var tokenInfo map[string]any
	if err := json.Unmarshal(response, &tokenInfo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal token response: %w", err)
	}

	if msg, ok := tokenInfo["Msg"].(string); ok && msg == "over max connect" {
		return nil, errors.New(msg)
	}

	return tokenInfo, nil
}

func (t *WhatsminerAccessToken) generateCipherAndSign(tokenInfo map[string]any, adminPassword string) error {
	msg, ok := tokenInfo["Msg"].(map[string]any)
	if !ok {
		return errors.New("invalid token format in 'Msg' field")
	}

	salt, _ := msg["salt"].(string)
	newsalt, _ := msg["newsalt"].(string)
	tokenTime, _ := msg["time"].(float64)

	fullSalt := fmt.Sprintf("$1$%s$", salt)
	r := regexp.MustCompile(`\s*\$(\d+)\$([\w\./]*)\$`)
	if !r.MatchString(fullSalt) {
		return errors.New("salt format is not correct")
	}

	m := md5_crypt.New()
	hash, err := m.Generate([]byte(adminPassword), []byte(fullSalt))
	if err != nil {
		return fmt.Errorf("error while generating password hash: %w", err)
	}

	key := strings.Split(hash, "$")[3]
	aesKey := sha256.Sum256([]byte(key))

	t.Cipher, err = aes.NewCipher(aesKey[:])
	if err != nil {
		return fmt.Errorf("failed to create cipher: %w", err)
	}

	fullNewSalt := fmt.Sprintf("$1$%s$", newsalt)
	signHash, err := m.Generate([]byte(key+fmt.Sprint(tokenTime)), []byte(fullNewSalt))
	if err != nil {
		return fmt.Errorf("error while generating sign hash: %w", err)
	}

	t.Sign = strings.Split(signHash, "$")[3]
	return nil
}

// initializeWriteAccess initializes write access for the token. The caller must hold the mutex.
func (t *WhatsminerAccessToken) initializeWriteAccess(adminPassword string) error {

	tokenInfo, err := t.getTokenInfo()
	if err != nil {
		return fmt.Errorf("failed to get token info: %w", err)
	}

	if err := t.generateCipherAndSign(tokenInfo, adminPassword); err != nil {
		return fmt.Errorf("failed to generate cipher and sign: %w", err)
	}

	t.Created = time.Now()
	return nil
}

// monitorToken runs in the background to keep the token fresh.
func (t *WhatsminerAccessToken) monitorToken() {
	for {
		t.mu.Lock()
		// Calculate duration until next refresh
		nextRefreshTime := t.Created.Add(25 * time.Minute)
		duration := time.Until(nextRefreshTime)
		t.mu.Unlock()

		if duration < 0 {
			duration = 0
		}

		timer := time.NewTimer(duration)

		select {
		case <-timer.C:
			t.mu.Lock()
			// Check if a refresh is still needed, as another thread might have done it.
			if time.Since(t.Created).Minutes() >= 25 {
				if err := t.initializeWriteAccess(t.AdminPassword); err != nil {
					if t.OnRefreshError != nil {
						t.OnRefreshError(err)
					} else {
						log.Printf("background token refresh failed: %v", err)
					}
				}
			}
			t.mu.Unlock()
		case <-t.stop:
			timer.Stop()
			return
		}
	}
}

// EnableWriteAccess enables write access for the token and starts the background refresh.
func (t *WhatsminerAccessToken) EnableWriteAccess(adminPassword string) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.AdminPassword = adminPassword
	if err := t.initializeWriteAccess(adminPassword); err != nil {
		return fmt.Errorf("error enabling write access: %w", err)
	}
	go t.monitorToken()
	return nil
}

// HasWriteAccess checks write access and refreshes the token if necessary.
func (t *WhatsminerAccessToken) HasWriteAccess() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.AdminPassword == "" {
		return errors.New("admin password is not set")
	}

	if time.Since(t.Created).Minutes() > 30 {
		// Writeable token has expired; reinitialize
		if err := t.initializeWriteAccess(t.AdminPassword); err != nil {
			return fmt.Errorf("error trying to renew write access: %w", err)
		}
	}

	return nil
}

// WhatsminerAPI represents a stateless class with only class methods for read/write API calls.
type WhatsminerAPI struct{}

// GetReadOnlyInfo sends a READ-ONLY API command.
func (w *WhatsminerAPI) GetReadOnlyInfo(accessToken *WhatsminerAccessToken, cmd string, additionalParams map[string]any) (map[string]any, error) {
	jsonCmd := map[string]any{"cmd": cmd}
	maps.Copy(jsonCmd, additionalParams)

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(accessToken.IPAddress, fmt.Sprintf("%d", accessToken.Port)), 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to miner: %w", err)
	}
	defer conn.Close()

	if err := json.NewEncoder(conn).Encode(jsonCmd); err != nil {
		return nil, fmt.Errorf("failed to encode command: %w", err)
	}

	resp, err := io.ReadAll(conn)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	sanitizedResp := sanitizeJSONResponse(string(resp))

	var result map[string]any
	if err := json.Unmarshal([]byte(sanitizedResp), &result); err != nil {
		return nil, fmt.Errorf("error while trying to unmarshal resp: %w", err)
	}

	return result, nil
}

// ExecCommand sends a WRITEABLE API command.
func (w *WhatsminerAPI) ExecCommand(accessToken *WhatsminerAccessToken, cmd string, additionalParams map[string]any) (map[string]any, error) {
	if err := accessToken.HasWriteAccess(); err != nil {
		return nil, fmt.Errorf("token has no write access: %w", err)
	}

	accessToken.mu.Lock()
	defer accessToken.mu.Unlock()

	if accessToken.Cipher == nil {
		return nil, errors.New("cipher not initialized - write access may have failed")
	}

	jsonCmd := map[string]any{"cmd": cmd, "token": accessToken.Sign}
	maps.Copy(jsonCmd, additionalParams)

	apiCmd, err := json.Marshal(jsonCmd)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal command: %w", err)
	}

	paddedCmd := addTo16(apiCmd)
	dst := make([]byte, len(paddedCmd))
	mode := ecb.NewECBEncrypter(accessToken.Cipher)
	mode.CryptBlocks(dst, paddedCmd)

	encStr := base64.StdEncoding.EncodeToString(dst)
	encStr = strings.ReplaceAll(encStr, "\n", "")

	dataEnc := map[string]any{
		"enc":  1,
		"data": encStr,
	}

	conn, err := net.Dial("tcp", net.JoinHostPort(accessToken.IPAddress, fmt.Sprintf("%d", accessToken.Port)))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to miner: %w", err)
	}
	defer conn.Close()

	if err := json.NewEncoder(conn).Encode(dataEnc); err != nil {
		return nil, fmt.Errorf("error encoding data: %w", err)
	}

	resp, err := io.ReadAll(conn)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	sanitizedResp := sanitizeJSONResponse(string(resp))

	var result map[string]any
	if err := json.Unmarshal([]byte(sanitizedResp), &result); err != nil {
		return nil, fmt.Errorf("error while trying to unmarshal resp: %w", err)
	}

	if status, ok := result["STATUS"].(string); ok && status == "E" {
		if msg, ok := result["Msg"].(string); ok {
			return nil, fmt.Errorf("miner API error: %s", msg)
		}
		return nil, errors.New("unknown miner API error")
	}

	encResult, ok := result["enc"].(string)
	if !ok {
		return nil, errors.New("encrypted response not found")
	}

	respCiphertext, err := base64.StdEncoding.DecodeString(encResult)
	if err != nil {
		return nil, fmt.Errorf("failed to decode encrypted response: %w", err)
	}

	respPlaintext, err := decrypt(string(respCiphertext), accessToken.Cipher)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt response: %w", err)
	}

	respFinal := strings.Split(respPlaintext, "\x00")[0]
	if err := json.Unmarshal([]byte(respFinal), &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal decrypted response: %w", err)
	}

	return result, nil
}

func decrypt(cipherstring string, block cipher.Block) (string, error) {
	ciphertext := []byte(cipherstring)

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext is too short")
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	mode := ecb.NewECBDecrypter(block)
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = PKCS5UnPadding(ciphertext)

	return string(ciphertext), nil
}

// sanitizeJSONResponse replaces non-standard JSON values from the miner response.
func sanitizeJSONResponse(resp string) string {
	resp = strings.ReplaceAll(resp, "inf", "999")
	resp = strings.ReplaceAll(resp, "nan", "0")
	return resp
}

// addTo16 pads a byte slice to a multiple of 16 bytes with null bytes.
func addTo16(b []byte) []byte {
	padSize := 16 - (len(b) % 16)
	if padSize == 16 {
		return b
	}
	pad := make([]byte, padSize)
	return append(b, pad...)
}

// PKCS5UnPadding removes PKCS5 padding from a byte slice.
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	if length == 0 {
		return nil
	}
	unpadding := int(src[length-1])
	if unpadding > length {
		return nil
	}
	return src[:(length - unpadding)]
}
