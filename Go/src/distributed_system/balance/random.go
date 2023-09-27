package balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
	curIndex int
	macList  []string
}

var _ RandomBalance

// Add 添加后台机器
func (r *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least")
	}
	//可以在赋值之前加一个输入参数正确性判断
	r.macList = append(r.macList, params...)
	return nil
}

// Next 随机化负载均衡
func (r *RandomBalance) Next() string {
	if len(r.macList) == 0 {
		return ""
	}
	r.curIndex = rand.Intn(len(r.macList))
	return r.macList[r.curIndex]
}

// Get 提供随机化接口
func (r *RandomBalance) Get() (res string, error error) {
	mac := r.Next()
	if mac == "" {
		return "", errors.New("no machine backend")
	}
	return mac, nil
}
