package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/api/ticTacToe"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("api running on port 3000...")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", ticTacToe.HomeLink)
	router.HandleFunc("/playermove", ticTacToe.PlayerMove).Methods("POST", "OPTIONS")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":3000",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
