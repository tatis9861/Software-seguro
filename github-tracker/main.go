package main // Paquete principal de la aplicación

import (
	"fmt"      // Para imprimir en la consola
	"io"       // Para leer el cuerpo de la petición
	"net/http" // Para manejar el servidor HTTP

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
