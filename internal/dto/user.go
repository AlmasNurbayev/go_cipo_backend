package dto

import (
	"github.com/guregu/null/v5"
)

type UserRequestParams struct {
	Id int64 `validate:"required,gte=0" example:"5"`
}

type UserRequestQueryParams struct {
	Name string `query:"name" validate:"omitempty" example:"almas"`
}

type UserResponse struct {
	Id    int64       `json:"id" db:"id"`
	Email string      `json:"email" db:"email"`
	Name  null.String `json:"name" db:"name"`
	Role  null.String `json:"role" db:"role"`
}
