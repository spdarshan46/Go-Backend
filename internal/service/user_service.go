package service

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"user-management-api/internal/models"
	"user-management-api/internal/repository"
	"user-management-api/internal/repository/sqlc"
	"user-management-api/internal/utils"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (sqlc.User, error) {
	if err := utils.ValidateStruct(req); err != nil {
		return sqlc.User{}, err
	}
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return sqlc.User{}, errors.New("invalid dob format, use YYYY-MM-DD")
	}

	if dob.After(time.Now()) {
		return sqlc.User{}, errors.New("date of birth cannot be in the future")
	}

	minDate := time.Now().AddDate(-150, 0, 0)
	if dob.Before(minDate) {
		return sqlc.User{}, errors.New("date of birth is too far in the past (max 150 years)")
	}

	return s.repo.CreateUser(ctx, req.Name, dob)
}

func (s *UserService) GetUserByID(ctx context.Context, id int32) (sqlc.User, error) {
	if id <= 0 {
		return sqlc.User{}, errors.New("invalid user id")
	}

	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlc.User{}, errors.New("user not found")
		}
		return sqlc.User{}, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, req *models.UpdateUserRequest) (sqlc.User, error) {
	if id <= 0 {
		return sqlc.User{}, errors.New("invalid user id")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return sqlc.User{}, err
	}

	exists, err := s.repo.CheckUserExists(ctx, id)
	if err != nil {
		return sqlc.User{}, err
	}
	if !exists {
		return sqlc.User{}, errors.New("user not found")
	}

	name := req.Name
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return sqlc.User{}, errors.New("invalid dob format, use YYYY-MM-DD")
	}

	user, err := s.repo.UpdateUser(ctx, id, name, dob)
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	if id <= 0 {
		return errors.New("invalid user id")
	}

	exists, err := s.repo.CheckUserExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("user not found")
	}

	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context, page, pageSize int) ([]sqlc.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := int32((page - 1) * pageSize)
	limit := int32(pageSize)

	users, err := s.repo.ListUsers(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountUsers(ctx)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (s *UserService) CalculateAgeForUser(user *sqlc.User) int {
	if user == nil {
		return 0
	}
	return utils.CalculateAge(user.Dob)
}