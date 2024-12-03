// Пакет с логером
package logger

import (
    "go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error
    logger, err = zap.NewProduction()
    if err != nil {
        panic(err)
    }
    defer logger.Sync()
}

// Info - логгирование на уровне info
func Info(message string) {
	logger.Info(message)
}

// Panic - логгирование с вызовом паники
func Panic(err error) {
	logger.Panic(err.Error())
}
