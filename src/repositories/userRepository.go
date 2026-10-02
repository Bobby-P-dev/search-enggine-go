package repositories

import (
	"context"

	"github.com/Bobby-P-dev/search-enggine-go/src/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetUser() ([]models.User, error)
	CreateUser(ctx context.Context, req models.User) (models.User, error)
}

type UserRepositoryImpl struct {
	DB *pgxpool.Pool
}

func NewUserRepositoryImpl(db *pgxpool.Pool) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (u *UserRepositoryImpl) GetUser() ([]models.User, error) {
	var users []models.User

	query := `
		SELECT id, nik, name, department, position 
		FROM employees 
		ORDER BY id DESC LIMIT 10
	`

	rows, err := u.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Nik,
			&user.Name,
			&user.Departemen,
			&user.Position,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepositoryImpl) CreateUser(ctx context.Context, req models.User) (models.User, error) {
	query := `
		INSERT INTO employees (nik, name, department, position) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, nik, name, department, position
	`
	var result models.User

	err := u.DB.QueryRow(
		ctx,
		query,
		req.Nik,
		req.Name,
		req.Departemen,
		req.Position,
	).Scan(
		&result.ID,
		&result.Nik,
		&result.Name,
		&result.Departemen,
		&result.Position,
	)
	if err != nil {
		return result, err
	}
	return result, nil
}
