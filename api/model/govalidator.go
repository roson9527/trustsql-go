package model

import (
	"github.com/asaskevich/govalidator"
	"strconv"
	"time"
)

func init() {
	var (
		state24678     = []uint32{2, 4, 6, 7, 8}
		type1234       = []uint32{1, 2, 3, 4}
		state02348     = []uint32{0, 2, 3, 4, 8}
		state023478    = []uint32{0, 2, 3, 4, 7, 8}
		assetTypeArray = []uint32{1, 2, 3, 4, 5, 6, 7, 8}
	)

	govalidator.CustomTypeTagMap.Set("state24678", buildValid(state24678))
	govalidator.CustomTypeTagMap.Set("type1234", buildValid(type1234))
	govalidator.CustomTypeTagMap.Set("state02348", buildValid(state02348))
	govalidator.CustomTypeTagMap.Set("state023478", buildValid(state023478))
	govalidator.CustomTypeTagMap.Set("assetTypeArray", buildValid(assetTypeArray))
	govalidator.CustomTypeTagMap.Set("blockHeight", isBlockHeight)
	govalidator.CustomTypeTagMap.Set("blockTime", isBlockTime)
}

func buildValid(params []uint32) func(i interface{}, o interface{}) bool {
	return func(i interface{}, o interface{}) bool {
		switch i.(type) {
		case []uint32:
			for _, item := range i.([]uint32) {
				if !isIn(item, params...) {
					return false
				}
			}
		default:
			return false
		}
		return true
	}
}

func isIn(str uint32, params ...uint32) bool {
	for _, param := range params {
		if str == param {
			return true
		}
	}

	return false
}

func isBlockHeight(i interface{}, o interface{}) bool {
	switch i.(type) {
	case []string:
		strList := i.([]string)
		if len(strList) != 2 {
			return false
		}
		start, _ := strconv.ParseUint(strList[0], 10, 64)
		end, _ := strconv.ParseUint(strList[1], 10, 64)
		if start > end {
			return false
		}
	default:
		return false
	}
	return true
}

func isBlockTime(i interface{}, o interface{}) bool {
	switch i.(type) {
	case []string:
		strList := i.([]string)
		if len(strList) != 2 {
			return false
		}
		from, _ := time.Parse(strList[0], "2006-01-02 15:04:05")
		to, _ := time.Parse(strList[1], "2006-01-02 15:04:05")
		if from.After(to) {
			return false
		}
	default:
		return false
	}
	return true
}
