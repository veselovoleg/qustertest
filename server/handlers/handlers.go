package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"sitetest/server/models"
)

//GetData handler
func GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data := models.GetData()
	jsonData, err := json.Marshal(&data)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, string(jsonData))
}
