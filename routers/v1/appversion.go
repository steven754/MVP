package v1

import (
	"MVP/pkg/e"
	"github.com/gin-gonic/gin"
)

//app更新接口
func GetAppVersionTest(c *gin.Context) {

	c.JSON(e.SUCCESS, gin.H{
		"Code": e.SUCCESS,
		"Msg":  e.GetMsg(e.SUCCESS),
		"Data": "返回数据成功",
	})
}
