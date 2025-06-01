package main

// Importa el paquete de testing de Go y la librería testify
import (
	"testing" // Paquete estándar para pruebas en Go

	"github.com/stretchr/testify/require" // Testify: librería que ofrece aserciones más legibles
)

func TestDummy(t *testing.T) {
	// Se crea una instancia de require ligada al test actual.
	// Esto permite usar métodos como c.Equal(...) con fallos que detienen inmediatamente la ejecución del test.
	c := require.New(t)

	// Resultado simulado de una operación (puede ser una función, lógica, etc.)
	result := 22

	// Verifica que el valor obtenido (result) sea igual a 22.
	// Si no lo es, el test falla inmediatamente.
	c.Equal(22, result)
}
