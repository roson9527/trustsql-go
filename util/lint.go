package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/buger/jsonparser"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func Lint(ety interface{}) (string, error) {
	signMap := make(map[string]string)

	err := checkString(&signMap, reflect.ValueOf(ety))
	if err != nil {
		return "", err
	}

	return lint(&signMap), nil
}

func checkString(m *map[string]string, v reflect.Value) error {
	for i := 0; i < v.NumField(); i += 1 {
		tagAttr := v.Type().Field(i).Tag.Get("json")
		tag := strings.Split(tagAttr, ",")[0]

		if "sign" == tag {
			continue
		}

		switch v.Field(i).Kind() {
		case reflect.Uint64:
			(*m)[tag] = strconv.FormatUint(v.Field(i).Interface().(uint64), 10)

		case reflect.Uint:
			(*m)[tag] = strconv.FormatUint(uint64(v.Field(i).Interface().(uint)), 10)

		case reflect.Int64:
			(*m)[tag] = strconv.FormatInt(v.Field(i).Interface().(int64), 10)

		case reflect.Int:
			(*m)[tag] = strconv.Itoa(v.Field(i).Interface().(int))

		case reflect.Map:
			mapField := v.Field(i).Interface().(map[string]interface{})
			if mapField == nil {

			}
			value, err := json.Marshal(mapField)
			if err != nil {
				return err
			}
			(*m)[tag] = string(value)

		case reflect.Struct:
			// 处理组合问题
			if len(tag) == 0 {
				err := checkString(m, v.Field(i))
				if err != nil {
					return err
				}
			} else {
				b, err := json.Marshal(v.Field(i).Interface())
				if err != nil {
					return err
				}
				(*m)[tag] = string(b)
			}
		default:
			(*m)[tag] = v.Field(i).Interface().(string)
		}
	}

	return nil
}

func LintJson(data string) (string, error) {
	signMap := make(map[string]string)

	jsonparser.ObjectEach([]byte(data), func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		signMap[string(key)] = string(value)
		return nil
	})

	return lint(&signMap), nil
}

func lint(m *map[string]string) string {
	keys := make([]string, 0)
	for k := range *m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	lintStr := ""

	for _, k := range keys {
		val := (*m)[k]
		if k == "" || k == "mch_sign" {
			continue
		}

		if len(lintStr) != 0 {
			lintStr += "&"
		}
		lintStr += fmt.Sprintf("%s=%s", k, val)
	}

	return lintStr
}

func contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}
