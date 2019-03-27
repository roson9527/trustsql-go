package browser

import (
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"github.com/roson9527/trustsql-go/api/easy"
	"github.com/roson9527/trustsql-go/api/model"
	"github.com/roson9527/trustsql-go/httprequest"
	"github.com/roson9527/trustsql-go/util"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

var prvKey = "IDohU64iE3y4b6ideHUQpsTqOh3+1GgBNeqsKMgYYv8="

func printJSON(body []byte) {
	bodyJson, _ := prettyjson.Format([]byte(body))
	fmt.Println(string(bodyJson))
}

func goRegMiddleware() httprequest.IHttpRequest {
	signer, _ := util.NewSigner(util.StringToBytes(prvKey))

	return httprequest.NewMiddleware(
		httprequest.NewGoReq(),
		easy.AutoFillHandler,
		easy.NewSignHandler(signer).Auto,
		easy.ValidHandler)
}

func commonInit(common interface{}) {
	v := reflect.ValueOf(common)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	v.FieldByName("TimeStamp").SetUint(uint64(time.Now().Unix()))
	v.FieldByName("MchId").SetString("gbaH9TmX8zcYjWXF3P")
	//v.FieldByName("ChainId").SetString("ch_tencent_testchain")
	//v.FieldByName("MchPubKey").SetString("Ar275qWzKyJMy+wnCQBDCz11gduAweRJUsyoxnRsFXuA")
}

func TestBrowser_ChainByNode(t *testing.T) {
	a := New(
		"http://119.29.157.123:15909",
		goRegMiddleware())
	_ = a

	data:=new(model.ChainByNodeRequest)
	commonInit(data)
	data.NodeId = "ndsjNkjjEt6D9nu5YE"

	body, errs := a.ChainByNode(data)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		printJSON([]byte(body))
	}
}

func TestBrowser_ChainInfo(t *testing.T) {
	a := New(
		"http://119.29.157.123:15909",
		goRegMiddleware())
	_ = a

	data:=new(model.ChainInfoRequest)
	commonInit(data)
	data.ChainId = "ch_tencent_testchain"

	body, errs := a.ChainInfo(data)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		printJSON([]byte(body))
	}
}

func TestBrowser_TxInfoByHeight(t *testing.T) {
	a := New(
		"http://119.29.157.123:15909",
		goRegMiddleware())
	_ = a

	data:=new(model.TxInfoByHeight)
	commonInit(data)
	data.ChainId = "ch_tencent_testchain"
	data.BeginHeight = 93080
	data.EndHeight = 93085
	data.MchPubKey = "Ar275qWzKyJMy+wnCQBDCz11gduAweRJUsyoxnRsFXuA"

	body, errs := a.TxInfoByHeight(data)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		printJSON([]byte(body))
	}
}