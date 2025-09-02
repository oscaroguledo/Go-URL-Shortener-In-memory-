package core

type SettingStruct struct {
	// Placeholder for future settings
	Domain                     string   `envconfig:"APP_DOMAIN" default:"localhost"`
	Environment                string   `envconfig:"APP_ENVIRONMENT" default:"local"`
	Port                       string   `envconfig:"APP_PORT" default:"8080"`
	RateLimitRequestsPerMinute int      `envconfig:"RATE_LIMIT_REQUESTS_PER_MINUTE" default:"60"`
	RateLimitWhitelist         []string `envconfig:"RATE_LIMIT_WHITELIST" default:""`
	DatabaseURL                string   `envconfig:"DATABASE_URL" default:"urls.db"`
	EnableDetailedLogging      bool     `envconfig:"ENABLE_DETAILED_LOGGING" default:"false"`
	EnableCORS                 bool     `envconfig:"ENABLE_CORS" default:"true"`
	CORSAllowedOrigins         []string `envconfig:"CORS_ALLOWED_ORIGINS" default:"*"`
	ShortCodeLength            int      `envconfig:"SHORT_CODE_LENGTH" default:"6"`
	URLExpirationDays          int      `envconfig:"URL_EXPIRATION_DAYS" default:"30"`
	EnableHTTPS                bool     `envconfig:"ENABLE_HTTPS" default:"false"`
}

func Settings() *SettingStruct {
	var settings SettingStruct
	return &settings
}

// Example usage:
// func main() {
//     settings := Settings()
//     // use settings here
// }
