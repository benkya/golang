/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2024-11-04 16:35:34
 * @LastEditors: ben ben.chaung@hotmail.com
 * @LastEditTime: 2024-11-06 15:26:44
 * @FilePath: /undefined/home/ben/Documents/go/test02/goapi.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义一个学生结构体
type Student struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Deparment string `json:"deparment"`
	Level     string `json:"level"`
}

// 学生数据库
var students = []Student{
	{"1001", "Tom", "Computer Science", "Undergraduate"},
	{"1002", "Jerry", "Mathematics", "Undergraduate"},
	{"1003", "Alice", "Physics", "Undergraduate"},
}

func main() {
	r := gin.Default()
	r.GET("/", WelcomeMessage)
	r.POST("/createstudents", CreateStudent)
	r.POST("/handleform", handleForm) // 处理form表单数据
	r.GET("/students", ShowAllStudents)
	r.DELETE("/deletestudents", DeleteStudent)
	r.PUT("/updatestudents", UpdateStudent)
	r.Run()
}

func WelcomeMessage(c *gin.Context) {
	message := map[string]string{"message": "Welcome to My 1stGo API"}
	c.JSON(http.StatusOK, message)
}

// 新建一个学生
func CreateStudent(c *gin.Context) { //指针接收
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	students = append(students, student)
	c.JSON(http.StatusCreated, student)
}

// 根据id删除学生
// 注意: 这里的删除操作并没有真正删除, 而是将学生信息标记为删除, 实际删除操作需要在业务逻辑层实现
// 这里只是简单的实现了删除功能, 实际业务中需要考虑到数据完整性, 权限等问题, 并提供相应的接口
func DeleteStudent(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, s := range students {
		if s.ID == student.ID {
			students[i].Level = "Deleted"
			break
		}
	}
	c.JSON(http.StatusOK, student)
}

// 根据id修改学生信息
func UpdateStudent(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, s := range students {
		if s.ID == student.ID {
			students[i] = student
			break
		}
	}
	c.JSON(http.StatusOK, student)
}

// 显示所有学生
func ShowAllStudents(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

// 接受form表单数据
func handleForm(c *gin.Context) {

	name := c.PostForm("name") //	获取表单数据
	Deparment := c.PostForm("Deparment")

	var student Student // 定义一个学生结构体

	if c.ShouldBind(&student) == nil {
		// do something with student
		c.JSON(http.StatusOK, gin.H{
			"message":   "Student received",
			"name":      name,
			"Deparment": Deparment,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student data"})
	}
}
