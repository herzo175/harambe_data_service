package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../schemas"
	"../util"
	"github.com/gorilla/mux"
)

type BrokerRouter struct {
	Router *mux.Router
	Schema schemas.BrokerSchema
}

func (br *BrokerRouter) MakeRouter() {
	subrouter := br.Router.PathPrefix("/brokers").Subrouter()
	// TODO: option route
	subrouter.HandleFunc("", br.GetAll).Methods("GET")
	subrouter.HandleFunc("", br.Create).Methods("POST")
	subrouter.HandleFunc("/{id}", br.GetById).Methods("GET")
	subrouter.HandleFunc("/{id}", br.Update).Methods("PUT", "PATCH")
	subrouter.HandleFunc("/{id}", br.Delete).Methods("DELETE")
}

func (br *BrokerRouter) GetAll(w http.ResponseWriter, r *http.Request) {
	brokers, err := br.Schema.GetAll(util.GenerateQueryFromURLQuery(r.URL.Query(), schemas.Broker{}))

	if err != nil {
		http.Error(w, err.Error(), 400)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data, err := json.Marshal(brokers)

	if err != nil {
		http.Error(w, err.Error(), 500)
		panic(err)
	}

	w.Write(data)
}

func (br *BrokerRouter) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	broker, err := br.Schema.GetById(params["id"])

	if err != nil {
		http.Error(w, err.Error(), 400)
		// TODO: advanced logging
		fmt.Println("BrokerRouter.GetById " + err.Error() + " " + params["id"])
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data, err := json.Marshal(broker)

	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Println("BrokerRouter.GetById " + err.Error() + " " + params["id"])
		return
	}

	w.Write(data)
}

func (br *BrokerRouter) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body schemas.Broker
	err := decoder.Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("BrokerRouter.Create " + err.Error())
		return
	}

	err = br.Schema.Create(&body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("BrokerRouter.Create " + err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data, err := json.Marshal(body)

	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Println("BrokerRouter.Create " + err.Error())
		return
	}

	w.Write(data)
}

func (br *BrokerRouter) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var body schemas.Broker
	err := decoder.Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("BrokerRouter.Update " + err.Error() + " " + params["id"])
		return
	}

	err = br.Schema.Update(params["id"], &body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("BrokerRouter.Update " + err.Error() + " " + params["id"])
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data, err := json.Marshal(body)

	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Println("BrokerRouter.Update " + err.Error())
		return
	}

	w.Write(data)
}

func (br *BrokerRouter) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)

	err := br.Schema.Delete(params["id"])

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("BrokerRouter.Delete " + err.Error())
		return
	}
}
