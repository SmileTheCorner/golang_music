package utils

import (
	"go_study/src/core/singer/entity"
	"strconv"
)

func GroupByList(data []entity.Singer) map[string][]entity.Singer {
	if data == nil {
		return nil
	}
	// 创建一个初始map
	dataMap := make(map[string][]entity.Singer)
	for _, item := range data {
		//获取歌手分类
		classify := item.SingerClassify
		key := strconv.Itoa(int(classify))
		if list, ok := dataMap[key]; ok {
			list = append(list, item)
			dataMap[key] = list
		} else {
			singers := make([]entity.Singer, 0)
			singers = append(singers, item)
			dataMap[key] = singers
		}
	}
	return dataMap
}
