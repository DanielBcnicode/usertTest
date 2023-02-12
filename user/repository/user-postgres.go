package repository

import (
	"context"

	"usertest.com/user"
)

type UserRepository struct {
	Connection *PostgresConn
}

func NewUserPostgresRepository(conn *PostgresConn) UserRepository {
	return UserRepository{Connection: conn}
}

func (u *UserRepository) Save(ctx context.Context, user *user.User) error {
	q := `
    INSERT INTO user (id, first_name, last_name, nickname, password, email, country, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id;
    `
	row := u.Connection.db.QueryRowContext(
		ctx, q, user.ID.String(), user.FirstName, user.LastName, user.NickName,
		user.Password, user.Email, user.Country, user.CreatedAt.String(),
		user.UpdatedAt.String(),
	)
	row.Scan()
	return row.Err()
}

func (u *UserRepository) Update(ctx context.Context, user *user.User) error {
	return nil
}

func (u *UserRepository) FindByFilter(ctx context.Context, filter user.RepositoryFilter, paginaror *user.Paginator) ([]user.User, error) {
	return nil, nil
}
