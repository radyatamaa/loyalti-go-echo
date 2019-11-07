package domain

import "time"

type GeneralModels struct {
	Created_by string
	Updated_by string
	Deleted_by string
	Row_status int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
