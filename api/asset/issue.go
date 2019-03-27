package asset

import "github.com/roson9527/trustsql-go/api/model"

func (a *Asset) IssueApply(data *model.IssueApplyRequest) (string, []error) {
	return a.basePost(IssueApply, data)
}

func (a *Asset) IssueSubmit(data *model.SubmitRequest) (string, []error) {
	return a.basePost(IssueSubmit, data)
}
