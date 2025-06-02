# Obtiene información sobre el archivo ZIP de la Lambda desde un bucket S3
data "aws_s3_object" "handle_github_webhook" {
  bucket = var.lambda_bucket                      # Nombre del bucket S3 definido en una variable
  key    = "handle-github-webhook.zip"            # Clave del objeto ZIP en el bucket
}

# Define la función Lambda llamada "handle-github-webhook"
resource "aws_lambda_function" "handle_github_webhook" {
  function_name    = "handle-github-webhook"       # Nombre de la función Lambda
  handler          = "bootstrap"                   # Punto de entrada del ejecutable (Go requiere "bootstrap")
  runtime          = "provided.al2"                # Usa el runtime personalizado basado en Amazon Linux 2
  s3_bucket        = var.lambda_bucket             # Bucket donde está el código ZIP
  timeout          = 300                           # Tiempo máximo de ejecución: 300 segundos
  s3_key           = "handle-github-webhook.zip"   # Nombre del archivo ZIP en el bucket
  role             = var.repo_collector_role_arn   # Rol IAM que Lambda utilizará para permisos
  source_code_hash = data.aws_s3_object.handle_github_webhook.version_id
  # Usamos el `version_id` para que Terraform detecte cambios en el ZIP automáticamente

  # Variables de entorno que la Lambda puede usar
  environment {
    variables = {
      DEMO = "DEMO"  # Variable de entorno de ejemplo (puedes agregar más según necesites)
    }
  }
}

# Output: ARN de invocación de la Lambda, útil para integraciones (como API Gateway o EventBridge)
output "handle_github_webhook_invoke_arn" {
  value = aws_lambda_function.handle_github_webhook.invoke_arn
}

# Output: Nombre de la función Lambda
output "handle_github_webhook_lambda_name" {
  value = aws_lambda_function.handle_github_webhook.function_name
}
