package ai

import (
	"Go-Git/gitlog"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    GptModel  `json:"model"`
	Messages []Message `json:"messages"`
}

func RequestToSummarizeGitLogs(commits []gitlog.Commit, reports []string) string {
	content := fmt.Sprintf("%+v\n\n%+v\n", commits, reports)

	request := Request{
		Model: Gpt4,
		Messages: []Message{
			{Role: RoleSystem, Content: "You should read the Git commit logs and daily report, and make a summary for the daily report. The final format of the summary is markdown, and the example is as follows.\n- Functional development\n  - How did you develop a function called a.\n- Debugging\n  - Fixed a bug.\n- - Refactor\n  - I did a refactor of some code.\n- - Test\n  - I did some kind of test.\n- Deployment\n  - A service was distributed.\n\nThe following is a list of git commit logs for the work the worker has done today. You make summary in Korean."},
			{Role: RoleUser, Content: content},
		},
	}
	respBody := SendChatRequest(request)
	resp, _ := ReadChatResponseBody(respBody)
	return resp
}

func SendChatRequest(chatRequest Request) []byte {
	url := getOpenAiApiUrl()
	apiKey := getOpenAiApiKey()

	jsonData, err := json.Marshal(chatRequest)
	if err != nil {
		fmt.Printf("Error marshalling request data: %v\n", err)
		return nil
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error createing request: %v\n", err)
		return nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erorr making request: %v\n", err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body :%v\n", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil
	}
	return body
}
