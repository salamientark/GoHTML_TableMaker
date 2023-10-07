package main

import (
	"fmt"
	"net/http"
	"encoding/csv"
	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.LoadHTMLGlob("templates/*")
  r.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
		"new_title" : "title",
	})
  })
	r.POST("/namek", func(c *gin.Context) {
		file, err := c.FormFile("file_bom")
		if err != nil {
			fmt.Println(err)
			return
		}
		file_desc, err := file.Open()

		if err != nil {
			fmt.Println(err)
			return
		}

		reader := csv.NewReader(file_desc)
		reader.Comma = ';'
    	Records, _ := reader.ReadAll()
    	fmt.Println(Records)

		// fmt.Println(avatar)
		
		c.HTML(http.StatusOK, "index.html", nil)
		c.HTML(200, "index.html", gin.H{
			"Records" : Records,
		})
    //c.HTML(http.StatusOK, "index.html", nil)
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}