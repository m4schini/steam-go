package internal

import "go.uber.org/zap"

func Logger(name string) *zap.Logger {
	return zap.L().Named("steam").Named(name)
}
