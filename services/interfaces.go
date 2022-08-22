package services

import "github.com/Sparapani/streetmarketapi/models"

type smCrud struct{}

type InsertSM interface {
	AddSM(sm models.StreetMarket) (id string, err error)
}

type RemoveSM interface {
	DeleteSM(sm models.StreetMarket) (err error)
}

type UpdSM interface {
	UpdateSM(sm models.StreetMarket) (err error)
}

type SearchSM interface {
	GetSM(smIn models.StreetMarket) (smOut []models.StreetMarket, err error)
}

type SMCrud interface {
	InsertSM
	RemoveSM
	UpdSM
	SearchSM
}

func (s smCrud) AddSM(sm models.StreetMarket) (id string, err error) {
	return AddSM(sm)
}

func (s smCrud) DeleteSM(sm models.StreetMarket) (err error) {
	return DeleteSM(sm)
}

func (s smCrud) UpdateSM(sm models.StreetMarket) (err error) {
	return UpdateSM(sm)
}

func (s smCrud) GetSM(smIn models.StreetMarket) (smOut []models.StreetMarket, err error) {
	return GetSM(smIn)
}

var SMRequest SMCrud

func init() {
	SMRequest = smCrud{}
}
