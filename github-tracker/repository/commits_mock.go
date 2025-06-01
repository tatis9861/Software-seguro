package repository

import (
	"context"
	"github-tracker/github-tracker/repository/entity"

	"github.com/stretchr/testify/mock" // Librería que facilita la creación de mocks para testing
)

// MockCommit es una implementación mock del repositorio Commit.
// Se utiliza en pruebas unitarias para simular el comportamiento del repositorio real.
type MockCommit struct {
	*mock.Mock // Embebe el struct Mock de testify, que contiene métodos para definir comportamientos simulados.
	Commit     // También embebe la interfaz original, por compatibilidad si se necesita.
}

// Insert es una implementación simulada del método Insert del repositorio Commit.
// Llama internamente al método Called(...) de testify para capturar la invocación y retornar valores preconfigurados.
func (m MockCommit) Insert(ctx context.Context, commit *entity.Commit) (err error) {
	// Llama a Called con los argumentos reales. Este método permite testear si fue llamado y con qué parámetros.
	results := m.Called(ctx, commit)

	// Devuelve el primer valor como error (índice 0)
	return results.Error(0)
}

// GetCommitByAuthorEmail es la versión simulada del método que devuelve commits según un email.
// Se usa en pruebas para evitar ir a la base de datos.
func (m MockCommit) GetCommitByAuthorEmail(ctx context.Context, email string) (commits []entity.Commit, err error) {
	// Llama a Called con los argumentos esperados en el test
	results := m.Called(ctx, email)

	// Devuelve una lista de commits simulados (índice 0) y un posible error (índice 1)
	return results.Get(0).([]entity.Commit), results.Error(1)
}
