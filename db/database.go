package db

import (
	"database/sql"
	"fmt"

	"github.com/Sparapani/streetmarketapi/models"

	_ "github.com/lib/pq"
)

func Setup() *sql.DB {

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbName)
	dbConn, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	return dbConn
}

func InsertStreetMarket(sm models.StreetMarket) (id string, err error) {
	db := DBRequest.Setup()
	defer db.Close()

	sqlStatement := `insert into streetmarket(longitude, latitude, set_census, area_deliberation, cod_district, district, cod_sub_city_hall,
		sub_city_hall, region_5, region_8, name_sm, registration, address, number_address, neighbourhood,
	  reference) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) returning id`

	err = db.QueryRow(sqlStatement, sm.Longitude, sm.Latitude, sm.SetCensus, sm.AreaDeliberation, sm.CodDistrict,
		sm.District, sm.CobSubCityHall, sm.SubCityHall, sm.Region5, sm.Region8, sm.NameSM,
		sm.Registration, sm.Address, sm.NumberAddress, sm.Neighbourhood, sm.Reference).Scan(&id)

	return
}

func DeleteStreetMarket(registerID string) (err error) {
	db := DBRequest.Setup()
	defer db.Close()

	_, err = db.Exec("delete from streetmarket where id =$1", registerID)

	return
}

func UpdateStretMarket(sm models.StreetMarket) (err error) {
	db := DBRequest.Setup()
	defer db.Close()

	sqlStatement := `update streetmarket set longitude = $2, latitude = $3, set_census = $4, area_deliberation = $5, cod_district = $6, district = $7,
	                 cod_sub_city_hall = $8, sub_city_hall = $9, region_5 = $10, region_8 = $11, name_sm = $12, registration = $13, address = $14, 
					 number_address = $15, neighbourhood = $16, reference = $17 where id = $1 returning id`

	err = db.QueryRow(sqlStatement, sm.ID, sm.Longitude, sm.Latitude, sm.SetCensus, sm.AreaDeliberation, sm.CodDistrict,
		sm.District, sm.CobSubCityHall, sm.SubCityHall, sm.Region5, sm.Region8, sm.NameSM,
		sm.Registration, sm.Address, sm.NumberAddress, sm.Neighbourhood).Scan(&sm.ID)

	return
}

func SearchStreetMarket(sm models.StreetMarket) (smFound []models.StreetMarket, err error) {
	db := DBRequest.Setup()
	defer db.Close()

	sqlStatement := `select * from streetmarket where district = coalesce($1,district) and region_5 = coalesce($2,region_5) 
		             and name_sm = coalesce($3,name_sm) and neighbourhood = coalesce($4,neighbourhood)`

	sqlReturn, err := db.Query(sqlStatement, sm.District, sm.Region5, sm.NameSM, sm.Neighbourhood)

	if err != nil {
		return
	}

	for sqlReturn.Next() {
		var smRegister models.StreetMarket
		sqlReturn.Scan(&smRegister.ID, &smRegister.Longitude, &smRegister.Latitude, &smRegister.SetCensus, &smRegister.AreaDeliberation, &smRegister.CodDistrict,
			&smRegister.District, &smRegister.CobSubCityHall, &smRegister.SubCityHall, &smRegister.Region5, &smRegister.Region8, &smRegister.NameSM,
			&smRegister.Registration, &smRegister.Address, &smRegister.NumberAddress, &smRegister.Neighbourhood)
		smFound = append(smFound, smRegister)
	}

	return
}
