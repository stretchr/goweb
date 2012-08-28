package goweb

import (
	"crypto/sha1"
	"fmt"
)

// HashSecret represents a secret string that is used to hash cookies.
// 
// For true security, this should be changed for each application.
var HashSecret string = "cX8Os0wfB6uCGZZSZHIi6rKsy7b0scE9"

// Hash one-way hashes a string with the private HashSecret value.
func Hash(s string) string {

	hash := sha1.New()
	hash.Write([]byte(s))
	hash.Write([]byte(HashSecret))
	return fmt.Sprintf("%x", hash.Sum(nil))

}
