package yaba

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSegmenter_LoadDictionary(t *testing.T) {
	seg := NewSegmenter()
	if err := seg.LoadDictionary(""); err != nil {
		assert.NoError(t, err)
	}
}
