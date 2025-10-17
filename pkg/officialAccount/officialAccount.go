package officialaccount

import (
	"wecahtdemo/conf"

	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	log "github.com/sirupsen/logrus"
)

// ExampleOfficialAccount 公众号操作样例
type ExampleOfficialAccount struct {
	wc              *wechat.Wechat
	officialAccount *officialaccount.OfficialAccount
}

// ExampleOfficialAccount new
func NewExampleOfficialAccount(wc *wechat.Wechat) *ExampleOfficialAccount {
	//init config
	globalCfg := conf.GetConfig()
	offCfg := &offConfig.Config{
		AppID:          globalCfg.AppID,
		AppSecret:      globalCfg.AppSecret,
		Token:          globalCfg.Token,
		EncodingAESKey: globalCfg.EncodingAESKey,
	}
	log.Debugf("offCfg=%+v", offCfg)
	officialAccount := wc.GetOfficialAccount(offCfg)
	return &ExampleOfficialAccount{
		wc:              wc,
		officialAccount: officialAccount,
	}
}

// Serve 处理消息
func (ex *ExampleOfficialAccount) Serve(c *gin.Context) {
	// 传入request和responseWriter
	server := ex.officialAccount.GetServer(c.Request, c.Writer)
	server.SkipValidate(true)
	server.SetMessageHandler(txtMessageHandler)

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		log.Errorf("Serve Error, err=%v", err)
		return
	}
	//发送回复的消息
	err = server.Send()
	if err != nil {
		log.Errorf("Send Error, err=%+v", err)
		return
	}
}
