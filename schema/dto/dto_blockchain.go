package dto

// DevPopulateOut
type DevPopulateOut struct {
	TxId     string
	Identity string
}

type GetRequest struct {
	ID string `json:"id"`
}

// asset de chaincode-go
type Asset struct {
	AppraisedValue int    `json:"AppraisedValue"`
	Color          string `json:"Color"`
	ID             string `json:"ID"`
	Owner          string `json:"Owner"`
	Size           int    `json:"Size"`
}