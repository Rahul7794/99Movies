package config

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Provider defines config functions to be implemented
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var defaultConfig *viper.Viper

// Config return config object
func Config() Provider {
	return defaultConfig
}

func init() {
	defaultConfig = readViperConfig()
}

func loadDotEnv() {
	// Earlier files take precedence
	_ = godotenv.Load(".env", ".env-dist")
}

func readViperConfig() *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()

	loadDotEnv()

	return v
}
