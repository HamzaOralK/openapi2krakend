package models

import (
	"fmt"
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
	fmt.Println(timeout)
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

type Logging struct {
	Level  string `json:"level"`
	Prefix string `json:"Prefix"`
	Syslog bool   `json:"cache_ttl"`
	Stdout bool   `json:"output_encoding"`
	Format string `json:"name"`
}

func NewLogging() Logging {
	return Logging{
		Level:  "WARNING",
		Prefix: "[KRAKEND]",
		Syslog: false,
		Stdout: true,
		Format: "default",
	}
}

type Configuration struct {
	Version        string                 `json:"version"`
	Timeout        string                 `json:"timeout"`
	CacheTtl       string                 `json:"cache_ttl"`
	OutputEncoding string                 `json:"output_encoding"`
	Name           string                 `json:"name"`
	Endpoints      []Endpoint             `json:"endpoints"`
	ExtraConfig    map[string]interface{} `json:"extra_config"`
}

func NewConfiguration(outputEncoding string, timeout string) Configuration {
	return Configuration{
		Version:        "2",
		Timeout:        timeout,
		CacheTtl:       "300s",
		OutputEncoding: outputEncoding,
		Name:           "Tenera API",
		Endpoints:      []Endpoint{},
		ExtraConfig: map[string]interface{}{
			"github_com/devopsfaith/krakend-gologging": NewLogging(),
		},
	}
}

func (c *Configuration) InsertEndpoint(endpoint Endpoint) {
	c.Endpoints = append(c.Endpoints, endpoint)
}
