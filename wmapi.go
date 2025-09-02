// Package wmapi provides a high-level interface for interacting with Whatsminer ASICs.
// This file contains the middleware layer that simplifies interaction with the API.
package wmapi

import (
	"fmt"
	"wmapi/client"
	"wmapi/transport"
)

type WhatsminerMiddleware struct {
	API         *transport.WhatsminerAPI
	AccessToken *transport.WhatsminerAccessToken
	Read        *client.ReadAPI
	Write       *client.WriteAPI
}

func NewWhatsminerAPI(ipAddress string, port int, adminPassword string) (*WhatsminerMiddleware, error) {
	token, err := transport.NewWhatsminerAccessToken(ipAddress, port, adminPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}

	api := &transport.WhatsminerAPI{}

	mw := &WhatsminerMiddleware{
		API:         api,
		AccessToken: token,
		Read:        &client.ReadAPI{API: api, Token: token},
		Write:       &client.WriteAPI{API: api, Token: token},
	}

	return mw, nil
}