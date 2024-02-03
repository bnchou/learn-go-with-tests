package racer

import (
	"net/http"
	"time"
)

func Racer(u1, u2 string) string {
	start1 := time.Now()
	http.Get(u1)
	duration1 := time.Since(start1)

	start2 := time.Now()
	http.Get(u2)
	duration2 := time.Since(start2)

	if duration1 > duration2 {
		return u2
	}

	return u1
}
