package service

import (
	"exciseGo/define"
	"exciseGo/helper"
	"exciseGo/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

//Login
// @Tags 公共方法
// @Summary 登录
// @Description Login
// @Param name formData string false "name"
// @Param password formData string false "password"
// @Success 200 {string} json:"{"code":200,"data":{},"msg":"ok"}"
// @Router /login [post]
func Login(c *gin.Context) {
	//record
	name := c.PostForm("name")
	password := c.PostForm("password")
	if name == "" || password == "" {
		c.JSON(200, gin.H{
			"code": -1,
			"msg": "name or password is empty",
		})
		return
	}
	//add md5
	password = helper.GetMd5(password)
	data := new(models.UserBasic)
	println("name",name)
	println("password",password)
	err := models.DB.Where("name = ? and password = ?", name, password).First(data).Error
	if err != nil {
		if err==gorm.ErrRecordNotFound{
			c.JSON(200, gin.H{
				"code": -1,
				"msg": "name or password is wrong",
			})
			return
		}
		c.JSON(200, gin.H{
			"code": -1,
			"msg": "login error"+err.Error(),
		})
	}
	generateJwtToken,err:= helper.GenerateJwtToken(data.Identity, name,data.IsAdmin)
	if err!=nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg": "generate jwt token error"+err.Error(),
		})
	}

	//right condition
	c.JSON(200, gin.H{
		"code": 0,
		"data": map[string]string{
			"token":generateJwtToken,
		},
	})
}

//SendCode
// @Tags 公共方法
// @Summary 发送验证码
// @Description SendCode
// @Param email formData string true "email"
// @Success 200 {string} json:"{"code":200,"data":{},"msg":"ok"}"
// @Router /send-code [post]
func SendCode(c *gin.Context) {
	//get email
	postForm := c.PostForm("email")
	fmt.Println("postForm",postForm)
	if postForm == "" {
		//return c.JSON
		c.JSON(200, gin.H{
			"code": -1,
			"msg": "email is empty",
		})
		return
	}

	code:=helper.GetRand(6)
	models.RDB.Set(c, postForm, code, time.Second*30000)
	err := helper.SendCode(postForm, code)
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg": "send code error"+err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg": "send code success",
	})
}

//Register function
// @Tags 公共方法
// @Summary 注册
// @Description Register
// @Param mail formData string true "email"
// @Param name formData string true "name"
// @Param code formData string true "code"
// @Param password formData string true "password"
// @Param phone formData string false "phone"
// @Success 200 {string} json:"{"code":200,"data":{},"msg":"ok"}"
// @Router /register [post]
func Register(c *gin.Context) {
	mail := c.PostForm("mail")
	userCode := c.PostForm("code")
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	fmt.Println("mail",mail)
	fmt.Println("userCode",userCode)
	fmt.Println("name",name)
	fmt.Println("password",password)
	fmt.Println("phone",phone)
	if mail == "" || userCode == "" || name == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}

	//check code is or not right
	result, err := models.RDB.Get(c,mail).Result()
	if result !=userCode {
		//code is not right
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确",
		})
		return
	}

	//mail must be unique
	var cnt_name int64
	err = models.DB.Where("name = ?", name).Model(new(models.UserBasic)).Count(&cnt_name).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get user ERROR"+err.Error(),
		})
		return
	}

	//mail must be unique
	var cnt int64
	err = models.DB.Where("mail = ?", mail).Model(new(models.UserBasic)).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get user ERROR"+err.Error(),
		})
		return
	}

	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱已存在",
		})
		return
	}

	//insert data
	user := models.UserBasic{
		Identity: helper.GetUUID(),
		Name: name,
		PassWord: password,
		Phone: phone,
		Mail: mail,
		DeletedAt: time.Now(),
	}
	fmt.Println("user",user.DeletedAt)
	err = models.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "insert user ERROR"+err.Error(),
		})
		return
	}

	//produce token
	generateJwtToken,err:= helper.GenerateJwtToken(user.Identity, name,user.IsAdmin)
	if err!=nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "generate jwt token error"+err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]string{
			"token":generateJwtToken,
		},
	})
}
//
////new GetUserRankList function
//// @Tags 公共方法
//// @Summary 获取排行榜
//// @Param page formData int false "page"
//// @Param pageSize formData int false "pageSize"
//// @Description GetUserRankList
//// @Success 200 {string} json:"{"code":200,"data":{},"msg":"ok"}"
//// @Router /user-rank-list [get]
//func GetUserRankList(c *gin.Context) {
//	//same code
//	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
//	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
//	if err!= nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": -1,
//			"msg":  "参数不正确",
//		})
//		return
//	}
//	page = (page-1)*size
//	var count int64
//	//define the model and query
//	list:=make([]*models.UserBasic,0)
//	fmt.Println("page",page)
//	fmt.Println("size",size)
//	//err = models.DB.Model(&models.UserBasic{}).Count(&count).Find(&list).Error
//	err = models.DB.Model(&models.UserBasic{}).Select("name", "phone").Find(&list).Error
//	// SELECT `id`, `name` FROM `users` LIMIT 10
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": -1,
//			"msg":  "Get Rank List ERROR"+err.Error(),
//		})
//		return
//	}
//	fmt.Println("list",list)
//	c.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"data": map[string]interface{}{
//			"list": list,
//			"count": count,
//		},
//	})
//}


// GetRankList
// @Tags 公共方法
// @Summary 用户排行榜
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user-rank-list [get]
func GetUserRankList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv Error:", err)
		return
	}
	page = (page - 1) * size

	var count int64
	list := make([]*models.UserBasic, 0)
	err = models.DB.Model(new(models.UserBasic)).Count(&count).Order("pass_num DESC, submit_num ASC").
		Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Rank List Error:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}


