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

// CreatePost godoc
// @Summary Create Post
// @Description Create Post
// @Tags posts
// @Accept  json
// @Produce  json
// @Param post body dtos.PostDTO true "Post Data"
// @Success 201 {object} entities.Post
//     Responses:
//       201: body:PositionResponseBody
// @Router /posts [post]
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

// PostGetByID godoc
// @Summary Post Get By ID
// @Description  Post Get By ID
// @Tags posts
// @Accept  json
// @Produce  json
// @Param id path string true "Post ID"
// @Success 200 {object} object
// @Router /posts/{id} [get]
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

// PostGetAll godoc
// @Summary Post List
// @Description Post List
// @Tags posts
// @Accept  json
// @Produce  json
// @Success 200 {object} object
// @Router /posts [get]
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

// PostUpdate godoc
// @Summary Post Update
// @Description Post Update
// @Tags posts
// @Accept  json
// @Produce  json
// @Param id path string true "Post ID"
// @Param post body dtos.PostDTO true "Post Data"
// @Success 200 {object} entities.Post
// @Router /posts/{id} [put]
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

// PostDelete godoc
// @Summary Post Delete
// @Description Post Delete
// @Tags posts
// @Accept  json
// @Produce  json
// @Param id path string true "Post ID"
// @Success 200 "Deleted"
// @Router /posts/{id} [delete]
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
