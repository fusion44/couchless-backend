package context

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config contains all configuration options
type Config struct {
	AppName string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret   string
	JWTExpireIn time.Duration

	DebugMode bool
	LogFormat string

	ServerPort string

	FileStoragePath string
	Fit2JSONPath    string
}

// LoadConfig loads the configuration file to memory
func LoadConfig(path string) *Config {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error context file: %s \n", err)
	}

	return &Config{
		AppName: config.Get("app-name").(string),

		DBHost:     config.Get("db.host").(string),
		DBPort:     config.Get("db.port").(string),
		DBUser:     config.Get("db.user").(string),
		DBPassword: config.Get("db.password").(string),
		DBName:     config.Get("db.dbname").(string),

		JWTSecret:   config.Get("auth.jwt-secret").(string),
		JWTExpireIn: config.GetDuration("auth.jwt-expire-in"),

		DebugMode: config.Get("log.debug-mode").(bool),
		LogFormat: config.Get("log.log-format").(string),

		ServerPort: config.Get("server.port").(string),

		FileStoragePath: config.Get("system.file-storage-path").(string),
		Fit2JSONPath:    config.Get("system.fit2json-path").(string),
	}
}
