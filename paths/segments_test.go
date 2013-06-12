package paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSegmentType(t *testing.T) {

	assert.Equal(t, segmentType(segmentTypeLiteral), getSegmentType("people"))
	assert.Equal(t, segmentType(segmentTypeDynamic), getSegmentType("{id}"))
	assert.Equal(t, segmentType(segmentTypeDynamicOptional), getSegmentType("[id]"))
	assert.Equal(t, segmentType(segmentTypeWildcard), getSegmentType(segmentWildcard))
	assert.Equal(t, segmentType(segmentTypeCatchall), getSegmentType(segmentCatchAll))

}

func TestCleanSegmentName(t *testing.T) {

	assert.Equal(t, "id", cleanSegmentName("id"))
	assert.Equal(t, "id", cleanSegmentName("{id}"))
	assert.Equal(t, "id", cleanSegmentName("[id]"))

}
