package types

type SlackBaseMessage struct {
	Text   string      `json:"text"`
	Blocks interface{} `json:"blocks"`
}

type SlackSectionElement struct {
	Type string           `json:"type"`
	Text SlackTextElement `json:"text"`
}

type SlackTextElement struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type SlackDivider struct {
	Type string `json:"type"`
}

type SlackContextElement struct {
	Type     string      `json:"type"`
	Elements interface{} `json:"elements"`
}

type SlackImageElement struct {
	Type  string `json:"type"`
	Image string `json:"image_url"`
	Alt   string `json:"alt_text"`
}
