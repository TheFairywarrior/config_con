package pipe

type Pipe struct {
}

type PipeConfig struct {
	Name     string `yaml:"name"`
	Consumer string `yaml:"consumer"`
}
