package middlewires

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goAdminStudy/config"
	"goAdminStudy/modules"
	"log"
	"net/http"
	"time"
)

// 自定义jwt claims
type Claims struct {
	Id       int32  `json:"id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// 加密获取token
// @Param user 用户
// @Return string, error
func GenerateToken(user modules.User) (string, error) {

	now := time.Now()
	// 过期时间
	expireAt := now.Add(60 * time.Minute)

	claims := Claims{
		Id:       user.Id,
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt.Unix(),
			Issuer:    "goAdmin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成签名token，
	return token.SignedString([]byte(config.JwtConfig.GetSecret()))

}

// 解析token
func ParseToken(token string) (*Claims, error) {

	claims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtConfig.GetSecret()), nil
	})

	if claims != nil {
		if c, ok := claims.Claims.(*Claims); ok && claims.Valid {
			return c, nil
		}
	}

	return nil, err

}

// jwt验证
func JwtMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		log.Printf("Token from header: %s \n", token)

		url := c.Request.URL
		log.Printf("Path: %s \t Rawpath: %s \n", url.Path, url.RawPath)
		code := http.StatusOK
		message := ""
		if token == "" {
			code = http.StatusUnauthorized
			message = "Token is empty."
		} else {
			if parseToken, err := ParseToken(token); err != nil {
				code = http.StatusUnauthorized
				message = "Token is invalid."
			} else if time.Now().Unix() > parseToken.ExpiresAt {
				code = http.StatusInternalServerError
				message = "Token is expired."
			}
		}

		if code != http.StatusOK {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
				"date":    time.Now().Format("2006-01-02 15:04:05"),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
