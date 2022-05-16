package lib

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func UniqueStrings(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

func DecodePayload(payload []byte) interface{} {
	// first attempt to parse for JSON, if not successful then just decode to string
	var structured interface{}
	err := json.Unmarshal(payload, &structured)
	if err != nil {
		return string(payload)
	}
	return structured
}