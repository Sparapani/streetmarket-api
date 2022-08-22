package db

import (
	"database/sql"

	"github.com/Sparapani/streetmarketapi/models"
)

type db struct{}

type SetupDB interface {
	Setup() *sql.DB
}

type InsertSMDB interface {
	InsertStreetMarket(sm models.StreetMarket) (id string, err error)
}

type RemoveSMDB interface {
	DeleteStreetMarket(registerID string) (err error)
}

type UpdateSMDB interface {
	UpdateStretMarket(sm models.StreetMarket) (err error)
}

type SearchSMDB interface {
	SearchStreetMarket(sm models.StreetMarket) (smFound []models.StreetMarket, err error)
}

type DB interface {
	SetupDB
	InsertSMDB
	RemoveSMDB
	UpdateSMDB
	SearchSMDB
}

func (d db) Setup() *sql.DB {
	return Setup()
}

func (d db) InsertStreetMarket(sm models.StreetMarket) (id string, err error) {
	return InsertStreetMarket(sm)
}

func (d db) DeleteStreetMarket(registerID string) (err error) {
	return DeleteStreetMarket(registerID)
}

func (d db) UpdateStretMarket(sm models.StreetMarket) (err error) {
	return UpdateStretMarket(sm)
}

func (d db) SearchStreetMarket(sm models.StreetMarket) (smFound []models.StreetMarket, err error) {
	return SearchStreetMarket(sm)
}

var DBRequest DB

func init() {
	DBRequest = db{}
}
