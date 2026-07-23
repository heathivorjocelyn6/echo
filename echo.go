package echo

import (
	"net/http"
)

func (e *Echo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := e.pool.Get().(*context)
	c.Reset(r, w)

	// ... execution logic ...

	// Check if hijacked or otherwise unsafe to return to pool
	if !c.response.hijacked {
		e.pool.Put(c)
	}
}