package browser

import "github.com/roson9527/trustsql-go/api/model"

// 没有权限
func (a *Browser) ChainByNode(data *model.ChainByNodeRequest) (string, []error) {
	_, body, errs := a.http.Post(ChainByNode, data)
	return body, errs
}

func (a *Browser) ChainInfo(data *model.ChainInfoRequest) (string, []error) {
	_, body, errs := a.http.Post(ChainInfo, data)
	return body, errs
}

// 没有权限
func (a *Browser) TxInfoByHeight(data *model.TxInfoByHeight) (string, []error) {
	return a.basePost(TxInfoByHeight, data)
}
