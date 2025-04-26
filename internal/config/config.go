package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AmountTolerance   float64 `json:"amount_tolerance"`
	DateToleranceDays int     `json:"date_tolerance_days"`
	DecimalSeparator  string  `json:"decimal_separator"`
}

var AppConfig Config

var defaultConfig = Config{
	AmountTolerance:   0.10,
	DateToleranceDays: 10,
	DecimalSeparator:  ".",
}

func LoadConfig(path string) error {
	file, err := os.Open(path)

	if os.IsNotExist(err) {
		if err = createDefaultConfig(path); err != nil {
			return fmt.Errorf("error creando config por defecto: %w", err)
		}

		fmt.Println("Config no encontrada. Se creó un archivo de configuración por defecto.")

		file, err = os.Open(path)
		if err != nil {
			return fmt.Errorf("error abriendo config creada: %w", err)
		}

	} else if err != nil {
		return fmt.Errorf("error abriendo config: %w", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		return fmt.Errorf("error leyendo config: %w", err)
	}

	return nil
}

func createDefaultConfig(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Para que quede más lindo el JSON
	return encoder.Encode(defaultConfig)
}
