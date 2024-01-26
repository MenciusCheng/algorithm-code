package utils

import (
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"go.uber.org/zap"
	"testing"
	"time"
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

func TestCheckGoPanic_panic_for_tick(t *testing.T) {
	//defer CheckGoPanic()
	tick := time.NewTicker(1 * time.Second)
	i := 0
	for {
		select {
		case <-tick.C:
			func() {
				// 需要在里面捕获，如果只是最外面捕获，则tick不再继续
				defer CheckGoPanic()
				if i == 1 {
					panic("do panic")
				}
				log.Info("bi", zap.Int("i", i))
				i++
			}()
		}
	}
	log.Info("finish", zap.Int("i", i))
}
