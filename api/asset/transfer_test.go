package asset

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	"github.com/roson9527/trustsql-go/api/easy"
	"github.com/roson9527/trustsql-go/api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTransferApply() *model.TransferApplyRequest {
	t := new(model.TransferApplyRequest)
	commonInit(t)
	// From 付款方
	t.SrcUid = "ownerid_jWXF3P"
	t.SrcAccount = defSigner().Account()
	t.SrcAccountPubKey = defSigner().PublicKeyStr()

	// To 收钱方
	t.DstUid = "ownerid_to"
	t.DstAccount = tarSigner().Account()
	t.DstAccountPubKey = tarSigner().PublicKeyStr()

	// Fee 手续费
	t.FeeUid = "ownerid_fee"
	t.FeeAccount = midSigner().Account()
	t.FeeAccountPubKey = midSigner().PublicKeyStr()

	t.AssetType = 1
	t.Amount = 50
	t.FeeAmount = 2

	return t
}

func newTransferSubmit(ia *model.TransferApplyRequest, body string) *model.TransferSubmitRequest { //解析出需要签名的部分并签名
	data, _, _, _ := jsonparser.Get([]byte(body), "sign_str_list")
	//fmt.Println(string(data))
	var signList []model.Sign
	json.Unmarshal(data, &signList)

	easy.SignList(&signList, signers())

	is := new(model.TransferSubmitRequest)
	is.MchPubKey = ia.MchPubKey
	is.MchId = ia.MchId
	is.AssetType = ia.AssetType
	is.ChainId = ia.ChainId
	is.SignList = signList
	is.TransactionId, _ = jsonparser.GetString([]byte(body), "transaction_id")
	return is
}

func TestAsset_TransferApply(t *testing.T) {
	a := New(
		"http://119.29.157.123:15910",
		goRegMiddleware())
	_ = a

	// 测试用例只能使用一次，因为是UTXO模型
	assetId := newAsset()
	ta := newTransferApply()

	ta.SrcAssetList = assetId
	body, errs := a.TransferApply(ta)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		printJSON([]byte(body))
	}
}

func TestAsset_TransferSubmit(t *testing.T) {
	a := New(
		"http://119.29.157.123:15910",
		goRegMiddleware())
	_ = a

	// 测试用例只能使用一次，因为是UTXO模型
	assetId := newAsset()
	ta := newTransferApply()

	ta.SrcAssetList = assetId
	body, errs := a.TransferApply(ta)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		printJSON([]byte(body))
	}

	sb := newTransferSubmit(ta, body)
	body, errs = a.TransferSubmit(sb)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		printJSON([]byte(body))
	}
}
