package router

import (
	"cmlabs-backend-crawler-freelance-test/application"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const port = "30001"

type IRouter interface {
	Run()
}

type Router struct {
	Engine      *gin.Engine
	Application application.IApplication
	Port        string
}

func NewRouter(app application.IApplication) IRouter {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"Origin, X-Requested-With, Content-Type, Accept, Authorization, Access-Control-Allow-Headers, Accept-Encoding, X-CSRF-Token"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	return &Router{
		Engine:      router,
		Application: app,
		Port:        port,
	}
}

func (r *Router) registerEndpoint() {
	r.Engine.POST("/api/crawl", r.Application.Crawl)
}

func (r *Router) Run() {
	r.registerEndpoint()
	r.Engine.Run(r.exposePort())
}

func (r *Router) exposePort() string {
	return fmt.Sprintf(":%s", r.Port)
}
