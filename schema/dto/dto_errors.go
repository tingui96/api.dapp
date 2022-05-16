package dto

// Problem api documentation
type Problem struct {
	Status uint   `example:"503"`
	Title  string `example:"err_code"`
	Detail string `example:"Some error details"`
}

// NewProblem construct a new api error struct and return a pointer to it
//
// - s [uint] ~ HTTP status tu respond
//
// - t [string] ~ Title of the error
//
// - d [string] ~ Description or detail of the error
func NewProblem(s uint, t string, d string) *Problem {
	return &Problem{Status: s, Title: t, Detail: d}
}
