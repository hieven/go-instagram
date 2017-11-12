package text

type TextManager interface {
	ExtractURL(text string) string
}
