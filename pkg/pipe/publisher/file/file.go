package file

import "os"

type FilePublisher struct {
	// contains filtered or unexported fields
	name     string `yaml:"name"`
	filePath string `yaml:"filePath"`
	fileMode int    `yaml:"fileMode"`
}

func (publisher FilePublisher) Name() string {
	return publisher.name
}

func (publisher FilePublisher) FilePath() string {
	return publisher.filePath
}

func (publisher FilePublisher) FileMode() int {
	return publisher.fileMode
}

func NewFilePublisher(name, filePath string, fileMode int) FilePublisher {
	return FilePublisher{
		name:     name,
		filePath: filePath,
		fileMode: fileMode,
	}
}

func (publisher FilePublisher) Publish(data []byte) error {
	file, err := os.OpenFile(publisher.FilePath(), os.O_APPEND|os.O_CREATE|os.O_RDWR, os.FileMode(publisher.FileMode()))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}
