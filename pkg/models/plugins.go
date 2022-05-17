package models

import (
	"errors"
	"github.com/okhuz/openapi2krakend/pkg/utility"
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
	LogLevel := strings.ToUpper(utility.GetEnv("LOG_LEVEL", "WARNING"))

	for i := range LogLevels {
		if LogLevels[i] == LogLevel {
			return LogLevel, nil
		}
	}

	return "", errors.New("log level values is wrong")
}

func getLogPrefix() string {
	return utility.GetEnv("LOG_PREFIX", "[KRAKEND]")
}

func getSysLog() bool {
	if strings.ToLower(utility.GetEnv("LOG_SYSLOG", "true")) == "true" {
		return true
	} else {
		return false
	}
}

func getStdout() bool {
	if strings.ToLower(utility.GetEnv("LOG_STDOUT", "true")) == "true" {
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
	allowOrigins := strings.Split(utility.GetEnv("ALLOWED_ORIGINS", "*"), ",")
	var allowedMethods []string
	methodsEnv := utility.GetEnv("ALLOWED_METHODS", "")
	if methodsEnv != "" {
		allowedMethods = strings.Split(methodsEnv, ",")
	} else {
		allowedMethods = []string{
			"GET",
			"HEAD",
			"POST",
			"PUT",
			"DELETE",
			"CONNECT",
			"OPTIONS",
			"TRACE",
			"PATCH",
		}
	}

	return Cors{
		AllowOrigins:  allowOrigins,
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        "12h",
		AllowMethods:  allowedMethods,
	}
}
