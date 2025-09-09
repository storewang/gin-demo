package officialaccount

import (
	"github.com/silenceper/wechat/v2/officialaccount/message"
	log "github.com/sirupsen/logrus"
)

//article1 := message.NewArticle("测试图文1", "图文描述", "", "")
//articles := []*message.Article{article1}
//news := message.NewNews(articles)
//return &message.Reply{MsgType: message.MsgTypeNews, MsgData: news}

//voice := message.NewVoice(mediaID)
//return &message.Reply{MsgType: message.MsgTypeVoice, MsgData: voice}

//
//video := message.NewVideo(mediaID, "标题", "描述")
//return &message.Reply{MsgType: message.MsgTypeVideo, MsgData: video}

//music := message.NewMusic("标题", "描述", "音乐链接", "HQMusicUrl", "缩略图的媒体id")
//return &message.Reply{MsgType: message.MsgTypeMusic, MsgData: music}

// 多客服消息转发
// transferCustomer := message.NewTransferCustomer("")
// return &message.Reply{MsgType: message.MsgTypeTransfer, MsgData: transferCustomer}
func txtMessageHandler(msg *message.MixMessage) *message.Reply {
	text := message.NewText(msg.Content)
	log.Infof("receive txt msg, content=%s", msg.Content)
	log.Infof("receive msg, msgType=%s", msg.MsgType)
	if msg.MsgType == message.MsgTypeImage {
		text.Content = "你发送的是图片"
		log.Infof("receive img msg, url=%s", msg.PicURL)
	}
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}
