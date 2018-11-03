package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDigit(t *testing.T) {
	assert.Equal(t, false, IsDigit('x'))
	assert.Equal(t, true, IsDigit('4'))
}
