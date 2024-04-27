package utils

import (
	"go_study/src/common"
)

func PageToObject(page *common.Page) map[string]interface{} {
	//获取到当前页数
	offset := (page.PageNum - 1) * page.PageSize
	if offset < 0 {
		page.Offset = 0
	} else {
		page.Offset = offset
	}
	word := page.KeyWord
	queryMap := make(map[string]interface{})
	if word != nil {
		wordMap := word.(map[string]interface{})
		queryMap = MapKeyConvUnderline(wordMap)
	}
	return queryMap
}
