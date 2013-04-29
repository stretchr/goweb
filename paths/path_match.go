package paths

type PathMatch struct {
	Matches    bool
	Parameters map[string]string
}

var PathDoesntMatch *PathMatch = new(PathMatch)
