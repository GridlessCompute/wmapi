package client

import (
	"encoding/json"
	"fmt"
	"strconv"
	"wmapi/transport"
)

const (
	LowPower    = "set_low_power"
	NormalPower = "set_normal_power"
	HighPower   = "set_high_power"
)

type WriteAPI struct {
	API   *transport.WhatsminerAPI
	Token *transport.WhatsminerAccessToken
}

type CustomLedSettings struct {
	Color    string `json:"color"`
	Period   int    `json:"period"`
	Duration int    `json:"duration"`
	Start    int    `json:"start"`
}

type CustomNetworkSettings struct {
	Ip   string
	Mask string
	Gate string
	Dns  string
	Host string
}

func (w *WriteAPI) Pools(pools ...Pool) (*CommandResponse, error) {
	if len(pools) == 0 || len(pools) > 3 {
		return nil, fmt.Errorf("you must provide between 1 and 3 pools")
	}

	params := make(map[string]any)
	for i, p := range pools {
		if p.URL == "" || p.Worker == "" {
			return nil, fmt.Errorf("pool URL and worker cannot be empty for pool %d", i+1)
		}
		params[fmt.Sprintf("pool%d", i+1)] = p.URL
		params[fmt.Sprintf("user%d", i+1)] = p.Worker
		params[fmt.Sprintf("passwd%d", i+1)] = p.Password
	}

	result, err := w.API.ExecCommand(w.Token, "pools", params)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

// Reboot initiates a reboot of the miner
func (w *WriteAPI) Restart() (*CommandResponse, error) {
	result, err := w.API.ExecCommand(w.Token, "restart_btminer", nil)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) PowerOffHashboard() (*CommandResponse, error) {
	result, err := w.API.ExecCommand(w.Token, "power_off", nil)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) PowerOnHashboard() (*CommandResponse, error) {
	result, err := w.API.ExecCommand(w.Token, "power_on", nil)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) ManageLedRestore(mode string) (*CommandResponse, error) {
	param := map[string]any{"param": mode}

	result, err := w.API.ExecCommand(w.Token, "set_led", param)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) ManageLedCustom(settings CustomLedSettings) (*CommandResponse, error) {
	param := map[string]any{
		"color":    settings.Color,
		"period":   settings.Period,
		"duration": settings.Duration,
		"start":    settings.Start,
	}
	result, err := w.API.ExecCommand(w.Token, "set_led", param)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) SwitchPowerMode(mode string) (*CommandResponse, error) {
	result, err := w.API.ExecCommand(w.Token, mode, nil)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) RebootSystem() (*CommandResponse, error) {
	result, err := w.API.ExecCommand(w.Token, "reboot", nil)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) RestoreFactorySettings() (*CommandResponse, error) {
	result, err := w.API.ExecCommand(w.Token, "factory_reset", nil)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) ModifyPassword(oldPwd, newPwd string) (*CommandResponse, error) {
	param := map[string]any{
		"old": oldPwd,
		"new": newPwd,
	}

	result, err := w.API.ExecCommand(w.Token, "factory_reset", param)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) NetworkSetDHCP() (*CommandResponse, error) {
	param := map[string]any{"param": "dhcp"}
	result, err := w.API.ExecCommand(w.Token, "factory_reset", param)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) NetworkSetCustom(conf CustomNetworkSettings) (*CommandResponse, error) {
	param := map[string]any{
		"ip":   conf.Ip,
		"mask": conf.Mask,
		"gate": conf.Gate,
		"dns":  conf.Dns,
		"host": conf.Host,
	}

	result, err := w.API.ExecCommand(w.Token, "factory_reset", param)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) TargetFreq(tgt int) (*CommandResponse, error) {
	tgt = min(tgt, 100)
	tgt = max(tgt, -100)
	param := map[string]any{"percent": tgt}
	result, err := w.API.ExecCommand(w.Token, "factory_reset", param)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) EnableFastboot() (*CommandResponse, error) {
	result, err := w.API.ExecCommand(w.Token, "enable_btminer_fast_boot", nil)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) Disablefastboot() (*CommandResponse, error) {
	result, err := w.API.ExecCommand(w.Token, "disable_btminer_fast_boot", nil)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) EnableWebPools() (*CommandResponse, error) {
	result, err := w.API.ExecCommand(w.Token, "enable_web_pools", nil)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) DisableWebPools() (*CommandResponse, error) {
	result, err := w.API.ExecCommand(w.Token, "disable_web_pools", nil)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) ChangeHostName(name string) (*CommandResponse, error) {
	param := map[string]any{"hostname": name}
	result, err := w.API.ExecCommand(w.Token, "set_hostname", param)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) PowerPercent(pct int) (*CommandResponse, error) {
	pct = min(pct, 100)
	pct = max(pct, 0)

	pctStr := strconv.Itoa(pct)

	param := map[string]any{"percent": pctStr}
	result, err := w.API.ExecCommand(w.Token, "set_power_pct", param)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) PowerPercentV2(pct int) (*CommandResponse, error) {
	pct = min(pct, 100)
	pct = max(pct, 0)

	pctStr := strconv.Itoa(pct)

	param := map[string]any{"percent": pctStr}
	result, err := w.API.ExecCommand(w.Token, "set_power_pct_v2", param)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) TempOffset(offset int) (*CommandResponse, error) {
	offset = min(offset, 0)
	offset = max(offset, -30)

	pctStr := strconv.Itoa(offset)

	param := map[string]any{"temp_offset": pctStr}
	result, err := w.API.ExecCommand(w.Token, "set_temp_offset", param)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) AdjPowerLimit(limit int) (*CommandResponse, error) {
	limit = min(limit, 99999)
	limit = max(limit, 0)

	pctStr := strconv.Itoa(limit)

	param := map[string]any{"power_limit": pctStr}
	result, err := w.API.ExecCommand(w.Token, "adjust_power_limit", param)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) AdjUpfreqSpeed(speed int) (*CommandResponse, error) {
	speed = min(speed, 9)
	speed = max(speed, 0)

	pctStr := strconv.Itoa(speed)

	param := map[string]any{"upfreq_speed": pctStr}
	result, err := w.API.ExecCommand(w.Token, "adjust_upfreq_speed", param)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) PowerOffCool(cool bool) (*CommandResponse, error) {
	c := "0"

	if cool {
		c = "1"
	}

	param := map[string]any{"poweroff_cool": c}
	result, err := w.API.ExecCommand(w.Token, "set_poweroff_cool", param)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) FanZeroSpeed(zero bool) (*CommandResponse, error) {
	z := "0"

	if zero {
		z = "1"
	}

	param := map[string]any{"fan_zero_speed": z}
	result, err := w.API.ExecCommand(w.Token, "set_fan_zero_speed", param)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) DisableBTMinerInit() (*CommandResponse, error) {
	result, err := w.API.ExecCommand(w.Token, "disbale_btminer_init", nil)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}

func (w *WriteAPI) EnableBTMinerInit() (*CommandResponse, error) {
	result, err := w.API.ExecCommand(w.Token, "enable_btminer_init", nil)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to json: %w", err)
	}

	var cmdResponse CommandResponse
	if err := json.Unmarshal(jsonBytes, &cmdResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json to CommandResponse: %w", err)
	}

	return &cmdResponse, nil
}