package api

import (
	"errors"
	"testing"

	"github.com/Sparapani/streetmarketapi/models"
	"github.com/Sparapani/streetmarketapi/services"

	"github.com/Sparapani/streetmarketapi/db"
	"github.com/Sparapani/streetmarketapi/utils"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Name    string
	InFile  string
	ErrorDB error
}

func TestGetSM(t *testing.T) {
	testCases := []testCase{
		{
			Name:   "test 1 - search - ok",
			InFile: "./tests/search/test1/data.json",
		},
		{
			Name:    "test 2 - search - not all fields empty but other other",
			InFile:  "./tests/search/test2/data.json",
			ErrorDB: errors.New("other error"),
		},
		{
			Name:    "test 3 - search - all fields empty",
			InFile:  "./tests/search/test3/data.json",
			ErrorDB: services.ErrRequiredFieldsEmpty,
		},
	}

	for _, tCase := range testCases {
		testGetSM(t, tCase)
	}

}

func testGetSM(t *testing.T, tCase testCase) {
	db.DBRequest = DBMock{}
	services.SMRequest = SMCrudMock{}

	searchStreetMarketfuncMock = func(sm models.StreetMarket) (smFound []models.StreetMarket, err error) {
		err = tCase.ErrorDB
		return
	}

	getSMMock = func(smIn models.StreetMarket) (smOut []models.StreetMarket, err error) {
		if smIn.District == "" && smIn.Region5 == "" && smIn.NameSM == "" && smIn.Neighbourhood == "" {
			err = services.ErrRequiredFieldsEmpty
			return
		}

		smOut, err = db.DBRequest.SearchStreetMarket(smIn)
		return
	}

	var sm models.StreetMarket

	utils.ReadJSONFile(t, tCase.InFile, &sm)

	_, errGot := services.SMRequest.GetSM(sm)

	expected := tCase.ErrorDB

	assert.Equal(t, expected, errGot)
}

func TestDeleteSM(t *testing.T) {
	testCases := []testCase{
		{
			Name:   "test 1 - delete - ok",
			InFile: "./tests/delete/test1/data.json",
		},
		{
			Name:    "test 2 - delete - nok other",
			InFile:  "./tests/delete/test2/data.json",
			ErrorDB: errors.New("other error"),
		},
		{
			Name:    "test 3 - delete - nok empty",
			InFile:  "./tests/delete/test3/data.json",
			ErrorDB: services.ErrRequiredFieldsEmpty,
		},
	}

	for _, tCase := range testCases {
		testDeleteSM(t, tCase)
	}

}

func testDeleteSM(t *testing.T, tCase testCase) {
	db.DBRequest = DBMock{}
	services.SMRequest = SMCrudMock{}

	deleteStreetMarketMock = func(registerID string) (err error) {
		err = tCase.ErrorDB
		return
	}

	deleteSMMock = func(sm models.StreetMarket) (err error) {
		if sm.ID == "" {
			err = services.ErrRequiredFieldsEmpty
			return
		}
		err = db.DBRequest.DeleteStreetMarket(sm.ID)
		return
	}

	var sm models.StreetMarket

	utils.ReadJSONFile(t, tCase.InFile, &sm)

	errGot := services.SMRequest.DeleteSM(sm)

	expected := tCase.ErrorDB

	assert.Equal(t, expected, errGot)
}

func TestAddSM(t *testing.T) {
	testCases := []testCase{
		{
			Name:   "test 1 - insert - ok",
			InFile: "./tests/insert/test1/data.json",
		},
		{
			Name:    "test 2 - insert - nok other",
			InFile:  "./tests/insert/test2/data.json",
			ErrorDB: errors.New("other error"),
		},
	}

	for _, tCase := range testCases {
		testAddSM(t, tCase)
	}

}

func testAddSM(t *testing.T, tCase testCase) {
	db.DBRequest = DBMock{}
	services.SMRequest = SMCrudMock{}

	insertStreetMarketMock = func(sm models.StreetMarket) (id string, err error) {
		err = tCase.ErrorDB
		return
	}

	addSMMock = func(sm models.StreetMarket) (id string, err error) {
		id, err = db.DBRequest.InsertStreetMarket(sm)
		return
	}

	var sm models.StreetMarket

	utils.ReadJSONFile(t, tCase.InFile, &sm)

	_, errGot := services.SMRequest.AddSM(sm)

	expected := tCase.ErrorDB

	assert.Equal(t, expected, errGot)

}

func TestUpdateSM(t *testing.T) {
	testCases := []testCase{
		{
			Name:   "test 1 - update - ok",
			InFile: "./tests/update/test1/data.json",
		},
		{
			Name:    "test 2 - update - nok",
			InFile:  "./tests/update/test2/data.json",
			ErrorDB: errors.New("other error"),
		},
	}

	for _, tCase := range testCases {
		testUpdateSM(t, tCase)
	}

}

func testUpdateSM(t *testing.T, tCase testCase) {
	db.DBRequest = DBMock{}
	services.SMRequest = SMCrudMock{}

	updateStretMarketMock = func(sm models.StreetMarket) (err error) {
		err = tCase.ErrorDB
		return
	}

	updateSMMock = func(sm models.StreetMarket) (err error) {
		err = db.DBRequest.UpdateStretMarket(sm)
		return err
	}

	var sm models.StreetMarket

	utils.ReadJSONFile(t, tCase.InFile, &sm)

	errGot := services.SMRequest.UpdateSM(sm)

	expected := tCase.ErrorDB

	assert.Equal(t, expected, errGot)
}
