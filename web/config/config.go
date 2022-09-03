package config

type Config struct {
	AppName  string
	Server   ServerConfig
	Database DatabaseConfig
}
type ServerConfig struct {
	Host string
	SsL  bool
}
type DatabaseConfig struct {
	Driver   string
	Host     string
	UserName string
	Passwd   string
	Schemal  string
	Pool     PoolConfig
}
type PoolConfig struct {
	IdleMax        int32
	Idlemin        int32
	ReadTimeOut    int64
	ConnectTimeout int64
}
