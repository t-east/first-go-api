package main

import (
	"github.com/t-east/first-go-api/article"

	"github.com/t-east/first-go-api/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	article := article.New()
	r := gin.Default()
	r.GET("/article", handler.ArticlesGet(article))
	r.POST("/article", handler.ArticlePost(article))

	r.Run() // listen and serve on 0.0.0.0:8080

}
