package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/andresfontan90/afip-compare/internal/config"
	"github.com/shopspring/decimal"
)

func StringToNumber(value string) (float64, error) {
	var valueStr string

	if strings.EqualFold(config.AppConfig.DecimalSeparator, ".") {
		valueStr = strings.Replace(value, ",", ".", 1)
	}

	valueF, err := decimal.NewFromString(valueStr)
	if err != nil {
		return 0, fmt.Errorf("error parseando el importe %s", value)
	}
	result, _ := valueF.Float64()
	return result, nil
}

func IsEmptyString(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

func NormalizeString(value string) string {
	trimedValue := strings.ToLower(strings.TrimSpace(value))
	replacer := strings.NewReplacer("á", "a", "é", "e", "í", "i", "ó", "o", "ú", "u")
	return replacer.Replace(trimedValue)
}

func StringToDate(value string) (time.Time, error) {
	result, err := time.Parse("02/01/2006", value)
	if err == nil {
		return result, nil
	}
	result, err = time.Parse("01/02/2006", value)
	if err == nil {
		return result, nil
	}
	result, err = time.Parse(time.DateOnly, value)
	if err == nil {
		return result, nil
	}

	return time.Time{}, errors.New("error parsing date")

}
