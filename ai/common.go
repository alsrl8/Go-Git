package ai

import "os"

type Role string

const (
	RoleSystem Role = "system"
	RoleUser   Role = "user"
)

type GptModel string

const (
	Gpt3 GptModel = "gpt-3.5-turbo-0125"
	Gpt4 GptModel = "gpt-4-0125-preview"
)

func getOpenAiApiUrl() string {
	return "https://api.openai.com/v1/chat/completions"
}

func getOpenAiApiKey() string {
	return os.Getenv("OPENAI_API_KEY")
}
