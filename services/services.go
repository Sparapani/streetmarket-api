package services

import (
	"github.com/Sparapani/streetmarketapi/db"
	"github.com/Sparapani/streetmarketapi/models"
)

func GetSM(smIn models.StreetMarket) (smOut []models.StreetMarket, err error) {

	if smIn.District == "" && smIn.Region5 == "" && smIn.NameSM == "" && smIn.Neighbourhood == "" {
		err = ErrRequiredFieldsEmpty
		logChannel <- logType{smIn, "GET", "NOK - REQUIRED FIELDS IS EMPTY"}
		return
	}

	smOut, err = db.DBRequest.SearchStreetMarket(smIn)

	if err != nil {
		logChannel <- logType{smIn, "GET", "NOK - OTHER ERROR"}
		return
	}

	logChannel <- logType{smIn, "GET", "OK"}
	return
}

func AddSM(sm models.StreetMarket) (id string, err error) {
	id, err = db.DBRequest.InsertStreetMarket(sm)
	if err != nil {
		logChannel <- logType{sm, "POST", "NOK - OTHER ERROR"}
		return
	}
	sm.ID = id
	logChannel <- logType{sm, "POST", "OK"}
	return
}

func DeleteSM(sm models.StreetMarket) (err error) {
	if sm.ID == "" {
		err = ErrRequiredFieldsEmpty
		logChannel <- logType{sm, "DELETE", "NOK - REQUIRED FIELDS IS EMPTY"}
		return
	}

	err = db.DBRequest.DeleteStreetMarket(sm.ID)
	if err != nil {
		logChannel <- logType{sm, "DELETE", "NOK - OTHER ERROR"}
	}

	logChannel <- logType{sm, "DELETE", "OK"}
	return
}

func UpdateSM(sm models.StreetMarket) (err error) {
	err = db.DBRequest.UpdateStretMarket(sm)
	if err != nil {
		logChannel <- logType{sm, "PUT", "NOK - OTHER ERROR"}
	}

	logChannel <- logType{sm, "PUT", "OK"}
	return err
}
