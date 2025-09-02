// Package whatsminer provides a high-level interface for interacting with Whatsminer ASICs.
// This file contains the middleware layer that simplifies interaction with the API.
package wmapi

import (
	"fmt"
)

type Pool struct {
	URL      string
	Worker   string
	Password string
}

type WhatsminerMiddleware struct {
	api         *WhatsminerAPI
	accessToken *WhatsminerAccessToken
	Read        *ReadAPI
	Write       *WriteAPI
}

func NewWhatsminerAPI(ipAddress string, port int, adminPassword string) (*WhatsminerMiddleware, error) {
	token, err := NewWhatsminerAccessToken(ipAddress, port, adminPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}

	mw := &WhatsminerMiddleware{
		api:         &WhatsminerAPI{},
		accessToken: token,
	}

	mw.Read = &ReadAPI{mw: mw}
	mw.Write = &WriteAPI{mw: mw}

	return mw, nil
}
