package transformer

type Step interface {
	Process(any) (any, error)
}

// StepConfig is the holder for the specific step configuration.
type StepConfig struct {

}

type Transformer struct {
	Name  string
	Steps []Step
}
