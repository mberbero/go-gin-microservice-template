package server

import (
	"fmt"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mberbero/go-microservice-template/app/routes"
	"github.com/mberbero/go-microservice-template/pkg/domains/post"
	"github.com/mberbero/go-microservice-template/pkg/infra/database"
)

func Run(host, port string) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(log gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s\"\n",
			log.ClientIP,
			log.TimeStamp.Format(time.UnixDate),
			log.Method,
			log.Path,
			log.Request.Proto,
			log.StatusCode,
			log.Latency,
		)
	}))
	r.Use(gin.Recovery())

	mongo := database.GetDB()

	postCollection := mongo.Db.Collection("posts")

	postRepo := post.NewRepo(postCollection)
	postService := post.NewService(postRepo)

	routes.PostRoutes(&r.RouterGroup, postService)
	r.Run(net.JoinHostPort(host, port))
}
