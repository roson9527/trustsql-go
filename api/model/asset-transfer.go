package model

type TransferApplyRequest struct {
	Common
	AssetTypeCommon
	AccountAttrs
	FeeUid       string      `json:"fee_uid" valid:"required,length(1|64)"`
	SrcAssetList string      `json:"src_asset_list" valid:"required,length(1|1024)"`
	Amount       uint32      `json:"amount"`
	FeeAmount    uint32      `json:"fee_amount,omitempty"`
	ExtraInfo    interface{} `json:"extra_info,omitempty"`
}

type TransferSubmitRequest struct {
	Common
	AssetTypeCommon
	TransactionId string `json:"transaction_id" valid:"required,length(1|32)"`
	SignList      []Sign `json:"sign_list"`
}
