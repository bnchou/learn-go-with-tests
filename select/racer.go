package racer

import (
	"errors"
	"net/http"
	"time"
)

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", errors.New("timeout")
	}
}

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, 10*time.Second)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
