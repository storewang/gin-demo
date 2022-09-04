package config

type Config struct {
	AppName  string
	Server   GinConf
	Database DbConf
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
