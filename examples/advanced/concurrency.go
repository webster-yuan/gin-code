package advanced

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

func helloWorld() {
	fmt.Println("hello world from helloWorld function")
}

func ConcurrencyMain() {
	fmt.Println("concurrency main")
	go helloWorld()
	go func() {
		fmt.Println("hello world from anonymous function")
	}()
	// 让主协程等一会
	time.Sleep(1 * time.Second)
	// 主协程继续执行
	fmt.Println("concurrency main continue")
}

// ChannelMain channel用于协程之间的消息同步
func ChannelMain() {
	// make 是创建管道的唯一方式
	// chan 是 channel类型，int 是管道中元素的类型，10 是管道的缓冲区大小
	ch := make(chan int, 10)
	defer close(ch)

	// 向管道中写入数据
	ch <- 1
	ch <- 2
	ch <- 3

	// 从管道中读取数据
	// 管道中的数据流动方式与队列一样，即先进先出（FIFO）
	// 在某一个时刻，只有一个协程能够对其写入数据，同时也只有一个协程能够读取管道中的数据.
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 无缓冲管道
	// 正常是一个协程写入数据，另一个协程读取数据
	// 但是无缓冲管道会阻塞写入操作，直到有另一个协程读取数据
	// 也会阻塞读取操作，直到有另一个协程写入数据
	unbufferedChan := make(chan int)
	defer close(unbufferedChan)
	go func() {
		unbufferedChan <- 42
	}()
	n := <-unbufferedChan
	fmt.Println(n)
	// 以上操作均很危险，因为如果缓冲区满了或者空了，就会永远阻塞下去

	// 调用示例如下
	// ControlByUnbufferedChanDemo()
}

func ControlByUnbufferedChanDemo() {
	ch := make(chan int, 10)
	chW := make(chan struct{})
	chR := make(chan struct{})
	defer func() {
		close(ch)
		close(chW)
		close(chR)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("写入", i)
		}
		chW <- struct{}{}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			// 这里delay 1ms 是为了确保写入操作完成，并不可靠，只是为了演示
			time.Sleep(1 * time.Millisecond)
			fmt.Println("读取", <-ch)
		}
		chR <- struct{}{}
	}()
	// 主线程从 chW 通道中读取数据，等待写入操作完成
	fmt.Println("写入完毕", <-chW)
	// 主线程从 chR 通道中读取数据，等待读取操作完成
	fmt.Println("读取完毕", <-chR)
}

// Logger 日志结构体
type Logger struct {
	Msg string
}

// NewLogger 创建一个“只写”通道，并启动扇出 goroutine
func NewLogger() chan<- Logger { // 编译器强制要求只能写入通道
	ch := make(chan Logger, 10)
	go fanout(ch)
	return ch
}

// fanout 扇出，内部持有“只读”视图，启动 3 个 worker 消费，数量只在 fanout 函数中确定
func fanout(ch <-chan Logger) {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ { // 启动 3 个 worker 消费日志
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for log := range ch { // for range 会不断读取管道中的元素
				// 模拟不同下游：这里简单打印 worker id 和日志内容
				fmt.Printf("[worker-%d] %s\n", id, log.Msg)
			}
		}(i)
	}
	// 等待所有 worker 完成
	wg.Wait()
}

func OneWayChannelMain() {
	// 拿到只写通道
	logger := NewLogger()

	// 模拟产生 10 条日志
	for i := 0; i < 10; i++ {
		logger <- Logger{Msg: fmt.Sprintf("hello world %d", i)}
	}

	// 等待 worker 把队列里剩余日志处理完
	// 简单做法：睡一小会儿；生产代码可用 sync.WaitGroup 或 close 通道
	// 主 goroutine 如果提前退出，后台 fanout 会被强制杀死，所以必须有个同步点
	// 这里简单用 sleep 代替
	time.Sleep(100 * time.Millisecond)
}

func ForRangeMain() {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 手动关闭管道，避免主协程因为读取不到数据，一直阻塞等待
		// 尽量在管道发送方去关闭管道
		close(ch)
	}()
	for n := range ch { // 遍历管道只会有一个返回值，读取管道中的元素，当空或者无缓冲区时会阻塞等待
		fmt.Println(n)
	}
	for i := 0; i < 10; i++ {
		n, ok := <-ch // 读取管道是有两个返回值的，ok代表能否读取数据成功，而不是管道是否关闭
		fmt.Println(n, ok)
	}
}

