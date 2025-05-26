package sistema

import (
	"time"
	TDAHeap "tp2/tdas/cola_prioridad"
	TDADiccionario "tp2/tdas/diccionario"
	"tp2/vuelo"
)

// SistemaConsultasDeVuelos define la interfaz para el sistema de consultas de vuelos.
type SistemaConsultasDeVuelos interface {
	AgregarArchivo(nombre string)
	VerTablero(k int, modo, desde, hasta string)
	InfoVuelo(codigo string)
	PrioridadVuelos(k int)
	SiguienteVuelo(origen, destino, fecha string)
	Borrar(desde, hasta string)
}

// SistemaConsulta es una implementación básica de SistemaConsultasDeVuelos.
type SistemaConsulta struct {
	vuelosPorCodigo *TDADiccionario.Diccionario[string, *vuelo.Vuelo]                                                 // Hash para acceso O(1)
	vuelosPorFecha  *TDADiccionario.DiccionarioOrdenado[time.Time, *vuelo.Vuelo]                                      // ABB por fechas
	prioridadHeap   *TDAHeap.ColaPrioridad[*vuelo.Vuelo]                                                              // Heap de prioridad
	conexiones      *TDADiccionario.Diccionario[string, *TDADiccionario.DiccionarioOrdenado[time.Time, *vuelo.Vuelo]] // Hash de ABBs
}

func CrearSistema() {
	//...
}

func (s *SistemaConsulta) AgregarArchivo(nombre string) {
	// ...
}

func (s *SistemaConsulta) VerTablero(k int, modo, desde, hasta string) {
	// ...
}

func (s *SistemaConsulta) InfoVuelo(codigo string) {
	// ...
}

func (s *SistemaConsulta) PrioridadVuelos(k int) {
	// ...
}

func (s *SistemaConsulta) SiguienteVuelo(origen, destino, fecha string) {
	// ...
}

func (s *SistemaConsulta) Borrar(desde, hasta string) {
	// ...
}
