package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"ls.com/kit/pkg/service"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// 显示表单页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 定义处理函数，用于处理表单提交
	r.POST("/generate/code", func(c *gin.Context) {
		// 获取表单提交的数据
		funcName := c.PostForm("func_name")
		codeType := c.PostForm("code_type")
		log.Println("funcName: ", funcName, "codeType: ", codeType)
		// 调用处理函数
		code := service.Generate(funcName, codeType)
		// 将结果传递给HTML模板
		data := gin.H{
			"func_name": funcName,
			"type":      codeType,
			"code":      code,
		}
		// 显示结果页面
		c.HTML(http.StatusOK, "index.html", data)
	})

	_ = r.Run(":8989")
}
