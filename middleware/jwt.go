package middleware

import (
	"errors"
	"fmt"
	"ginblog/utils/errormessage"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)
//生成token

//
//var MySecret = []byte("夏天夏天悄悄过去")
var  code int
var Jwtkey =[]byte("7e31q1ew")
type Myciaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
func GenToken(username string) (string, int) {
	// 创建一个我们自己的声明
	expireTime:=time.Now().Add(10 *time.Hour)
	SetClaims :=Myciaims{
		Username:      username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:expireTime.Unix(),
			Issuer: "ginblog",
		},

	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	 newtoken,err :=token.SignedString(Jwtkey)
	 if err !=nil{
	 	return "",errormessage.ERROR
	 }
	 return newtoken,errormessage.SUCCESS


}

//验证token
func ParseToken(tokenString string) (*Myciaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Myciaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Jwtkey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Myciaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func Jwttoken()gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": errormessage.ERROR_TOKEN_EXISTS,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": errormessage.ERROR_TOKEN_TYPR_ERONG,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": errormessage.ERROR_TOKEN_WRONG,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		if time.Now().Unix() >mc.ExpiresAt{
			fmt.Println("token 已经过期了")
			c.Abort()
		}
		c.Set("username", mc.Username)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}


	}
