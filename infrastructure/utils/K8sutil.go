package utils

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/shitou/go-demo-gin/web/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type K8sUtil struct {
	K8sClient *kubernetes.Clientset
}

func NewK8sUtil(k8sConfig *config.Config) *K8sUtil {
	var tlsConfig = &tls.Config{
		InsecureSkipVerify: true, // 忽略证书验证
	}

	var transport http.RoundTripper = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
		DisableCompression:    true,
	}
	log.Println("k8s.host:", k8sConfig.K8sConf.Host)
	log.Println("k8s.token:", k8sConfig.K8sConf.BearerToken)

	var config = &rest.Config{
		Host:        k8sConfig.K8sConf.Host,
		BearerToken: k8sConfig.K8sConf.BearerToken,
		Transport:   transport,
	}

	c, err := kubernetes.NewForConfig(config)

	if err != nil {
		log.Fatal(err)
	}

	return &K8sUtil{K8sClient: c}
}
