package main

import (
	"fmt"

	"os"

	"github.com/andresfontan90/afip-compare/internal/process"
	"tawesoft.co.uk/go/dialog"
)

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
