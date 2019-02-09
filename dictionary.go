package yaba

import "github.com/dghubble/trie"

type Dictionary struct {
	trie trie.Trier
}

func NewDictionary(tokens []*Token) *Dictionary {
	t := trie.NewRuneTrie()
	for _, token := range tokens {
		t.Put(token.Text, token)
	}
	return &Dictionary{t}
}

func (d *Dictionary) Get(key string) *Token {
	return d.trie.Get(key).(*Token)
}
