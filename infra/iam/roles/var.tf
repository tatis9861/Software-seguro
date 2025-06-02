# --------------------------------------
# Variables de entrada para el módulo
# --------------------------------------

# ARN (Amazon Resource Name) de una política de acceso para RDS
# Esta variable espera que se le pase un string con el ARN de la política que da permisos sobre RDS
variable "rds_policy_arn" {
  type = string
}

# ARN de una política de acceso para CloudWatch Logs u otro sistema de logs
# Se espera un string con el ARN de una política que otorga permisos para registrar o leer logs
variable "log_policy_arn" {
  type = string
}
