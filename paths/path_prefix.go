package paths

import (
	stewstrings "github.com/stretchr/stew/strings"
	"reflect"
	"strings"
)

const controllerSuffix string = "Controller"

// PathPrefixForClass gets the path prefix by reflecting on the name of the
// class of the specified object.
//
// E.g.
//     LovelyPeopleController == "lovely-people"
func PathPrefixForClass(controller interface{}) string {

	name := reflect.TypeOf(controller).Elem().Name()

	// trim off "Controller" suffix
	if strings.HasSuffix(name, controllerSuffix) {
		name = name[0 : len(name)-len(controllerSuffix)]
	}

	segments := stewstrings.SplitByCamelCase(name)
	name = strings.ToLower(strings.Join(segments, "-"))
	return name
}
