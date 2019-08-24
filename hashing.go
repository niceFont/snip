package main

import (
	"crypto/sha1"
	"fmt"
)

//HashUrl hashes the url string
func HashUrl(url string) string {

	hash := sha1.New()
	hash.Write([]byte(url))
	result := hash.Sum(nil)

	return fmt.Sprintf("%x", result)[:5]
}
