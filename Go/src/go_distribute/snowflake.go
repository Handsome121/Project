/*
雪花算法原理就是生成一个的64位比特位的 long 类型的唯一 id。
最高1位固定值0，因为生成的 id 是正整数，如果是1就是负数了。
接下来41位存储毫秒级时间戳，2^41/(1000606024365)=69，大概可以使用69年。
再接下10位存储机器码，包括5位 datacenterId 和5位 workerId。最多可以部署2^10=1024台机器。
最后12位存储序列号。同一毫秒时间戳时，通过这个递增的序列号来区分。即对于同一台机器而言，同一毫秒时间戳下，可以生成2^12=4096个不重复 id。
*/

package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	// 起始时间戳（2020-01-01 00:00:00 UTC）
	startTime = 1577836800000 // 2020-01-01 in milliseconds

	// ID 构成部分位数
	timestampBits = 41 // 时间戳占41位（约69年）
	machineIDBits = 10 // 机器ID占10位（支持1024台机器）
	sequenceBits  = 12 // 序列号占12位（每毫秒4096个ID）

	// 最大值常量
	maxTimestamp = (1 << timestampBits) - 1
	maxMachineID = (1 << machineIDBits) - 1
	maxSequence  = (1 << sequenceBits) - 1
)

type Snowflake struct {
	mu            sync.Mutex // 互斥锁（可选，需要高并发时使用）
	machineID     int64      // 机器ID
	sequence      int64      // 当前序列号
	lastTimestamp int64      // 最后生成ID的时间戳
}

// NewSnowflake 创建新的雪花实例
func NewSnowflake(machineID int64) (*Snowflake, error) {
	if machineID < 0 || machineID > maxMachineID {
		return nil, fmt.Errorf("machineID invalid: %d, must be between 0 and %d", machineID, maxMachineID)
	}

	return &Snowflake{
		machineID:     machineID,
		sequence:      0,
		lastTimestamp: -1,
	}, nil
}

// NextID 生成下一个唯一ID
func (sf *Snowflake) NextID() (int64, error) {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	// 获取当前时间戳（毫秒）
	now := time.Now().UnixMilli()

	// 处理时间回拨
	if now < sf.lastTimestamp {
		return 0, errors.New("clock moved backwards")
	}

	// 如果是同一毫秒
	if now == sf.lastTimestamp {
		sf.sequence++
		if sf.sequence > maxSequence {
			return 0, errors.New("too many requests in the same millisecond")
		}
	} else {
		// 新毫秒，重置序列号
		sf.sequence = 0
		sf.lastTimestamp = now
	}

	// 检查时间戳是否溢出
	if now > maxTimestamp {
		return 0, errors.New("timestamp overflow")
	}

	// 组合ID
	id := (now << (machineIDBits + sequenceBits)) |
		(sf.machineID << sequenceBits) |
		sf.sequence

	return id, nil
}

func main() {
	// 创建两个不同节点的雪花实例
	sf1, _ := NewSnowflake(1)
	sf2, _ := NewSnowflake(2)

	var wg sync.WaitGroup
	wg.Add(2)

	// 测试连续生成
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			id, err := sf1.NextID()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Printf("Node1: %d\n", id)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			id, err := sf2.NextID()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Printf("Node2: %d\n", id)
		}
	}()

	wg.Wait()
}
