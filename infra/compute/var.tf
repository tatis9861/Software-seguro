# Define una variable llamada "lambda_bucket" de tipo string
# Esta variable contendrá el nombre del bucket S3 donde están almacenados los archivos ZIP de las funciones Lambda
variable "lambda_bucket" {
  type = string
}

# Define una variable llamada "repo_collector_role_arn" de tipo string
# Esta variable contendrá el ARN del rol IAM que las funciones Lambda usarán para ejecutarse
variable "repo_collector_role_arn" {
  type = string
}
