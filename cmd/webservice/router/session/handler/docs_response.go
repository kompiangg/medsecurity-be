package handler

import "medsecurity/pkg/http"

// Default docs response
// Copas to make your work easier

type HTTPBaseResp struct {
	Error *http.HTTPErrorBaseResponse `json:"error"`
}

type HTTPErrResp struct {
	Error HTTPBaseResp
	Data  interface{} `json:"data"`
}
