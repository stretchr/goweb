package paths

import (
	"strings"
)

type segmentType int

const (
	segmentTypeLiteral segmentType = iota
	segmentTypeDynamic
	segmentTypeDynamicOptional
	segmentTypeWildcard
	segmentTypeCatchall
)

func getSegmentType(segment string) segmentType {

	if strings.HasPrefix(segment, "{") && strings.HasSuffix(segment, "}") {
		return segmentTypeDynamic
	}

	if strings.HasPrefix(segment, "[") && strings.HasSuffix(segment, "]") {
		return segmentTypeDynamicOptional
	}

	if segment == "*" {
		return segmentTypeWildcard
	}

	if segment == "..." {
		return segmentTypeCatchall
	}

	return segmentTypeLiteral

}
