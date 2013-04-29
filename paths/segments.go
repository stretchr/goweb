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

	if strings.HasPrefix(segment, segmentDynamicPrefix) && strings.HasSuffix(segment, segmentDynamicSuffix) {
		return segmentTypeDynamic
	}

	if strings.HasPrefix(segment, segmentOptionalDynamicPrefix) && strings.HasSuffix(segment, segmentOptionalDynamicSuffix) {
		return segmentTypeDynamicOptional
	}

	if segment == segmentWildcard {
		return segmentTypeWildcard
	}

	if segment == segmentCatchAll {
		return segmentTypeCatchall
	}

	return segmentTypeLiteral

}

func cleanSegmentName(segment string) string {
	return strings.Trim(segment, "{}[]")
}
