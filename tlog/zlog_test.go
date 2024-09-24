package main

import (
	"go.uber.org/zap"
	"testing"
)

func TestInfo(t *testing.T) {
	logger, err := zap.NewProduction()
	if err != nil {
		return
	}
	logger.Info("test", zap.String("key", "value"))
}
