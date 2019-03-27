package browser

import (
	"fmt"
	"github.com/roson9527/trustsql-go/httprequest"
)

type Browser struct {
	host string
	http httprequest.IHttpRequest
}

func New(host string, h httprequest.IHttpRequest) *Browser {
	a := new(Browser)
	a.http = h
	a.host = host
	return a
}

func (a *Browser) basePost(path string, data interface{}) (string, []error) {
	_, body, errs := a.http.Post(fmt.Sprintf("%s/%s", a.host, path), data)
	return body, errs
}
