package easy

import (
	"encoding/json"
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"github.com/roson9527/trustsql-go/api/model"
	"github.com/roson9527/trustsql-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func printJSON(body []byte) {
	bodyJson, _ := prettyjson.Format([]byte(body))
	fmt.Println(string(bodyJson))
}

func TestEasySignList(t *testing.T) {
	easyPrvKey := "IDohU64iE3y4b6ideHUQpsTqOh3+1GgBNeqsKMgYYv8="
	testStr := `[{"id":1,"account":"1EbTFMeUVkVRWHvUtmajs4PDhFeRethfPU","sign_str":"ce302c22abbb9f3238fc57a21f635499bec86562ad4271add946cfb981ee9afe"}]`
	signer, err := util.NewSigner(util.StringToBytes(easyPrvKey))
	signerMap := make(map[string]*util.Signer)
	signerMap[signer.Account()] = signer
	assert.NoError(t, err, "it should be fine")

	var signList []model.Sign
	json.Unmarshal([]byte(testStr), &signList)

	SignList(&signList, signerMap)
	ret, _ := json.Marshal(signList)
	fmt.Println(string(ret))
}
