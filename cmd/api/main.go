package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"go-devops/api/routes"
	"go-devops/config"
	"go-devops/internal/models"
)

func main() {
	// Khởi tạo logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Đọc cấu hình
	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("Không thể đọc cấu hình: %v", err)
	}

	// Kết nối database
	db, err := config.InitDB(cfg)
	if err != nil {
		logger.Fatalf("Không thể kết nối database: %v", err)
	}

	// Auto migrate
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		logger.Fatalf("Không thể migrate database: %v", err)
	}

	// Thiết lập router
	router := routes.SetupRouter(db, logger)

	// Khởi động server
	serverAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	logger.Infof("Server đang chạy tại %s", serverAddr)

	if err := router.Run(serverAddr); err != nil {
		logger.Fatalf("Không thể khởi động server: %v", err)
	}
}
