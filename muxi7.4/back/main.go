package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
	"errors"
	"github.com/dgrijalva/jwt-go"
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
// Book 图书信息
// @Description 图书信息
type Book struct {
	ID     string `json:"id" gorm:"primary_key" binding:"required,alphanum"`
	Title  string `json:"title" binding:"required,min=2"`
	Author string `json:"author" binding:"required"`
	Stock  int    `json:"stock" binding:"gte=0"`
  }

var (
	db    *gorm.DB
	books = make(map[string]Book)
	users = map[string]string{"admin": "admin123"}
)

func initDB() {
    // 从环境变量读取DSN，默认使用Docker Compose的服务名
    dsn := os.Getenv("MYSQL_DSN")
    if dsn == "" {
        dsn = "root:20041017@tcp(127.0.0.1:3306)/Book_Management?charset=utf8mb4&parseTime=True&loc=Local"
    }

    // 添加重试逻辑
    var err error
    maxAttempts := 5
    for i := 1; i <= maxAttempts; i++ {
        db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err == nil {
            break
        }
        log.Printf("数据库连接尝试 %d/%d 失败: %v", i, maxAttempts, err)
        time.Sleep(5 * time.Second)
    }

    if err != nil {
        log.Fatal("无法连接数据库:", err) // 改用log.Fatal避免无限重启
    }

    if err := db.AutoMigrate(&Book{}); err != nil {
        log.Fatal("数据库迁移失败:", err)
    }
}

// AddBook 添加新书
// @Description 添加新书
// @Description 添加一本新书到图书库
// @Tags 图书管理
// @Accept json
// @Produce json
// @Parm book body Book true "图书信息"
// @Success 201 {object} Book "成功添加图书"
// @Failure 400 {object} object "请求参数错误"
// @Failure 409 {object} object "图书ID已存在"
// @Router /books [post]
func AddBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查ID是否已存在
	var existingBook Book
	result := db.First(&existingBook, "id = ?", newBook.ID)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Book with this ID already exists"})
		return
	}

	// 添加新书
	if err := db.Create(&newBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newBook)
}

// DeleteBook 删除图书
// @Description 删除图书
// @Description 根据id删除图书
// @Tags 图书管理
// @Accept json
// @Produce json
// @Parm book body Book true "图书ID"
// @Success 200 {object} Book "成功删除图书"
// @Failure 404 {object} object "未找到对应图书"
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	// 检查图书是否存在
	var book Book
	result := db.First(&book, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// 删除图书
	if err := db.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// UpdateBook 更新图书
// @Description 更新图书
// @Description 根据ID更新对应图书
// @Tags 图书管理
// @Accept json
// @Produce json
// @Parm book body Book true "图书ID"
// @Success 200 {object} Book "成功更新图书"
// @Failure 400 {object} object "请求参数错误"
// @Failure 404 {object} object "未找到对应图书"
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	// 检查图书是否存在
	var book Book
	result := db.First(&book, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新图书信息
	updatedBook.ID = id // 确保ID不被修改
	if err := db.Save(&updatedBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}

// GetBookByID 获取图书
// @Description 获取
// @Description 根据ID获取图书
// @Tags 图书管理
// @Produce json
// @Parm book body Book true "图书ID"
// @Success 201 {object} Book "成功获取图书"
// @Failure 404 {object} object "为找到对应图书"
// @Router /books/{id} [get]
func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	// 查找图书
	var book Book
	result := db.First(&book, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// GetAllBooks 获取所有图书
// @Description 获取所有图书
// @Description 获取所有图书信息
// @Tags 图书管理
// @Produce json
// @Success 200 {array} Book "图书列表"
// @Router /books [get]
func GetAllBooks(c *gin.Context) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// @title 图书管理API
// @version 1.0
// @description 这是一个简单的图书管理系统API
// @contact.name API支持
// @contact.email support@example.com
// @host localhost:8080
// @BasePath /
func main() {
	initDB()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/login", AuthLoginHandler)
	// 注册用户
r.POST("/register", func(c *gin.Context) {
	var user Profile
	if err := c.ShouldBindJSON(&user); err != nil {
	  c.JSON(400, gin.H{"error": err.Error()})
	  return
	}
	users[user.Username] = user.Password // 实际应加密存储
	c.JSON(201, gin.H{"message": "用户注册成功"})
  })
	// 设置路由
	authGroup := r.Group("/")
	authGroup.Use(AuthMiddleware())
	{
		authGroup.POST("/books", AddBook)          // 添加图书
		authGroup.DELETE("/books/:id", DeleteBook) // 删除图书
		authGroup.PUT("/books/:id", UpdateBook)    // 更新图书
	}

	r.GET("/books/:id", GetBookByID) // 查询单个图书
	r.GET("/books", GetAllBooks)     // 查询所有图书

	if err := r.Run(":8080"); err != nil {
        log.Fatal("服务器启动失败:", err)
    }

}
