package asset

import "github.com/roson9527/trustsql-go/api/model"

func (a *Asset) MultTransferMidApply(data *model.MultTransferMidApplyRequest) (string, []error) {
	return a.basePost(MultTransferMidApply, data)
}

func (a *Asset) MultTransferMidSubmit(data *model.MultTransferMidSubmitRequest) (string, []error) {
	return a.basePost(MultTransferMidSubmit, data)
}

func (a *Asset) MultSignInApply(data *model.MultSignInApplyRequest) (string, []error) {
	return a.basePost(MultSignInApply, data)
}

func (a *Asset) MultSignInSubmit(data *model.MultSignInSubmitRequest) (string, []error) {
	return a.basePost(MultSignInSubmit, data)
}
