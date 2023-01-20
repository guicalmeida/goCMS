package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guicalmeida/goCMS/database"
)

func main() {
	database.ConnectToDB()
	fmt.Println("iniciando servidor REST com Go")
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":9254")
}
