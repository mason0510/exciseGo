package service

import "github.com/gin-gonic/gin"

// Ping ...
// @Tags 公共方法
// @Summary 获取服务器连通
// @Description Ping
// @Param ping query string false "query"
// @Success 200 {string} json:"{"code":200,"data":{},"msg":"ok"}"
// @Router /problem-list [get]
func Ping(c *gin.Context)  {
	 c.JSON(200, gin.H{
		"message": "pong",
	})
}
