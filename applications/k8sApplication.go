package applications

import (
	"github.com/shitou/go-demo-gin/applications/dto"
	"github.com/shitou/go-demo-gin/modles"
)

type K8sApplication struct {
	svc *modles.K8sService
}

func NewK8sApplication(svc *modles.K8sService) *K8sApplication {
	return &K8sApplication{svc: svc}
}

func (app *K8sApplication) GetAllNamespace() []string {
	return app.svc.GetAllNamespace()
}

func (app *K8sApplication) GetSvcListByNs(ns string) []dto.ServiceDTO {

	return app.svc.GetSvcListByNs(ns)
}
