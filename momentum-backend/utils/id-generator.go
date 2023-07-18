package utils

import (
	"crypto"
	"encoding/hex"
)

var hash = crypto.SHA3_256.New()

// Generates new id from node by generating a hash from src-string
// src must be unique in its domain to guarantee collision resistance.
// If one property is not enough, combine multiple into the src string.
func GenerateId(src string) (string, error) {

	hash.Reset()

	_, err := hash.Write([]byte(src))
	if err != nil {
		return "", err
	}

	result := make([]byte, 0)
	result = hash.Sum(result)
	if err != nil {
		return "", nil
	}

	// fmt.Println("created hash:", result, "/ str:", hex.EncodeToString(result))

	id := hex.EncodeToString(result)
	id = id[:len(id)/4]
	return id, nil
}
