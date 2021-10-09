package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mberbero/go-microservice-template/pkg/domains/post"
	"github.com/mberbero/go-microservice-template/pkg/dtos"
	"github.com/mberbero/go-microservice-template/pkg/entities"
)

func PostRoutes(r *gin.RouterGroup, s post.Service) {
	r.GET("/posts", PostGetAll(s))
	r.POST("/posts", PostCreate(s))
	r.GET("/posts/:id", PostGetByID(s))
	r.PUT("/posts/:id", PostUpdate(s))
	r.DELETE("/posts/:id", PostDelete(s))
}

func PostCreate(s post.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload *dtos.PostDTO
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		post := &entities.Post{
			Title:   payload.Title,
			Content: payload.Content,
			Tags:    payload.Tags,
		}
		if err := s.Create(post); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, post)
	}
}

func PostGetByID(s post.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		post, err := s.Get(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, post)
	}
}

func PostGetAll(s post.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, _ := strconv.Atoi(c.Query("page"))
		perPage, _ := strconv.Atoi(c.Query("per_page"))
		if page == 0 {
			page = 1
		}
		if perPage == 0 {
			perPage = 10
		}

		posts, err := s.GetAll(int64(page), int64(perPage))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, posts)
	}
}

func PostUpdate(s post.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var post *dtos.PostDTO
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := s.Update(id, post); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, post)
	}
}

func PostDelete(s post.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := s.Delete(id); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Post deleted"})
	}
}
