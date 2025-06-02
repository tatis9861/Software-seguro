# ------------------------------------------------------------------------------
# Definición de la política de confianza para el rol (trust relationship)
# Permite que el servicio Lambda pueda asumir este rol IAM
# ------------------------------------------------------------------------------
data "aws_iam_policy_document" "instance_assume_role_policy" {
  statement {
    effect = "Allow" # Permitir la acción definida abajo

    principals {
      type        = "Service"                # El tipo de principal es un servicio AWS
      identifiers = ["lambda.amazonaws.com"] # Específicamente el servicio AWS Lambda
    }

    actions = ["sts:AssumeRole"] # Permiso para asumir el rol
  }
}

# ------------------------------------------------------------------------------
# Creación del rol IAM que será asumido por una función Lambda
# Este rol se usará para acceder a RDS y registrar logs en CloudWatch
# ------------------------------------------------------------------------------
resource "aws_iam_role" "repo_collector" {
  name               = "repo-collector-platzi" # Nombre legible del rol
  path               = "/"                     # Ruta raíz del rol
  assume_role_policy = data.aws_iam_policy_document.instance_assume_role_policy.json
  # Asocia la política de confianza definida arriba
}

# ------------------------------------------------------------------------------
# Asociación de la política que permite conectarse a RDS
# Esta política es pasada como variable al módulo
# ------------------------------------------------------------------------------
resource "aws_iam_role_policy_attachment" "can_access_rds" {
  role       = aws_iam_role.repo_collector.name # Nombre del rol a asociar
  policy_arn = var.rds_policy_arn               # ARN de la política de RDS (entrada como variable)
}

# ------------------------------------------------------------------------------
# Asociación de la política que permite registrar logs en CloudWatch
# Esta política también se pasa como variable externa
# ------------------------------------------------------------------------------
resource "aws_iam_role_policy_attachment" "can_log" {
  role       = aws_iam_role.repo_collector.name # Nombre del rol a asociar
  policy_arn = var.log_policy_arn               # ARN de la política de logs (entrada como variable)
}

# ------------------------------------------------------------------------------
# Output: Exporta el ARN del rol creado para que otros módulos lo usen
# ------------------------------------------------------------------------------
output "repo_collector_role_arn" {
  value = aws_iam_role.repo_collector.arn # Devuelve el ARN del rol IAM creado
}