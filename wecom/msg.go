package wecom

type Message interface {
	GetType() string
}

type BaseMessage struct {
	MsgType string `json:"msgtype"`
}

func (m *BaseMessage) GetType() string {
	return m.MsgType
}

type Text struct {
	Content string `json:"content"`
}
type TextMessage struct {
	BaseMessage
	Text `json:"text"`
}

type Markdown struct {
	Content string `json:"content"`
}
type MarkdownMessage struct {
	BaseMessage
	Markdown `json:"markdown"`
}

type NewsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

type News struct {
	NewsArticle []*NewsArticle `json:"articles"`
}
type NewsMessage struct {
	BaseMessage
	News `json:"news"`
}
