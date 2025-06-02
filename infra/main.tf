# --------------------------------------
# Configuración de Terraform y backend
# --------------------------------------
terraform {
  backend "s3" {
    bucket = "ciberseguridad-bucket" # Nombre del bucket S3 donde se almacenará el archivo de estado remoto (terraform.tfstate)
    key    = "terraform.tfstate"     # Ruta (clave) del archivo dentro del bucket
    region = "us-east-2"             # Región AWS donde está ubicado el bucket
  }
}

# --------------------------------------
# Proveedor de infraestructura: AWS
# --------------------------------------
provider "aws" {
  region = "us-east-2" # Región por defecto donde se desplegarán los recursos
}

# --------------------------------------
# Módulo IAM
# --------------------------------------
module "iam" {
  source = "./iam" # Ruta al módulo local que contiene recursos relacionados con IAM (roles, políticas, etc.)
}

# --------------------------------------
# Módulo S3
# --------------------------------------
module "s3" {
  source = "./s3" # Ruta al módulo local que gestiona buckets S3
}

# --------------------------------------
# Módulo Compute (por ejemplo: Lambda)
# --------------------------------------
module "compute" {
  source                  = "./compute"                        # Ruta al módulo que contiene recursos computacionales (Lambda, EC2, etc.)
  lambda_bucket           = module.s3.lambda_bucket            # Usa como input el bucket creado por el módulo S3
  repo_collector_role_arn = module.iam.repo_collector_role_arn # Usa como input el ARN del rol creado por el módulo IAM
}
