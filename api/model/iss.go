package model

type IssAppendRequest struct {
	Common
	Account   string      `json:"account" valid:"required,length(1|64)"`
	PublicKey string      `json:"public_key" valid:"required,length(1|256)"`
	Content   interface{} `json:"content" valid:"required"`
	ExtraInfo interface{} `json:"extra_info" valid:"required"`
	// 请求留空本字段，会返回待签名串
	// 将返回的待签名串进行签名
	// 再次请求本协议来真正写入
	Sign string `json:"sign,omitempty" valid:"length(0|256)"`
}

type IssQueryRequest struct {
	Common
	Pages
	Account     string      `json:"account,omitempty" valid:"length(1|64)"`
	THash       string      `json:"t_hash,omitempty" valid:"length(1|256)"`
	BlockHeight []string    `json:"b_height,omitempty" valid:"blockHeight"`
	BlockTime   []string    `json:"b_time,omitempty" valid:"blockTime"`
	Content     interface{} `json:"content,omitempty" valid:"required"`
	ExtraInfo   interface{} `json:"extra_info,omitempty" valid:"required"`
}
