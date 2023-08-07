package sortv

type StepType uint8

const (
	StepLess StepType = iota
	StepSwap
)

type Step struct {
	StepType
	i, j int
}

func NewStep(t StepType, i int, j int) *Step {
	return &Step{t, i, j}
}
