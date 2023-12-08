package async_queue

import (
	"fmt"
	"github.com/MenciusCheng/algorithm-code/utils/log"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"reflect"
	"strings"
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

// 测试发布任务
func TestAsyncQueue_Client(t *testing.T) {
	client := InitClient()

	queueEntity := NewAsyncQueue("TestQueueA", client)

	// 挂起任务
	for i := 0; i < 3; i++ {
		data := fmt.Sprintf("{\"id\":\"n%d\",\"type\":\"pending\"}", i)
		res, err := queueEntity.Publish([]byte(data))
		log.Info("Publish pending task", zap.Any("res", res), zap.Error(err))
	}

	// 延时任务
	for i := 0; i < 2; i++ {
		data := fmt.Sprintf("{\"id\":\"s%d\",\"type\":\"schedule\"}", i)
		res, err := queueEntity.Publish([]byte(data), ConfigTaskDelaySecond(10))
		log.Info("Publish schedule task", zap.Any("res", res), zap.Error(err))
	}

	// DAG任务
	tasks := make([][]byte, 0)
	for i := 0; i < 3; i++ {
		data := fmt.Sprintf("{\"id\":\"dag%d\",\"type\":\"daging\"}", i)
		tasks = append(tasks, []byte(data))
	}
	res, err := queueEntity.PublishListDAG(tasks, ConfigTaskPriority(QueuePriorityHigh))
	log.Info("Publish dag task", zap.Any("res", res), zap.Error(err))
}

// 测试消费任务
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

	time.Sleep(120 * time.Second)
}

