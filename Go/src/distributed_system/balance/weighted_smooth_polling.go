package balance

import (
	"errors"
	"strconv"
)

type WeightRandomBalance struct {
	curIndex int
	macList  []*weightNode
}
type weightNode struct {
	addr            string
	weight          int //初始化的权重
	curWeight       int //更新的临时权重
	effectiveWeight int //节点有效权重
}

func (r *WeightRandomBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("params len need 2")
	}
	weight, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}
	node := &weightNode{
		addr:   params[0],
		weight: int(weight),
	}
	node.effectiveWeight = node.weight
	r.macList = append(r.macList, node)
	return nil
}

func (r *WeightRandomBalance) Next() string {
	if len(r.macList) == 0 {
		return ""
	}
	var best *weightNode
	total := 0
	for i := 0; i < len(r.macList); i++ {
		w := r.macList[i]
		//1,获取权重总和
		total += w.effectiveWeight
		//2,更新curWeight
		w.curWeight += w.effectiveWeight
		//3,
		if w.effectiveWeight < w.weight {
			w.effectiveWeight++
		}
		//4,选中临时权重最大的节点
		if best == nil || best.curWeight < w.curWeight {
			best = w
			r.curIndex = i
		}
	}
	if best == nil {
		return ""
	}
	//5,修改临时权重为临时权重-total
	best.curWeight = best.curWeight - total
	return best.addr
}

func (r *WeightRandomBalance) Get() (string, error) {
	mac := r.Next()
	if mac == "" {
		return mac, errors.New("no machine backend")
	}
	return mac, nil
}
