package paths

import (
	"strings"
)

/*
  Path represents the path segment of a URL.
*/
type Path struct {
	RawPath string

	segments  []string
	extension string
}

/*
  NewPath creates a new Path with the given raw string.
*/
func NewPath(rawPath string) *Path {

	p := new(Path)
	p.RawPath = cleanPath(rawPath)
	p.Segments()
	return p

}

/*
	cleanPath cleans returns the cleaned version of the specified path.
*/
func cleanPath(path string) string {
	return strings.TrimRight(strings.TrimLeft(path, PathSeperator), PathSeperator)
}

/*
  Segments gets the segments for this path broken up by the PathSeparator.
*/
func (p *Path) Segments() []string {

	if len(p.segments) == 0 {

		p.segments = strings.Split(p.RawPath, "/")

		// handle the extension in the last segment
		lastSegment := p.segments[len(p.segments)-1]
		lastDot := strings.LastIndex(lastSegment, FileExtensionSeparator)
		if lastDot > -1 {
			extsegs := strings.Split(lastSegment, FileExtensionSeparator)
			p.segments[len(p.segments)-1] = extsegs[0]
			p.extension = extsegs[1]
		}

	}

	return p.segments
}
