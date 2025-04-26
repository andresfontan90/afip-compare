package main

import (
	"fmt"

	"os"

	"github.com/andresfontan90/afip-compare/internal/config"
	"github.com/andresfontan90/afip-compare/internal/process"
	"github.com/andresfontan90/afip-compare/internal/utils"
)

func main() {
	utils.ClearConsole()

	if err := config.LoadConfig(config.ConfigFileName); err != nil {
		fmt.Println("Error cargando configuración inicial:", err)
		os.Exit(1)
	}

	for {
		showMenu()

		var option int
		_, err := fmt.Scanln(&option)

		utils.ClearConsole()

		if err != nil {
			showInvalidOption()
			continue
		}

		switch option {
		case 1:
			err := startProcess()
			fmt.Println("")
			fmt.Println("")
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Proceso finalizado correctamente")
			}
		case 2:
			showInstructions()
		case 3:
			exitProgram()
		default:
			showInvalidOption()
		}
	}
}

func startProcess() error {
	if err := process.Process(); err != nil {
		return err
	}
	return nil
}

func exitProgram() {
	fmt.Println("Saliendo del programa...")
	fmt.Println("")
	os.Exit(0)
}

func showMenu() {
	fmt.Println("")
	fmt.Println("╔══════════════════════════════════════════════════╗")
	fmt.Println("║         Cruce de información entre hojas         ║")
	fmt.Println("╠══════════════════════════════════════════════════╣")
	fmt.Println("║ 1. Iniciar proceso                               ║")
	fmt.Println("║ 2. Ver instrucciones                             ║")
	fmt.Println("║ 3. Salir                                         ║")
	fmt.Println("╚══════════════════════════════════════════════════╝")
	fmt.Print("Selecciona una opción: ")
	fmt.Println("")
}

func showInvalidOption() {
	fmt.Println("Opción no válida. Intente de nuevo.")
}

func showInstructions() {
	fmt.Println("**** INSTRUCCIONES ****")
	fmt.Println("")
	fmt.Println("Tenga en cuenta que el proceso fallará si el mismo no encuentra las columnas que se detallan a continuación (el orden no importa)")
	fmt.Println("HOJA 1: 'cuit informante' | 'fecha comprobante' | 'punto' | 'nro. comprobante' | 'importe total operacion' | 'importe neto' | 'impuesto liquidado' | 'fuente'")
	fmt.Println("HOJA 2: 'cuit emisor' | 'fecha emision comprobante' | 'punto' | 'cpbte desde' | 'importe total' | 'importe neto' | 'impuesto liquidado' | 'fuente'")
	fmt.Println("")
	fmt.Println("Las columnas de fecha deben tener FORMATO FECHA")
	fmt.Println("Las columnas de montos deben tener FORMATO NÚMERICO con dos decimales")
	fmt.Println("La columna de nro de comprobante debe tener FORMATO NÚMERICO con cero decimales")
	fmt.Println("")
	fmt.Println("La tolerancia para cruzar los montos es de 0.10$")
	fmt.Println("La tolerancia para cruzar las fechas es de 10 días")
	fmt.Println()
}
