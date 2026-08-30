package main

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ TodoRepository = (*PostgreSQLTodoRepository)(nil)

type PostgreSQLTodoRepository struct {
	db *pgxpool.Pool
}

func NewPostgreSQLTodoRepository(
	db *pgxpool.Pool,
) *PostgreSQLTodoRepository {
	return &PostgreSQLTodoRepository{
		db: db,
	}
}

func (r *PostgreSQLTodoRepository) FindAll(
	ctx context.Context,
) ([]Todo, error) {
	rows, err := r.db.Query(
		ctx,
		`
		SELECT id, title, status
		FROM todos
		ORDER BY id
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]Todo, 0)

	for rows.Next() {
		var todo Todo

		if err := rows.Scan(
			&todo.Id,
			&todo.Title,
			&todo.Status,
		); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *PostgreSQLTodoRepository) FindByID (
	ctx	context.Context,
	id	int,
) (Todo, bool, error) {
	var todo Todo

	err := r.db.QueryRow(
		ctx,
		`
		SELECT id, title, status
		FROM todos
		WHERE id = $1
		`,
		id,
	).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Status,
	)
	
	if errors.Is(err, pgx.ErrNoRows) {
		return Todo{}, false, nil
	}

	if err != nil {
		return Todo{}, false, err
	}

	return todo, true, nil
}

func (r *PostgreSQLTodoRepository) Create (
	ctx context.Context,
	todo Todo,
) (Todo, error) {
	err := r.db.QueryRow(
		ctx,
		`
		INSERT INTO todos (title, status)
		VALUES ($1, $2)
		RETURNING id
		`, 
		todo.Title, 
		todo.Status,
	).Scan(&todo.Id)

	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func (r *PostgreSQLTodoRepository) Update (
	ctx context.Context,
	todo Todo,
) (Todo, bool, error) {
	err := r.db.QueryRow(
		ctx,
		`
		UPDATE todos
		SET status = $1
		WHERE id = $2
		RETURNING id, title, status
		`,
		todo.Status,
		todo.Id,
	).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Status,
	)

	if errors.Is(err, pgx.ErrNoRows){
		return Todo{}, false, nil
	}
	
	if err != nil {
		return Todo{}, false, err 
	}

	return todo, true, nil
}

func (r *PostgreSQLTodoRepository) Delete (
	ctx context.Context,
	id	int,
) (bool, error) {
	result, err := r.db.Exec(
		ctx,
		`
		DELETE FROM todos
		WHERE id = $1
		`,
		id,
	)

	if err != nil {
		return false, err
	}

	if result.RowsAffected() == 0 {
		return false, nil
	}

	return true, nil
}