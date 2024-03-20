package controllers

import (
	"gocrud/initializers"
	"gocrud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//ambil data
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Buat post

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//ambil posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//respon post
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//get id from url
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	//respon post
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//ambil id dari url
	id := c.Param("id")

	//ambil data
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//liat post update
	var post models.Post
	initializers.DB.First(&post, id)

	//update post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//respon
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)
	c.Status(200)

}
