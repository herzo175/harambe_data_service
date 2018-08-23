package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../schemas"
	"../util"
	"github.com/gorilla/mux"
)

type AlgorithmRouter struct {
	Router *mux.Router
	Schema schemas.AlgorithmSchema
}

func (ar *AlgorithmRouter) MakeRouter() {
	subrouter := ar.Router.PathPrefix("/algorithms").Subrouter()
	// TODO: option route
	subrouter.HandleFunc("", ar.GetAll).Methods("GET")
	subrouter.HandleFunc("", ar.Create).Methods("POST")
	subrouter.HandleFunc("/{id}", ar.GetById).Methods("GET")
	subrouter.HandleFunc("/{id}", ar.Update).Methods("PUT", "PATCH")
}

func (ar *AlgorithmRouter) GetAll(w http.ResponseWriter, r *http.Request) {
	algorithms, err := ar.Schema.GetAll(util.GenerateQueryFromURLQuery(r.URL.Query(), schemas.Algorithm{}))

	if err != nil {
		http.Error(w, err.Error(), 400)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data, err := json.Marshal(algorithms)

	if err != nil {
		http.Error(w, err.Error(), 500)
		panic(err)
	}

	w.Write(data)
}

func (ar *AlgorithmRouter) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	algorithm, err := ar.Schema.GetById(params["id"])

	if err != nil {
		http.Error(w, err.Error(), 400)
		// TODO: advanced logging
		fmt.Println("AlgorithmRouter.GetById " + err.Error() + " " + params["id"])
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data, err := json.Marshal(algorithm)

	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Println("AlgorithmRouter.GetById " + err.Error() + " " + params["id"])
		return
	}

	w.Write(data)
}

func (ar *AlgorithmRouter) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body schemas.Algorithm
	err := decoder.Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("AlgorithmRouter.Create " + err.Error())
		return
	}

	err = ar.Schema.Create(&body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("AlgorithmRouter.Create " + err.Error())
		return
	}
}

func (ar *AlgorithmRouter) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var body schemas.Algorithm
	err := decoder.Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("AlgorithmRouter.Update " + err.Error() + " " + params["id"])
		return
	}

	err = ar.Schema.Update(params["id"], &body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("AlgorithmRouter.Update " + err.Error() + " " + params["id"])
		return
	}
}
