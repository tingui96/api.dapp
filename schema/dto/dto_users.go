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

type Candidates struct {
	URL       string `json:"url"`
	SectionID string `json:"sectionID"`
	Section   string `json:"section"`
	User
}

var CandidatesIDName = map[string]string{"candidate1":"Muhilo Perez Andorra", "candidate2":"Alexander Rubio Lisipo",
	"candidate3":"Ruben de la paz", "candidate4":"Alfonso Montes De Oca"}

var CandidatesJSON = `[
    {
        "id": "candidate1",
        "url": "candidate1",
        "sectionID": "-",
        "section": "-",
        "username": "candidate1@gmail.com",
        "name": "Muhilo Perez Andorra",
        "university": "Universidad de Mx",
        "province": "-"
    },
    {
        "id": "candidate2",
        "url": "candidate2",
        "sectionID": "-",
        "section": "-",
        "username": "candidate2@gmail.com",
        "name": "Alexander Rubio Lisipo",
        "university": "Universidad de Mx",
        "province": "-"
    },
    {
        "id": "candidate3",
        "url": "candidate3",
        "sectionID": "-",
        "section": "-",
        "username": "candidate3@gmail.com",
        "name": "Ruben de la paz",
        "university": "Universidad de Mx",
        "province": "-"
    },
    {
        "id": "candidate4",
        "url": "candidate4",
        "sectionID": "-",
        "section": "-",
        "username": "candidate4@gmail.com",
        "name": "Alfonso Montes De Oca",
        "university": "Universidad de Mx",
        "province": "-"
    }
]`

type GetUser struct {
	User User `json:"user"`
}

// User define yml file based users
//type User struct {
//	Username   string `yaml:"Username"`
//	Passphrase string `yaml:"Passphrase"`
//	Province   string `yaml:"Province"`
//	DNI        string `yaml:"DNI"`
//	University string `yaml:"University"`
//}

// UsersSpec
type UsersSpec struct {
	Count int `yaml:"Count"`
}