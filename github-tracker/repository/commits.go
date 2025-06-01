package repository

import (
	"context"
	"database/sql"
	"github-tracker/github-tracker/repository/entity"
)

// Commit es la interfaz que define las operaciones disponibles
// sobre la entidad Commit en la capa de persistencia (repositorio).
type Commit interface {
	// Insert guarda un nuevo commit en la base de datos.
	Insert(ctx context.Context, commit *entity.Commit) (err error)

	// GetCommitsByAuthorEmail obtiene todos los commits hechos por un autor según su email.
	GetCommitsByAuthorEmail(ctx context.Context, email string) (commits []entity.Commit, err error)
}

// commit es la implementación concreta del repositorio que
// contiene una conexión a la base de datos SQL.
type commit struct {
	Conn *sql.DB // Conexión a la base de datos usando database/sql
}

// NewCommit es un constructor que retorna una instancia del repositorio commit,
// inyectando una conexión a la base de datos.
func NewCommit(conn *sql.DB) commit {
	return commit{
		Conn: conn,
	}
}

// Insert guarda un nuevo registro de tipo Commit en la base de datos.
// Usa PrepareContext + QueryRowContext, lo cual no es óptimo aquí,
// ya que no se espera devolver ningún resultado del INSERT.
func (m commit) Insert(ctx context.Context, commit *entity.Commit) (err error) {
	// Consulta SQL con placeholders posicionales ($1, $2, ...) — seguro contra SQL Injection
	query := `
		INSERT INTO commits (
			repo_name, commit_id, commit_message,
			author_username, author_email,
			payload, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	// Prepara el statement SQL. Esto compila la consulta una sola vez para evitar repetir análisis.
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err // Retorna si ocurre un error al preparar la consulta
	}
	defer stmt.Close() // Asegura que se liberen los recursos del statement

	// Ejecuta la consulta con los valores correspondientes.
	// Aunque se usa QueryRowContext, este método se suele utilizar para consultas SELECT.
	// No se espera un valor de retorno aquí, por lo que ExecContext sería más apropiado.
	err = stmt.QueryRowContext(
		ctx,
		commit.RepoName,
		commit.CommitID,
		commit.CommitMessage,
		commit.AuthorUsername,
		commit.AuthorEmail,
		commit.Payload,
		commit.CreatedAt,
		commit.UpdatedAt,
	).Err()

	return err
}

/*
Mejora sugerida (más simple y eficiente):
Usar ExecContext directamente en lugar de preparar el statement y usar QueryRowContext.

func (m commit) Insert(ctx context.Context, commit *entity.Commit) (err error) {
	query := `
		INSERT INTO commits (
			repo_name, commit_id, commit_message,
			author_username, author_email,
			payload, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err = m.Conn.ExecContext(ctx, query,
		commit.RepoName,
		commit.CommitID,
		commit.CommitMessage,
		commit.AuthorUsername,
		commit.AuthorEmail,
		commit.Payload,
		commit.CreatedAt,
		commit.UpdatedAt,
	)
	return err
}
*/

/*
Alternativa aún más limpia si usas sqlx y etiquetas `db` en entity.Commit:
import "github.com/jmoiron/sqlx"

type commit struct {
	DB *sqlx.DB
}

func (m commit) Insert(ctx context.Context, commit *entity.Commit) (err error) {
	query := `
		INSERT INTO commits (
			repo_name, commit_id, commit_message,
			author_username, author_email,
			payload, created_at, updated_at
		)
		VALUES (
			:repo_name, :commit_id, :commit_message,
			:author_username, :author_email,
			:payload, :created_at, :updated_at
		)
	`
	_, err = m.DB.NamedExecContext(ctx, query, commit)
	return err
}
*/

// GetCommitsByAuthorEmail busca y retorna todos los commits que coincidan
// con un autor específico, filtrando por email.
func (m commit) GetCommitsByAuthorEmail(ctx context.Context, email string) ([]entity.Commit, error) {
	// Consulta segura con placeholder para evitar SQL Injection
	query := `
        SELECT *
        FROM commits
        WHERE author_email = $1
    `

	// Ejecuta la consulta pasando el email como parámetro
	rows, err := m.Conn.QueryContext(ctx, query, email)
	if err != nil {
		return nil, err // Retorna error si falla la ejecución de la consulta
	}
	defer rows.Close() // Cierra las filas una vez finalizado el procesamiento

	var commits []entity.Commit

	// Itera sobre cada fila del resultado
	for rows.Next() {
		var commit entity.Commit

		// Escanea los valores de la fila actual en la estructura commit.
		// El orden debe coincidir con el orden de columnas en la tabla.
		err := rows.Scan(
			&commit.ID,
			&commit.RepoName,
			&commit.CommitID,
			&commit.CommitMessage,
			&commit.AuthorUsername,
			&commit.AuthorEmail,
			&commit.Payload,
			&commit.CreatedAt,
			&commit.UpdatedAt,
		)
		if err != nil {
			return nil, err // Si ocurre un error al escanear, se detiene y retorna el error
		}

		commits = append(commits, commit) // Agrega el commit a la lista
	}

	return commits, nil // Retorna la lista de commits encontrada
}
