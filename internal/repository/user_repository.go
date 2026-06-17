package repository

import (
	"context"
	"database/sql"
	"time"

	"user-management-api/internal/repository/sqlc"
)

type UserRepository struct {
	queries *sqlc.Queries
	db      *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		queries: sqlc.New(db),
		db:      db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, name string, dob time.Time) (sqlc.User, error) {
	return r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int32) (sqlc.User, error) {
	return r.queries.GetUserByID(ctx, id)
}

func (r *UserRepository) UpdateUser(ctx context.Context, id int32, name string, dob time.Time) (sqlc.User, error) {
	return r.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}

func (r *UserRepository) ListUsers(ctx context.Context, limit, offset int32) ([]sqlc.User, error) {
	return r.queries.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
}

func (r *UserRepository) CountUsers(ctx context.Context) (int64, error) {
	return r.queries.CountUsers(ctx)
}

func (r *UserRepository) CheckUserExists(ctx context.Context, id int32) (bool, error) {
	return r.queries.CheckUserExists(ctx, id)
}