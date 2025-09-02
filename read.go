package wmapi

import (
	"encoding/json"
	"fmt"
)

type ReadAPI struct {
	mw *WhatsminerMiddleware
}

// Summary retrieves the miner's summary information
func (r *ReadAPI) Summary() (*SummaryResponse, error) {
	data, err := r.mw.api.GetReadOnlyInfo(r.mw.accessToken, "summary", nil)
	if err != nil {
		return nil, err
	}
	var summary SummaryResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}
	if err := json.Unmarshal(jsonData, &summary); err != nil {
		return nil, fmt.Errorf("failed to unmarshal summary: %w", err)
	}
	return &summary, nil
}

// Pools retrieves the configured mining pools
func (r *ReadAPI) Pools() (*PoolsResponse, error) {
	data, err := r.mw.api.GetReadOnlyInfo(r.mw.accessToken, "pools", nil)
	if err != nil {
		return nil, err
	}
	var pools PoolsResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}
	if err := json.Unmarshal(jsonData, &pools); err != nil {
		return nil, fmt.Errorf("failed to unmarshal pools: %w", err)
	}
	return &pools, nil
}

func (r *ReadAPI) Edevs() (*EdevsResponse, error) {
	data, err := r.mw.api.GetReadOnlyInfo(r.mw.accessToken, "edevs", nil)
	if err != nil {
		return nil, err
	}
	var edevs EdevsResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}
	if err := json.Unmarshal(jsonData, &edevs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal edevs: %w", err)
	}
	return &edevs, nil
}

func (r *ReadAPI) DevDetails() (*DevdetailsResponse, error) {
	data, err := r.mw.api.GetReadOnlyInfo(r.mw.accessToken, "devdetails", nil)
	if err != nil {
		return nil, err
	}
	var devdetails DevdetailsResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}
	if err := json.Unmarshal(jsonData, &devdetails); err != nil {
		return nil, fmt.Errorf("failed to unmarshal devdetails: %w", err)
	}
	return &devdetails, nil
}

func (r *ReadAPI) PSU() (*PSUResponse, error) {
	data, err := r.mw.api.GetReadOnlyInfo(r.mw.accessToken, "get_psu", nil)
	if err != nil {
		return nil, err
	}
	var psu PSUResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}
	if err := json.Unmarshal(jsonData, &psu); err != nil {
		return nil, fmt.Errorf("failed to unmarshal psu: %w", err)
	}
	return &psu, nil
}

func (r *ReadAPI) Version() (*VersionResponse, error) {
	data, err := r.mw.api.GetReadOnlyInfo(r.mw.accessToken, "get_version", nil)
	if err != nil {
		return nil, err
	}
	var version VersionResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}
	if err := json.Unmarshal(jsonData, &version); err != nil {
		return nil, fmt.Errorf("failed to unmarshal version: %w", err)
	}
	return &version, nil
}

func (r *ReadAPI) Status() (*StatusResponse, error) {
	data, err := r.mw.api.GetReadOnlyInfo(r.mw.accessToken, "status", nil)
	if err != nil {
		return nil, err
	}
	var status StatusResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}
	if err := json.Unmarshal(jsonData, &status); err != nil {
		return nil, fmt.Errorf("failed to unmarshal status: %w", err)
	}
	return &status, nil
}

func (r *ReadAPI) ErrorCode() (*ErrorResponse, error) {
	data, err := r.mw.api.GetReadOnlyInfo(r.mw.accessToken, "get_error_code", nil)
	if err != nil {
		return nil, err
	}
	var errorResponse ErrorResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}
	if err := json.Unmarshal(jsonData, &errorResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal error response: %w", err)
	}
	return &errorResponse, nil
}
