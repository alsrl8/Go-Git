package ai

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Response struct {
	ID          string   `json:"id"`
	Object      string   `json:"object"`
	Created     int      `json:"created"`
	Model       string   `json:"model"`
	Choices     []Choice `json:"choices"`
	Usage       Usage    `json:"usage"`
	FingerPrint string   `json:"system_fingerprint"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	LogProbs     *string `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func ReadChatResponseBody(respBody []byte) (string, int) {
	var response Response
	if err := json.Unmarshal(respBody, &response); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return "", -1
	}

	if len(response.Choices) > 0 {
		firstContent := response.Choices[0].Message.Content
		lines := strings.Split(firstContent, "\n")
		indent := "    "
		for i, line := range lines {
			lines[i] = indent + line
		}
		indentedContent := strings.Join(lines, "\n")
		return indentedContent, response.Usage.TotalTokens
	} else {
		fmt.Println("No content found")
		return "", -1
	}
}
