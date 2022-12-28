package main 

import (
	"net/http"
	"github.com/gin-gonic/gin" 
)

type Blog struct {
	ID int `json: "id"`
	Title string `json: "title"`
	Description string `json: "description"`
	Category string `json: "category"` 
}

var blogs []Blog 

func getIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func getBlog(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, blogs) 
}

func postBlog(c *gin.Context) {
	var newBlog Blog 

	if err := c.BindJSON(&newBlog); err != nil {
		return 
	}

	blogs = append(blogs, newBlog) 
	c.IndentedJSON(http.StatusCreated, newBlog) 
}

func main() {
	r := gin.Default()
	r.GET("/", getIndex) 
	r.GET("/api/v1/blogs", getBlog)  
	r.POST("/api/v1/blogs", postBlog) 
	r.Run(":8000") 
}