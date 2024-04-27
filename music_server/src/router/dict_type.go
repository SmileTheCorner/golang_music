package router

import (
	"github.com/gin-gonic/gin"
	"go_study/src/sys/dict/api"
)

func InitDictTypeRouter() {
	RegistRoute(func(rgPblicRoute *gin.RouterGroup, rgAuthRoute *gin.RouterGroup) {
		dictTypeApi := api.DictTypeApi{}
		rgSingerAuth := rgAuthRoute.Group("dictType")
		{
			//获取所有字典类型
			rgSingerAuth.GET("/getAllType", dictTypeApi.GetAllDictType)
			//根据字典类型查询字典值
			rgSingerAuth.GET("/getByDictType", dictTypeApi.GetDictDataByDictType)
			//字典类型树
			rgSingerAuth.GET("/dictTypeTree", dictTypeApi.GetDictTypeTree)
			//根据id查询字典类型信息
			rgSingerAuth.GET("/getById", dictTypeApi.GetDictTypeById)
			//新增字典类型
			rgSingerAuth.POST("/add", dictTypeApi.AddDictType)
			//修改字典类型
			rgSingerAuth.PUT("/update", dictTypeApi.UpdateDictType)
			//根据id切片删除字典类型
			rgSingerAuth.DELETE("/deleteByIds", dictTypeApi.DropDictTypeByIds)
		}
		dictData := rgAuthRoute.Group("dictData")
		{
			//更新字典值
			dictData.PUT("/update", dictTypeApi.UpdateDictValueData)
			//新增字典值
			dictData.POST("/add", dictTypeApi.AddDictValueData)
			//根据id切片删除字典值
			dictData.DELETE("/dropByIds", dictTypeApi.DropDictValueDataByIds)
		}
	})
}
