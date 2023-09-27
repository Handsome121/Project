package balance

import "errors"

// PollingBalance 轮询负载均衡
type PollingBalance struct {
	curIndex int
	macList  []string
}

func (p *PollingBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least")
	}
	p.macList = append(p.macList, params...)
	return nil
}

func (p *PollingBalance) Next() string {
	if len(p.macList) == 0 {
		return ""
	}
	lens := len(p.macList)
	if p.curIndex >= lens {
		p.curIndex = 0
	}
	mac := p.macList[p.curIndex]
	p.curIndex = (p.curIndex + 1) % lens

	return mac
}

func (p *PollingBalance) Get() (string, error) {
	mac := p.Next()
	if mac == "" {
		return mac, errors.New("no machine backend")
	}

	return mac, nil
}
