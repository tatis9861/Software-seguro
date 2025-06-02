// Paquete principal para la aplicación
package main

// Importa los paquetes necesarios
import (
	"context"  // Proporciona el tipo de contexto para manejar tiempos de espera, cancelación, etc.
	"net/http" // Proporciona constantes como http.StatusOK

	"github.com/aws/aws-lambda-go/events" // Proporciona tipos para eventos de AWS como API Gateway
	"github.com/aws/aws-lambda-go/lambda" // Permite ejecutar funciones Lambda
)

// Función que maneja las solicitudes entrantes desde API Gateway
func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		IsBase64Encoded: false,                    // Indica que el cuerpo de la respuesta no está codificado en base64
		StatusCode:      http.StatusOK,            // Devuelve un código HTTP 200 OK
		Body:            "hello from get-metrics", // Texto plano como respuesta
	}, nil // No se devuelve ningún error
}

// Punto de entrada principal de la función Lambda
func main() {
	lambda.Start(handleRequest) // Inicia la función Lambda usando handleRequest como controlador
}
