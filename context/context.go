package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := store.Fetch(r.Context())
		fmt.Fprint(w, data)
	}
}

// Implementation before propagating context
type StoreLegacy interface {
	Fetch() string
	Cancel()
}

func ServerLegacy(store StoreLegacy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// context has a method Done() which returns a channel
		// which gets sent a signal when the context is "done" or "cancelled".
		// We want to listen to that signal and call store.Cancel if we get it
		// but we want to ignore it if our Store manages to Fetch before it.
		//
		// To manage this we run Fetch in a goroutine and it will write the result into a new channel data.
		// We then use select to effectively race to the two asynchronous processes
		// and then we either write a response or Cancel.

		ctx := r.Context()

		data := make(chan string)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
