package async_queue

import (
	"fmt"
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
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

	queueEntity := NewAsyncQueue("TestQueueA", client)
	total := 5

	for i := 0; i < total; i++ {
		data := fmt.Sprintf("{\"id\":%d}", i)
		queueEntity.Publish([]byte(data))
	}
}

func TestAsyncQueue_Server(t *testing.T) {
	client := InitClient()

	myHandler := func(param *TaskInfo) error {
		log.Info("处理任务", zap.ByteString("data", param.Payload))
		return nil
	}

	queueEntity := NewAsyncQueue("TestQueueA", client,
		ConfigConcurrency(3),
		ConfigHandler(HandlerFunc(myHandler)),
	)

	// 初始化调用队列数据
	err := queueEntity.StartConsuming()
	if err != nil {
		t.Errorf("StartConsuming err: %v", err)
		return
	}

	time.Sleep(10 * time.Second)
}

func TestAsyncQueue_Schedule(t *testing.T) {
	client := InitClient()

	wg := &sync.WaitGroup{}
	myHandler := func(param *TaskInfo) error {
		defer func() {
			wg.Done()
		}()
		log.Info("处理任务", zap.ByteString("data", param.Payload))
		return nil
	}
	total := 10
	wg.Add(total)

	queueEntity := NewAsyncQueue("TestQueueA", client,
		ConfigConcurrency(3),
		ConfigHandler(HandlerFunc(myHandler)),
	)

	processAt := time.Now().Unix()
	for i := 0; i < total; i++ {
		data := fmt.Sprintf("{\"id\":%d}", i)
		queueEntity.Publish([]byte(data), ConfigTaskProcessAt(processAt))
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
	log.Info("finish")
	time.Sleep(2 * time.Second)
}

func TestAsyncQueue_List(t *testing.T) {
	client := InitClient()

	wg := &sync.WaitGroup{}
	myHandler := func(param *TaskInfo) error {
		defer func() {
			wg.Done()
		}()
		log.Info("处理任务", zap.ByteString("data", param.Payload))
		return nil
	}
	total := 10
	wg.Add(total)

	queueEntity := NewAsyncQueue("TestQueueA", client,
		ConfigConcurrency(3),
		ConfigHandler(HandlerFunc(myHandler)),
	)

	tasks := make([][]byte, 0, total)
	for i := 0; i < total; i++ {
		data := fmt.Sprintf("{\"id\":%d}", i)
		tasks = append(tasks, []byte(data))
	}
	queueEntity.PublishListDAG(tasks)
	time.Sleep(1 * time.Second)

	// 初始化调用队列数据
	err := queueEntity.StartConsuming()
	if err != nil {
		t.Errorf("StartConsuming err: %v", err)
		return
	}

	wg.Wait()
	log.Info("finish")
	time.Sleep(2 * time.Second)
}

func TestAsyncQueue_Retry(t *testing.T) {
	client := InitClient()

	cnt := make(map[string]int)
	wg := &sync.WaitGroup{}
	lock := &sync.Mutex{}
	myHandler := func(param *TaskInfo) error {
		lock.Lock()
		defer func() {
			lock.Unlock()
		}()
		if cnt[string(param.Payload)] == 0 {
			cnt[string(param.Payload)] += 1
			log.Info("重试任务", zap.ByteString("data", param.Payload))
			param.RetryDelay(2)
			return nil
		}

		log.Info("处理任务", zap.ByteString("data", param.Payload))
		wg.Done()
		return nil
	}
	total := 10
	wg.Add(total)

	queueEntity := NewAsyncQueue("TestQueueA", client,
		ConfigConcurrency(3),
		ConfigHandler(HandlerFunc(myHandler)),
		ConfigRetryMax(1),
	)

	processAt := time.Now().Unix()
	for i := 0; i < total; i++ {
		data := fmt.Sprintf("{\"id\":%d}", i)
		queueEntity.Publish([]byte(data), ConfigTaskProcessAt(processAt))
		processAt += 1
	}
	time.Sleep(1 * time.Second)

	// 初始化调用队列数据
	err := queueEntity.StartConsuming()
	if err != nil {
		t.Errorf("StartConsuming err: %v", err)
		return
	}

	wg.Wait()
	log.Info("finish")
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
		log.Info("处理任务", zap.ByteString("data", param.Payload))
		lock.Lock()
		got[string(param.Payload)] = true
		lock.Unlock()
		time.Sleep(10 * time.Millisecond)
		return nil
	}
	total := 10
	wg.Add(total)

	queueEntity := NewAsyncQueue("TestQueueA", client,
		ConfigConcurrency(3),
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
			queueEntity.Publish([]byte(data))
			want[data] = true
			time.Sleep(10 * time.Millisecond)
		}
		log.Info("推送消息完成")
	}()

	wg.Wait()
	time.Sleep(2 * time.Second)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
		return
	}
}
