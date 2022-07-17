package router

import (
	_ "exciseGo/docs"
	"exciseGo/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func  Router() *gin.Engine {
	r := gin.Default()
	//docs.SwaggerInfo.BasePath = "/api/v1"
	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//test
	r.GET("/ping", service.Ping)

	//problem
	r.GET("/problem-list", service.GetDefaultProblemList)
	r.GET("/problem-detail", service.GetProblemDetail)

	//submit
	r.GET("/submit-list", service.GetSubmitList)

	// user add /login
	r.POST("/login", service.Login)

	r.POST("/send-code", service.SendCode)
	//register
	r.POST("/register", service.Register)

	//user rank list
	r.GET("/user-rank-list", service.GetUserRankList)

	//admin private ways
	//test token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgxMzc4MDcsImlzcyI6ImV4Y2lzZUdvIiwibmFtZSI6IjdjNmJmMGU2LTM2YzUtNDM4YS1iOTNmLTY0YWM4ZTUzODcxNyIsImlkZW50aXR5IjoiemhhbmciLCJpc19hZG1pbiI6MH0.NenDydBQbGbgGmhenk4bKFIvgayIKgSi_kbhiTK_dkA
	r.POST("/problem-create", service.ProblemCreate)

	return r
}