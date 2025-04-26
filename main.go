package main

import (
	"fmt"

	"os"

	"github.com/andresfontan90/afip-compare/internal/process"
	"github.com/andresfontan90/afip-compare/internal/utils"
)

func main() {
	for {
		fmt.Println("")
		fmt.Println("**** Cruce de información entre dos hojas de excel ****")
		fmt.Println("")
		fmt.Println("1. Iniciar proceso")
		fmt.Println("2. Ver instrucciones")
		fmt.Println("3. Salir")
		fmt.Print("Selecciona una opción: ")
		fmt.Println("")

		var option int
		_, err := fmt.Scanln(&option)

		utils.ClearConsole()

		if err != nil {
			fmt.Println("Opción no válida. Intente de nuevo.")
			continue
		}

		switch option {
		case 1:
			err := startComparison()
			if err != nil {
				fmt.Println("Error en el proceso ", err.Error())
			} else {
				fmt.Println("Proceso finalizado correctamente")
			}

			continue
		case 2:
			showInstructions()
		case 3:
			fmt.Println("Saliendo del programa...")
			os.Exit(0)
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

func startComparison() error {
	if err := process.Process(); err != nil {
		return err
	}
	return nil
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

/*
func displayInstructions() {
	fmt.Println("")
	fmt.Println("*** Cruce de información entre dos hojas de excel ***")
	fmt.Println("")
	fmt.Println("INSTRUCCIONES")
	fmt.Println("Tenga en cuenta que el proceso fallará si el mismo no encuentra las columnas que se detallan a continuación (el orden no importa)")
	fmt.Println("HOJA 1: 'cuit informante' | 'fecha comprobante' | 'punto' | 'nro. comprobante' | 'importe total operacion' | 'importe neto' | 'impuesto liquidado' | 'fuente'")
	fmt.Println("HOJA 2: 'cuit emisor' | 'fecha emision comprobante' | 'punto' | 'cpbte desde' | 'importe total' | 'importe neto' | 'impuesto liquidado' | 'fuente'")
	fmt.Println("")
	fmt.Println("las columnas de fecha deben tener formato de fecha")
	fmt.Println("las columnas de montos deben tener formato númerico con dos decimales")
	fmt.Println("la columna de nro de comprobante debe tener formato númerico con cero decimales")
	fmt.Println("")
	fmt.Println("la tolerancia para cruzar los montos es de 0.10$")
	fmt.Println("la tolerancia para cruzar las fechas es de 10 días")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("1 - Seleccionar el archivo excel a procesar")
	fmt.Println("2 - Recibirá un mensaje al finalizar el proceso")
	fmt.Println("")
	fmt.Println("Presiona una tecla para comenzar")
}

func startProcess() error {
	if err := process.Process(); err != nil {
		return err
	}
	return nil
}

func showCompletionMessage() {
	dialog.Alert("Proceso finalizado correctamente")
}

func showErrorMessage(err error) {
	dialog.Alert("Archivos generados con inconsistencias.\nError: %s", err.Error())
}

func main() {
	displayInstructions()

	// Esperamos que el usuario presione una tecla para iniciar el proceso
	fmt.Scanln()

	// Iniciar el proceso
	if err := startProcess(); err != nil {
		showErrorMessage(err)
		os.Exit(1)
	}

	// Mostrar mensaje de éxito
	showCompletionMessage()
}
*/
