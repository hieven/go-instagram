package text

import "mvdan.cc/xurls"

type textManager struct {
}

func New() (TextManager, error) {
	text := &textManager{}
	return text, nil
}

func (textManager *textManager) ExtractURL(text string) string {
	return xurls.Strict.FindString(text)
}
