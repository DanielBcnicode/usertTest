package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"usertest.com/user"
)

var FilterFields = map[string]string{
	"first_name": `"first_name"`,
	"last_name":  `"last_name"`,
	"nickname":   `"nickname"`,
	"email":      `"email"`,
	"country":    `"country"`,
}

type UserRepository struct {
	Connection *PostgresConn
}

func NewUserPostgresRepository(conn *PostgresConn) UserRepository {
	return UserRepository{Connection: conn}
}

func (u *UserRepository) Save(ctx context.Context, user *user.User) error {
	q := `
    INSERT INTO "user" ("id", "first_name", "last_name", "nick_name", "password", "email",
	    "country", "created_at", "updated_at")
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
    `
	row := u.Connection.db.QueryRowContext(
		ctx, q, user.ID, user.FirstName, user.LastName, user.NickName,
		user.Password, user.Email, user.Country, user.CreatedAt, user.UpdatedAt,
	)

	return row.Err()
}

func (u *UserRepository) Update(ctx context.Context, user *user.User) error {
	q := `
    UPDATE "user" SET "first_name"=$1, "last_name"=$2, "nick_name"=$3, "password"=$4, "email"=$5,
	    "country"=$6, "updated_at"=$7 WHERE "id"=$8;
    `
	res, err := u.Connection.db.ExecContext(ctx, q, user.FirstName, user.LastName, user.NickName,
		user.Password, user.Email, user.Country, user.UpdatedAt,
		user.ID)

	if err != nil {
		log.Printf("ERROR: updating user row: %s\n", err)
		return err
	}

	i, err := res.RowsAffected()
	if i != 1 {
		log.Printf("ERROR: user with id = %s not found\n", user.ID)
		return errors.New("user to update not found in database")
	}

	return err
}

func (u *UserRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	q := `
    DELETE FROM "user" WHERE "id"=$1;
    `
	res, err := u.Connection.db.ExecContext(
		ctx, q, userID,
	)

	if err != nil {
		log.Printf("ERROR: deleting user row: %s\n", err)
		return err
	}

	i, _ := res.RowsAffected()

	if i != 1 {
		log.Printf("ERROR: user with id = %s not found\n", userID)
		return errors.New("user to delete not found in database")
	}

	return err
}

func (u *UserRepository) FindByFilter(ctx context.Context, filter user.RepositoryFilter, paginator *user.Paginator) ([]user.User, error) {
	d := make([]user.User, 0)
	limit := 10
	offset := 0
	if paginator != nil {
		if paginator.PagSize > 0 {
			limit = paginator.PagSize
		}
		offset = limit * paginator.CurrentPage
	}
	f, err := generateUserFilterQuery(filter)
	if err != nil {
		return d, err
	}

	q := ` SELECT "id", "first_name", "last_name", "nick_name", "password", "email",
	"country", "created_at", "updated_at" FROM "user"`
	l := ` LIMIT $1 OFFSET $2;`
	q += f + l

	rows, err := u.Connection.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		log.Printf("Error FindByFilter: %s\n", err)
		return d, err
	}

	for rows.Next() {
		user := user.User{}
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.NickName,
			&user.Password,
			&user.Email,
			&user.Country,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			log.Printf("Error Scanning row : %s\n", err)
			continue
		}
		d = append(d, user)
	}

	return d, nil
}

func generateUserFilterQuery(filter user.RepositoryFilter) (string, error) {
	s := ""
	i := 0
	for k, v := range filter.Filters {
		//Add ´WHERE´ claus
		if i == 0 {
			s = " WHERE "
		}
		f, ok := FilterFields[k]
		if !ok {
			return "", fmt.Errorf("filter '%s' not allowed", k)
		}
		//Add ´AND´ clausule it there are more than one filter 
		if i > 0 {
			s += "AND "
		}
		s += f + ` = "` + v + `" `
		i++
	}

	return s, nil
}
