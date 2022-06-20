package dto

// docType
const (
	ElectionDocType = "election"
	BallotDocType   = "ballot"
	SectionDocType  = "section"
	UserDocType     = "user"
	ElectionID      = "election-ID1"
	ElectionIndex   = ElectionDocType+":"+ElectionID
)

type UserRegister struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// UserList lUser list of yml file based users
type UserList struct {
	Users []User `yaml:"Users"`
}

// User struct
type User struct {
	Id         string `json:"id"`
	Passphrase string `json:"passphrase"`
	Clear      string `json:"clear"`
	Username   string `json:"username"`
	Name       string `json:"name"`
}

type GetUser struct {
	User User `json:"user"`
}