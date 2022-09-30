package file

import (
	"io/ioutil"
	"os"
	"reflect"
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

func TestFilePublisher_Getters(t *testing.T) {
	publisher := FilePublisher{
		name:     "test",
		filePath: "test.txt",
		fileMode: 0644,
	}
	assert.Equal(t, "test", publisher.Name())
	assert.Equal(t, "test.txt", publisher.FilePath())
	assert.Equal(t, 0644, publisher.FileMode())
}

func TestNewFilePublisher(t *testing.T) {
	type args struct {
		name     string
		filePath string
		fileMode int
	}
	tests := []struct {
		name string
		args args
		want FilePublisher
	}{
		{
			name: "TestNewFilePublisher",
			args: args{
				name:     "test",
				filePath: "test.txt",
				fileMode: 0644,
			},
			want: FilePublisher{
				name:     "test",
				filePath: "test.txt",
				fileMode: 0644,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilePublisher(tt.args.name, tt.args.filePath, tt.args.fileMode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilePublisher() = %v, want %v", got, tt.want)
			}
		})
	}
}
