package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const Secret = "114514"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Profile struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 生成token
func GenerateToken(username string, password string) (string, error) {
	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(time.Hour * 24).Unix(),
			Issuer:    "book-manager"},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Secret))
}

// 处理登录请求
func AuthLoginHandler(c *gin.Context) {
	var user Profile
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "请求参数错误",
			"error": err.Error(),
		})
		return
	}
	if user.Username != "admin" || user.Password != "admin123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "用户名或密码错误",
		})
		return
	}
	token, err := GenerateToken(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"msg":   "生成Token失败",
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "登录成功",
		"token": token,
	})
}

// 验证解析token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(Secret), nil // 转换为字节切片
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Authorization header 不能为空",
			})
			c.Abort()
			return
		}
		// 检查Bearer格式
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.JSON(401, gin.H{"error": "Authorization header格式必须是Bearer <token>"})
			c.Abort()
			return
		}
		tokenString := authHeader[7:]
		claims, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  401,
				"msg":   "无效token",
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}
