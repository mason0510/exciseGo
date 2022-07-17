package service

import (
	"exciseGo/define"
	"exciseGo/models"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	"strconv"
)

// GetSubmitList
// @Tags 公共方法
// @Summary 提交列表
// @Description GetSubmitList
// @Param page query int false "page"
// @Param size 	query int false "size"
// @Param problem_identity query string false "problem_identity"
// @Param user_identity query string false "user_identity"
// @Param status query string false "status"
// @Success 200 {string} json:"{"code":200,"data":{},"msg":"ok"}"
// @Router /submit-list [get]
func GetSubmitList(c *gin.Context) {
	//record
	size,err:=strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	if err!=nil{
		log.Error("GetProblemlist size error",err)
		return
	}

	//string转换int
	page,err:=strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err!=nil{
		log.Error("GetProblemlist page error",err)
		return
	}
	//-->1--->0
	page=(page-1)*size
	var count int64

	problemIdentity := c.Query("problem_identity")
	userIdentity := c.Query("user_identity")
	status,err:=strconv.Atoi(c.DefaultQuery("status", define.DefaultStatus))
	if err!=nil{
		log.Error("GetSubmitList status error",err)
		return
	}
	list:=make([]models.SubmitBasic,0)

	tx:=models.GetDefaultSubmitList(problemIdentity, userIdentity,status)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err!=nil{
		c.JSON(200, gin.H{
			"code": -1,
			"msg": "GetSubmitList status error"+err.Error(),
		})
		return
	}

	//right condition
	c.JSON(200, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list": list,
			"count": count,
		},
		"msg": "ok",
	})
}