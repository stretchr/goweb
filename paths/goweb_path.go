package paths

import (
	"github.com/stretchrcom/stew/objects"
	"strings"
)

/*
  PathPattern represents a path that can contain special matching segments.

  Valid paths include:

    /*** - Matches everything
    /literal
    /{placeholder}
    /[optional placeholder]
    /* - matches like a placeholder but doesn't care what it is
    /something/*** - Matches the start plus anything after it

*/
type PathPattern struct {

	// RawPath is the raw path.
	RawPath string

	path *Path
}

func NewPathPattern(path string) (*PathPattern, error) {

	p := new(PathPattern)
	p.RawPath = path
	p.path = NewPath(path)

	return p, nil
}

/*
	GetPathMatch gets the PathMatch object describing the match or otherwise
	between this PathPattern and the specified Path.
*/
func (p *PathPattern) GetPathMatch(path *Path) *PathMatch {

	pathMatch := new(PathMatch)
	pathMatch.Matches = true

	// if this is the root catch all, just return yes
	if p.RawPath == segmentCatchAll {
		return pathMatch
	}

	pathMatch.Parameters = make(objects.Map)

	checkSegments := p.path.Segments()
	pathSegments := path.Segments()

	// make sure the segments match in length, or there is a catchall
	// at the end of the check path
	if len(checkSegments) < len(pathSegments) {

		// is the last segment a catchall?
		if getSegmentType(checkSegments[len(checkSegments)-1]) != segmentTypeCatchall {
			return PathDoesntMatch
		}

	} else if len(checkSegments) > len(pathSegments) {

		if getSegmentType(checkSegments[len(checkSegments)-1]) != segmentTypeDynamicOptional {
			return PathDoesntMatch
		}

	}

	// check each segment
	for segmentIndex, checkSegment := range checkSegments {

		switch getSegmentType(checkSegment) {
		case segmentTypeLiteral:

			pathSegment := pathSegments[segmentIndex]
			if strings.ToLower(pathSegment) != strings.ToLower(checkSegment) {
				return PathDoesntMatch
			}

		case segmentTypeDynamic, segmentTypeDynamicOptional:

			if segmentIndex < len(pathSegments) {
				pathMatch.Parameters[cleanSegmentName(checkSegment)] = pathSegments[segmentIndex]
			}

		}

	}

	return pathMatch

}
