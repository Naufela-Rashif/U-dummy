package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
	// Menggunakan viper untuk memuat file konfigurasi
	viper.SetConfigFile("config.yaml") // Tentukan file konfigurasi
	viper.AddConfigPath(".")           // Cari di direktori saat ini

	// Membaca file konfigurasi
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error config file: %s \n", err)
	}

	// Membaca konfigurasi database ke dalam struct
	var dbConfig DatabaseConfig
	if err := viper.UnmarshalKey("database", &dbConfig); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %v", err)
	}

	return &dbConfig, nil
}
