package officialaccount

import (
	"context"
	"log"

	"wechatdemo/conf"

	"google.golang.org/genai"
)

var GeminiClient *GeminiApi

type GeminiApi struct {
	client *genai.Client
	ctx    context.Context
}

func NewClient() (*GeminiApi, error) {
	cfg := conf.GetConfig()
	ctx := context.Background()
	log.Printf("Using Gemini APIKey: %s", cfg.Gemini.APIKey)
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  cfg.Gemini.APIKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}
	GeminiClient = &GeminiApi{client: client, ctx: ctx}
	return GeminiClient, nil
}

func (api *GeminiApi) Text(text string) string {
	temp := float32(0.9)
	topP := float32(0.5)
	topK := float32(20.0)

	config := &genai.GenerateContentConfig{
		Temperature:      &temp,
		TopP:             &topP,
		TopK:             &topK,
		ResponseMIMEType: "application/json",
	}
	cfg := conf.GetConfig()
	log.Printf("Using Gemini Module: %s", cfg.Gemini.ModuleName)
	log.Printf("Using Gemini client: %s", api.client)
	result, err := api.client.Models.GenerateContent(
		api.ctx,
		cfg.Gemini.ModuleName,
		genai.Text(text),
		config,
	)
	if err != nil {
		return "接口调用失败"
	}
	return result.Text()
}
func (api *GeminiApi) ChartWithText(text string, imageBytes []byte) string {
	parts := []*genai.Part{
		genai.NewPartFromBytes(imageBytes, "image/jpeg"),
		genai.NewPartFromText(text),
	}
	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}
	cfg := conf.GetConfig()
	result, err := api.client.Models.GenerateContent(
		api.ctx,
		cfg.Gemini.ModuleName,
		contents,
		config,
	)
	if err != nil {
		return "接口调用失败"
	}
	return result.Text()
}

func (api *GeminiApi) Text2Audio(text string) []byte {
	voiceConfig := &genai.VoiceConfig{
		PrebuiltVoiceConfig: &genai.PrebuiltVoiceConfig{
			VoiceName: "Gacrux",
		},
	}
	spconfig := &genai.SpeechConfig{
		VoiceConfig: voiceConfig,
	}
	config := &genai.GenerateContentConfig{
		ResponseModalities: []string{"AUDIO"},
		SpeechConfig:       spconfig,
	}

	spart := &genai.Part{
		Text: text,
	}
	content := &genai.Content{
		Parts: []*genai.Part{spart},
	}

	result, err := api.client.Models.GenerateContent(
		api.ctx,
		"gemini-2.5-flash-preview-tts",
		[]*genai.Content{content},
		config,
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	resp := result.Candidates
	if len(resp) == 0 || resp[0].Content == nil ||
		len(resp[0].Content.Parts) == 0 || resp[0].Content.Parts[0].InlineData == nil {
		log.Fatal("未找到音频数据")
		return nil
	}
	base64Data := resp[0].Content.Parts[0].InlineData.Data

	return base64Data
}
