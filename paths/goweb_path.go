package paths

import (
	"strings"
)

/*
  GowebPath represents a path that can contain special
  matching segments.

  Valid paths include:

    /literal
    /{placeholder}
    /[optional placeholder]
    /people/* - matches like a placeholder but doesn't care what it is
    /something/... - Matches the start plus anything after it
    /people/{id:int} - specific types
    /people/{id:string}
*/
type GowebPath struct {
	RawPath string
	path    *Path
}

func NewGowebPath(path string) (*GowebPath, error) {

	p := new(GowebPath)
	p.RawPath = path
	p.path = NewPath(path)

	return p, nil
}

func (p *GowebPath) GetPathMatch(path *Path) *PathMatch {

	pathMatch := new(PathMatch)
	pathMatch.Matches = true
	pathMatch.Parameters = make(map[string]string)

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
