package test

import (
	"testing"

	"github.com/beautwc/pkg"
	"github.com/stretchr/testify/assert"
)

func TestLines(t *testing.T) {
	lines := pkg.LinesCount{}
	err := lines.GetLinesCount("../examples/text1.txt")
	assert.NoError(t, err)
	assert.Equal(t, lines.TotalLines, 93)
}

func TestLinesFailure(t *testing.T) {
	lines := pkg.LinesCount{}
	err := lines.GetLinesCount("dummytext.txt")
	assert.Error(t, err)
}

func TestMultipleFilesLines(t *testing.T) {
	lines := pkg.LinesCount{}
	err := lines.GetLinesCount("../examples/text1.txt", "../examples/text2.txt", "../examples/text3.txt")
	assert.NoError(t, err)
	totalLineCount := lines.TotalLines
	linesPerFile := lines.PerFileLines[0].Count + lines.PerFileLines[1].Count + lines.PerFileLines[2].Count
	assert.Equal(t, totalLineCount, linesPerFile)
}
