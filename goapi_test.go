/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2024-11-05 16:56:08
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2024-11-05 17:00:07
 * @FilePath: /undefined/home/ben/Documents/go/stud-api/goapi_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// 使用gin框架进行API开发
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}

//运行程序，在浏览器中访问http://localhost:8080/ping，可以看到返回的结果：	{"message":"pong"}

//查询商品信息	GET /goods/:id
//创建商品	POST /goods
//更新商品	PUT /goods/:id
//删除商品	DELETE /goods/:id
