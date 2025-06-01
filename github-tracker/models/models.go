package models

// GitHubWebhook representa un evento de tipo "push" enviado por GitHub a un webhook.
// Solo incluye los campos que interesan en este caso: el repositorio y el commit principal (head commit).
type GitHubWebhook struct {
	Repository Repository `json:"repository"`  // Información básica del repositorio donde ocurrió el push
	HeadCommit Commit     `json:"head_commit"` // El último commit incluido en el push (el más reciente)
}

// Repository representa una estructura simplificada del repositorio GitHub.
type Repository struct {
	FullName string `json:"full_name"` // Nombre completo del repositorio (ej: "usuario/repositorio")
}

// Commit representa los datos esenciales de un commit en un evento de push.
type Commit struct {
	ID      string     `json:"id"`      // SHA del commit
	Message string     `json:"message"` // Mensaje del commit (lo que escribió el desarrollador)
	Author  CommitUser `json:"author"`  // Información del autor del commit
}

// CommitUser contiene los datos del autor del commit.
type CommitUser struct {
	Email    string `json:"email"`    // Email del autor del commit
	Username string `json:"username"` // Nombre de usuario de GitHub del autor
}
