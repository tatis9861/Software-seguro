package entity

import "time"

// Commit representa una entidad de base de datos que modela un commit registrado desde un webhook de GitHub.
// Este struct se utiliza en la capa de persistencia (repositorio), generalmente para guardar commits en una tabla SQL.
type Commit struct {
	ID             int       `db:"id"`              // Identificador único del commit en la base de datos (clave primaria, autoincremental)
	RepoName       string    `db:"repo_name"`       // Nombre completo del repositorio (por ejemplo: "usuario/repositorio")
	CommitID       string    `db:"commit_id"`       // SHA del commit (hash único del commit en Git)
	CommitMessage  string    `db:"commit_message"`  // Mensaje del commit escrito por el desarrollador
	AuthorUsername string    `db:"author_username"` // Nombre de usuario de GitHub del autor del commit
	AuthorEmail    string    `db:"author_email"`    // Email del autor del commit
	Payload        string    `db:"payload"`         // JSON completo del webhook recibido (opcionalmente para auditoría o depuración)
	CreatedAt      time.Time `db:"created_at"`      // Fecha de creación del registro en la base de datos
	UpdatedAt      time.Time `db:"updated_at"`      // Fecha de la última actualización del registro
}
