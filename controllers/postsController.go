package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zeimedee/slightlytechie/models"
	"github.com/zeimedee/slightlytechie/services"
)

func AddPost(c *gin.Context) {
	c.JSON(200, "add post")
}

func ReadPost(c *gin.Context) {
	id := c.Param("id")

	data, err := services.ReadPost(id)
	if err != nil {
		panic(err)
	}
	var userr []models.Post
	var user models.Post

	for data.Next() {
		if err := data.Scan(&user.Id, &user.Title, &user.Body, &user.Created_at); err != nil {
			log.Println(err.Error())
		}

		userr = append(userr, user)
		if err != nil {
			panic(err)
		}

		c.JSON(200, userr)
	}
}

func ReadAllPost(c *gin.Context) {
	data, err := services.ReadAllPost()
	if err != nil {
		panic(err)
	}

	var users []models.Post
	var user models.Post

	for data.Next() {
		if err := data.Scan(&user.Id, &user.Title, &user.Body, &user.Created_at); err != nil {
			log.Println(err.Error())
		}

		users = append(users, user)
		if err != nil {
			panic(err)
		}

		c.JSON(200, users)
	}
}

func UpdatePost(c *gin.Context) {
	c.JSON(200, "update post")
}

func DeletePost(c *gin.Context) {
	c.JSON(200, "delete post")
}
