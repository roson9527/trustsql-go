package iss

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/roson9527/trustsql-go/api/model"
	"github.com/roson9527/trustsql-go/httprequest"
)

type Iss struct {
	host string
	http httprequest.IHttpRequest
}

func New(host string, h httprequest.IHttpRequest) *Iss {
	a := new(Iss)
	a.http = h
	a.host = host
	return a
}

func (a *Iss) basePost(path string, data interface{}) (string, []error) {
	_, body, errs := a.http.Post(fmt.Sprintf("%s/%s", a.host, path), data)
	return body, errs
}

func (a *Iss) Append(data *model.IssAppendRequest) (string, []error) {
	return a.basePost(IssAppend, data)
}

func (a *Iss) EasyAppend(data *model.IssAppendRequest) (string, []error) {
	body, errs := a.Append(data)
	if errs != nil {
		return body, errs
	}

	signStr, err := jsonparser.GetString([]byte(body), "sign_str")
	if err != nil {
		return body, []error{err}
	}

	data.Sign = signStr
	return a.Append(data)
	//
	//chars, err := hex.DecodeString(signStr)
	//if err!=nil {
	//	return signStr, []error{err}
	//}
}

func (a *Iss) Query(data *model.IssQueryRequest) (string, []error) {
	return a.basePost(IssQuery, data)
}
