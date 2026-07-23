package echo

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestContextNoStaleValues(t *testing.T) {
	pool := &ContextPool{}
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	c := pool.Acquire(w, r)
	c.Set("user", "alice")
	c.Set("token", "abc123")
	pool.Release(c)

	c2 := pool.Acquire(w, r)
	if val := c2.Get("user"); val != nil {
		t.Errorf("expected nil for 'user', got %v", val)
	}
	if val := c2.Get("token"); val != nil {
		t.Errorf("expected nil for 'token', got %v", val)
	}
}

func TestContextConcurrentSafe(t *testing.T) {
	pool := &ContextPool{}
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	c := pool.Acquire(w, r)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.Set("key", i)
			c.Get("key")
		}(i)
	}
	wg.Wait()
}
