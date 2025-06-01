package main // Paquete principal de la aplicación

import (
	"context"
	"fmt" // Para imprimir en la consola
	"github-tracker/github-tracker/models"
	"github-tracker/github-tracker/repository"
	"github-tracker/github-tracker/repository/entity"
	"io"       // Para leer el cuerpo de la petición
	"net/http" // Para manejar el servidor HTTP
	"time"

	"github.com/gorilla/mux" // Librería externa para manejar rutas de forma más avanzada
)

// postHandler maneja las peticiones POST que llegan a /hello
func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a POST request!") // Imprime un mensaje cuando se recibe una petición
	defer r.Body.Close()                    // Asegura que el cuerpo de la petición se cierre después de leerlo
	// Lee todo el cuerpo de la petición
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading the request") // Si hay error al leer, lo muestra
		return
	}
	// Imprime el contenido recibido en el cuerpo de la petición
	fmt.Println(string(body))
}

// insertGitHubWebhook toma los datos de un webhook de GitHub y los inserta como un registro de commit en la base de datos.
func insertGitHubWebhook(
	ctx context.Context, // Contexto para manejo de cancelación, timeout o traceo en la operación
	repo repository.Commit, // Interfaz del repositorio que define el método Insert para guardar commits
	webhook models.GitHubWebhook, // Estructura con los datos deserializados del webhook recibido (modelo del paquete models)
	body string, // Cadena con el JSON original del webhook (se guarda como Payload)
	createdTime time.Time, // Marca de tiempo a usar para CreatedAt y UpdatedAt
) error {
	// Se construye una entidad Commit usando los datos del webhook recibido.
	// Esta estructura será insertada en la base de datos.
	commit := entity.Commit{
		RepoName:       webhook.Repository.FullName,        // Nombre completo del repo (e.g., "usuario/repositorio")
		CommitID:       webhook.HeadCommit.ID,              // SHA del commit
		CommitMessage:  webhook.HeadCommit.Message,         // Mensaje del commit
		AuthorUsername: webhook.HeadCommit.Author.Username, // Usuario de GitHub que hizo el commit
		AuthorEmail:    webhook.HeadCommit.Author.Email,    // Email del autor del commit
		Payload:        body,                               // Cuerpo completo del webhook (JSON) como string
		CreatedAt:      createdTime,                        // Fecha de creación (posiblemente now)
		UpdatedAt:      createdTime,                        // Fecha de actualización (igual que createdTime al inicio)
	}

	// Inserta el commit en la base de datos usando el repositorio.
	// El método Insert se espera que maneje errores internamente (validaciones, SQL, etc.)
	err := repo.Insert(ctx, &commit)

	// Devuelve el error si ocurrió (puede ser nil si todo salió bien).
	return err
}

func main() {
	// Crea un nuevo router usando la librería Gorilla Mux
	router := mux.NewRouter()
	// Asocia la ruta "/hello" con la función postHandler, solo para métodos POST
	router.HandleFunc("/hello", postHandler).Methods("POST")
	// Informa que el servidor está corriendo
	fmt.Println("Server listening on port 8080")
	// Inicia el servidor HTTP en el puerto 8080 usando el router definido
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		// Si ocurre un error al iniciar el servidor, lo imprime
		fmt.Println(err.Error())
	}
}
