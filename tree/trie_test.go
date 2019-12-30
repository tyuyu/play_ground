package tree

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	t.Log(trie.Search("apple"))
	t.Log(trie.Search("app"))
	t.Log(trie.StartsWith("app"))
	trie.Insert("app")
	t.Log(trie.Search("app"))
}
