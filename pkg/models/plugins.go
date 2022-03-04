package models

import (
	"errors"
	"log"
	"strings"
)

type Logging struct {
	Level  string `json:"level"`
	Prefix string `json:"prefix"`
	Syslog bool   `json:"syslog"`
	Stdout bool   `json:"stdout"`
	Format string `json:"format"`
}

func NewLogging() Logging {
	logLevel, e := getLogLevel()
	if e != nil {
		log.Fatalln(e)
	}
	return Logging{
		Level:  logLevel,
		Prefix: getLogPrefix(),
		Syslog: getSysLog(),
		Stdout: getStdout(),
		Format: "default",
	}
}

func getLogLevel() (string, error) {
	LogLevels := []string{"CRITICAL", "DEBUG", "ERROR", "INFO", "WARNING"}
	LogLevel := strings.ToUpper(getEnv("LOG_LEVEL", "WARNING"))

	for i := range LogLevels {
		if LogLevels[i] == LogLevel {
			return LogLevel, nil
		}
	}

	return "", errors.New("log level values is wrong")
}

func getLogPrefix() string {
	return getEnv("LOG_PREFIX", "[KRAKEND]")
}

func getSysLog() bool {
	if strings.ToLower(getEnv("LOG_SYSLOG", "true")) == "true" {
		return true
	} else {
		return false
	}
}

func getStdout() bool {
	if strings.ToLower(getEnv("LOG_STDOUT", "true")) == "true" {
		return true
	} else {
		return false
	}
}

type Cors struct {
	AllowOrigins  []string `json:"allow_origins"`
	ExposeHeaders []string `json:"expose_headers"`
	MaxAge        string   `json:"max_age"`
	AllowMethods  []string `json:"allow_methods"`
}

func NewCors() Cors {
	allowOrigins := strings.Split(getEnv("ALLOWED_ORIGINS", "*"), ",")
	var allowedMethods []string
	methodsEnv := getEnv("ALLOWED_METHODS", "")
	if methodsEnv != "" {
		allowedMethods = strings.Split(methodsEnv, ",")
	} else {
		allowedMethods = []string{
			"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
		}
	}

	return Cors{
		AllowOrigins:  allowOrigins,
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        "12h",
		AllowMethods:  allowedMethods,
	}
}
