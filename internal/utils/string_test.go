package utils_test

import (
	"os"
	"testing"

	"github.com/andresfontan90/afip-compare/internal/config"
	"github.com/andresfontan90/afip-compare/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestStringToNumber(t *testing.T) {
	assert := assert.New(t)

	t.Run("StringToNumber OK", func(_ *testing.T) {
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

		expected := 1234.56

		value, err := utils.StringToNumber("1234,56")

		assert.NoError(err)
		assert.Equal(expected, value)
	})

}
