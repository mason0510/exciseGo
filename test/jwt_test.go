package test

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

//define byte key
var jwtKey =  []byte("gin-gorm-oj-key")

//define struct for jwt token
type JwtToken struct {
	jwt.StandardClaims
	Name string `json:"name"`
	Identity string `json:"identity"`
}

//define generate jwt token function
func TestGenerateJwtToken(t *testing.T) {
	//define JwtToken
	token := JwtToken{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "exciseGo",
		},
		Name: "username",
		Identity: "admin",
	}
	newWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	tokenString, err := newWithClaims.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenString)
}

//parsing token
func TestAnalysetoken(t *testing.T){
	tokenString:="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6InVzZXJfMSIsIm5hbWUiOiJHZXQifQ.4inO9HZINmKFYO9qEF2SYYPHk0GuuA-qUdwIhUa8USE"
	token:=new(JwtToken)
	parseWithClaims, err := jwt.ParseWithClaims(tokenString, token, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		panic(err)
	}
	if parseWithClaims.Valid {
		fmt.Println(token.Name)
		fmt.Println(token.Identity)
	}
}
