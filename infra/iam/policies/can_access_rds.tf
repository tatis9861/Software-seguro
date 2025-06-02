# ------------------------------------------
# Recurso: Política IAM personalizada para RDS
# ------------------------------------------
resource "aws_iam_policy" "can_access_rds" {
  name        = "can-access-rds"                         # Nombre legible para la política IAM
  path        = "/"                                      # Ruta de la política en la jerarquía de IAM
  description = "Allow manage RDS databases for queries" # Descripción de la política

  # Política IAM definida en formato JSON
  policy = jsonencode(
    {
      Version : "2012-10-17", # Versión del documento de política (requerido por AWS)
      Statement : [
        {
          Effect : "Allow", # Permite las acciones especificadas
          Action : [
            "rds-db:connect" # Permiso específico: permite conectarse a bases de datos RDS
          ],
          Resource : [
            "arn:aws:rds-db:us-east-2:*:*" # Recursos permitidos: todos los recursos RDS en la región us-east-2
          ]
        }
      ]
    }
  )
}
# ------------------------------------------
# Output: ARN de la política creada
# ------------------------------------------
output "can_access_rds_arn" {
  value = aws_iam_policy.can_access_rds.arn # Exporta el ARN de la política para poder ser usado desde otros módulos o el módulo raíz
}
