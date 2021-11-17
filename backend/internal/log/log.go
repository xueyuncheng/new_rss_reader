package log

import (
	"go.uber.org/zap"
)

var Sugar *zap.SugaredLogger

func InitLog() {
	// logger, err := zap.NewProduction()
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()
	Sugar = logger.Sugar()
}