// 测试延时任务
func TestAsyncQueue_Schedule(t *testing.T) {
	client := InitClient()

	wg := &sync.WaitGroup{}
	myHandler := func(param *TaskInfo) error {
		defer func() {
			wg.Done()
		}()
		log.Info("处理任务", zap.ByteString("data", param.Payload))
		time.Sleep(1 * time.Second)
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
		processAt += 2
		data := fmt.Sprintf("{\"id\":%d}", i)
		queueEntity.Publish([]byte(data), ConfigTaskProcessAt(processAt))
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

// 测试DAG任务
func TestAsyncQueue_List(t *testing.T) {
	client := InitClient()

	gotDoneK1 := make([]string, 0)
	gotDoneK2 := make([]string, 0)
	wg := &sync.WaitGroup{}
	myHandler := func(param *TaskInfo) error {
		defer func() {
			wg.Done()
		}()
		log.Info("处理任务", zap.ByteString("data", param.Payload))
		if strings.Contains(string(param.Payload), "k1") {
			gotDoneK1 = append(gotDoneK1, string(param.Payload))
		} else {
			gotDoneK2 = append(gotDoneK2, string(param.Payload))
		}
		time.Sleep(500 * time.Millisecond)
		return nil
	}
	total := 5

	queueEntity := NewAsyncQueue("TestQueueA", client,
		ConfigConcurrency(3),
		ConfigHandler(HandlerFunc(myHandler)),
	)

	wantDoneK1 := make([]string, 0)
	tasks := make([][]byte, 0, total)
	for i := 0; i < total; i++ {
		data := fmt.Sprintf("{\"id\":%d,\"type\":\"k1\"}", i)
		tasks = append(tasks, []byte(data))
		wantDoneK1 = append(wantDoneK1, data)
	}
	queueEntity.PublishListDAG(tasks)
	wg.Add(total)

	wantDoneK2 := make([]string, 0)
	tasks = make([][]byte, 0, total)
	for i := 0; i < total; i++ {
		data := fmt.Sprintf("{\"id\":%d,\"type\":\"k2\"}", i)
		tasks = append(tasks, []byte(data))
		wantDoneK2 = append(wantDoneK2, data)
	}
	queueEntity.PublishListDAG(tasks)
	wg.Add(total)

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
	if !reflect.DeepEqual(gotDoneK1, wantDoneK1) {
		t.Errorf("gotDoneK1 = %v, wantDoneK1 = %v", gotDoneK1, wantDoneK1)
		return
	}
	if !reflect.DeepEqual(gotDoneK2, wantDoneK2) {
		t.Errorf("gotDoneK2 = %v, wantDoneK2 = %v", gotDoneK2, wantDoneK2)
		return
	}
}

func TestAsyncQueue_Retry(t *testing.T) {
	client := InitClient()

	gotRetryCnt := make(map[string]int)
	gotDoneCnt := make(map[string]int)
	wg := &sync.WaitGroup{}
	lock := &sync.Mutex{}
	myHandler := func(param *TaskInfo) error {
		lock.Lock()
		defer func() {
			lock.Unlock()
		}()
		if gotRetryCnt[string(param.Payload)] < 2 {
			gotRetryCnt[string(param.Payload)] += 1
			log.Info("重试任务", zap.ByteString("data", param.Payload))
			param.RetryDelay(2)
			return nil
		}

		log.Info("处理任务", zap.ByteString("data", param.Payload))
		gotDoneCnt[string(param.Payload)] += 1
		wg.Done()
		return nil
	}
	total := 10
	wg.Add(total)

	queueEntity := NewAsyncQueue("TestQueueA", client,
		ConfigConcurrency(3),
		ConfigHandler(HandlerFunc(myHandler)),
		ConfigRetryMax(2),
	)

	wantRetryCnt := make(map[string]int)
	wantDoneCnt := make(map[string]int)
	processAt := time.Now().Unix()
	for i := 0; i < total; i++ {
		data := fmt.Sprintf("{\"id\":%d}", i)
		queueEntity.Publish([]byte(data), ConfigTaskProcessAt(processAt))
		processAt += 1
		wantRetryCnt[data] = 2
		wantDoneCnt[data] = 1
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
	if !reflect.DeepEqual(gotRetryCnt, wantRetryCnt) {
		t.Errorf("gotRetryCnt = %v, wantRetryCnt = %v", gotRetryCnt, wantRetryCnt)
		return
	}
	if !reflect.DeepEqual(gotDoneCnt, wantDoneCnt) {
		t.Errorf("gotDoneCnt = %v, wantDoneCnt = %v", gotDoneCnt, wantDoneCnt)
		return
	}
}

// 全类型任务混合测试
func TestTaskQueueEntity_Run(t *testing.T) {
	client := InitClient()

	gotPendingCnt := make(map[string]int)
	gotScheduleCnt := make(map[string]int)
	gotDagingCnt := make(map[string]int)
	gotRetryCnt := make(map[string]int)
	gotRetryDoneCnt := make(map[string]int)
	wg := &sync.WaitGroup{}
	lock := &sync.Mutex{}
	myHandler := func(param *TaskInfo) error {
		lock.Lock()
		defer func() {
			lock.Unlock()
		}()
		data := string(param.Payload)
		if strings.Contains(data, "pending") {
			gotPendingCnt[data] += 1
		} else if strings.Contains(data, "schedule") {
			gotScheduleCnt[data] += 1
		} else if strings.Contains(data, "daging") {
			gotDagingCnt[data] += 1
		} else if strings.Contains(data, "retry") {
			if gotRetryCnt[string(param.Payload)] < 2 {
				gotRetryCnt[string(param.Payload)] += 1
				param.RetryDelay(2)
				log.Info("重试任务", zap.ByteString("data", param.Payload), zap.Any("param", param))
				return nil
			} else {
				gotRetryDoneCnt[string(param.Payload)] += 1
			}
		}
		log.Info("处理任务", zap.ByteString("data", param.Payload), zap.Any("param", param))
		wg.Done()
		return nil
	}
	total := 10

	// 初始化队列
	queueEntity := NewAsyncQueue("TestQueueA", client,
		ConfigConcurrency(3),
		ConfigHandler(HandlerFunc(myHandler)),
		ConfigRetryMax(2),
	)

	wantPendingCnt := make(map[string]int)
	wantScheduleCnt := make(map[string]int)
	wantDagingCnt := make(map[string]int)
	wantRetryCnt := make(map[string]int)
	wantRetryDoneCnt := make(map[string]int)
	go func() {
		// 发布任务
		i := 0
		for i < total {
			// 挂起任务
			data := fmt.Sprintf("{\"id\":\"n%d\",\"type\":\"pending\"}", i)
			queueEntity.Publish([]byte(data), ConfigTaskPriority(QueuePriorityHigh))
			wantPendingCnt[data] = 1
			wg.Add(1)

			// 延时任务
			data = fmt.Sprintf("{\"id\":\"s%d\",\"type\":\"schedule\"}", i)
			queueEntity.Publish([]byte(data), ConfigTaskDelaySecond(5))
			wantScheduleCnt[data] = 1
			wg.Add(1)

			// DAG任务
			tasks := make([][]byte, 0)
			for j := 0; j < 3; j++ {
				data = fmt.Sprintf("{\"id\":\"dag%d-%d\",\"type\":\"daging\"}", i, j)
				tasks = append(tasks, []byte(data))
				wantDagingCnt[data] = 1
				wg.Add(1)
			}
			queueEntity.PublishListDAG(tasks)

			// 重试任务
			data = fmt.Sprintf("{\"id\":\"re%d\",\"type\":\"retry\"}", i)
			queueEntity.Publish([]byte(data), ConfigTaskPriority(QueuePriorityLow))
			wantRetryCnt[data] = 2
			wantRetryDoneCnt[data] = 1
			wg.Add(1)

			i++
			time.Sleep(2 * time.Second)
		}
	}()
	time.Sleep(2 * time.Second)

	// 初始化调用队列数据
	err := queueEntity.StartConsuming()
	if err != nil {
		t.Errorf("StartConsuming err: %v", err)
		return
	}

	wg.Wait()
	log.Info("finish")
	time.Sleep(2 * time.Second)
	if !reflect.DeepEqual(gotPendingCnt, wantPendingCnt) {
		t.Errorf("gotPendingCnt = %v, wantPendingCnt = %v", gotPendingCnt, wantPendingCnt)
		return
	}
	if !reflect.DeepEqual(gotScheduleCnt, wantScheduleCnt) {
		t.Errorf("gotScheduleCnt = %v, wantScheduleCnt = %v", gotScheduleCnt, wantScheduleCnt)
		return
	}
	if !reflect.DeepEqual(gotDagingCnt, wantDagingCnt) {
		t.Errorf("gotDagingCnt = %v, wantDagingCnt = %v", gotDagingCnt, wantDagingCnt)
		return
	}
	if !reflect.DeepEqual(gotRetryCnt, wantRetryCnt) {
		t.Errorf("gotRetryCnt = %v, wantRetryCnt = %v", gotRetryCnt, wantRetryCnt)
		return
	}
	if !reflect.DeepEqual(gotRetryDoneCnt, wantRetryDoneCnt) {
		t.Errorf("gotRetryDoneCnt = %v, wantRetryDoneCnt = %v", gotRetryDoneCnt, wantRetryDoneCnt)
		return
	}
}
