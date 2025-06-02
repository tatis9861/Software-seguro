# Crea un bucket de S3 llamado "ciberseguridad-lambda-platzi"
resource "aws_s3_bucket" "lambda_bucket" {
  bucket = "ciberseguridad-lambda-platzi"  # Nombre único del bucket
}

# Habilita el versionamiento para el bucket creado arriba
resource "aws_s3_bucket_versioning" "lambda_bucket" {
  bucket = aws_s3_bucket.lambda_bucket.id  # ID del bucket al que se aplicará el versionamiento

  versioning_configuration {
    status = "Enabled"  # Activa el versionamiento para este bucket
  }
}

# Define una salida que mostrará el nombre del bucket creado
output "lambda_bucket" {
  value = aws_s3_bucket.lambda_bucket.bucket  # Retorna el nombre del bucket como output
}