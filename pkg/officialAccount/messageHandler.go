package officialaccount

import (
	"context"
	"wechatdemo/conf"

	"github.com/google/generative-ai-go/genai"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
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
	log.Infof("receive txt msg, content=%s", msg.Content)
	log.Infof("receive msg, msgType=%s", msg.MsgType)

	if msg.MsgType == message.MsgTypeImage {
		text := message.NewText("你发送的是图片")
		log.Infof("receive img msg, url=%s", msg.PicURL)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	}

	// 使用Gemini生成回复
	replyContent, err := generateGeminiReply(msg.Content)
	if err != nil {
		log.Errorf("generate gemini reply error: %v", err)
		text := message.NewText("抱歉，我暂时无法回复，请稍后再试。")
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	}

	text := message.NewText(replyContent)
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}

// generateGeminiReply 使用Gemini AI生成回复
func generateGeminiReply(userMessage string) (string, error) {
	cfg := conf.GetConfig()
	if cfg.Gemini.APIKey == "" {
		return "Gemini API密钥未配置", nil
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(cfg.Gemini.APIKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("请作为微信公众号助手回复用户消息："+userMessage))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		if text, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
			return string(text), nil
		}
	}

	return "无法生成回复", nil
}
