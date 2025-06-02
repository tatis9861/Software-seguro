# Llama al módulo "lambda" desde el subdirectorio ./lambda
module "lambda" {
  source = "./lambda"  # Ruta local donde se encuentra el módulo

  # Pasa el nombre del bucket S3 donde están los archivos ZIP de las Lambdas
  lambda_bucket = var.lambda_bucket

  # Pasa el ARN del rol IAM que utilizarán las funciones Lambda dentro del módulo
  repo_collector_role_arn = var.repo_collector_role_arn
}
