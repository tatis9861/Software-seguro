package main

// Importa el paquete de testing de Go y la librería testify
import (
	// Paquetes estándar de Go
	"context"       // Manejo de contextos: permite cancelar operaciones, definir timeouts, y propagar información entre capas.
	"encoding/json" // Para serializar/deserializar estructuras a JSON (en este caso, convertir el webhook en string).
	"testing"       // Paquete estándar de Go para escribir y ejecutar pruebas unitarias.
	"time"          // Permite trabajar con fechas, como generar timestamps para `CreatedAt` y `UpdatedAt`.

	// Paquetes propios del proyecto
	"github-tracker/github-tracker/models"            // Contiene estructuras relacionadas con GitHub, como el webhook (`GitHubWebhook`), commits, repos, etc.
	"github-tracker/github-tracker/repository"        // Contiene interfaces y mocks del repositorio (como `MockCommit`) usados para pruebas unitarias.
	"github-tracker/github-tracker/repository/entity" // Contiene las entidades del dominio usadas para persistencia, como `entity.Commit`.

	// Paquetes externos (de terceros)
	"github.com/stretchr/testify/mock"    // Proporciona herramientas para crear y usar mocks fácilmente (definir comportamiento simulado, verificar llamadas, etc.).
	"github.com/stretchr/testify/require" // Librería de aserciones: permite validar condiciones en tests de forma clara y detener el test si falla.
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

func TestInsert(t *testing.T) {
	// Inicializa el asertor de testify. `c` nos permite hacer validaciones con `require`.
	c := require.New(t)

	// Simula un webhook de GitHub con datos ficticios.
	webhook := models.GitHubWebhook{
		Repository: models.Repository{
			FullName: "tatis9861/Software-seguro", // Nombre del repositorio
		},
		HeadCommit: models.Commit{
			ID:      "9da3ed5d641d46dd1401d0768bc9dde90e86e1cb",       // SHA del commit
			Message: "Add sample code test for handle-github-webhook", // Mensaje del commit
			Author: models.CommitUser{
				Email:    "tatis9861@gmail.com", // Email del autor
				Username: "tatis9861",           // Username del autor
			},
		},
	}

	// Convierte el webhook a JSON (string) para guardarlo como "payload"
	body, err := json.Marshal(webhook)
	c.NoError(err) // Asegura que no haya error al serializar

	// Define el timestamp que se usará como CreatedAt y UpdatedAt
	createdTime := time.Now()

	// Crea un mock de CommitRepository usando testify/mock
	m := mock.Mock{}
	mockCommit := repository.MockCommit{Mock: &m}

	// Crea la entidad que se espera insertar en la base de datos
	commit := entity.Commit{
		RepoName:       webhook.Repository.FullName,
		CommitID:       webhook.HeadCommit.ID,
		CommitMessage:  webhook.HeadCommit.Message,
		AuthorUsername: webhook.HeadCommit.Author.Username,
		AuthorEmail:    webhook.HeadCommit.Author.Email,
		Payload:        string(body), // Guarda el JSON completo como string
		CreatedAt:      createdTime,
		UpdatedAt:      createdTime,
	}

	// Crea el contexto de ejecución (puede llevar timeout, cancelación, etc.)
	ctx := context.Background()

	// Define la expectativa: el método Insert será llamado con `ctx` y `commit`, y devolverá nil (sin error)
	mockCommit.On("Insert", ctx, &commit).Return(nil)

	// Ejecuta la función real que se está probando, con el mock inyectado
	err = insertGitHubWebhook(ctx, mockCommit, webhook, string(body), createdTime)

	// Verifica que no haya ocurrido error al ejecutar la función
	c.NoError(err)

	// Verifica que se cumplieron todas las expectativas configuradas en el mock
	m.AssertExpectations(t)
}
