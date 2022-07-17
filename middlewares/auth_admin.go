package middlewares

import (
	"exciseGo/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

//define AuthAdminCheck fuction
func AuthAdminCheck() gin.HandlerFunc {
	//return fuc
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		jwtToken, err := helper.Analysetoken(header)
		if  err!= nil {
			c.JSON(401, gin.H{
				"message": "Unauthorized :Admin Token Error",
				"code": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
		if jwtToken.IsAdmin != 1 {//is not admin
			//return error
			c.JSON(401, gin.H{
				"message": "Unauthorized :Is Not Admin ",
				"code": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
	}
}