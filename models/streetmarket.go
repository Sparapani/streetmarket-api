package models

type StreetMarket struct {
	ID               string `json:"id,omitempty"`
	Longitude        string `json:"longitude,omitempty"`
	Latitude         string `json:"latitude,omitempty"`
	SetCensus        string `json:"set_census,omitempty"`
	AreaDeliberation string `json:"area_deliberation,omitempty"`
	CodDistrict      string `json:"cod_district,omitempty"`
	District         string `json:"district,omitempty"`
	CobSubCityHall   string `json:"cob_sub_city_hall,omitempty"`
	SubCityHall      string `json:"sub_city_hall,omitempty"`
	Region5          string `json:"region_5,omitempty"`
	Region8          string `json:"region_8,omitempty"`
	NameSM           string `json:"name_sm,omitempty"`
	Registration     string `json:"registration,omitempty"`
	Address          string `json:"address,omitempty"`
	NumberAddress    string `json:"number_address,omitempty"`
	Neighbourhood    string `json:"neighbourhood,omitempty"`
	Reference        string `json:"reference,omitempty"`
}

func NewStreetMarketByArray(line []string) (newSM StreetMarket) {
	newSM = StreetMarket{
		Longitude:        line[1],
		Latitude:         line[2],
		SetCensus:        line[3],
		AreaDeliberation: line[4],
		CodDistrict:      line[5],
		District:         line[6],
		CobSubCityHall:   line[7],
		SubCityHall:      line[8],
		Region5:          line[9],
		Region8:          line[10],
		NameSM:           line[11],
		Registration:     line[12],
		Address:          line[13],
		NumberAddress:    line[14],
		Neighbourhood:    line[15],
		Reference:        line[16],
	}
	return
}
