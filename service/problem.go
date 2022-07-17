package service

import (
	//"encoding/json"
	"exciseGo/define"
	"exciseGo/helper"
	"exciseGo/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)
// GetDefaultProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Description GetProblemlist
// @Param page query int false "page"
// @Param size 	query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json:"{"code":200,"data":{},"msg":"ok"}"
// @Router /problem-list [get]
func GetDefaultProblemList(c *gin.Context)  {
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
	keyword := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")
	list:=make([]*models.ProblemBasic,0)
	tx:=models.GetDefaultProblemList(keyword, categoryIdentity)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error

	c.JSON(200, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list": list,
			"count": count,
		},
		"msg": "ok",
	})
}

// GetDefaultProblemList
// @Tags 公共方法
// @Summary 问题详情
// @Description GetProblemDetail
// @Param identity query string false "problem_identity"
// @Success 200 {string} json:"{"code":200,"data":{},"msg":"ok"}"
// @Router /problem-detail [get]
func GetProblemDetail(c *gin.Context)  {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题唯一标识不能为空",
		})
		return
	}
	//自定义必须指定
	data:=new(models.ProblemBasic)
	err := models.DB.Where("identity = ?", identity).Preload("ProblemCategory").Preload("ProblemCategory.CategoryBasic").First(&data).Error
	if err != nil {
		//提示信息
		if err== gorm.ErrRecordNotFound {
			//返回信息
			c.JSON(200, gin.H{
				"code": 400,
				"data": nil,
				"msg": "problem not found",
			})
			return
		}

		c.JSON(200, gin.H{
				"code": -1,
				"msg": "problem datail not found"+err.Error(),
		})
			return

	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}

//AddProblem function
// @Tags
// @Summary 添加问题
// @Description AddProblem
// @Param authorization header string true "authorization"
// @Param title formData string false "title"
// @Param content formData string false "content"
// @Param max_mem formData int false "max_mem"
// @Param max_runtime formData int false "max_runtime"
// @Param category_ids formData array false "category_ids"
// @Param test_cases formData array false "test_cases"
// @Success 200 {string} json:"{"code":200,"data":{},"msg":"ok"}"
// @Router /problem-add [post]
func AddProblem(c *gin.Context)  {
	title := c.PostForm("title")
	content:= c.PostForm("content")
	maxRuntime,_:= strconv.Atoi(c.PostForm("max_runtime"))
	maxMemory,_:= strconv.Atoi(c.PostForm("max_memory"))
	Categoryids:= c.PostForm("category_ids")
	testCases := c.PostForm("test_cases")
	if title == "" || content == ""  || testCases == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}

	identity := helper.GetUUID()
	data := &models.ProblemBasic{
		Identity:   identity,
		Title:      title,
		Content:    content,
		MaxRunTime: maxRuntime,
		MaxMem:     maxMemory,
		PassNum:   0,
		SubmitNum: 100,
	}
	//插入数据
	fmt.Printf("data=%+v\n", data)

	//处理分类
	categoryBasics := make([]*models.ProblemCategory, 0)
	for _, id := range strings.Split(Categoryids, ",") {
		atoi, _ := strconv.Atoi(id)
		categoryBasics = append(categoryBasics, &models.ProblemCategory{
			ProblemId: int64(data.ID),
			CategoryId: int64(atoi),
		})
	}
	data.ProblemCategory = categoryBasics

	//处理测试用例
	//testCaseBasics := make([]*models.TestCase, 0)
	//range
	//for _,testcase := range strings.Split(testCases, ",") {
	////for _,testcase := range  testCases {
	//	//{input:"1 2\n","output":"3\n"}
	//	caseMap := make(map[string]string)
	//	err := json.Unmarshal([]byte(testcase), &caseMap)
	//	if err != nil {
	//		c.JSON(http.StatusOK, gin.H{
	//			"code": -1,
	//			"msg":  "测试用例格式错误",
	//		})
	//		return
	//	}
	//
	//	if _,ok := caseMap["input"]; !ok {
	//		c.JSON(http.StatusOK, gin.H{
	//			"code": -1,
	//			"msg":  "测试用例格式错误",
	//		})
	//		return
	//	}
	//	if _,ok := caseMap["output"]; !ok {
	//		c.JSON(http.StatusOK, gin.H{
	//			"code": -1,
	//			"msg":  "测试用例格式错误",
	//		})
	//		return
	//	}
	//
	//	testCaseBasic := &models.TestCase{
	//		Identity:        helper.GetUUID(),
	//		ProblemIdentity: identity,
	//		Input: caseMap["input"],
	//		Output: caseMap["output"],
	//	}
	//	fmt.Printf("testCaseBasic=%+v\n", testCaseBasic)
	//	testCaseBasics = append(testCaseBasics, testCaseBasic)
	//}

	fmt.Println("identity=:",identity)
	// 创建问题
	err:= models.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Problem Create Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"identity": data.Identity,
		},
	})
}

// ProblemCreate
// @Tags 管理员私有方法
// @Summary 问题创建
// @Accept json
// @Param authorization header string true "authorization"
// @Param data body define.ProblemBasic true "ProblemBasic"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-create [post]
func ProblemCreate(c *gin.Context) {
	fmt.Println("ProblemCreate")
	in := new(define.ProblemBasic)
	err := c.ShouldBindJSON(in)
	if err != nil {
		println("[JsonBind Error] : ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
		return
	}

	if in.Title == "" || in.Content == "" || len(in.ProblemCategories) == 0 || len(in.TestCases) == 0 || in.MaxRuntime == 0 || in.MaxMem == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}

	identity := helper.GetUUID()
	data := &models.ProblemBasic{
		DeletedAt: time.Now(),
		Identity:   identity,
		Title:      in.Title,
		Content:    in.Content,
		MaxRunTime: in.MaxRuntime,
		MaxMem:     in.MaxMem,
	}
	fmt.Printf("data=%+v\n", data)
	// 处理分类
	categoryBasics := make([]*models.ProblemCategory, 0)
	for _, id := range in.ProblemCategories {
		categoryBasics = append(categoryBasics, &models.ProblemCategory{
			ProblemId:  data.ID,
			CategoryId: int64(uint(id)),
		})
	}
	data.ProblemCategory = categoryBasics
	// 处理测试用例
	testCaseBasics := make([]*models.TestCase, 0)
	for _, v := range in.TestCases {
		// 举个例子 {"input":"1 2\n","output":"3\n"}
		testCaseBasic := &models.TestCase{
			Identity:        helper.GetUUID(),
			ProblemIdentity: identity,
			Input:           v.Input,
			Output:          v.Output,
		}
		testCaseBasics = append(testCaseBasics, testCaseBasic)
	}
	data.TestCase = testCaseBasics

	// 创建问题
	err = models.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Problem Create Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"identity": data.Identity,
		},
	})
}
