package api

import (
	"database/sql"

	"github.com/Sparapani/streetmarketapi/models"
)

type DBMock struct{}
type SMCrudMock struct{}

var setupMock func() *sql.DB
var insertStreetMarketMock func(sm models.StreetMarket) (id string, err error)
var deleteStreetMarketMock func(registerID string) (err error)
var updateStretMarketMock func(sm models.StreetMarket) (err error)
var searchStreetMarketfuncMock func(sm models.StreetMarket) (smFound []models.StreetMarket, err error)

var addSMMock func(sm models.StreetMarket) (id string, err error)
var deleteSMMock func(sm models.StreetMarket) (err error)
var updateSMMock func(sm models.StreetMarket) (err error)
var getSMMock func(smIn models.StreetMarket) (smOut []models.StreetMarket, err error)

func (d DBMock) Setup() *sql.DB {
	return setupMock()
}

func (d DBMock) InsertStreetMarket(sm models.StreetMarket) (id string, err error) {
	return insertStreetMarketMock(sm)
}

func (d DBMock) DeleteStreetMarket(registerID string) (err error) {
	return deleteStreetMarketMock(registerID)
}

func (d DBMock) UpdateStretMarket(sm models.StreetMarket) (err error) {
	return updateStretMarketMock(sm)
}

func (d DBMock) SearchStreetMarket(sm models.StreetMarket) (smFound []models.StreetMarket, err error) {
	return searchStreetMarketfuncMock(sm)
}

func (s SMCrudMock) AddSM(sm models.StreetMarket) (id string, err error) {
	return addSMMock(sm)
}

func (s SMCrudMock) DeleteSM(sm models.StreetMarket) (err error) {
	return deleteSMMock(sm)
}

func (s SMCrudMock) UpdateSM(sm models.StreetMarket) (err error) {
	return updateSMMock(sm)
}

func (s SMCrudMock) GetSM(smIn models.StreetMarket) (smOut []models.StreetMarket, err error) {
	return getSMMock(smIn)
}
