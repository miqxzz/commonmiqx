package common

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger инициализирует логгер с настройками для вывода в файл и консоль
func InitLogger() (*zap.Logger, error) {
	// Создаем директорию для логов, если она не существует
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, err
	}

	// Создаем конфигурацию для логгера
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Настраиваем вывод в файл
	config.OutputPaths = []string{
		filepath.Join("logs", "app.log"),
		"stdout",
	}

	// Создаем логгер
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
