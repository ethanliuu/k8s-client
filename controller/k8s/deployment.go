package k8s

import (
	"github.com/ethanliuuu/k8s-client/controller"
	"github.com/ethanliuuu/k8s-client/controller/response"
	"github.com/ethanliuuu/k8s-client/models/k8s"
	"github.com/ethanliuuu/k8s-client/pkg/k8s/client"
	"github.com/ethanliuuu/k8s-client/pkg/k8s/deployment"
	"github.com/gin-gonic/gin"
)

func RestartDeploymentController(c *gin.Context) {
	clientset, err := client.ClientSet()
	if err != nil {
		response.FailHasMsg(response.InternalServerError, err.Error(), c)
		return
	}

	var rd k8s.RestartDeployment
	err = controller.CheckParams(c, &rd)
	if err != nil {
		response.FailHasMsg(response.ParamError, err.Error(), c)
		return
	}

	err = deployment.RestartDeployment(clientset, rd.DeploymentName, rd.Namespace)
	if err != nil {
		response.FailHasMsg(response.InternalServerError, err.Error(), c)
		return
	}
	response.OK(c)
	return
}

func ScaleDeploymentController(c *gin.Context)  {
	clientset, err := client.ClientSet()
	if err != nil {
		response.FailHasMsg(response.InternalServerError,err.Error(),c)
	}
	var scaleData k8s.ScaleDeployment
	err = controller.CheckParams(c,&scaleData)
	if err != nil {
		response.FailHasMsg(response.ParamError,err.Error(),c)
	}
	err = deployment.ScaleDeployment(clientset, scaleData.Namespace, scaleData.DeploymentName, *scaleData.ScaleNumber)
	if err != nil {
		response.FailHasMsg(response.InternalServerError,err.Error(),c)
	}
	response.OK(c)
	return
}