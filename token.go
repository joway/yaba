package yaba

type Token struct {
	Text      string
	Start     int
	Frequency int
	Pos       string
}

func NewToken(text string, start int, freq int, pos string) *Token {
	return &Token{
		Text:      text,
		Start:     start,
		Frequency: freq,
		Pos:       pos,
	}
}

func (t *Token) Length() int {
	return len(t.Text)
}
