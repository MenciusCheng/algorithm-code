package utils

import (
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"go.uber.org/zap"
	"testing"
)

func TestCheckGoPanic_mid_panic(t *testing.T) {
	defer CheckGoPanic()
	log.Info("do first")
	panic("do panic")
	log.Info("do second")
}

func TestCheckGoPanic_panic_for(t *testing.T) {
	defer CheckGoPanic()
	for i := 0; i < 10; i++ {
		log.Info("do first", zap.Int("i", i))
		if i == 0 {
			// panic 后直接返回，循环终止
			panic("do panic")
		}
		log.Info("do second", zap.Int("i", i))
	}
}

func TestCheckGoPanic_panic_for_sub(t *testing.T) {
	for i := 0; i < 10; i++ {
		func() {
			defer CheckGoPanic()

			log.Info("do first", zap.Int("i", i))
			if i == 0 {
				// panic 后直接返回子函数，循环继续
				panic("do panic")
			}
			log.Info("do second", zap.Int("i", i))
		}()
	}
}
