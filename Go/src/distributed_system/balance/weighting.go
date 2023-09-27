package balance

import (
	"errors"
	"math/rand"
	"sort"
)

type WeightRandomBanlance struct {
	addrs   []WeightRandomNode
	weights []int //记录排序后的权重前缀和
	max     int   //记录权重的总和
}
type WeightRandomNode struct {
	addr   string
	weight int
}

func NewWeightRandomBanlance(w ...WeightRandomNode) WeightRandomBanlance {
	//将节点按权重从小到大排序
	sort.Slice(w, func(i, j int) bool {
		return w[i].weight < w[j].weight
	})

	weights := make([]int, 0)
	total := 0
	for i := 0; i < len(w); i++ {
		total += w[i].weight
		weights = append(weights, total)
	}

	return WeightRandomBanlance{
		addrs:   w,
		weights: weights,
		max:     total,
	}
}
func NewWeightRandomNode(addr string, weight int) WeightRandomNode {
	return WeightRandomNode{
		addr:   addr,
		weight: weight,
	}
}

// Add 添加节点
func (w *WeightRandomBanlance) Add(wn ...WeightRandomNode) {
	nodes := append(w.addrs, wn...)
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].weight < nodes[j].weight
	})
	weights := make([]int, 0)
	total := 0
	for i := 0; i < len(nodes); i++ {
		total += nodes[i].weight
		weights = append(weights, total)
	}
	w.addrs = nodes
	w.max = total
	w.weights = weights
}
func (w *WeightRandomBanlance) Next() string {
	if len(w.weights) == 0 {
		return ""
	}
	r := rand.Intn(w.max) + 1 //从1开始计数
	index := sort.SearchInts(w.weights, r)
	return w.addrs[index].addr
}

func (w *WeightRandomBanlance) Get() (string, error) {
	mac := w.Next()
	if mac == "" {
		return mac, errors.New("no node backend")
	}
	return mac, nil
}
