package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/leong-y/go-gin-example/pkg/setting"
)

// 算出当前页数的偏移量
/* 	 例如：
\	   分页数为10，当前为第2页则
\	   result = (2-1)*10
\	   即result = 10，意为从第10条开始查找
*/
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}
