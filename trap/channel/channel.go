package channel

import (
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"go.uber.org/zap"
	"time"
)

func ExampleForSelect() {
	ch := make(chan int, 100)

	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
			//log.Info("write ch", zap.Any("value", i))
		}
		time.Sleep(1 * time.Second)
		close(ch)
		log.Info("ch close")
	}()
	time.Sleep(1 * time.Millisecond)
	//res := make([]int, 0)
	exist := true
	for exist {
		//time.Sleep(1 * time.Millisecond)
		exist = false
		select {
		case i, ok := <-ch:
			log.Info("read ch", zap.Any("value", i), zap.Bool("ok", ok))
			exist = true
			//res = append(res, i)
		default:
			log.Info("default case")
		}
	}
	log.Info("finish")
	time.Sleep(1 * time.Second)
}

// 不带缓冲区，交替打印
func ExampleFor() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
			log.Info("write ch", zap.Any("value", i))
		}
		time.Sleep(1 * time.Second)
		close(ch)
		log.Info("ch close")
	}()
	for i := range ch {
		log.Info("read ch", zap.Any("value", i))
	}
	log.Info("finish")
}

// 带缓冲区，混合打印
func ExampleForBuff() {
	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
			log.Info("write ch", zap.Any("value", i))
		}
		time.Sleep(1 * time.Second)
		close(ch)
		log.Info("ch close")
	}()
	for i := range ch {
		log.Info("read ch", zap.Any("value", i))
	}
	log.Info("finish")
}
