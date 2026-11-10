package notes

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, note *Note) error
	GetAll(ctx context.Context) ([]Note, error)
	GetByID(ctx context.Context, id int64) (*Note, error)
	Update(ctx context.Context, id int64, note *Note) (*Note, error)
	Delete(ctx context.Context, id int64) error
}

type PostgresRepository struct {
	DB *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &PostgresRepository{
		DB: db,
	}
}

func (r *PostgresRepository) Create(ctx context.Context, note *Note) error {
	query := `INSERT INTO notes(title, content) VALUES($1, $2) RETURNING id, created_at, updated_at`

	err := r.DB.QueryRow(
		ctx,
		query,
		note.Title,
		note.Content,
	).Scan(
		&note.ID,
		&note.CreatedAt,
		&note.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil

}

func (r *PostgresRepository) GetAll(ctx context.Context) ([]Note, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM notes ORDER BY created_at DESC`
	rows, err := r.DB.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var notes []Note

	for rows.Next() {
		var note Note

		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)

		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return notes, nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id int64) (*Note, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM notes WHERE id = $1`

	var note Note

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&note.ID,
		&note.Title,
		&note.Content,
		&note.CreatedAt,
		&note.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNoteNotFound
		}
		return nil, err
	}
	return &note, nil

}

func (r *PostgresRepository) Update(ctx context.Context, id int64, note *Note) (*Note, error) {
	query := `UPDATE notes SET title = $1, content = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3 RETURNING id, title, content, created_at, updated_at`

	var updatedNote Note

	err := r.DB.QueryRow(ctx, query, note.Title, note.Content, id).Scan(
		&updatedNote.ID,
		&updatedNote.Title,
		&updatedNote.Content,
		&updatedNote.CreatedAt,
		&updatedNote.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNoteNotFound
		}

		return nil, err
	}

	return &updatedNote, nil
}

func (r *PostgresRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM notes WHERE id = $1`

	result, err := r.DB.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return ErrNoteNotFound
	}

	return nil
}
