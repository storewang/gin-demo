package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shitou/go-demo-gin/applications"
	"github.com/shitou/go-demo-gin/infrastructure/utils"
	"github.com/shitou/go-demo-gin/modles"
)

type K8sHandler struct {
	k8sSvc *applications.K8sApplication
}

func NewK8sHandler(k8suril *utils.K8sUtil) *K8sHandler {
	k8sSvc := modles.NewK8sService(k8suril)
	k8sApp := applications.NewK8sApplication(k8sSvc)

	return &K8sHandler{k8sSvc: k8sApp}
}

func (k *K8sHandler) GetAllNamespace(c *gin.Context) {
	namespaces := k.k8sSvc.GetAllNamespace()

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": namespaces,
	})
}

func (k *K8sHandler) GetSvcListByNs(c *gin.Context) {
	ns := c.DefaultQuery("ns", "g1")
	log.Println("查询服务:", ns)
	svcList := k.k8sSvc.GetSvcListByNs(ns)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": svcList,
	})
}
