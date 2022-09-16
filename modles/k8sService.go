package modles

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/shitou/go-demo-gin/applications/dto"
	"github.com/shitou/go-demo-gin/infrastructure/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type K8sService struct {
	k8suril *utils.K8sUtil
}

func NewK8sService(k8suril *utils.K8sUtil) *K8sService {
	return &K8sService{k8suril: k8suril}
}

func (svc *K8sService) GetAllNamespace() []string {
	ns, err := svc.k8suril.K8sClient.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("获取命名空间失败:", err.Error())
	}

	arr := make([]string, 0, len(ns.Items))
	for _, item := range ns.Items {
		fmt.Printf("namespace: %v\n", item.Name)
		if item.Name != "" {
			arr = append(arr, item.Name)
		}
	}

	return arr
}

func (svc *K8sService) GetSvcListByNs(ns string) []dto.ServiceDTO {
	svcList, err := svc.k8suril.K8sClient.CoreV1().Services(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("获取服务列表失败:", err.Error())
	}

	arr := make([]dto.ServiceDTO, 0, len(svcList.Items))
	for _, item := range svcList.Items {
		svcname := item.Name
		svcspec := item.Spec
		//fmt.Printf("svc.name: %v , time:%v, selector:%v \n", svcname, item.CreationTimestamp, svcspec.Selector)
		arr = append(arr,
			dto.ServiceDTO{
				Name:       svcname,
				Type:       string(svcspec.Type),
				ClusterIp:  svcspec.ClusterIP,
				ExternalIp: strings.Join(svcspec.ExternalIPs, ","),
				Ports:      svc.ports2str(svcspec.Ports),
				Selector:   svc.map2str(svcspec.Selector),
			})

	}

	return arr
}
func (svc *K8sService) map2str(selector map[string]string) string {
	keys := make([]string, 0, len(selector))
	for k := range selector {
		keys = append(keys, fmt.Sprintf("%s=%s", k, selector[k]))
	}
	return strings.Join(keys, ",")
}
func (svc *K8sService) ports2str(ports []v1.ServicePort) string {
	arr := make([]string, 0, len(ports))

	for _, item := range ports {
		arr = append(arr, fmt.Sprintf("%d/%s", item.Port, item.Protocol))
	}

	return strings.Join(arr, ",")
}
