package models

import (
	"strings"
)

type Backend struct {
	UrlPattern          string   `json:"url_pattern"`
	Encoding            string   `json:"encoding"`
	Method              string   `json:"method"`
	Host                []string `json:"host"`
	DisableHostSanitize bool     `json:"disable_host_sanitize"`
}

func NewBackend(host string, endpoint string, method string, outputEncoding string) Backend {
	return Backend{
		UrlPattern:          endpoint,
		Encoding:            outputEncoding,
		Method:              strings.ToUpper(method),
		Host:                []string{host},
		DisableHostSanitize: false,
	}
}

type Endpoint struct {
	Endpoint          string    `json:"endpoint"`
	Method            string    `json:"method"`
	OutputEncoding    string    `json:"output_encoding"`
	Timeout           string    `json:"timeout"`
	QuerystringParams []string  `json:"querystring_params"`
	Backend           []Backend `json:"backend"`
	HeadersToPass     []string  `json:"headers_to_pass"`
}

func NewEndpoint(host string, endpoint string, backendEndpoint string, method string, outputEncoding string, timeout string) Endpoint {
	backend := NewBackend(host, backendEndpoint, method, outputEncoding)
	return Endpoint{
		Endpoint:          endpoint,
		Method:            strings.ToUpper(method),
		OutputEncoding:    outputEncoding,
		Timeout:           timeout,
		QuerystringParams: []string{},
		Backend:           []Backend{backend},
		HeadersToPass:     []string{"Content-Type"},
	}
}

func (e *Endpoint) InsertQuerystringParams(param string) {
	e.QuerystringParams = append(e.QuerystringParams, param)
}

func (e *Endpoint) InsertHeadersToPass(header string) {
	e.HeadersToPass = append(e.HeadersToPass, header)
}

type Configuration struct {
	Version        string                 `json:"version"`
	Timeout        string                 `json:"timeout"`
	CacheTtl       string                 `json:"cache_ttl"`
	OutputEncoding string                 `json:"output_encoding"`
	Name           string                 `json:"name"`
	Endpoints      []Endpoint             `json:"endpoints"`
	ExtraConfig    map[string]interface{} `json:"extra_config,omitempty"`
}

func NewConfiguration(outputEncoding string, timeout string) Configuration {
	var extraConfig = make(map[string]interface{}, 15)

	if getEnv("ENABLE_CORS", "false") == "true" {
		extraConfig["github_com/devopsfaith/krakend-cors"] = NewCors()
	}
	if getEnv("ENABLE_LOGGING", "false") == "true" {
		extraConfig["github_com/devopsfaith/krakend-gologging"] = NewLogging()
	}

	return Configuration{
		Version:        "2",
		Timeout:        timeout,
		CacheTtl:       "300s",
		OutputEncoding: outputEncoding,
		Name:           "Tenera API",
		Endpoints:      []Endpoint{},
		ExtraConfig:    extraConfig,
	}
}

func (c *Configuration) InsertEndpoint(endpoint Endpoint) {
	c.Endpoints = append(c.Endpoints, endpoint)
}
