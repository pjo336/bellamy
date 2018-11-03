package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWhitespace(t *testing.T) {
	assert.Equal(t, true, IsWhitespace(' '))
	assert.Equal(t, true, IsWhitespace('\n'))
	assert.Equal(t, true, IsWhitespace('\t'))
	assert.Equal(t, true, IsWhitespace('\r'))
	assert.Equal(t, false, IsWhitespace('4'))
}
