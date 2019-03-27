package asset

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	"github.com/roson9527/trustsql-go/api/easy"
	"github.com/roson9527/trustsql-go/api/model"
	"github.com/roson9527/trustsql-go/util"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func autoFill(tar interface{}, src interface{}) {
	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := reflect.ValueOf(tar)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < v.NumField(); i += 1 {
		name := v.Type().Field(i).Name
		field := t.FieldByName(name)
		if field.CanSet() {
			field.Set(v.Field(i))
		}
	}
}

func defSigner() *util.Signer {
	signer, _ := util.NewSigner(util.StringToBytes(prvKey))
	return signer
}

func tarSigner() *util.Signer {
	signer, _ := util.NewSigner(util.StringToBytes(tarPrvKey))
	return signer
}

func midSigner() *util.Signer {
	signer, _ := util.NewSigner(util.StringToBytes(midPrvKey))
	return signer
}

func signers() map[string]*util.Signer {
	ss := make(map[string]*util.Signer)
	ss[defSigner().Account()] = defSigner()
	ss[tarSigner().Account()] = tarSigner()
	ss[midSigner().Account()] = midSigner()
	return ss
}

var (
	tarPrvKey = "vsFauL3irj+KLfHXPcPd9MM8TAKL4oT2rdeSjdikQfo="
	midPrvKey = "wyCtuGSGPmae/IP5puxAy7AW4DbKaX5fyOfvJlSAZRI="
)

func newMultMid() *model.MultTransferMidApplyRequest {
	m := new(model.MultTransferMidApplyRequest)
	commonInit(m)

	// From 付款方
	m.SrcUid = "ownerid_jWXF3P"
	m.SrcAccount = defSigner().Account()
	m.SrcAccountPubKey = defSigner().PublicKeyStr()

	// To 收钱方
	m.DstUid = "ownerid_to"
	m.DstAccount = tarSigner().Account()
	m.DstAccountPubKey = tarSigner().PublicKeyStr()

	// Fee 手续费
	m.FeeAccount = midSigner().Account()
	m.FeeAccountPubKey = midSigner().PublicKeyStr()

	m.MultAssets = make([]model.MultAsset, 0)
	//m.MultAssets = append(m.MultAssets, model.MultAsset{})

	return m
}

func newMultSubmit(ia *model.MultTransferMidApplyRequest, body string) *model.MultTransferMidSubmitRequest { //解析出需要签名的部分并签名
	data, _, _, _ := jsonparser.Get([]byte(body), "sign_str_list")
	//fmt.Println(string(data))
	var signList []model.Sign
	json.Unmarshal(data, &signList)

	easy.SignList(&signList, signers())

	is := new(model.MultTransferMidSubmitRequest)
	is.MchPubKey = ia.MchPubKey
	is.MchId = ia.MchId
	is.AssetType = ia.AssetType
	is.ChainId = ia.ChainId
	is.SignList = signList
	is.TransactionId, _ = jsonparser.GetString([]byte(body), "transaction_id")
	return is
}

func newSignInSubmit(ia *model.MultSignInApplyRequest, body string) *model.MultSignInSubmitRequest { //解析出需要签名的部分并签名
	data, _, _, _ := jsonparser.Get([]byte(body), "sign_str_list")
	//fmt.Println(string(data))
	var signList []model.Sign
	json.Unmarshal(data, &signList)

	easy.SignList(&signList, signers())

	is := new(model.MultSignInSubmitRequest)
	is.MchPubKey = ia.MchPubKey
	is.MchId = ia.MchId
	is.AssetType = ia.AssetType
	is.ChainId = ia.ChainId
	is.SignList = signList
	is.TransactionId, _ = jsonparser.GetString([]byte(body), "transaction_id")
	return is
}

func TestAsset_MultTransferMidApply(t *testing.T) {
	a := New(
		"http://119.29.157.123:15910",
		goRegMiddleware())
	_ = a

	assetId := newAsset()
	mm := newMultMid()
	mm.AssetType = 1
	mm.MultAssets = append(mm.MultAssets, model.MultAsset{
		AssetId:   assetId,
		DstAmount: 1,
		FeeAmount: 1,
	})

	body, errs := a.MultTransferMidApply(mm)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
	} else {
		printJSON([]byte(body))
	}
}

func TestAsset_MultTransferMidSubmit(t *testing.T) {
	a := New(
		"http://119.29.157.123:15910",
		goRegMiddleware())
	_ = a

	// 测试用例只能使用一次，因为是UTXO模型
	//assetId := "26aRp8fn4EpxRknbNwQBTFtLkwMsFnVZwY83uF6aUQCVSrd"
	assetId := newAsset()
	mm := newMultMid()
	mm.AssetType = 1
	mm.MultAssets = append(mm.MultAssets, model.MultAsset{
		AssetId:   assetId,
		DstAmount: 1,
		FeeAmount: 1,
	})

	body, errs := a.MultTransferMidApply(mm)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		printJSON([]byte(body))
	}

	is := newMultSubmit(mm, body)
	body, errs = a.MultTransferMidSubmit(is)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		printJSON([]byte(body))
	}
}

func TestAsset_MultSignInApplyAndSubmit(t *testing.T) {
	a := New(
		"http://119.29.157.123:15910",
		goRegMiddleware())
	_ = a

	// 测试用例只能使用一次，因为是UTXO模型
	//assetId := "26aRp8fn4EpxRknbNwQBTFtLkwMsFnVZwY83uF6aUQCVSrd"
	assetId := newAsset()
	mm := newMultMid()
	mm.AssetType = 1
	mm.MultAssets = append(mm.MultAssets, model.MultAsset{
		AssetId:   assetId,
		DstAmount: 1,
		FeeAmount: 1,
	})

	body, errs := a.MultTransferMidApply(mm)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		printJSON([]byte(body))
	}

	is := newMultSubmit(mm, body)
	body, errs = a.MultTransferMidSubmit(is)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		printJSON([]byte(body))
	}

	signInApply := new(model.MultSignInApplyRequest)
	commonInit(signInApply)
	signInApply.AssetType = mm.AssetType
	signInApply.MidAssetId, _ = jsonparser.GetString([]byte(body), "mid_asset_id")
	signInApply.OpCode = 1

	body, errs = a.MultSignInApply(signInApply)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		printJSON([]byte(body))
	}

	signSubmit := newSignInSubmit(signInApply, body)
	body, errs = a.MultSignInSubmit(signSubmit)
	if !assert.Len(t, errs, 0, "it should be zero") {
		assert.Error(t, errs[0])
		return
	} else {
		printJSON([]byte(body))
	}
}
