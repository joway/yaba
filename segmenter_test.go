package yaba

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createSegment(t *testing.T) *Segmenter {
	seg := NewSegmenter()
	if err := seg.LoadStopWords(""); err != nil {
		assert.NoError(t, err)
	}
	if err := seg.LoadDictionary(""); err != nil {
		assert.NoError(t, err)
	}
	return seg
}

func TestSegmenter_Segment(t *testing.T) {
	seg := createSegment(t)
	text := "我来到了北京清华大学旁边的北京大学旁边的清华大学"
	tokens := seg.Segment(text)

	got := ""
	for _, token := range tokens {
		got += token.Text
		fmt.Println(token.Text)
	}
	assert.Equal(t, text, got)
}
