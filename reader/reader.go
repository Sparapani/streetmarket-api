package reader

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/Sparapani/streetmarketapi/models"
	"github.com/Sparapani/streetmarketapi/services"
)

func ReadStreetMartketUsingCSV(fileName string) {
	csvFile, err := os.Open(fileName)

	if err != nil {
		log.Printf("could not open the CSV file. err [%s]", err)
		return
	}

	defer csvFile.Close()

	csvLines := csv.NewReader(csvFile)

	csvLines.FieldsPerRecord = -1

	count := 0

	for {
		line, err := csvLines.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("could not read the line. err [%s]", err)
			continue
		}

		if count == 0 {
			count++
			continue
		}

		smRegistration := models.NewStreetMarketByArray(line)

		_, err = services.SMRequest.AddSM(smRegistration)

		if err != nil {
			continue
		}
	}

}
