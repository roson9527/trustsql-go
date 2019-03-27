package easy

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/hokaccha/go-prettyjson"
	"github.com/pkg/errors"
	"github.com/roson9527/trustsql-go/api/model"
	"github.com/roson9527/trustsql-go/httprequest"
	"github.com/roson9527/trustsql-go/util"
	"reflect"
)

func Sign(ety interface{}, signer *util.Signer) {
	l, err := util.Lint(ety)
	fmt.Println(l)
	_ = err

	v := reflect.ValueOf(ety)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	field := v.FieldByName("MchSign")
	if field.CanSet() {
		field.SetString(signer.Signature([]byte(l), false))
	}

	b, _ := json.Marshal(ety)
	bodyJson, _ := prettyjson.Format([]byte(b))
	fmt.Println(string(bodyJson))
}

func SignRenString(signStr string, signer *util.Signer) (string, error) {
	chars, err := hex.DecodeString(signStr)
	if err != nil {
		return signStr, err
	}

	return signer.Signature(chars, true), nil
}

type SignHandler struct {
	signer *util.Signer
}

func NewSignHandler(signer *util.Signer) *SignHandler {
	esh := new(SignHandler)
	esh.signer = signer

	return esh
}

func (esh *SignHandler) Auto(c *httprequest.ChainContext) {
	Sign(c.Data, esh.signer)
}

func (esh *SignHandler) Sign(c *httprequest.ChainContext) {
	switch c.Data.(type) {
	case *model.IssAppendRequest:
		ia := c.Data.(*model.IssAppendRequest)
		if len(ia.Sign) == 0 {
			break
		}

		ia.Sign, _ = SignRenString(ia.Sign, esh.signer)
	}
}

func SignList(ety *[]model.Sign, signerMap map[string]*util.Signer) error {
	for i := 0; i < len(*ety); i++ {
		account := (*ety)[i].Account
		if signer, ok := signerMap[account]; !ok {
			return errors.New(fmt.Sprintf("can't find signer:Account:[%s]", account))
		} else {
			chars, err := hex.DecodeString((*ety)[i].SignStr)
			if err != nil {
				return err
			}

			(*ety)[i].SignRet = signer.Signature(chars, true)
		}
	}

	return nil
}

func AutoFillHandler(c *httprequest.ChainContext) {
	if c.Data == nil {
		return
	}

	switch c.Data.(type) {
	case model.IAutoFill:
		(c.Data).(model.IAutoFill).AutoFill()
	}
}

func ValidHandler(c *httprequest.ChainContext) {
	data := c.Data
	valid, err := govalidator.ValidateStruct(data)
	if !valid && err != nil {
		//fmt.Println("EasyValidHandler Failed:", err)
		c.Error(err)
	}
}
