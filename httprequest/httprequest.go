package httprequest

import (
	"github.com/parnurzeal/gorequest"
)

type IHttpRequest interface {
	Post(targetUrl string, data interface{}) (gorequest.Response, string, []error)
}

type GoReq struct {
}

func NewGoReq() *GoReq {
	return new(GoReq)
}

func (g *GoReq) Post(targetUrl string, data interface{}) (gorequest.Response, string, []error) {
	gr := gorequest.New()
	return gr.Post(targetUrl).Send(data).End()
}

type Middleware struct {
	h     IHttpRequest
	chain Chain
}

func NewMiddleware(h IHttpRequest, mid ...HandlerFunc) *Middleware {
	grc := new(Middleware)
	mid = append(mid, func(c *ChainContext) {
		c.Response, c.Body, c.Errors = h.Post(c.TargetUrl, c.Data)
	})
	grc.chain = NewChain(mid...)

	return grc
}

func (grc *Middleware) Post(targetUrl string, data interface{}) (gorequest.Response, string, []error) {
	c := grc.chain.NewContext()
	c.TargetUrl = targetUrl
	c.Data = data

	grc.chain.Run(c)
	return c.Response, c.Body, c.Errors
}
