package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type RequestMessageBlock struct {
	Type     string                  `json:"type"`
	Elements []RequestMessageElement `json:"elements"`
}

type RequestMessageElement struct {
	Type     string                `json:"type"`
	Style    RichTextStyle         `json:"style"`
	Indent   int                   `json:"indent"`
	Elements []RichTextListElement `json:"elements"`
}

type RichTextListElement struct {
	Type     string                   `json:"type"`
	Elements []RichTextSectionElement `json:"elements"`
}

type RichTextSectionElement struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func getUserToken() (userToken string) {
	userToken = os.Getenv("X_Slack_User_OAuth_Token")
	return
}

func getBotUserToken() (botUserToken string) {
	botUserToken = os.Getenv("X_Slack_Bot_User_Oauth_Token")
	return
}

func PostMessage(channelId, message, threadTimestamp string) {
	token := getUserToken()

	lines := strings.Split(message, "\n")
	var elements []RequestMessageElement
	for _, line := range lines {
		if strings.HasPrefix(line, "- ") {
			line = strings.TrimPrefix(line, "- ")
			elements = append(elements, RequestMessageElement{
				Type:   "rich_text_list",
				Style:  Bullet,
				Indent: 0,
				Elements: []RichTextListElement{
					{
						Type: "rich_text_section",
						Elements: []RichTextSectionElement{{
							Type: "text",
							Text: line,
						}},
					},
				},
			})
		} else if strings.HasPrefix(line, "  - ") {
			line = strings.TrimPrefix(line, "  - ")
			elements = append(elements, RequestMessageElement{
				Type:   "rich_text_list",
				Style:  Bullet,
				Indent: 1,
				Elements: []RichTextListElement{
					{
						Type: "rich_text_section",
						Elements: []RichTextSectionElement{{
							Type: "text",
							Text: line,
						}},
					},
				},
			})
		}
	}

	var blocks []RequestMessageBlock
	blocks = append(blocks, RequestMessageBlock{
		Type:     "rich_text",
		Elements: elements,
	})

	msg := map[string]interface{}{
		"channel":   channelId,
		"text":      message,
		"thread_ts": threadTimestamp, // 이 파라미터가 스레드에 메시지를 연결합니다.
		"blocks":    blocks,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error marshalling message:", err)
		return
	}

	req, err := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", bytes.NewBuffer(msgBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
}

func FetchChannelMessages(channelId string) (ChannelMessageResponse, error) {
	token := getUserToken()
	url := fmt.Sprintf("https://slack.com/api/conversations.history?channel=%s&limit=1", channelId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var response ChannelMessageResponse
	err = json.Unmarshal(body, &response)
	return response, err
}
