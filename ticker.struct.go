package xzticker

import "time"

type Ticker interface {
	Run()
	Stop()
}

type ticker struct {
	interval time.Duration // 定时间隔时间
	handler  func()        // 定时执行函数
	stop     chan bool     // 发送停止信号
}

func (t *ticker) Run() {
	ticker := time.NewTicker(t.interval)
	for {
		select {
		case <-ticker.C:
			// 定时器执行函数
			if t.handler != nil {
				t.handler()
			}
		case <-t.stop:
			ticker.Stop()
			return
		}
	}
}

func (t *ticker) Stop() {
	t.stop <- true
}
