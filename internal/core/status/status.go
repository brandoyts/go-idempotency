package status

import "fmt"

type Status int

const (
	Processing Status = iota
	Done
)

func (s Status) ToString() string {
	switch s {
	case Processing:
		return "Processing"
	case Done:
		return "Done"
	default:
		return fmt.Sprintf("unknown status (%v)", s)
	}
}
