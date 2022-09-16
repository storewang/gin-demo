package dto

type ServiceDTO struct {
	Name       string
	Type       string
	ClusterIp  string
	ExternalIp string
	Ports      string
	Age        string
	Selector   string
}
