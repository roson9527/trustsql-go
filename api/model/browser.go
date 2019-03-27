package model

type ChainByNodeRequest struct {
	NodeId    string `json:"node_id" valid:"required, length(1|32)"`
	MchId     string `json:"mch_id" valid:"required, length(1|32)"`
	TimeStamp uint32 `json:"timestamp" valid:"required"`
	MchSign   string `json:"mch_sign,omitempty" valid:"required, length(1|256)"`
}

type ChainInfoRequest struct {
	ChainId   string `json:"chain_id" valid:"required, length(1|32)"`
	MchId     string `json:"mch_id" valid:"required, length(1|32)"`
	TimeStamp uint32 `json:"timestamp" valid:"required"`
	MchSign   string `json:"mch_sign,omitempty" valid:"required, length(1|256)"`
}

type TxInfoByHeight struct {
	NodeId      string `json:"node_id,omitempty" valid:"length(1|32)"`
	ChainId     string `json:"chain_id" valid:"required, length(1|32)"`
	BeginHeight uint32 `json:"begin_height" valid:"required"`
	EndHeight   uint32 `json:"end_height" valid:"required"`
	MchId       string `json:"mch_id" valid:"required, length(1|32)"`
	TimeStamp   uint32 `json:"timestamp" valid:"required"`
	MchSign     string `json:"mch_sign,omitempty" valid:"required, length(1|256)"`
	MchPubKey   string `json:"mch_pubkey" valid:"required, length(1|64)"`
}
