package helper

import (
	"crypto/tls"
	"fmt"
	uuid "github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"os"
	"time"
)


//define jwt token
var jwtKey =  []byte("gin-gorm-oj-key")


//GetMd5 function
func GetMd5(str string) string {
	//sprintf := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	//return sprintf
	return str
}
type JwtToken struct {
	jwt.StandardClaims
	Name string `json:"name"`
	Identity string `json:"identity"`
	IsAdmin int64 `json:"is_admin"`
}


//define generattetoken function
func GenerateJwtToken(name, identity string,isAdmin int64) (string ,error){
	//define JwtToken
	token := JwtToken{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "exciseGo",
		},
		Name: name,
		IsAdmin: isAdmin,
		Identity: identity,
	}
	newWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	tokenString, err := newWithClaims.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}
	return tokenString,err
}

func Analysetoken(tokenString string) (*JwtToken, error) {
	token:=new(JwtToken)
	parseWithClaims, err := jwt.ParseWithClaims(tokenString, token, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		panic(err)
	}
	if !parseWithClaims.Valid {
		fmt.Println(token.Name)
		fmt.Println(token.Identity)
		return nil, err
	}
	return token, err
}

//send mail code function
func SendCode(toUseremail, code string) error {
	e := email.NewEmail()
	e.From = " Mason <zhangke_2021@126.com>"
	e.To = []string{toUseremail}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("验证码:<b>"+code+"</b>")
	//tsl port 587  common port 25
	return e.SendWithTLS("smtp.126.com:587", smtp.PlainAuth("", "zhangke_2021@126.com", "YWKWXQBYMWCDJIHV", "smtp.126.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.126.com"})
}

// GetUUID
func GetUUID()  string{
	newV4, _ := uuid.NewV4()
	s := newV4.String()
	return s
}

// GetRand code function
func GetRand(n int) string {
	var letter = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

//codeSave function
func CodeSave(code []byte) (string,error) {
	dirName := "code/" + GetUUID()
	path := dirName + "/runner.go"
	err := os.Mkdir(dirName, 0777)
	if err != nil {
		return "", err
	}
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	f.Write(code)
	defer f.Close()
	return path, nil
}

