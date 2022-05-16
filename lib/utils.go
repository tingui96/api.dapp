package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"
)

// Checksum returns the checksum of some data, using a specified algorithm.
// It only returns an error when an invalid algorithm is used. The valid ones
// are SHA256
func Checksum(algorithm string, data []byte) (checksum string, err error) {
	// default
	var _hash hash.Hash
	switch strings.ToUpper(algorithm) {
	case "SHA256":
		_hash = sha256.New()
	default:
		msg := "invalid algorithm parameter passed go Checksum: %s"
		return checksum, fmt.Errorf(msg, algorithm)
	}
	_hash.Write(data)
	str := hex.EncodeToString(_hash.Sum(nil))
	return str, nil
}
