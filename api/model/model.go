package model

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type UintArray []uint

type IAutoFill interface {
	AutoFill()
}

type CommonBase struct {
	SignType  string `json:"sign_type" valid:"required, in(ECDSA)"`
	MchSign   string `json:"mch_sign,omitempty" valid:"required, length(1|256)"`
	Version   string `json:"version" valid:"required, in(2.0)"`
	MchId     string `json:"mch_id" valid:"required, length(1|32)"`
	ChainId   string `json:"chain_id" valid:"required, length(1|32)"`
	MchPubKey string `json:"mch_pubkey" valid:"required, length(1|64)"`
}

type Common struct {
	CommonBase
	TimeStamp uint64 `json:"timestamp" valid:"required"`
}

type STSCommon struct {
	CommonBase
	TimeStamp uint64 `json:"timestamp,string" valid:"required"`
}

func (ib *CommonBase) AutoFill() {
	ib.Version = "2.0"
	ib.SignType = "ECDSA"
	if len(ib.MchId) == 0 {
		ib.MchId = viper.GetString("info.mch_id")
	}
	if len(ib.MchPubKey) == 0 {
		ib.MchPubKey = viper.GetString("info.mch_pubkey")
	}
}

func (ib *Common) AutoFill() {
	ib.CommonBase.AutoFill()
	ib.TimeStamp = uint64(time.Now().Unix())
}

func (ib *STSCommon) AutoFill() {
	ib.CommonBase.AutoFill()
	ib.TimeStamp = uint64(time.Now().Unix())
	fmt.Println(ib.TimeStamp)
}

// 这里的json属性序列化后顺序需要注意
type Sign struct {
	Account string `json:"account"`
	Id      int32  `json:"id"`
	SignRet string `json:"sign"`
	SignStr string `json:"sign_str"`
}

type MultAsset struct {
	AssetId   string `json:"asset_id" valid:"required, length(0|64)"`
	DstAmount uint64 `json:"dst_amount"`
	FeeAmount uint64 `json:"fee_amount,omitempty"`
}

type Pages struct {
	PageLimit uint32 `json:"page_limit" valid:"required,range(0|20)"`
	PageNo    uint32 `json:"page_no" valid:"required,range(1|100000000)"`
}
