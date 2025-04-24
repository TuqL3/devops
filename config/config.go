package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config cấu trúc cấu hình ứng dụng
type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"server"`

	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		SSLMode  string `mapstructure:"sslmode"`
	} `mapstructure:"database"`
}

// Load đọc cấu hình từ file hoặc biến môi trường
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	// Đọc từ file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Không thể đọc file cấu hình, sử dụng biến môi trường")
	}

	// Ghi đè bằng biến môi trường
	if os.Getenv("DB_HOST") != "" {
		viper.Set("database.host", os.Getenv("DB_HOST"))
	}
	if os.Getenv("DB_PORT") != "" {
		viper.Set("database.port", os.Getenv("DB_PORT"))
	}
	if os.Getenv("DB_USER") != "" {
		viper.Set("database.user", os.Getenv("DB_USER"))
	}
	if os.Getenv("DB_PASSWORD") != "" {
		viper.Set("database.password", os.Getenv("DB_PASSWORD"))
	}
	if os.Getenv("DB_NAME") != "" {
		viper.Set("database.name", os.Getenv("DB_NAME"))
	}
	if os.Getenv("SERVER_HOST") != "" {
		viper.Set("server.host", os.Getenv("SERVER_HOST"))
	}
	if os.Getenv("SERVER_PORT") != "" {
		viper.Set("server.port", os.Getenv("SERVER_PORT"))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// InitDB khởi tạo kết nối database
func InitDB(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
