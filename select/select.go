package sel

import (
	"errors"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {
	// 	Waiting for values to be sent to a channel with myVar := <-ch is a blocking call
	// `select` allows you to wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", errors.New("request took too long and timed out")
	}
}

// - A `chan struct{}` is the smallest data type available from a memory perspective so we get no memory allocation
// - Always **make** channels!
// - Rather than say `var ch chan struct{}`, when you use `var` the variable will be initialised with the "zero" value of the type. So for `string` it is `""`, `int` it is `0`, etc.
//   - For `channels` the zero value is `nil` and if you try and send to it with `<-` it will block forever because you cannot send to `nil` channels
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
