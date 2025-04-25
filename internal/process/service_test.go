package process_test

import (
	"testing"

	"github.com/andresfontan90/afip-compare/internal/process"
	"github.com/stretchr/testify/assert"
)

func TestPRocess(t *testing.T) {
	err := process.Process()
	assert.Nil(t, err)
}
