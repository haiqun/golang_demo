package demo3

import "time"

type servicePanel struct {

}

// 此处截取自研的熔断器代码中的限流实现，这是非阻塞的实现
func (sp *servicePanel) incLimit() error {
	// 如果大于限制的条件则返回错误
	if sp.currentLimitCount.Load() > sp.currLimitFunc(nil) {
		return ErrCurrentLimit
	}
	sp.currentLimitCount.Inc()
	return nil
}

func (sp *servicePanel) clearLimit() {
	// 定期每秒重置计数器，从而达到每秒限制的并发数
	// 比如限制1000req/s，在这里指每秒清理1000的计数值
	// 令牌桶是定期放，这里是逆思维，每秒清空，实现不仅占用内存低而且效率高
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			sp.currentLimitCount.Store(0)
		}
	}
}
