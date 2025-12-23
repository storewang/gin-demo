package conf

import (
	"flag"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	cfgFile = flag.String("config", "./config.yaml", "配置文件路径")

	cfg *Config
)

type Config struct {
	Listen string `yaml:"listen"`
	Redis  struct {
		Host        string `yaml:"host"`
		Password    string `yaml:"password"`
		Database    int    `yaml:"database"`
		MaxActive   int    `yaml:"maxActive"`
		MaxIdle     int    `yaml:"maxIdle"`
		IdleTimeout int    `yaml:"idleTimeout"`
	} `yaml:"redis"`
	*OfficialAccountConfig `yaml:"officialAccountConfig"`
	Gemini                 struct {
		APIKey     string `yaml:"apiKey"`
		ModuleName string `yaml:"moduleName"`
	} `yaml:"gemini"`
}

// OfficialAccountConfig 公众号相关配置
type OfficialAccountConfig struct {
	AppID          string `yaml:"appID"`
	AppSecret      string `yaml:"appSecret"`
	Token          string `yaml:"token"`
	EncodingAESKey string `yaml:"encodingAESKey"`
}

// GetConfig 获取配置
func GetConfig() *Config {
	if cfg != nil {
		return cfg
	}
	bytes, err := os.ReadFile(*cfgFile)
	if err != nil {
		panic(err)
	}

	cfgData := &Config{}
	err = yaml.Unmarshal(bytes, cfgData)
	if err != nil {
		panic(err)
	}
	cfg = cfgData
	return cfg
}
