// Package main implements the all functions to make an API runnable
package main

import (
	"github.com/Sparapani/streetmarketapi/api"
	"github.com/Sparapani/streetmarketapi/services"
)

func init() {
	services.InitLog()
}

func main() {
	//reader.ReadStreetMartketUsingCSV("DEINFO_AB_FEIRASLIVRES_2014.csv")
	api.StartRouter()
	services.FinishLog()
}
