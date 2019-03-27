package asset

import (
	"encoding/json"
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"github.com/roson9527/trustsql-go/api/easy"
	"github.com/roson9527/trustsql-go/api/model"
	"github.com/roson9527/trustsql-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/timest/go-pprint"
	"reflect"
	"testing"
)

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

// 资产批量查询
// Done!
func TestAsset_AccountQuery(t *testing.T) {
	aq := new(model.AccountQueryRequest)
	aq.MchId = "gbaH9TmX8zcYjWXF3P"
	aq.ChainId = "ch_tencent_testchain"
	aq.MchPubKey = "Ar275qWzKyJMy+wnCQBDCz11gduAweRJUsyoxnRsFXuA"
	aq.OwnerUid = "ownerid_jWXF3P"
	aq.OwnerAccount, _ = util.GenerateAddrByPubKey(util.StringToBytes(aq.MchPubKey))
	aq.AssetId = "26aN6NjtqUed6e9zxNAVwcwgx5vt1XgeEteZHHWAV6YL6dU"
	aq.State = []uint32{0, 2, 3}
	aq.PageLimit = 5
	aq.PageNo = 1
	pprint.Format(aq)

	a := New(
		"http://119.29.157.123:15910",
		goRegMiddleware())
	_ = a

	body, errs := a.AccountQuery(aq)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
	} else {
		printJSON([]byte(body))
	}
}

// 交易批量查询
// Done！已经修正，原因是lint时omitempty属性字段也被打包进去
func TestAsset_TransBatchQuery(t *testing.T) {
	aq := new(model.TransBatchQueryRequest)
	commonInit(aq)
	aq.TransactionId = "201811150007074552"
	aq.PageLimit = 5
	aq.PageNo = 1

	a := New(
		"http://119.29.157.123:15910",
		goRegMiddleware())
	_ = a

	body, errs := a.TransBatchQuery(aq)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
	} else {
		printJSON([]byte(body))
	}
}

// 资产汇总对账
func TestAsset_RecSumQuery(t *testing.T) {
	es := new(model.RecSumQueryRequest)
	commonInit(es)
	es.State = []uint32{8}
	es.AssetType = []uint32{1}
	es.Date = "2018-11-15"
	//es.AutoFill()
	//easy.Sign(es, defSigner())

	a := New(
		"http://119.29.157.123:15910",
		goRegMiddleware())
	_ = a

	body, errs := a.RecSumQuery(es)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
	} else {
		printJSON([]byte(body))
	}
}

// 资产明细对账
// Done！
func TestAsset_RecDetailQuery(t *testing.T) {
	es := new(model.RecDetailQueryRequest)
	commonInit(es)
	es.State = []uint32{0}
	es.AssetType = []uint32{1}
	es.Date = "2018-11-15"
	es.PageLimit = 5
	es.PageNo = 1

	//pprint.Format(es)

	a := New(
		"http://119.29.157.123:15910",
		goRegMiddleware())
	_ = a

	body, errs := a.RecDetailQuery(es)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
	} else {
		printJSON([]byte(body))
	}
}

// 交易汇总对账
// Done！
func TestAsset_TransRecSumQuery(t *testing.T) {
	es := new(model.TransRecSumQueryRequest)
	commonInit(es)
	es.State = []uint32{2, 4}
	es.Date = "2018-11-15"
	es.TransType = []uint32{1, 2, 3, 4}
	es.AutoFill()
	easy.Sign(es, defSigner())

	//pprint.Format(es)
	b, _ := json.Marshal(es)
	printJSON(b)

	a := New(
		"http://119.29.157.123:15910",
		goRegMiddleware())
	_ = a

	body, errs := a.TransRecSumQuery(es)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
	} else {
		printJSON([]byte(body))
	}
}

// 交易明细对账
// Done！已经修正，原因：timestamp是string而不是uint64 emmmmmm
func TestAsset_TransRecDetailQuery(t *testing.T) {
	ia := new(model.TransRecDetailQueryRequest)
	commonInit(ia)
	ia.State = []uint32{2, 4, 6, 7, 8}
	ia.AssetType = []uint32{1}
	ia.TransType = []uint32{1, 2, 3, 4}
	ia.Date = "2018-11-15"
	ia.PageLimit = 5
	ia.PageNo = 1
	//ia.TimeStamp = fmt.Sprintf("%d", uint64(time.Now().Unix()))

	a := New(
		"http://119.29.157.123:15910",
		goRegMiddleware())
	_ = a

	body, errs := a.TransRecDetailQuery(ia)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
	} else {
		printJSON([]byte(body))
	}
}
