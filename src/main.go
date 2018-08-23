package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	"../config"
	"./routers"
	"./schemas"
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

	userSchema := schemas.UserSchema{Session: session}
	userRouter := routers.UserRouter{Router: router, Schema: userSchema}
	userRouter.MakeRouter()

	port := os.Args[1]
	fmt.Println("Starting data service on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func mrHappy(w http.ResponseWriter, r *http.Request) {
	// TODO: return response object instead of string
	json.NewEncoder(w).Encode("happy")
}
