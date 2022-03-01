package main

import (
	"github.com/ethanliuuu/k8s-client/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	InitServer()
}

func InitServer() {
	r := gin.Default()
	Group := r.Group("/api/v1")
	{
		Group.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
		//Group.GET("namespace",k8s.GetNamespaceList)

		//Group.POST("/deployment/restart",k8s.RestartDeploymentController)
		routes.InitContainerRouter(Group)
	}
	r.Run(":8080")
}
