package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Justyer/JekoServer/http/router/ver"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	r.Static("/static", "http/static")
	r.LoadHTMLGlob("http/template/*")

	r.GET("/canvas", func(c *gin.Context) {
		c.HTML(http.StatusOK, "canvas.html", gin.H{})
	})
	r.GET("/flex", func(c *gin.Context) {
		c.HTML(http.StatusOK, "flex.html", gin.H{})
	})
	r.GET("/tool", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tool.html", gin.H{})
	})
	r.MaxMultipartMemory = 2 << 30 // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// single file
		file, err := c.FormFile("file")
		if err != nil {
			log.Println("err", err, file)
		}
		log.Println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, file.Filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	ver.APIV1(r.Group("v1"))
}
