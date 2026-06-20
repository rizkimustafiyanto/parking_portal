package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort string

	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	DBSSLMode string
	Timezone  string

	DBMaxOpenConns int
	DBMaxIdleConns int
	CORSOrigins    []string

	JWTSecret string

	AutoMigrate       bool
	SeedDatabase      bool
	SeedAdminName     string
	SeedAdminEmail    string
	SeedAdminPassword string
	SeedAdminRole     string
}

func Load() (*Config, error) {
	viper.AutomaticEnv()

	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_NAME", "mydb")
	viper.SetDefault("DB_SSLMODE", "disable")
	viper.SetDefault("JWT_SECRET", "secret")
	viper.SetDefault("DB_TIMEZONE", "UTC")
	viper.SetDefault("DB_MAX_OPEN_CONNS", 10)
	viper.SetDefault("DB_MAX_IDLE_CONNS", 10)
	viper.SetDefault("CORS_ORIGINS", []string{})
	viper.SetDefault("AUTO_MIGRATE", true)
	viper.SetDefault("SEED_DATABASE", false)
	viper.SetDefault("SEED_ADMIN_NAME", "Administrator")
	viper.SetDefault("SEED_ADMIN_EMAIL", "admin@example.com")
	viper.SetDefault("SEED_ADMIN_PASSWORD", "password123")
	viper.SetDefault("SEED_ADMIN_ROLE", "admin")

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("config file not found, using defaults: %v", err)
	}

	cfg := &Config{
		AppPort:           viper.GetString("APP_PORT"),
		DBHost:            viper.GetString("DB_HOST"),
		DBPort:            viper.GetString("DB_PORT"),
		DBUser:            viper.GetString("DB_USER"),
		DBPass:            viper.GetString("DB_PASSWORD"),
		DBName:            viper.GetString("DB_NAME"),
		DBSSLMode:         viper.GetString("DB_SSLMODE"),
		JWTSecret:         viper.GetString("JWT_SECRET"),
		Timezone:          viper.GetString("DB_TIMEZONE"),
		DBMaxOpenConns:    viper.GetInt("DB_MAX_OPEN_CONNS"),
		DBMaxIdleConns:    viper.GetInt("DB_MAX_IDLE_CONNS"),
		CORSOrigins:       viper.GetStringSlice("CORS_ORIGINS"),
		AutoMigrate:       viper.GetBool("AUTO_MIGRATE"),
		SeedDatabase:      viper.GetBool("SEED_DATABASE"),
		SeedAdminName:     viper.GetString("SEED_ADMIN_NAME"),
		SeedAdminEmail:    viper.GetString("SEED_ADMIN_EMAIL"),
		SeedAdminPassword: viper.GetString("SEED_ADMIN_PASSWORD"),
		SeedAdminRole:     viper.GetString("SEED_ADMIN_ROLE"),
	}

	return cfg, nil
}
