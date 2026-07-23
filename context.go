package echo

import (
	"net/http"
)

type context struct {
	request  *http.Request
	response *Response
	store    Map
	path     string
	pnames   []string
	pvalues  []string
	handler  HandlerFunc
}

func (c *context) Reset(r *http.Request, w http.ResponseWriter) {
	c.request = r
	c.response.reset(w)
	c.store = nil
	c.path = ""
	c.pnames = nil
	c.pvalues = nil
	c.handler = NotFoundHandler
}

func (c *context) Clone() Context {
	clone := &context{
		request:  c.request,
		response: c.response,
		path:     c.path,
		pnames:   c.pnames,
		pvalues:  c.pvalues,
		handler:  c.handler,
	}
	if c.store != nil {
		clone.store = make(Map)
		for k, v := range c.store {
			clone.store[k] = v
		}
	}
	return clone
}