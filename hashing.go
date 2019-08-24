package main

import (
	"crypto/sha1"
	"fmt"
)

//HashURL hashes the url string
func HashURL(url string) string {

	hash := sha1.New()
	hash.Write([]byte(url))
	result := hash.Sum(nil)

	return fmt.Sprintf("%x", result)[:5]
}
