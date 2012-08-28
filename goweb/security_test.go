package goweb

import (
	"testing"
)

func TestHash(t *testing.T) {

	assertEqual(t, Hash("123"), "4ac85a9a64ed1b5a4560721b2034dd6738c1153e", "Hash of 123")

}
