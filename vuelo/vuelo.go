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

// LeerArchivo lee un CSV y retorna una lista de Vuelos.
func LeerArchivo(nombre string) ([]*Vuelo, error) {
	// ...
}

// ParsearVuelo convierte una línea CSV en un struct Vuelo.
func ParsearVuelo(registro []string) (*Vuelo, error) {
	// ...
}