// WaitGroupMain 等所有后台任务都彻底结束，再往下走 就是WaitGroup的适用场景
// “并发干很多活” → “必须全部干完才能继续下一步”
// WaitGroup 就是官方给你提供的最轻量、无锁、零依赖的“大闸”
// WaitGroup：控制同级别子协程都正确执行
func WaitGroupMain() {
	var wait sync.WaitGroup
	// 程序初始化时，指定子协程的个数
	wait.Add(1)
	go func() {
		fmt.Println(1)
		// 每当一个协程执行完毕时调用 Done，计数就-1
		wait.Done()
	}()
	// 主协程调用 Wait 会一直阻塞直到全部计数减为 0，然后才会被唤醒
	wait.Wait()
	fmt.Println(2)

	var mainWaitGroup sync.WaitGroup
	// 计数10
	mainWaitGroup.Add(10)
	fmt.Println("mainWaitGroup start")
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i) // 子协程任务
			mainWaitGroup.Done()
		}(i)
	}
	mainWaitGroup.Wait()
	fmt.Println("mainWaitGroup end")
}

// WithCancelContextMain Context 控制子孙协程以及层级更深的协程
func WithCancelContextMain() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt) // 监听 Ctrl-C 信号
		<-c
		cancel()
	}()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker(ctx, i)
		}(i)
	}
	wg.Wait()
	fmt.Println("all workers returned")
}

func worker(ctx context.Context, id int) {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d done\n", id)
			return
		default:
			time.Sleep(300 * time.Millisecond)
			fmt.Printf("worker %d working step %d \n", id, i)
		}
	}
}

type ctxKey int

const requestIDKey ctxKey = 0

func WithValueContextMain() {
	http.HandleFunc("/user", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	id := request.Header.Get("X-Request-Id")
	if id == "" {
		id = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	ctx := context.WithValue(request.Context(), requestIDKey, id)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			requestId := ctx.Value(requestIDKey).(string)
			fmt.Printf("request id: %s\n", requestId)
		}(i)
	}
	wg.Wait()
	_, err := writer.Write([]byte("ok"))
	if err != nil {
		return
	}
}

// SelectMain 阻塞等待
func SelectMain() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)
	l := make(chan struct{})

	go Send(chA)
	go Send(chB)
	go Send(chC)
	// 另外启动一个协程：通过 for 循环，配合 select 来一直监测三个管道是否可以用
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case n, ok := <-chA:
				if ok {
					fmt.Println("A", n, ok)
				} else {
					chA = nil
				}
			case n, ok := <-chB:
				if ok {
					fmt.Println("B", n, ok)
				} else {
					chB = nil
				}
			case n, ok := <-chC:
				if ok {
					fmt.Println("C", n, ok)
				} else {
					chC = nil
				}
			case <-ctx.Done(): // 只判一次超时
				fmt.Println("timeout")
				l <- struct{}{}
				return
			}
			if chA == nil && chB == nil && chC == nil {
				fmt.Println("finished normally")
				l <- struct{}{}
				return
			}
		}
	}()
	// 读取管道，读取数据读取不到会一直阻塞等待
	<-l
}

func Send(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond)
		ch <- i
	}
}

func MutexMain() {
	var wg sync.WaitGroup
	var count = 0
	var lock sync.Mutex

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(data *int) {
			// 加锁
			lock.Lock()
			// 模拟访问耗时
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5000)))
			// 访问数据
			temp := *data
			// 模拟计算耗时
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5000)))
			ans := 1
			// 修改数据
			*data = temp + ans
			// 解锁
			lock.Unlock()
			fmt.Println(*data)
			wg.Done()
		}(&count)
	}
	wg.Wait()
	fmt.Println("最终结果：", count)
}

var wg sync.WaitGroup
var count = 0
var rw sync.RWMutex

// 条件变量
// 在创建条件变量时，因为在这里条件变量作用的是读协程，所以将读锁作为互斥锁传入
var cond = sync.NewCond(rw.RLocker())

func RWMutexMain() {
	wg.Add(12)
	// 读多写少
	go func() {
		for i := 0; i < 3; i++ {
			go Write(&count)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 7; i++ {
			go Read(&count)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("最终结果：", count)
}

func Read(i *int) {
	rw.RLock()
	fmt.Println("拿到读锁了")
	for *i < 3 { // 条件不满足就一直阻塞
		cond.Wait()
	}
	fmt.Println("释放读锁", *i)
	rw.RUnlock()
	wg.Done()
}

func Write(i *int) {
	rw.Lock()
	fmt.Println("拿到写锁了")
	temp := *i
	*i = temp + 1
	fmt.Println("释放写锁")
	rw.Unlock()
	wg.Done()
}

var (
	once  sync.Once
	index map[string]int
)

func Index() map[string]int {
	once.Do(func() {
		index = make(map[string]int)
		for i, w := range []string{"foo", "bar", "baz"} {
			index[w] = i
		}
	})
	return index
}

type Pool struct {
	ch   chan int
	once sync.Once
}

func (p *Pool) Close() {
	p.once.Do(func() {
		close(p.ch) // 只能关一次，否则 panic
	})
}

// AtomicMain 原子操作
func AtomicMain() {
	var aint64 atomic.Int64
	aint64.Store(123)
	aint64.Swap(432)
	aint64.Add(667)
	fmt.Println(aint64.Load())
}
