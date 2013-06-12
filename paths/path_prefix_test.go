package paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ChildController struct{}
type ChildWithMultipleWordsController struct{}

func TestPathPrefixForClass(t *testing.T) {

	c := new(ChildController)
	assert.Equal(t, PathPrefixForClass(c), "child")

	c2 := new(ChildWithMultipleWordsController)
	assert.Equal(t, PathPrefixForClass(c2), "child-with-multiple-words")

}
