package logging

import "go.uber.org/zap"

var logger *zap.Logger

func InitLogger() *zap.Logger {
	logger, _ = zap.NewProduction()
	return logger
}

func GetS() *zap.SugaredLogger {
	defer logger.Sync() // flushes buffer, if any
	return logger.Sugar()
}
