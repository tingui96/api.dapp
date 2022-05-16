package dto

// DevPopulateOut
type DevPopulateOut struct {
	TxId     string
	Identity string
}

type TestRequest struct {
	Label string `json:"label"`
	State State  `json:"state,omitempty"`
}

// State enum for Election state field
type State uint

const (
	UNINITIATED State = iota + 1 // no iniciado
	INITIATED                    // iniciado
	FINISHED                     // terminado
)
func (state State) String() string {
	names := []string{"UNINITIATED", "INITIATED", "FINISHED"}
	if state < UNINITIATED || state > FINISHED {
		return "UNKNOWN"
	}
	return names[state-1]
}

type GetRequest struct {
	ID string `json:"id"`
}