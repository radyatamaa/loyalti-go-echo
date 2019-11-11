package model

import (
	"time"
)

type GeneralModels struct {
	Id         int `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"`
	Created    time.Time
	CreatedBy  string
	Modified   time.Time
	ModifiedBy string
	Active     bool
	IsDeleted  bool
	Deleted    *time.Time
	Deleted_by string
}