package asset

import (
	"fmt"
	"github.com/roson9527/trustsql-go/httprequest"
)

type Asset struct {
	host string
	http httprequest.IHttpRequest
}

func New(host string, h httprequest.IHttpRequest) *Asset {
	a := new(Asset)
	a.http = h
	a.host = host
	return a
}

func (a *Asset) basePost(path string, data interface{}) (string, []error) {
	_, body, errs := a.http.Post(fmt.Sprintf("%s/%s", a.host, path), data)
	return body, errs
}
