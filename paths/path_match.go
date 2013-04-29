package paths

import (
	"github.com/stretchrcom/stew/objects"
)

type PathMatch struct {
	Matches    bool
	Parameters objects.Map
}

var PathDoesntMatch *PathMatch = new(PathMatch)
