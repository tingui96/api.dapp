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
	ID          string `json:"ID"`
	PatientName string `json:"PatientName"`
	Description string `json:"Description"`
	State       int    `json:"State"`
	Group       string `json:"Group"`
}
