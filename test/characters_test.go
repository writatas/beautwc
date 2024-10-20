package test

import (
	"testing"

	"github.com/beautwc/pkg"
	"github.com/stretchr/testify/assert"
)

func TestCharacterCount(t *testing.T) {
	characters := pkg.CharactersCount{}
	err := characters.GetCharactersCount("../examples/text2.txt")
	assert.NoError(t, err)
	assert.Equal(t, characters.TotalMatches, 12)
}

func TestCharacterCountFailer(t *testing.T) {
	characters := pkg.CharactersCount{}
	err := characters.GetCharactersCount("dummytext.txt")
	assert.Error(t, err)
}

func TestCharacterCountMultipleFiles(t *testing.T) {
	characters := pkg.CharactersCount{}
	err := characters.GetCharactersCount("../examples/text1.txt", "../examples/text2.txt", "../examples/text3.txt")
	assert.NoError(t, err)
	totalCharacters := characters.TotalMatches
	totalPerLineMatches := characters.PerFileMatches[0].Count + characters.PerFileMatches[1].Count + characters.PerFileMatches[2].Count
	assert.Equal(t, totalCharacters, totalPerLineMatches)
}
