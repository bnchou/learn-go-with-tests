package racer

import (
	"errors"
	"net/http"
	"time"
)

func Racer(a, b string) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-timeout():
		return "", errors.New("timeout")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func timeout() chan struct{} {
	ch := make(chan struct{})
	go func() {
		time.Sleep(10 * time.Second)
		close(ch)
	}()
	return ch
}
