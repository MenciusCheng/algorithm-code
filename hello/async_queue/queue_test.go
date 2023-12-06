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

	queueEntity := NewAsyncQueue(
		ConfigName("TestQueueA"),
		ConfigRedisConn(client),
	)
	total := 5

	for i := 0; i < total; i++ {
		data := fmt.Sprintf("{\"id\":%d}", i)
		_ = queueEntity.Publish(&TaskInfo{
			Payload: []byte(data),
		})
		time.Sleep(10 * time.Millisecond)
	}
}

func TestAsyncQueue_Server(t *testing.T) {
	client := InitClient()

	myHandler := func(param *TaskInfo) error {
		log.Println("处理任务", "data", string(param.Payload))
		return nil
	}

	queueEntity := NewAsyncQueue(
		ConfigName("TestQueueA"),
		ConfigConcurrency(3),
		ConfigRedisConn(client),
		ConfigHandler(HandlerFunc(myHandler)),
	)

	// 初始化调用队列数据
	err := queueEntity.StartConsuming()
	if err != nil {
		t.Errorf("StartConsuming err: %v", err)
		return
	}

	time.Sleep(60 * time.Second)
}

func TestAsyncQueue_Schedule(t *testing.T) {
	client := InitClient()

	wg := &sync.WaitGroup{}
	myHandler := func(param *TaskInfo) error {
		defer func() {
			wg.Done()
		}()
		log.Println("处理任务", "data", string(param.Payload))
		return nil
	}
	total := 20
	wg.Add(total)

	queueEntity := NewAsyncQueue(
		ConfigName("TestQueueA"),
		ConfigConcurrency(3),
		ConfigRedisConn(client),
		ConfigHandler(HandlerFunc(myHandler)),
	)

	processAt := time.Now().Unix()
	for i := 0; i < total; i++ {
		data := fmt.Sprintf("{\"id\":%d}", i)
		queueEntity.Publish(&TaskInfo{
			TaskMeta: TaskMeta{
				ProcessAt: processAt,
			},
			Payload: []byte(data),
		})
		processAt += 2
	}
	time.Sleep(1 * time.Second)

	// 初始化调用队列数据
	err := queueEntity.StartConsuming()
	if err != nil {
		t.Errorf("StartConsuming err: %v", err)
		return
	}

	wg.Wait()
	log.Println("finish")
	time.Sleep(2 * time.Second)
}

func TestAsyncQueue_List(t *testing.T) {
	client := InitClient()

	wg := &sync.WaitGroup{}
	myHandler := func(param *TaskInfo) error {
		defer func() {
			wg.Done()
		}()
		log.Println("处理任务", "data", string(param.Payload))
		return nil
	}
	total := 10
	wg.Add(total)

	queueEntity := NewAsyncQueue(
		ConfigName("TestQueueA"),
		ConfigConcurrency(3),
		ConfigRedisConn(client),
		ConfigHandler(HandlerFunc(myHandler)),
	)

	tasks := make([]*TaskInfo, 0, total)
	for i := 0; i < total; i++ {
		data := fmt.Sprintf("{\"id\":%d}", i)
		tasks = append(tasks, &TaskInfo{
			Payload: []byte(data),
		})
	}
	queueEntity.PublishList(tasks)
	time.Sleep(1 * time.Second)

	// 初始化调用队列数据
	err := queueEntity.StartConsuming()
	if err != nil {
		t.Errorf("StartConsuming err: %v", err)
		return
	}

	wg.Wait()
	log.Println("finish")
	time.Sleep(2 * time.Second)
}

func TestTaskQueueEntity_Run(t *testing.T) {
	client := InitClient()

	wg := &sync.WaitGroup{}
	lock := &sync.RWMutex{}
	got := make(map[string]bool)
	myHandler := func(param *TaskInfo) error {
		defer func() {
			wg.Done()
		}()
		log.Println("处理任务", "data", string(param.Payload))
		lock.Lock()
		got[string(param.Payload)] = true
		lock.Unlock()
		time.Sleep(10 * time.Millisecond)
		return nil
	}
	total := 10
	wg.Add(total)

	queueEntity := NewAsyncQueue(
		ConfigName("TestQueueA"),
		ConfigConcurrency(10),
		ConfigRedisConn(client),
		ConfigHandler(HandlerFunc(myHandler)),
	)

	// 初始化调用队列数据
	err := queueEntity.StartConsuming()
	if err != nil {
		t.Errorf("StartConsuming err: %v", err)
		return
	}

	want := make(map[string]bool)
	go func() {
		// 推送消息
		for i := 0; i < total; i++ {
			data := fmt.Sprintf("{\"id\":%d}", i)
			_ = queueEntity.Publish(&TaskInfo{
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
