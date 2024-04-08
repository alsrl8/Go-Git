package slack

type ChannelMessageResponse struct {
	Ok               bool             `json:"ok"`
	Messages         []ChannelMessage `json:"messages"`
	HasMore          bool             `json:"has_more"`
	PinCount         int              `json:"pin_count"`
	ResponseMetadata `json:"response_metadata"`
}

type ChannelMessage struct {
	Type string `json:"type"`
	User string `json:"user"`
	Text string `json:"text"`
	Ts   string `json:"ts"`
}

type ResponseMetadata struct {
	NextCursor string `json:"next_cursor"`
}
