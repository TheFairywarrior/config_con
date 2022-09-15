package file

import "os"

type FilePublisher struct {
	// contains filtered or unexported fields
	Name     string `yaml:"name"`
	FilePath string `yaml:"filePath"`
	FileMode int    `yaml:"fileMode"`
}

func (publisher FilePublisher) Publish(data []byte) error {
	file, err := os.OpenFile(publisher.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(publisher.FileMode))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}
