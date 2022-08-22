package utils

import (
	"encoding/json"
	"os"
	"testing"
)

// ReadJSONFile reads filename fName and unmarshall json to data
func ReadJSONFile(t *testing.T, fName string, data interface{}) {
	dataByte, err := os.ReadFile(fName)
	if err != nil {
		t.Fatalf("couldnt read file. error [%s]", err)
		return
	}
	err = json.Unmarshal(dataByte, data)
	if err != nil {
		t.Fatalf("couldnt unmarshal data to json. error [%s]", err)
	}
}
