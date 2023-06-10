package xzticker

import "time"

type Option func(t *ticker)

func WithInterval(interval time.Duration) Option {
	return func(t *ticker) {
		t.interval = interval
	}
}

func WithHandler(handler func()) Option {
	return func(t *ticker) {
		t.handler = handler
	}
}
