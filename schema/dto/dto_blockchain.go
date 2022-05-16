package dto

import "time"

// DevPopulateOut
type DevPopulateOut struct {
	TxId     string
	Identity string
}

// OrgSpec
type OrgSpec struct {
	Name          string       `yaml:"Name"`
	Domain        string       `yaml:"Domain"`
	EnableNodeOUs bool         `yaml:"EnableNodeOUs"`
	CA            NodeSpec     `yaml:"CA"`
	Template      NodeTemplate `yaml:"Template"`
	Specs         []NodeSpec   `yaml:"Specs"`
	Users         UsersSpec    `yaml:"Users"`
}

// NodeSpec
type NodeSpec struct {
	IsAdmin            bool
	Hostname           string   `yaml:"Hostname"`
	CommonName         string   `yaml:"CommonName"`
	Country            string   `yaml:"Country"`
	Province           string   `yaml:"Province"`
	Locality           string   `yaml:"Locality"`
	OrganizationalUnit string   `yaml:"OrganizationalUnit"`
	StreetAddress      string   `yaml:"StreetAddress"`
	PostalCode         string   `yaml:"PostalCode"`
	SANS               []string `yaml:"SANS"`
}

// NodeTemplate
type NodeTemplate struct {
	Count    int      `yaml:"Count"`
	Start    int      `yaml:"Start"`
	Hostname string   `yaml:"Hostname"`
	SANS     []string `yaml:"SANS"`
}

type ChartData struct {
	Labels   []string   `json:"labels"`
	Datasets []Datasets `json:"datasets"`
}

type Datasets struct {
	Label           string `json:"label"`
	Data            []Data `json:"data"`
	BackgroundColor string `json:"backgroundColor"`
}

type Data struct {
	X string `json:"x"`
	Y int    `json:"y"`
}

type Section struct {
	BackgroundColor string
}

type ElectionRequest struct {
	ElectionID string `json:"electionID"`
	StartTime  string `json:"startTime" metadata:",optional"`
	Hours      uint   `json:"hours" metadata:",optional"`
	State      State  `json:"state,omitempty"`
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


// GetRequest
type GetRequest struct {
	ID string `json:"id"`
	ElectionID  string `json:"electionID" metadata:",optional"`
}

type ElectionOpenRequest struct {
	ID        string        `json:"id"`
	StartTime time.Time     `json:"startTime,omitempty" metadata:",optional"` // ex: time.Date(time.Now().Year(), time.November, 10, 9, 30, 0, 0, time.FixedZone("America/Havana", 0))
	Duration  time.Duration `json:"duration"`                                 // duration time in minutes
}

// BallotCreateRequest
type BallotCreateRequest struct {
	ElectionID  string `json:"electionID"`
	VoterID     string `json:"voterID"`
	CandidateID string `json:"candidateID"`
}


type ResultResponse struct {
	Votes       int    `json:"votes"`
	CandidateID string `json:"candidateID"`
}

type ElectionCreateRequest struct {
	Name      string    `json:"name"`
	Country   string    `json:"country"`
	Locality  string    `json:"locality"`
	StartTime time.Time `json:"startTime,omitempty" metadata:",optional"` // ex: time.Date(time.Now().Year(), time.November, 10, 9, 30, 0, 0, time.FixedZone("America/Havana", 0))
}