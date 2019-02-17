package yaba

import (
	"github.com/joway/trie"
	"io/ioutil"
	"strconv"
	"strings"
)

type Dictionary struct {
	trie *trie.Trie
}

func NewDictionary() *Dictionary {
	t := trie.NewRoot()
	return &Dictionary{
		trie: t,
	}
}

func (d *Dictionary) Load(filename string) error {
	content, err := ioutil.ReadFile(filename)
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
		token := NewToken(word, -1, ifreq, pos)

		d.trie.AddWordString(token.Text, token)
	}

	return nil
}

func (d *Dictionary) Match(key string) *Token {
	_, val := d.trie.PrefixSearchString(key)
	if val == nil {
		return nil
	}
	return val.(*Token)
}
