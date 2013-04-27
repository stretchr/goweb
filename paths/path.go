package paths

import (
	"strings"
)

/**
  Path represents the path segment of a URL.
*/
type Path struct {
	RawPath string

	segments []string
}

/**
  NewPath creates a new Path with the given raw string.
*/
func NewPath(rawPath string) *Path {
	p := new(Path)
	p.RawPath = rawPath
	return p
}

/**
  Segments gets the segments for this path.
*/
func (p *Path) Segments() ([]string, error) {

	if len(p.segments) == 0 {
		p.segments = strings.Split(p.RawPath, "/")
	}

	return p.segments, nil
}
