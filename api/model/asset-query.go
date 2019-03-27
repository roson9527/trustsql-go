package model

//资产批量查询
type AccountQueryRequest struct {
	Common
	OwnerParams
	Pages
	AssetId string   `json:"asset_id" valid:"required, length(0|64)"`
	State   []uint32 `json:"state" valid:"state023478"`
}

// 交易批量查询
type TransBatchQueryRequest struct {
	Common
	Pages
	TransactionId string `json:"transaction_id" valid:"length(1|32)"`

	SrcAccount  string   `json:"src_account,omitempty" valid:"length(1|64)"`
	DstAccount  string   `json:"dst_account,omitempty" valid:"length(1|64)"`
	TransType   []uint32 `json:"trans_type,omitempty" valid:"type1234"`
	BlockHeight uint64   `json:"b_height,omitempty" valid:"length(1|64)"`
	TransHash   string   `json:"trans_hash,omitempty" valid:"length(1|64)"`
	State       []uint32 `json:"state,omitempty" valid:"state24678"`
}

// 资产汇总对账
type RecSumQueryRequest struct {
	STSCommon
	AssetType []uint32 `json:"asset_type" valid:"required"`
	Date      string   `json:"date" valid:"required,length(1|64)"`
	State     []uint32 `json:"state" valid:"state02348"`
}

// 资产明细对账
type RecDetailQueryRequest struct {
	STSCommon
	Pages
	AssetType []uint32 `json:"asset_type" valid:"required"`
	Date      string   `json:"date" valid:"required,length(1|64)"`
	State     []uint32 `json:"state" valid:"state02348"`
}

// 交易汇总对账
type TransRecSumQueryRequest struct {
	STSCommon
	Date      string   `json:"date" valid:"required, length(1|64)"`
	State     []uint32 `json:"state" valid:"required, state24678"`
	TransType []uint32 `json:"trans_type" valid:"required, type1234"`
}

// 交易明细对账
type TransRecDetailQueryRequest struct {
	STSCommon
	Pages
	AssetType []uint32 `json:"asset_type" valid:"required, assetTypeArray"`
	Date      string   `json:"date" valid:"required, length(1|64)"`
	State     []uint32 `json:"state" valid:"required, state24678"`
	TransType []uint32 `json:"trans_type" valid:"required, type1234"`
}
