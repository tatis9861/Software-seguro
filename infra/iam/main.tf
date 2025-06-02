# Define un módulo llamado "policies" que carga los recursos desde el subdirectorio ./policies
module "policies" {
  source = "./policies"
}

# Define un módulo llamado "roles" que carga los recursos desde el subdirectorio ./roles
# También le pasa dos variables: los ARN de las políticas creadas en el módulo "policies"
module "roles" {
  source         = "./roles"
  rds_policy_arn = module.policies.can_access_rds_arn  # Pasa el ARN de la política que da acceso a RDS
  log_policy_arn = module.policies.can_log_arn         # Pasa el ARN de la política que permite el acceso a logs
}

# Define una salida que muestra el ARN del rol "repo_collector" creado en el módulo "roles"
output "repo_collector_role_arn" {
  value = module.roles.repo_collector_role_arn
}