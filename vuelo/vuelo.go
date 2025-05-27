package vuelo

import (
	"time"
)

// Vuelo representa un registro de vuelo.
type Vuelo struct {
	Codigo      string
	Aerolinea   string
	Origen      string
	Destino     string
	NumeroAvion string
	Prioridad   int
	Fecha       time.Time
	Atraso      string
	TiempoVuelo string
	Cancelacion string
}
