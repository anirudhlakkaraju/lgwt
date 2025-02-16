package sel

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := getResponseTime(a)
	bDuration := getResponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func getResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
