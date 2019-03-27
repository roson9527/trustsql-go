package httprequest

import (
	"github.com/parnurzeal/gorequest"
)

type HandlerFunc func(*ChainContext)

// HandlersChain defines a HandlerFunc array.
type HandlersChain []HandlerFunc

type ChainContext struct {
	TargetUrl string
	Data      interface{}
	Response  gorequest.Response
	Body      string
	Errors    []error
	handlers  HandlersChain
	index     int8
}

func (c *ChainContext) Next() {
	c.index += 1
	for s := int8(len(c.handlers)); c.index < s && c.index >= 0; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *ChainContext) Error(err error) {
	if err != nil {
		c.Errors = append(c.Errors, err)
	}
	// 无法跳跃到的位置，直接返回
	c.index = 127
}

type Constructor func(c *ChainContext)

// 中间件链
type Chain struct {
	constructors []HandlerFunc
}

func NewChain(hfs ...HandlerFunc) Chain {
	return Chain{append(([]HandlerFunc)(nil), hfs...)}
}

func (ca *Chain) NewContext() *ChainContext {
	c := new(ChainContext)
	c.handlers = ca.constructors
	c.Errors = make([]error, 0)
	c.index = 0

	return c
}

func (ca *Chain) Run(c *ChainContext) {
	// 这样不需要在每一个handler中调用Next()
	for s := int8(len(c.handlers)); c.index < s && c.index >= 0; c.index++ {
		c.handlers[c.index](c)
	}
}
