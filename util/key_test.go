package util

import (
	"github.com/KleeTaurus/go-trustsql-sdk/tscec"
	"github.com/magiconair/properties/assert"
	"testing"
)

var pubKey = "Ar275qWzKyJMy+wnCQBDCz11gduAweRJUsyoxnRsFXuA"

func TestGenerateAddrByPubKey(t *testing.T) {
	pubBytes := StringToBytes(pubKey)
	ret, _ := GenerateAddrByPubKey(pubBytes)
	assert.Equal(t, ret, string(tscec.GenerateAddrByPubkey(pubBytes)), "then should be eq")
}
