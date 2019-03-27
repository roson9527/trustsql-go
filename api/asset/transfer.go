package asset

import "github.com/roson9527/trustsql-go/api/model"

func (a *Asset) TransferApply(data *model.TransferApplyRequest) (string, []error) {
	return a.basePost(TransferApply, data)
}

func (a *Asset) TransferSubmit(data *model.TransferSubmitRequest) (string, []error) {
	return a.basePost(TransferSubmit, data)
}
