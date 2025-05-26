package vuelo

import "time"

// LeerArchivo lee un CSV y retorna una lista de Vuelos.
func LeerArchivo(nombre string) ([]*Vuelo, error) {
	// ...
}

// ParsearFecha convierte una cadena en time.Time.
func ParsearFecha(fechaStr string) (time.Time, error) {
	// ...
}

// ParsearVuelo convierte una línea CSV en un struct Vuelo.
func ParsearVuelo(registro []string) (*Vuelo, error) {
	// ...
}
