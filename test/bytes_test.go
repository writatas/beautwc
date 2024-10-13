package test

import (
	"testing"

	"github.com/beautwc/pkg"
	"github.com/stretchr/testify/assert"
)

func TestBytes(t *testing.T) {
	bytes := pkg.Bytes{Bytes: 1024.0}

	bytes.CalculateBytesAndPrefixes()
	assert.Equal(t, bytes.Kilobytes, bytes.Bytes/1024)
	assert.Equal(t, bytes.Megabytes, bytes.Bytes/(1024*1024))
	assert.Equal(t, bytes.Gigabytes, bytes.Bytes/(1024*1024*1024))
}

func TestReadBytesFromFileSuccess(t *testing.T) {
	totalMemory, err := pkg.GetFilesSize("../examples/text1.txt", "../examples/text2.txt", "../examples/text3.txt")
	assert.NoError(t, err)
	t.Run("TestBytesCalculation", func(t *testing.T) {
		assert.Equal(t, totalMemory.TotalBytes.Kilobytes, totalMemory.TotalBytes.Bytes/1024)
		assert.Equal(t, totalMemory.TotalBytes.Megabytes, totalMemory.TotalBytes.Bytes/(1024*1024))
		assert.Equal(t, totalMemory.TotalBytes.Gigabytes, totalMemory.TotalBytes.Bytes/(1024*1024*1024))
	})

}

func TestReadBytesFromFileFailure(t *testing.T) {
	_, err := pkg.GetFilesSize("dummyfilepath.txt")
	assert.Error(t, err)
}
