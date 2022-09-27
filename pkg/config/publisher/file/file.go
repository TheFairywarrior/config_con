package file


type FilePublisherConfig struct {
	// contains filtered or unexported fields
	Name     string `yaml:"name"`
	FilePath string `yaml:"filePath"`
	FileMode int    `yaml:"fileMode"`
}
