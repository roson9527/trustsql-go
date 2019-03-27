package iss

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/hokaccha/go-prettyjson"
	"github.com/roson9527/trustsql-go/api/easy"
	"github.com/roson9527/trustsql-go/api/model"
	"github.com/roson9527/trustsql-go/httprequest"
	"github.com/roson9527/trustsql-go/util"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var prvKey = "IDohU64iE3y4b6ideHUQpsTqOh3+1GgBNeqsKMgYYv8="

type HelloX struct {
	Hello string `json:"hello"`
}

func goRegMiddleware() httprequest.IHttpRequest {
	signer := defSigner()

	return httprequest.NewMiddleware(
		httprequest.NewGoReq(),
		easy.AutoFillHandler,
		easy.NewSignHandler(signer).Sign,
		easy.NewSignHandler(signer).Auto,
		easy.ValidHandler)
}

func printJSON(body []byte) {
	bodyJson, _ := prettyjson.Format([]byte(body))
	fmt.Println(string(bodyJson))
}

func commonInit(common interface{}) {
	v := reflect.ValueOf(common)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	v.FieldByName("MchId").SetString("gbaH9TmX8zcYjWXF3P")
	v.FieldByName("ChainId").SetString("ch_tencent_testchain")
	v.FieldByName("MchPubKey").SetString("Ar275qWzKyJMy+wnCQBDCz11gduAweRJUsyoxnRsFXuA")
}

func defSigner() *util.Signer {
	signer, _ := util.NewSigner(util.StringToBytes(prvKey))
	return signer
}

func TestIss_Append(t *testing.T) {
	ia := new(model.IssAppendRequest)
	commonInit(ia)

	ia.Content = HelloX{
		Hello: "world",
	}
	ia.ExtraInfo = HelloX{
		Hello: "world",
	}
	ia.PublicKey = ia.MchPubKey
	ia.Account, _ = util.GenerateAddrByPubKey(util.StringToBytes(ia.PublicKey))

	a := New(
		"http://119.29.157.123:15903",
		goRegMiddleware())
	_ = a

	body, errs := a.Append(ia)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		//printJSON([]byte(body))
	}

	// 再次写入
	signStr, err := jsonparser.GetString([]byte(body), "sign_str")
	if !assert.NoError(t, err, "err should be nil") {
		return
	}

	ia.Sign = signStr
	//
	//chars, err := hex.DecodeString(signStr)
	//if !assert.NoError(t, err, "err should be nil") {
	//	return
	//}
	//
	//ia.Sign = defSigner().Signature(chars, true)

	body, errs = a.Append(ia)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		printJSON([]byte(body))
	}
}

func TestIss_EasyAppend(t *testing.T) {
	ia := new(model.IssAppendRequest)
	commonInit(ia)

	ia.Content = HelloX{
		Hello: "world",
	}
	ia.ExtraInfo = HelloX{
		Hello: "world",
	}
	ia.PublicKey = ia.MchPubKey
	ia.Account, _ = util.GenerateAddrByPubKey(util.StringToBytes(ia.PublicKey))

	a := New(
		"http://119.29.157.123:15903",
		goRegMiddleware())
	_ = a

	body, errs := a.EasyAppend(ia)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		printJSON([]byte(body))
	}
}

func TestIss_Query(t *testing.T) {
	iq := new(model.IssQueryRequest)
	commonInit(iq)
	//iq.Account, _ = util.GenerateAddrByPubKey(util.StringToBytes(iq.MchPubKey))
	iq.PageLimit = 10
	iq.PageNo = 1
	iq.THash = "8a83543b4a131b878d7ca212ebf47751bed2bcecb58e0ef32db4ef0d70ae824c"

	a := New(
		"http://119.29.157.123:15903",
		goRegMiddleware())
	_ = a

	body, errs := a.Query(iq)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		printJSON([]byte(body))
	}
}
