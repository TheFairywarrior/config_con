package file

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilePublisher_Publish(t *testing.T) {
	filePath := "test.txt"
	defer os.Remove(filePath)
	publisher := FilePublisher{
		filePath: filePath,
		fileMode: 0644,
	}

	err := publisher.Publish([]byte("test"))
	assert.NoError(t, err)

	fileData, err := ioutil.ReadFile(filePath)
	assert.NoError(t, err)
	assert.Equal(t, "test", string(fileData))
}
