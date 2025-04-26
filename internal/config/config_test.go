package config_test

import (
	"os"
	"testing"

	"github.com/andresfontan90/afip-compare/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	assert := assert.New(t)
	t.Run("LoadConfig OK - File does not exist", func(_ *testing.T) {
		err := config.LoadConfig("test-file")
		assert.NoError(err)
	})

	t.Run("LoadConfig Error - Parsing file", func(_ *testing.T) {
		// Temporal file
		tmpFile, err := os.CreateTemp("", "invalid-config-*.json")
		assert.NoError(err)
		defer os.Remove(tmpFile.Name())

		_, err = tmpFile.WriteString("{invalid-json")
		assert.NoError(err)
		tmpFile.Close()

		err = config.LoadConfig(tmpFile.Name())
		assert.ErrorContains(err, "error leyendo config")
	})

	t.Run("LoadConfig OK", func(_ *testing.T) {
		// Temporal file
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

		err = config.LoadConfig(tmpFile.Name())
		assert.NoError(err)
		assert.Equal(0.1, config.AppConfig.AmountTolerance)
		assert.Equal(10, config.AppConfig.DateToleranceDays)
		assert.Equal(".", config.AppConfig.DecimalSeparator)
	})
}
