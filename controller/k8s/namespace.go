package k8s

import (
	"github.com/ethanliuuu/k8s-client/controller/response"
	"github.com/ethanliuuu/k8s-client/pkg/k8s/client"
	"github.com/ethanliuuu/k8s-client/pkg/k8s/namespace"
	"github.com/gin-gonic/gin"
)

func GetNamespaceList(c *gin.Context) {
	clientset, err := client.ClientSet()
	if err != nil {
		response.FailHasMsg(response.InternalServerError, err.Error(), c)
	}
	namespaceList, err := namespace.GetNamespaceList(clientset)
	if err != nil {
		response.FailHasMsg(response.InternalServerError, err.Error(), c)
	}
	response.OKHasData(namespaceList, c)
}
