package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	"./config"
	"./src/routers"
	"./src/schemas"
)

func main() {
	session, err := mgo.Dial(config.DBConnectionString)

	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	// register routes
	router.HandleFunc("/up", mrHappy).Methods("GET")

	algorithmSchema := schemas.AlgorithmSchema{Session: session}
	algorithmRouter := routers.AlgorithmRouter{Router: router, Schema: algorithmSchema}
	algorithmRouter.MakeRouter()

	fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func mrHappy(w http.ResponseWriter, r *http.Request) {
	// TODO: return response object instead of string
	json.NewEncoder(w).Encode("happy")
}
