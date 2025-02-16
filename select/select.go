package sel

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	timeA := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	timeB := time.Since(startB)

	if timeA < timeB {
		return a
	}
	return b
}
