package echo

import (
	"net/http"
	"sync"
)

type Context struct {
	keys     map[string]any
	mu       sync.RWMutex
	response http.ResponseWriter
	request  *http.Request
}

type ContextPool struct {
	pool sync.Pool
}

func (cp *ContextPool) Acquire(w http.ResponseWriter, r *http.Request) *Context {
	got := cp.pool.Get()
	var c *Context
	if got == nil {
		c = &Context{}
	} else {
		c = got.(*Context)
	}
	c.response = w
	c.request = r
	c.keys = make(map[string]any)
	return c
}

func (cp *ContextPool) Release(c *Context) {
	c.keys = nil
	c.request = nil
	c.response = nil
	cp.pool.Put(c)
}

func (c *Context) Set(key string, val any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.keys == nil {
		c.keys = make(map[string]any)
	}
	c.keys[key] = val
}

func (c *Context) Get(key string) any {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.keys == nil {
		return nil
	}
	return c.keys[key]
}
