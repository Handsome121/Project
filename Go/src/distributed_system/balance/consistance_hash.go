package balance

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"sync"

	"github.com/spaolacci/murmur3"
)

const (
	TopWeight   = 100 //最大的权重
	minReplicas = 10  //最小的副本数量
	prime       = 16777619
)

type UInt64Slice []uint64

// Hash 默认的hash函数
func Hash(data []byte) uint64 {
	return murmur3.Sum64(data)
}
func (s UInt64Slice) Len() int {
	return len(s)
}

func (s UInt64Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s UInt64Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type (
	// Func 定义hash函数的形式
	Func func(data []byte) uint64

	ConsistentHash struct {
		hashFunc Func              //使用的hash函数
		replicas int               //每一个新增节点的副本数量
		keys     []uint64          //存放hash后的节点
		hashMap  map[uint64]string //存放节点的hash值与节点的对应信息
		mux      sync.RWMutex
	}
)

// NewConsistentHashBalance 新建一个一致性hash环
func NewConsistentHashBalance(replicas int, fn Func) *ConsistentHash {
	if replicas < minReplicas {
		replicas = minReplicas //限制最小的节点副本数量
	}
	if fn == nil {
		fn = Hash
	}
	return &ConsistentHash{
		replicas: replicas,
		hashFunc: fn,
		hashMap:  make(map[uint64]string),
	}
}

// Add 向环中添加节点
func (c *ConsistentHash) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least")
	}
	addr := params[0]
	//加锁进行操作
	c.mux.Lock()
	defer c.mux.Unlock()
	for i := 0; i < c.replicas; i++ {
		hash := c.hashFunc([]byte(addr + strconv.Itoa(i)))
		c.keys = append(c.keys, hash)
		c.hashMap[hash] = addr
	}

	sort.Slice(c.keys, func(i, j int) bool {
		return c.keys[i] < c.keys[j]
	})
	return nil
}
func (c *ConsistentHash) IsEmpty() bool {
	return len(c.keys) == 0
}

// Get 获取请求改发往的服务器
func (c *ConsistentHash) Get(key string) (string, error) {
	if c.IsEmpty() {
		return "", errors.New("no node backend")
	}
	c.mux.RLock()
	defer c.mux.RUnlock()
	hash := c.hashFunc([]byte(key))
	//用二分查找找到第一个hash值大于key的节点
	index := sort.Search(len(c.keys), func(i int) bool { return c.keys[i] >= hash }) % (len(c.keys))
	return c.hashMap[c.keys[index]], nil
}

// Remove 从consistent中删除节点
func (c *ConsistentHash) Remove(key string) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	if c.IsEmpty() {
		return errors.New("no node bakend")
	}
	for i := 0; i < c.replicas; i++ {
		hash := c.hashFunc([]byte(key + strconv.Itoa(i)))
		index := sort.Search(len(c.keys), func(i int) bool {
			return c.keys[i] >= hash
		})
		if index < len(c.keys) && c.keys[index] == hash {
			c.keys = append(c.keys[:index], c.keys[index+1:]...)
			delete(c.hashMap, hash)
		} else {
			return errors.New("node not exsts")
		}
	}
	return nil
}
func (c *ConsistentHash) Out() {
	for i := 0; i < len(c.keys); i++ {
		fmt.Println(c.keys[i])
	}
	for k, v := range c.hashMap {
		fmt.Println("hash", k, "value", v)
	}
}
