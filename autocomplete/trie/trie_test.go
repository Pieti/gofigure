package trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindWord(t *testing.T) {
	tr := NewTrie()
	tr.Add("foo")
	tr.Add("bar")

	assert.NotNil(t, tr.Find("foo"))
	assert.NotNil(t, tr.Find("bar"))
	assert.Nil(t, tr.Find("foobar"))
}

func TestFindNode(t *testing.T) {
	tr := NewTrie()
	tr.Add("foobar")

	assert.Nil(t, tr.Find("foo"))
	assert.Nil(t, tr.Find("bar"))
	assert.NotNil(t, tr.Find("foobar"))

	assert.NotNil(t, tr.FindNode("foo"))
	assert.Nil(t, tr.FindNode("bar"))
}

func TestFindWordsWithPrefix(t *testing.T) {
	tr := NewTrie()
	tr.Add("foo")
	tr.Add("bar")
	tr.Add("foobar")

	words := tr.FindWordsWithPrefix("fo")
	assert.ElementsMatch(t, words, []string{"foo", "foobar"})
}

func TestFindNoWordsWithPrefix(t *testing.T) {
	tr := NewTrie()
	tr.Add("foo")
	tr.Add("bar")
	tr.Add("foobar")

	words := tr.FindWordsWithPrefix("baz")
	assert.Empty(t, words)
}

func TestCount(t *testing.T) {
	tr := NewTrie()
	tr.Add("foo")
	tr.Add("bar")
	tr.Add("foobar")

	node := tr.FindNode("foo")
	assert.Equal(t, 1, node.Count())
}
