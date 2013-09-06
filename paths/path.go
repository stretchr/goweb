package paths

import (
	"fmt"
	"path"
	"strings"
)

// Path represents the path segment of a URL.
type Path struct {
	// RawPath is the original raw string of this path.
	RawPath string

	// segments holds the path segments.
	segments []string

	// extension holds the file extension of this path.
	extension string
}

// NewPath creates a new Path with the given raw string.
func NewPath(rawPath string) *Path {

	p := new(Path)
	p.RawPath = cleanPath(rawPath)
	p.Segments()
	return p

}

// cleanPath cleans returns the cleaned version of the specified path.
func cleanPath(rawPath string) string {
	return strings.TrimRight(strings.TrimLeft(path.Clean(rawPath), PathSeperator), PathSeperator)
}

// PathFromSegments turns the arguments into a path string.
func PathFromSegments(segments ...interface{}) string {

	var theStrings []string = make([]string, len(segments))

	for segIndex, seg := range segments {
		theStrings[segIndex] = fmt.Sprintf("%v", seg)
	}

	return strings.Join(theStrings, PathSeperator)
}

// Segments gets the segments for this path broken up by the PathSeparator.
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

// RealFilePath gets the real file path by assuming the current path is
// the public prefix, the urlPath is the actual request and the systemPath
// is the physical location where those files live.
func (p *Path) RealFilePath(systemPath, urlPath string) string {

	urlPath = cleanPath(urlPath)
	if strings.HasPrefix(urlPath, p.RawPath) {
		urlPath = strings.TrimPrefix(urlPath, p.RawPath)
	} else {
		panic(fmt.Sprintf("goweb.paths.Path: Cannot use RealFilePath when the urlPath doesn't start with the path in the first place. \"%s\" doesn't start with \"%s\".", urlPath, p.RawPath))
	}

	return fmt.Sprintf("%s%s", systemPath, urlPath)
}
