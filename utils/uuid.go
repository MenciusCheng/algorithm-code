package utils

import (
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// UUID v4版本 基于随机数
func GetUUIDV4() string {
	randomUUID, err := uuid.NewRandom()
	if err != nil {
		log.Error("GetUUIDV4 err", zap.Error(err))
		return ""
	}
	return randomUUID.String()
}
