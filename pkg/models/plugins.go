package models

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

type Cors struct {
	AllowOrigins  []string `json:"allow_origins"`
	ExposeHeaders []string `json:"expose_headers"`
	MaxAge        string   `json:"max_age"`
	AllowMethods  []string `json:"allow_methods"`
}

func NewCors(allowOrigins []string) Cors {
	return Cors{
		AllowOrigins:  allowOrigins,
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        "12h",
		AllowMethods: []string{
			"GET",
			"HEAD",
			"POST",
			"PUT",
			"DELETE",
			"CONNECT",
			"OPTIONS",
			"TRACE",
			"PATCH",
		},
	}
}
