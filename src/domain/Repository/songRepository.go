package repository

import (
	"github.com/radyatamaa/loyalti-go-echo/src/domain"
)

type SongRepository interface {
	GetAll()([] *domain.Song,error)
}
