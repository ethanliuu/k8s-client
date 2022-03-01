package controller

import (
	"fmt"
	"github.com/ethanliuuu/k8s-client/common"
	"github.com/gin-gonic/gin"
)

func CheckParams(c *gin.Context, ptr interface{}) error {
	if ptr == nil {
		return nil
	}
	switch t := ptr.(type) {
	case string:
		if t != "" {
			panic(t)
		}
	case error:
		panic(t.Error())
	}
	if err := c.ShouldBindJSON(&ptr); err != nil {
		common.LOG.Warn(fmt.Sprintf("解析参数失败：%v\n", err.Error()))
		return err
	}
	return nil
}
