package paths

import (
	"github.com/stretchr/stew/objects"
)

/*
  PathMatch holds details about whether a path matches a PathPattern or not.

  If it does match, the Parameters map will contain the values for any
  dynamic parameters discovered.
*/
type PathMatch struct {
	Matches    bool
	Parameters objects.Map
}

/*
  PathDoesntMatch is a special instance of PathMatch that indicates the paths
  do not match.
*/
var PathDoesntMatch *PathMatch = new(PathMatch)
