package config

type Config struct {
	AppName  string
	Server   GinConf
	Database DbConf
	K8sConf  K8sConf
}
type GinConf struct {
	Host string
	Port string
}
type DbConf struct {
	Host   string
	User   string
	Passwd string
	Schema string
}
type K8sConf struct {
	BearerToken string
	Host        string
}
