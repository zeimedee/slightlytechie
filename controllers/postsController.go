package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zeimedee/slightlytechie/models"
	"github.com/zeimedee/slightlytechie/services"
)

func AddPost(c *gin.Context) {
	var body models.Post

	err := c.ShouldBindJSON(&body)
	if err != nil {
		panic(err)
	}
	data, err := services.AddPost(body.Title, body.Body)
	if err != nil {
		panic(err)
	}
	c.JSON(200, data)
}

func ReadPost(c *gin.Context) {
	id := c.Param("id")

	data, err := services.ReadPost(id)
	if err != nil {
		panic(err)
	}
	var userr []models.Post

	for data.Next() {
		var user models.Post
		if err := data.Scan(&user.Id, &user.Title, &user.Body, &user.Created_at, &user.Updated_at); err != nil {
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
	users := []models.Post{}

	defer data.Close()
	for data.Next() {
		var user models.Post
		err := data.Scan(&user.Id, &user.Title, &user.Body, &user.Created_at, &user.Updated_at)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println(user)
		users = append(users, user)
		if err != nil {
			panic(err)
		}

	}

	c.JSON(200, users)
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var body models.Post

	err := c.ShouldBindJSON(&body)
	if err != nil {
		panic(err)
	}
	data, err := services.UpdatePost(id, body.Title, body.Body)
	if err != nil {
		panic(err)
	}
	c.JSON(200, data)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	data, err := services.DeletePost(id)
	if err != nil {
		panic(err)
	}

	c.JSON(200, data)

}
