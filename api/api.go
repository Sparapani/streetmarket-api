package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Sparapani/streetmarketapi/models"
	"github.com/Sparapani/streetmarketapi/services"

	"github.com/gorilla/mux"
)

func StartRouter() {

	router := mux.NewRouter()

	router.HandleFunc("/streetmarket/", getStreetMarket).Methods("GET")
	router.HandleFunc("/streetmarket/", addStreetMarket).Methods("POST")
	router.HandleFunc("/streetmarket/{id}", updateStreetMarket).Methods("PUT")
	router.HandleFunc("/streetmarket/{id}", deleteStreetMarket).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getStreetMarket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var smIn models.StreetMarket

	json.NewDecoder(r.Body).Decode(&smIn)

	smOut, err := services.SMRequest.GetSM(smIn)

	if err != nil {
		switch {
		case errors.Is(err, services.ErrRequiredFieldsEmpty):
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(models.StreetMarket{})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&smOut)

}

func addStreetMarket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var smIn models.StreetMarket

	json.NewDecoder(r.Body).Decode(&smIn)

	id, err := services.SMRequest.AddSM(smIn)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&id)
}

func deleteStreetMarket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var smIn models.StreetMarket

	json.NewDecoder(r.Body).Decode(&smIn)

	err := services.SMRequest.DeleteSM(smIn)

	if err != nil {
		switch {
		case errors.Is(err, services.ErrRequiredFieldsEmpty):
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode("NOK")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}

func updateStreetMarket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var smIn models.StreetMarket

	json.NewDecoder(r.Body).Decode(&smIn)

	err := services.SMRequest.UpdateSM(smIn)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("NOK")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}
