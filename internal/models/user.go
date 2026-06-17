package models

import (
	"time"
)

type User struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	DOB       string    `json:"dob"`
	Age       int       `json:"age,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CreateUserRequest struct {
	Name string    `json:"name" validate:"required,min=1,max=100"`
	DOB  string `json:"dob" validate:"required"`
}

type UpdateUserRequest struct {
	Name string    `json:"name" validate:"omitempty,min=1,max=100"`
	DOB  string `json:"dob" validate:"omitempty"`
}

type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}

type ListUsersResponse struct {
	Users      []UserResponse `json:"users"`
	TotalCount int            `json:"total_count"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}