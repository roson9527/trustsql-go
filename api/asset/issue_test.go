package asset

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/roson9527/trustsql-go/api/easy"
	"github.com/roson9527/trustsql-go/api/model"
	"github.com/roson9527/trustsql-go/httprequest"
	"github.com/roson9527/trustsql-go/util"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var prvKey = "IDohU64iE3y4b6ideHUQpsTqOh3+1GgBNeqsKMgYYv8="

func goRegMiddleware() httprequest.IHttpRequest {
	signer, _ := util.NewSigner(util.StringToBytes(prvKey))

	return httprequest.NewMiddleware(
		httprequest.NewGoReq(),
		easy.AutoFillHandler,
		easy.NewSignHandler(signer).Auto,
		easy.ValidHandler)
}

func newIssue() *model.IssueApplyRequest {
	ia := new(model.IssueApplyRequest)
	ia.MchId = "gbaH9TmX8zcYjWXF3P"
	ia.ChainId = "ch_tencent_testchain"
	ia.MchPubKey = "Ar275qWzKyJMy+wnCQBDCz11gduAweRJUsyoxnRsFXuA"
	ia.OwnerUid = "ownerid_jWXF3P"

	ia.OwnerAccountPubKey = ia.MchPubKey
	ia.OwnerAccount, _ = util.GenerateAddrByPubKey(util.StringToBytes(ia.MchPubKey))
	ia.AssetType = 1
	ia.Amount = 10000
	ia.Unit = "fen"
	//
	ia.AutoFill()
	//随机SourceId
	ia.SourceId = strconv.FormatInt(int64(ia.TimeStamp), 10)
	// 更方便的签名方式
	//Sign(ia, signer)
	return ia
}

func newIssueSubmit(ia *model.IssueApplyRequest, body string) *model.SubmitRequest {
	//解析出需要签名的部分并签名
	data, _, _, _ := jsonparser.Get([]byte(body), "sign_str_list")
	//fmt.Println(string(data))
	var signList []model.Sign
	json.Unmarshal(data, &signList)

	signer, _ := util.NewSigner(util.StringToBytes(prvKey))
	signMap := make(map[string]*util.Signer)
	signMap[signer.Account()] = signer
	easy.SignList(&signList, signMap)

	is := new(model.SubmitRequest)
	is.MchPubKey = ia.MchPubKey
	is.MchId = ia.MchId
	is.AssetType = ia.AssetType
	is.ChainId = ia.ChainId
	is.SignList = signList
	is.TransactionId, _ = jsonparser.GetString([]byte(body), "transaction_id")
	fmt.Println(is.TransactionId)
	return is
}

func TestAsset_IssueApply(t *testing.T) {
	a := New(
		"http://119.29.157.123:15910",
		goRegMiddleware())
	_ = a

	ia := newIssue()
	body, errs := a.IssueApply(ia)
	_ = body
	assert.Len(t, errs, 0, "it should be zero")
}

func TestAsset_IssueSubmit(t *testing.T) {
	a := New(
		"http://119.29.157.123:15910",
		goRegMiddleware())
	_ = a

	ia := newIssue()
	ia.Amount = 100000
	body, errs := a.IssueApply(ia)

	is := newIssueSubmit(ia, body)
	body, errs = a.IssueSubmit(is)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
	} else {
		printJSON([]byte(body))
	}
}

func newAsset() string {
	a := New(
		"http://119.29.157.123:15910",
		goRegMiddleware())
	_ = a

	ia := newIssue()
	ia.Amount = 100000
	body, errs := a.IssueApply(ia)

	is := newIssueSubmit(ia, body)
	body, errs = a.IssueSubmit(is)
	if len(errs) > 0 {
		return ""
	}

	data, _ := jsonparser.GetString([]byte(body), "asset_id")
	return data
}
