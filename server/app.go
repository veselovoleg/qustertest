package main

import (
	"fmt"
	"net/http"
	"sitetest/server/handlers"

	"github.com/gorilla/mux"
)

//PORTNUMBER - choose port
const PORTNUMBER = ":8081"

func main() {
	router := mux.NewRouter()

	//router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../static/"))))
	router.HandleFunc("/getdata", handlers.GetData).Methods("GET")
	fmt.Println("Server started in " + PORTNUMBER)
	http.ListenAndServe(PORTNUMBER, router)
}
