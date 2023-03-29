package user

import (
	"chat/db/cockroach"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"log"
)

type repository struct {
	db *cockroach.Cockroach
}

func NewRepository(tx *cockroach.Cockroach) Repository {
	return &repository{db: tx}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	sql, args, err := r.db.Builder.
		Insert("users").
		Columns("username, password, email").
		Values(user.UserName, user.Password, user.Email).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		log.Fatalf("could not initialize database connection:: %s", err)
		return nil, err
	}

	var id string
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return nil, err
	}

	user.ID = id
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	sql, args, err := r.db.Builder.
		Select("id, username, password, email").
		From("users").
		Where(squirrel.Eq{"email": email}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&u.ID, &u.UserName, &u.Password, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *repository) EmailExists(ctx context.Context, email string) (bool, error) {
	sql, args, err := r.db.Builder.
		Select("1").
		From("users").
		Where(squirrel.Eq{"email": email}).
		Limit(1).
		ToSql()

	if err != nil {
		return false, err
	}

	var tmp int64
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&tmp)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, err
		}
		return false, err
	}
	return tmp == 1, nil
}
