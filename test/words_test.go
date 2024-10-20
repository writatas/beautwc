package test

import (
	"fmt"
	"testing"

	"github.com/beautwc/pkg"
	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	words := pkg.WordsCount{}
	err := words.GetWordsCount("../examples/text2.txt")
	assert.NoError(t, err)
	assert.Equal(t, words.TotalWords, 2)
}

func TestWordCountFailure(t *testing.T) {
	words := pkg.WordsCount{}
	err := words.GetWordsCount("dummytext.txt")
	assert.Error(t, err)
}

func TestWordCountMultipleFiles(t *testing.T) {
	words := pkg.WordsCount{}
	err := words.GetWordsCount("../examples/text1.txt", "../examples/text2.txt", "../examples/text3.txt")
	assert.NoError(t, err)
	totalWordCount := words.TotalWords
	perFileTotalWordCount := words.WordsPerFile[0].Count + words.WordsPerFile[1].Count + words.WordsPerFile[2].Count
	assert.Equal(t, totalWordCount, perFileTotalWordCount)
	fmt.Println(words)
}
