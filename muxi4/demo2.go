package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  int    `json:"stock"` // 修改为int类型更合理
}

var books = make(map[string]Book)

func AddBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查ID是否已存在
	if _, exists := books[newBook.ID]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Book with this ID already exists"})
		return
	}

	// 添加新书
	books[newBook.ID] = newBook
	c.JSON(http.StatusCreated, newBook)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	// 检查图书是否存在
	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// 删除图书
	delete(books, id)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	// 检查图书是否存在
	if _, exists := books[id]; !exists {
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
	books[id] = updatedBook
	c.JSON(http.StatusOK, updatedBook)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	// 查找图书
	book, exists := books[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func GetAllBooks(c *gin.Context) {
	// 将map转换为slice
	var bookList []Book
	for _, book := range books {
		bookList = append(bookList, book)
	}

	c.JSON(http.StatusOK, bookList)
}

func main() {
	r := gin.Default()

	// 初始化一些示例数据
	books["1"] = Book{ID: "1", Title: "Go语言编程", Author: "许式伟", Stock: 10}
	books["2"] = Book{ID: "2", Title: "Clean Code", Author: "Robert C. Martin", Stock: 5}

	// 设置路由
	r.POST("/books", AddBook)          // 添加图书
	r.DELETE("/books/:id", DeleteBook) // 删除图书
	r.PUT("/books/:id", UpdateBook)    // 更新图书
	r.GET("/books/:id", GetBookByID)   // 查询单个图书
	r.GET("/books", GetAllBooks)       // 查询所有图书

	r.Run(":8080")
}
