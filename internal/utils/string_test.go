package utils_test

import (
	"testing"

	"github.com/andresfontan90/afip-compare/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestPRocess(t *testing.T) {
	value, err := utils.StringToNumber("1.234,56")
	assert.Nil(t, value)
	assert.Nil(t, err)
}
