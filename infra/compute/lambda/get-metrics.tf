# Obtiene información sobre un objeto existente en S3 (el archivo ZIP de la Lambda)
data "aws_s3_object" "get_metrics" {
  bucket = var.lambda_bucket         # Nombre del bucket S3 donde está el archivo ZIP
  key    = "get-metrics.zip"         # Nombre del archivo ZIP en el bucket
}

# Define una función Lambda llamada "get-metrics"
resource "aws_lambda_function" "get_metrics" {
  function_name    = "get-metrics"               # Nombre visible de la función Lambda
  handler          = "bootstrap"                 # Punto de entrada del binario (para runtimes personalizados)
  runtime          = "provided.al2"              # Runtime personalizado basado en Amazon Linux 2
  s3_bucket        = var.lambda_bucket           # Bucket donde se encuentra el ZIP
  timeout          = 300                         # Tiempo máximo de ejecución: 300 segundos (5 minutos)
  s3_key           = "get-metrics.zip"           # Nombre del archivo ZIP en el bucket
  role             = var.repo_collector_role_arn # ARN del rol IAM que Lambda usará para ejecutarse
  source_code_hash = data.aws_s3_object.get_metrics.version_id
  # 👆 Esto ayuda a que Terraform detecte cambios en el ZIP al usar el version_id del archivo en S3

  # 🌱 Variables de entorno para la función Lambda
  environment {
    variables = {
      DEMO = "DEMO"  # Puedes agregar aquí cualquier variable de entorno necesaria para tu Lambda
    }
  }
}

# Output: ARN de invocación de la Lambda (útil para integraciones, como API Gateway)
output "get_metrics_invoke_arn" {
  value = aws_lambda_function.get_metrics.invoke_arn
}

# Output: Nombre de la función Lambda
output "get_metrics_lambda_name" {
  value = aws_lambda_function.get_metrics.function_name
}
