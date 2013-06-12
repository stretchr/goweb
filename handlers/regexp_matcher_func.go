package handlers

import (
	"github.com/stretchr/goweb/context"
	"regexp"
)

// RegexPath returns a MatcherFunc that mathces the path based on the specified
// Regex pattern.
//
// For more information, see the goweb.RegexPath shortcut function.
func RegexPath(regexpPattern string) MatcherFunc {

	// compile the regex
	regex, regexErr := regexp.Compile(regexpPattern)

	// return a MatcherFunc that will check the regex against the path
	// and return the decisive MatcherFuncDecision.
	return func(ctx context.Context) (MatcherFuncDecision, error) {

		// if the regex fails, just return the error
		if regexErr != nil {
			return DontCare, regexErr
		}

		var decision MatcherFuncDecision

		if regex.MatchString(ctx.Path().RawPath) {
			decision = Match
		} else {
			decision = NoMatch
		}

		return decision, nil
	}

}
