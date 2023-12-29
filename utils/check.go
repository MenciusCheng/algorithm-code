package utils

import (
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"go.uber.org/zap"
)

func CheckGoPanic() {
	if r := recover(); r != nil {
		log.DPanic("panic recovered ", zap.Any("panic", r))
	}
}
