package handlers

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestMatcherFuncDecision(t *testing.T) {

	assert.Equal(t, MatcherFuncDecision(-1), DontCare)
	assert.Equal(t, MatcherFuncDecision(0), NoMatch)
	assert.Equal(t, MatcherFuncDecision(1), Match)

}
