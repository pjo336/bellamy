package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLetter(t *testing.T) {
	assert.Equal(t, true, IsLetter('x'))
	assert.Equal(t, false, IsLetter('4'))
}
