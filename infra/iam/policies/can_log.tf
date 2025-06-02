# ------------------------------------------
# Recurso: Política IAM personalizada para CloudWatch Logs
# ------------------------------------------
resource "aws_iam_policy" "can_log" {
  name        = "can-log"                 # Nombre de la política IAM
  path        = "/"                       # Ruta por defecto (raíz del namespace de IAM)
  description = "Allow log to Cloudwatch" # Descripción clara del propósito de la política

  # Definición de la política en formato JSON (codificada desde HCL con jsonencode)
  policy = jsonencode(
    {
      Version : "2012-10-17", # Versión requerida por AWS para políticas
      Statement : [
        {
          Effect : "Allow",         # Esta política PERMITE las siguientes acciones
          Action : [                # Acciones específicas permitidas
            "logs:CreateLogGroup",  # Crear grupos de logs (si no existen)
            "logs:CreateLogStream", # Crear flujos de logs (streams) dentro del grupo
            "logs:PutLogEvents",    # Enviar eventos de log a CloudWatch Logs
          ],
          Resource : "*" # Aplica a todos los recursos (grupos de logs, streams, etc.)
        }
      ]
    }
  )
}

# ------------------------------------------
# Output: ARN de la política can_log
# ------------------------------------------
output "can_log_arn" {
  value = aws_iam_policy.can_log.arn # Expone el ARN de la política como salida para usar en otros módulos o recursos
}