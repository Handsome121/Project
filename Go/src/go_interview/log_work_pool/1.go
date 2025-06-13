package log_work_pool

import (
	"bufio"
	"context"
	"io"
	"os"
	"sync"
	"time"
)

// LogProcessor 日志处理器结构体
type LogProcessor struct {
	workers    int                // worker数量
	bufferSize int                // 缓冲区大小
	logChan    chan string        // 日志通道
	wg         sync.WaitGroup     // 等待组
	cancel     context.CancelFunc // 取消函数
}

// NewLogProcessor 创建新的日志处理器
func NewLogProcessor(workers, bufferSize int) *LogProcessor {
	return &LogProcessor{
		workers:    workers,
		bufferSize: bufferSize,
		logChan:    make(chan string, bufferSize),
	}
}

// Start 启动日志处理器
func (lp *LogProcessor) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	lp.cancel = cancel

	// 启动worker池
	for i := 0; i < lp.workers; i++ {
		lp.wg.Add(1)
		go lp.worker(ctx, i)
	}
}

// Stop 停止日志处理器
func (lp *LogProcessor) Stop() {
	close(lp.logChan) // 关闭通道
	lp.cancel()       // 取消上下文
	lp.wg.Wait()      // 等待所有worker完成
}

// worker 处理日志的工作协程
func (lp *LogProcessor) worker(ctx context.Context, id int) {
	defer lp.wg.Done()

	for {
		select {
		case logEntry, ok := <-lp.logChan:
			if !ok {
				return // 通道已关闭
			}
			// 处理日志条目
			processLogEntry(logEntry)
		case <-ctx.Done():
			return // 上下文取消
		}
	}
}

// processLogEntry 处理单个日志条目
func processLogEntry(entry string) {
	// 这里实现实际的日志处理逻辑
	// 例如: 解析、过滤、分析、存储等
	// 为了性能考虑，应尽量减少这里的耗时操作
}

// ProcessFile 处理日志文件
func (lp *LogProcessor) ProcessFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// 将日志行发送到处理通道
		select {
		case lp.logChan <- line:
		case <-time.After(100 * time.Millisecond):
			// 处理通道满时的超时逻辑
			// 可以记录metrics或采取其他措施
		}
	}

	return nil
}

// ProcessStream 处理实时日志流
func (lp *LogProcessor) ProcessStream(stream io.Reader) error {
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		select {
		case lp.logChan <- scanner.Text():
		case <-time.After(100 * time.Millisecond):
			// 处理通道满时的超时逻辑
		}
	}
	return scanner.Err()
}
