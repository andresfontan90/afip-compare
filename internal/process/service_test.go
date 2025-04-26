package process_test

import (
	"os"
	"testing"

	"github.com/andresfontan90/afip-compare/internal/config"
	"github.com/andresfontan90/afip-compare/internal/process"
	"github.com/stretchr/testify/assert"
)

func TestPRocess(t *testing.T) {
	assert := assert.New(t)

	tmpFile, err := os.CreateTemp("", "temp-config-*.json")
	assert.NoError(err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(`{
			"amount_tolerance": 0.10,
			"date_tolerance_days": 10,
			"csv_separator": ";",
			"decimal_separator": "."
		}`)

	assert.NoError(err)
	tmpFile.Close()

	config.LoadConfig(tmpFile.Name())

	err = process.Process()
	assert.Nil(t, err)
}
