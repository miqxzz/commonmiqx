package commonmiqx

import (
	"go.uber.org/zap"
)

// Logger представляет собой обертку над zap.Logger
type Logger struct {
	*zap.Logger
}

// NewLogger создает новый экземпляр логгера
func NewLogger() (*Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &Logger{logger}, nil
}

// NewDevelopmentLogger создает логгер для разработки
func NewDevelopmentLogger() (*Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return &Logger{logger}, nil
}

// Sync синхронизирует буфер логгера
func (l *Logger) Sync() error {
	return l.Logger.Sync()
} 