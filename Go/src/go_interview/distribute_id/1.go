package distribute_id

import (
	"sync"
	"time"
)

const (
	epoch         int64 = 1609459200000 // 2021-01-01 00:00:00 UTC
	machineIDBits uint8 = 10
	sequenceBits  uint8 = 12
	maxMachineID  int64 = -1 ^ (-1 << machineIDBits)
	maxSequence   int64 = -1 ^ (-1 << sequenceBits)
)

type IDGenerator struct {
	machineID int64
	lastTime  int64
	sequence  int64
	pool      sync.Pool
	stopChan  chan struct{}
	mu        sync.Mutex
}

type IDBatch struct {
	ids   []int64
	index int
}

func NewIDGenerator(machineID int64) *IDGenerator {
	if machineID < 0 || machineID > maxMachineID {
		panic("invalid machine ID")
	}

	g := &IDGenerator{
		machineID: machineID,
		pool: sync.Pool{
			New: func() interface{} {
				return &IDBatch{ids: make([]int64, 1000), index: 0}
			},
		},
		stopChan: make(chan struct{}),
	}

	// 启动后台填充goroutine
	go g.fillPool()

	return g
}

func (g *IDGenerator) NextID() int64 {
	batch := g.pool.Get().(*IDBatch)
	defer func() {
		if batch.index >= len(batch.ids) {
			batch.ids = batch.ids[:0]
			batch.index = 0
			g.pool.Put(batch)
		}
	}()

	if batch.index < len(batch.ids) {
		id := batch.ids[batch.index]
		batch.index++
		return id
	}

	// 后备方案：直接生成
	return g.generateID()
}

func (g *IDGenerator) fillPool() {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			batch := g.pool.Get().(*IDBatch) // 先尝试获取旧对象
			if batch == nil {
				batch = &IDBatch{ids: make([]int64, 1000)} // 实在没有才新建
			}
			for i := 0; i < len(batch.ids); i++ {
				batch.ids[i] = g.generateID()
			}
			g.pool.Put(batch)
		case <-g.stopChan:
			return
		}
	}
}

func (g *IDGenerator) generateID() int64 {
	g.mu.Lock()
	defer g.mu.Unlock()

	now := time.Now().UnixNano() / 1e6
	if now < g.lastTime {
		panic("clock moved backwards")
	}

	if now == g.lastTime {
		g.sequence = (g.sequence + 1) & maxSequence
		if g.sequence == 0 {
			for now <= g.lastTime {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		g.sequence = 0
	}

	g.lastTime = now

	return ((now - epoch) << (machineIDBits + sequenceBits)) |
		(g.machineID << sequenceBits) |
		g.sequence
}

func (g *IDGenerator) Close() {
	close(g.stopChan)
}
