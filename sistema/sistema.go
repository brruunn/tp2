package sistema

// SistemaConsultasDeVuelos define la interfaz para el sistema de consultas de vuelos.
type SistemaConsultasDeVuelos interface {
	AgregarArchivo(nombre string)
	VerTablero(k int, modo, desde, hasta string)
	InfoVuelo(codigo string)
	PrioridadVuelos(k int)
	SiguienteVuelo(origen, destino, fecha string)
	Borrar(desde, hasta string)
}
