package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var jsondata = `
	{"inner":{"value1":10,"value2":22},
	"alsoInner":{"value1":20},
	"hello":"world",
	"age": 30
	}`

type HelloWorld struct {
	Hello string `json:"hello"`
}

type JsonTest struct {
	Inner     string `json:"inner"`
	AlsoInner Val1   `json:"alsoInner"`
	HelloWorld
	Age uint `json:"age"`
}

type Val1 struct {
	Value1 int `json:"value1"`
}

var ret = `age=30&alsoInner={"value1":20}&hello=world&inner={"value1":10,"value2":22}`

func TestLintJson(t *testing.T) {
	jRet, err := LintJson(jsondata)
	assert.Equal(t, jRet, ret, "then should be eq")
	assert.NoError(t, err)
}

func TestLint(t *testing.T) {
	jt := new(JsonTest)
	jt.Age = 30
	jt.AlsoInner.Value1 = 20
	jt.Hello = "world"
	jt.Inner = `{"value1":10,"value2":22}`
	jRet, err := Lint(*jt)

	assert.Equal(t, jRet, ret, "then should be eq")
	assert.NoError(t, err)
}
