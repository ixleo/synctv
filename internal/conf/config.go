package conf

import (
	"github.com/synctv-org/synctv/utils"
)

type Config struct {
	// Log
	Log LogConfig `yaml:"log"`

	// Server
	Server ServerConfig `yaml:"server"`

	// Jwt
	Jwt JwtConfig `yaml:"jwt"`

	// Database
	Database DatabaseConfig `yaml:"database"`

	// OAuth2
	OAuth2 OAuth2Config `yaml:"oauth2"`

	// RateLimit
	RateLimit RateLimitConfig `yaml:"rate_limit"`

	// Vendor
	Vendor VendorConfig `yaml:"vendor"`
}

func (c *Config) Save(file string) error {
	return utils.WriteYaml(file, c)
}

func DefaultConfig() *Config {
	return &Config{
		// Log
		Log: DefaultLogConfig(),

		// Server
		Server: DefaultServerConfig(),

		// Jwt
		Jwt: DefaultJwtConfig(),

		// Database
		Database: DefaultDatabaseConfig(),

		// OAuth2
		OAuth2: DefaultOAuth2Config(),

		// RateLimit
		RateLimit: DefaultRateLimitConfig(),

		// Vendor
		Vendor: DefaultVendorConfig(),
	}
}
