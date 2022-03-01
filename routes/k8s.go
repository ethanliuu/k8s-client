package routes

import (
	"github.com/ethanliuuu/k8s-client/controller/k8s"
	"github.com/gin-gonic/gin"
)

func InitContainerRouter(r *gin.RouterGroup) {
	k8sgroup := r.Group("k8s")
	{
		k8sgroup.GET("namespace", k8s.GetNamespaceList)

		k8sgroup.POST("deployment/restart", k8s.RestartDeploymentController)
		k8sgroup.POST("deployment/scale", k8s.ScaleDeploymentController)

	}
}
