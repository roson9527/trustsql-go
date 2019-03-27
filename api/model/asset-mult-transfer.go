package model

type AccountAttrs struct {
	SrcUid           string `json:"src_uid" valid:"required,length(1|64)"`
	SrcAccount       string `json:"src_account" valid:"required,length(1|64)"`
	SrcAccountPubKey string `json:"src_account_pubkey" valid:"required,length(1|64)"`
	DstUid           string `json:"dst_uid" valid:"required,length(1|64)"`
	DstAccount       string `json:"dst_account" valid:"required,length(1|64)"`
	DstAccountPubKey string `json:"dst_account_pubkey" valid:"required,length(1|64)"`
	FeeAccount       string `json:"fee_account" valid:"required,length(1|64)"`
	FeeAccountPubKey string `json:"fee_account_pubkey" valid:"required,length(1|64)"`
}

type MultTransferMidApplyRequest struct {
	Common
	AssetTypeCommon
	AccountAttrs
	MultAssets []MultAsset `json:"mult_assets" valid:"required"`
	SignInDate string      `json:"sign_in_date,omitempty" valid:"length(1|32)"`
	ExtraInfo  interface{} `json:"extra_info,omitempty"`
}

type MultTransferMidSubmitRequest struct {
	Common
	AssetTypeCommon
	TransactionId string `json:"transaction_id" valid:"required,length(1|32)"`
	SignList      []Sign `json:"sign_list"`
}

type MultSignInApplyRequest struct {
	Common
	AssetTypeCommon
	MidAssetId string      `json:"mid_asset_id" valid:"required,length(1|64)"`
	OpCode     uint32      `json:"op_code" valid:"required,range(1|3)"`
	ExtraInfo  interface{} `json:"extra_info,omitempty"`
}

type MultSignInSubmitRequest struct {
	Common
	AssetTypeCommon
	TransactionId string `json:"transaction_id" valid:"required,length(1|32)"`
	SignList      []Sign `json:"sign_list"`
}
