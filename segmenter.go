package yaba

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Segmenter struct {
	dict *Dictionary
}

func NewSegmenter() *Segmenter {
	return &Segmenter{}
}

func (s *Segmenter) Segment(text string) []Token {
	
}

func (seg *Segmenter) LoadDictionary(fileName string) error {
	if fileName == "" {
		fileName = "data/dict.txt"
	}

	tokens := make([]*Token, 0)

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		split := strings.Split(line, " ")
		word := split[0]
		freq := split[1]
		ifreq, _ := strconv.Atoi(freq)
		pos := split[2]
		tokens = append(tokens, NewToken(word, -1, ifreq, pos))
	}

	seg.dict = NewDictionary(tokens)

	return nil
}
