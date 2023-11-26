package async_queue

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"reflect"
	"sync"
	"testing"
	"time"
)

func InitClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码为空
		DB:       0,  // 使用默认数据库
	})
	return client
}

func TestAsyncQueue_Client(t *testing.T) {
	client := InitClient()

	queueEntity := NewAsyncQueueEntity(
		ConfigName("TestQueueA"),
		ConfigRedisConn(client),
	)
	total := 5

	for i := 0; i < total; i++ {
		data := fmt.Sprintf("{\"id\":%d}", i)
		_ = queueEntity.Publish(Task{
			Payload: []byte(data),
		})
		time.Sleep(10 * time.Millisecond)
	}
}

func TestAsyncQueue_Server(t *testing.T) {
	client := InitClient()

	myHandler := func(param *HandlerParam) {
		log.Println("处理任务", "data", string(param.Payload))
	}

	queueEntity := NewAsyncQueueEntity(
		ConfigName("TestQueueA"),
		ConfigMaxQueueConcurrencyNum(3),
		ConfigRedisConn(client),
		ConfigHandler(myHandler),
	)

	// 初始化调用队列数据
	err := queueEntity.InitQueue()
	if err != nil {
		t.Errorf("InitQueue err: %v", err)
		return
	}

	time.Sleep(60 * time.Second)
}

func TestAsyncQueue_DAG(t *testing.T) {
	client := InitClient()

	myHandler := func(param *HandlerParam) {
		log.Println("处理任务", "data", string(param.Payload))
	}

	queueEntity := NewAsyncQueueEntity(
		ConfigName("TestQueueA"),
		ConfigMaxQueueConcurrencyNum(3),
		ConfigRedisConn(client),
		ConfigHandler(myHandler),
	)

	total := 10
	tasks := make([]Task, 0, total)
	for i := 0; i < total; i++ {
		data := fmt.Sprintf("{\"id\":%d}", i)
		tasks = append(tasks, Task{
			Type:    TaskTypeDAG,
			Payload: []byte(data),
		})
	}
	queueEntity.PublishDAG(tasks)
	time.Sleep(1 * time.Second)

	// 初始化调用队列数据
	err := queueEntity.InitQueue()
	if err != nil {
		t.Errorf("InitQueue err: %v", err)
		return
	}

	time.Sleep(60 * time.Second)
}

func TestTaskQueueEntity_Run(t *testing.T) {
	client := InitClient()

	wg := &sync.WaitGroup{}
	lock := &sync.RWMutex{}
	got := make(map[string]bool)
	myHandler := func(param *HandlerParam) {
		defer func() {
			wg.Done()
		}()
		log.Println("处理任务", "data", string(param.Payload))
		lock.Lock()
		got[string(param.Payload)] = true
		lock.Unlock()
		time.Sleep(10 * time.Millisecond)
	}
	total := 10
	wg.Add(total)

	queueEntity := NewAsyncQueueEntity(
		ConfigName("TestQueueA"),
		ConfigMaxQueueConcurrencyNum(10),
		ConfigRedisConn(client),
		ConfigHandler(myHandler),
	)

	// 初始化调用队列数据
	err := queueEntity.InitQueue()
	if err != nil {
		t.Errorf("InitQueue err: %v", err)
		return
	}

	want := make(map[string]bool)
	go func() {
		// 推送消息
		for i := 0; i < total; i++ {
			data := fmt.Sprintf("{\"id\":%d}", i)
			_ = queueEntity.Publish(Task{
				Payload: []byte(data),
			})
			want[data] = true
			time.Sleep(10 * time.Millisecond)
		}
		log.Println("推送消息完成")
	}()

	wg.Wait()
	time.Sleep(2 * time.Second)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
		return
	}
	log.Println("finish", "got len", len(got), "want len", len(want))
}
