package racer

import (
	"net/http"
	"time"
)

func Racer(u1, u2 string) string {
	duration1 := measureGet(u1)
	duration2 := measureGet(u2)

	if duration1 > duration2 {
		return u2
	}

	return u1
}

func measureGet(url string) time.Duration {
	start1 := time.Now()
	http.Get(url)
	return time.Since(start1)
}
