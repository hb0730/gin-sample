package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
	ID   uint64
	Name string
}

func main() {
	users := []User{{ID: 1, Name: "小红"}, {ID: 2, Name: "小绿"}}
	r := gin.Default()
	// 新增用户
	r.POST("/add", func(c *gin.Context) {
		var json User
		c.BindJSON(&json)
		users = append(users, json)
	})
	// 获取用户集
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// 获取单个用户
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.GetInt64("id")
		i := uint64(id)
		var user User
		for _, v := range users {
			if v.ID == i {
				user = v
				break
			}
		}
		c.JSON(http.StatusOK, user)
	})

	// 更新用户
	r.PUT("/update/:id", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		var json User
		var index int
		for i, user := range users {
			if user.ID == id {
				index = i
				break
			}
		}
		c.BindJSON(&json)
		users[index] = json
		if json != (User{}) {
			c.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "更新成功",
				"data":    users,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "更新失败",
			})
		}
	})

	// 删除用户
	r.DELETE("/delete/:id", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		for i, v := range users {
			if v.ID == id {
				users = append(users[:i], users[i+1:]...)
			} else {
				i++
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"message": "删除成功",
			"data":    users,
		})

	})

	r.Run(":80")
}
