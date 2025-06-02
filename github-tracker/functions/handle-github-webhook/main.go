// Paquete principal de la aplicación
package main

// Importa los paquetes necesarios
import (
	"context"  // Proporciona contexto para manejar tiempo de ejecución, cancelación, etc.
	"net/http" // Proporciona constantes para estados HTTP como http.StatusOK

	"github.com/aws/aws-lambda-go/events" // Tipos definidos por AWS para eventos como los de API Gateway
	"github.com/aws/aws-lambda-go/lambda" // Permite que esta función se ejecute como una Lambda
)

// Función que maneja las solicitudes HTTP provenientes de API Gateway
func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		IsBase64Encoded: false,                                    // Indica que la respuesta no está codificada en base64
		StatusCode:      http.StatusOK,                            // Devuelve un HTTP 200 OK
		Body:            "hello from handle-github-notifications", // Cuerpo de la respuesta
	}, nil // No se devuelve error
}

// Función principal de entrada para Lambda
func main() {
	lambda.Start(handleRequest) // Inicia la ejecución del Lambda usando handleRequest como handler
}
