package model

type OwnerParams struct {
	OwnerUid     string `json:"owner_uid" valid:"required, length(1|64)"`
	OwnerAccount string `json:"owner_account" valid:"required, length(1|64)"`
}

type AssetTypeCommon struct {
	AssetType uint64 `json:"asset_type" valid:"range(0|32)"`
}

type IssueApplyRequest struct {
	Common
	OwnerParams
	AssetTypeCommon
	SourceId           string      `json:"source_id" valid:"required, length(1|64)"`
	OwnerAccountPubKey string      `json:"owner_account_pubkey" valid:"required, length(1|64)"`
	Amount             uint64      `json:"amount" valid:"required, range(1|10000000000000)"`
	Unit               string      `json:"unit" valid:"required, length(1|32)"`
	Content            interface{} `json:"content,omitempty"`
}

type SubmitRequest struct {
	Common
	AssetTypeCommon
	TransactionId string `json:"transaction_id" valid:"required"`
	SignList      []Sign `json:"sign_list"`
}
