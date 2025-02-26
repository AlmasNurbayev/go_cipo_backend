package models

import (
	"time"

	"github.com/guregu/null/v5"
)

type UserEntity struct {
	Id           int64       `json:"id" db:"id"`
	Email        string      `json:"email" db:"email"`
	Name         null.String `json:"name" db:"name"`
	Password     null.String `json:"password" db:"password"`
	Salt         null.String `json:"salt" db:"salt"`
	Role         null.String `json:"role" db:"role"`
	Changed_date time.Time   `json:"changed_date" db:"changed_date"`
	Create_date  time.Time   `json:"create_date" db:"create_date"`
}
