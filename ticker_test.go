package xzticker

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTicker(t *testing.T) {
	var _ticker Ticker
	count := 5
	_ticker = New(
		WithInterval(1*time.Second),
		WithHandler(func() {
			if count > 0 {
				count -= 1
			} else {
				_ticker.Stop()
			}
			fmt.Println(count)
		}),
	)
	go _ticker.Run()
	time.Sleep(time.Duration(count) * time.Second)
}
