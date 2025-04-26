package process

import "time"

type Entity struct {
	Cuit   string
	Date   time.Time
	Punto  string
	Comp   string
	Mount  float64
	Neto   float64
	Tax    float64
	Source string
}

const (
	cuitSheet1      = "cuit informante"
	cuitSheet2      = "cuit emisor"
	dateSheet1      = "fecha comprobante"
	dateSheet2      = "fecha emision comprobante"
	puntoSheet1     = "punto"
	puntoSheet2     = "punto"
	compSheet1      = "nro. comprobante"
	compSheet2      = "cpbte desde"
	mountSheet1     = "importe total operacion"
	mountSheet2     = "importe total"
	mountNetoSheet1 = "importe neto"
	mountNetoSheet2 = "importe neto"
	taxSheet1       = "impuesto liquidado"
	taxSheet2       = "impuesto liquidado"
	sourceSheet1    = "fuente"
	sourceSheet2    = "fuente"
)

var (
	columNames     = make(map[string]string)
	fileMap1       = make(map[string][]string)
	fileMap2       = make(map[string][]string)
	toleranceMount = 0.10
	toleranceDate  = 10
	csvSeparator   = ';'
)
