package xzticker

import (
	"runtime"
	"time"
)

func New(opts ...Option) *ticker {
	t := &ticker{
		interval: 60 * time.Second,
		handler:  nil,
		stop:     make(chan bool),
	}
	for _, opt := range opts {
		opt(t)
	}
	runtime.SetFinalizer(t, func(t *ticker) {
		t.Stop()
	})
	return t
}
