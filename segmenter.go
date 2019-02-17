package yaba

import (
	"io/ioutil"
	"strings"
)

type Segmenter struct {
	dict      *Dictionary
	stopWords []string
}

func NewSegmenter() *Segmenter {
	return &Segmenter{}
}

func (seg *Segmenter) Segment(text string) []*Token {
	tokens := make([]*Token, 0)
	start := 0
	total := len(text)
	for start < total {
		token := seg.dict.Match(text[start:])
		if token == nil {
			start++
			continue
		}
		tokens = append(
			tokens,
			NewToken(
				token.Text,
				start,
				token.Frequency,
				token.Pos,
			),
		)
		start += len(token.Text)
	}
	return tokens
}

func (seg *Segmenter) LoadStopWords(filename string) error {
	if filename == "" {
		filename = "data/stop.txt"
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	lines := strings.Split(string(content), "\n")
	seg.stopWords = make([]string, len(lines))
	for _, line := range lines {
		seg.stopWords = append(seg.stopWords, line)
	}
	return nil
}

func (seg *Segmenter) LoadDictionary(filename string) error {
	if filename == "" {
		filename = "data/dict.txt"
	}
	seg.dict = NewDictionary()
	return seg.dict.Load(filename)
}
