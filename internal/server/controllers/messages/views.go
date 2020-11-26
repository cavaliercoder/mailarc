package messages

import (
	"mailarc/build/gen/views"
	"mailarc/internal/mimecontent"
)

var tmplGetMessage = parseView("GetMessage", views.ViewGetMessage)

type ViewModelContent struct {
	Name    string
	Content *mimecontent.Content
}

type ViewMessage struct {
	UID     string
	Subject string
	From    string
	Message *mimecontent.Content
	Root    *mimecontent.Content
	Parts   []*ViewModelContent
}

func NewViewMessage(uid string, content *mimecontent.Content) *ViewMessage {
	v := &ViewMessage{
		UID:     uid,
		Subject: content.Headers.Get("Subject"),
		From:    content.Headers.Get("From"),
		Message: content,
		Root:    content.Root(),
		Parts:   getContent(content),
	}
	return v
}

func getContentName(content *mimecontent.Content) (s string) {
	if s = content.Filename(); s != "" {
		return s
	}
	return content.MediaType()
}

func getContent(content *mimecontent.Content) []*ViewModelContent {
	a := make([]*ViewModelContent, 0, 8)
	queue := make([]*mimecontent.Content, 1, 8)
	queue[0] = content
	for len(queue) > 0 {
		content = queue[0]
		queue = queue[1:]
		if content.IsMultipart() {
			for _, child := range content.Children {
				queue = append(queue, child)
			}
			continue
		}
		a = append(a, &ViewModelContent{
			Name:    getContentName(content),
			Content: content,
		})
	}
	return a
}
