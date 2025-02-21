package context

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			log.Printf("error occured: %q", err)
			return
		}

		fmt.Fprint(w, data)
	}
}

// Implementation before propagating context
type StoreLegacy interface {
	Fetch() string
	Cancel()
}

// ServerLegacy is the legacy server implentation
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
