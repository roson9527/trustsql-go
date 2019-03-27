package asset

import "github.com/roson9527/trustsql-go/api/model"

// 资产批量查询
func (a *Asset) AccountQuery(data *model.AccountQueryRequest) (string, []error) {
	return a.basePost(AccountQuery, data)
}

// 交易批量查询
func (a *Asset) TransBatchQuery(data *model.TransBatchQueryRequest) (string, []error) {
	//TODO test
	return a.basePost(TransBatchQuery, data)
}

// 资产汇总对账
func (a *Asset) RecSumQuery(data *model.RecSumQueryRequest) (string, []error) {
	//TODO test
	return a.basePost(RecSumQuery, data)
}

// 资产明细对账
func (a *Asset) RecDetailQuery(data *model.RecDetailQueryRequest) (string, []error) {
	//TODO test
	return a.basePost(RecDetailQuery, data)
}

// 交易汇总对账
func (a *Asset) TransRecSumQuery(data *model.TransRecSumQueryRequest) (string, []error) {
	//TODO test
	return a.basePost(TransRecSumQuery, data)
}

// 交易明细对账
func (a *Asset) TransRecDetailQuery(data *model.TransRecDetailQueryRequest) (string, []error) {
	//TODO test
	return a.basePost(TransRecDetailQuery, data)
}
