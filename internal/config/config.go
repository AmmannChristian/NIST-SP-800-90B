package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config holds the microservice configuration
type Config struct {
	// Server configuration
	ServerPort  int
	ServerHost  string
	GRPCEnabled bool
	GRPCPort    int

	// Logging
	LogLevel string

	// File upload limits
	MaxUploadSize int64 // in bytes

	// Request timeouts
	Timeout time.Duration

	// Metrics
	MetricsEnabled bool

	// Authentication
	AuthEnabled  bool
	AuthIssuer   string
	AuthAudience string
	AuthJWKSURL  string
}

// LoadConfig loads configuration from environment variables with defaults
func LoadConfig() (*Config, error) {
	config := &Config{
		// Defaults
		ServerPort:     getEnvAsInt("SERVER_PORT", 8080),
		ServerHost:     getEnv("SERVER_HOST", "0.0.0.0"),
		GRPCEnabled:    getEnvAsBool("GRPC_ENABLED", false),
		GRPCPort:       getEnvAsInt("GRPC_PORT", 9090),
		LogLevel:       getEnv("LOG_LEVEL", "info"),
		MaxUploadSize:  getEnvAsInt64("MAX_UPLOAD_SIZE", 100*1024*1024), // 100MB default
		Timeout:        getEnvAsDuration("TIMEOUT", 5*time.Minute),
		MetricsEnabled: getEnvAsBool("METRICS_ENABLED", true),
		AuthEnabled:    getEnvAsBool("AUTH_ENABLED", false),
		AuthIssuer:     getEnv("AUTH_ISSUER", ""),
		AuthAudience:   getEnv("AUTH_AUDIENCE", ""),
		AuthJWKSURL:    getEnv("AUTH_JWKS_URL", ""),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.ServerPort < 1 || c.ServerPort > 65535 {
		return fmt.Errorf("invalid server port: %d (must be 1-65535)", c.ServerPort)
	}

	if c.GRPCEnabled && (c.GRPCPort < 1 || c.GRPCPort > 65535) {
		return fmt.Errorf("invalid gRPC port: %d (must be 1-65535)", c.GRPCPort)
	}

	if c.MaxUploadSize < 1024 {
		return fmt.Errorf("max upload size too small: %d (must be at least 1024 bytes)", c.MaxUploadSize)
	}

	validLogLevels := map[string]bool{
		"debug": true,
		"info":  true,
		"warn":  true,
		"error": true,
	}
	if !validLogLevels[c.LogLevel] {
		return fmt.Errorf("invalid log level: %s (must be debug, info, warn, or error)", c.LogLevel)
	}

	if c.AuthEnabled {
		if !c.GRPCEnabled {
			return fmt.Errorf("authentication requires gRPC to be enabled")
		}
		if c.AuthIssuer == "" {
			return fmt.Errorf("invalid auth issuer: required when AUTH_ENABLED=true")
		}
		if c.AuthAudience == "" {
			return fmt.Errorf("invalid auth audience: required when AUTH_ENABLED=true")
		}
	}

	return nil
}

// Helper functions to read environment variables

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		return defaultValue
	}
	return value
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := time.ParseDuration(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
