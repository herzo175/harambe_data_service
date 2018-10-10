package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../schemas"
	"../util"
	"github.com/gorilla/mux"
)

type ReportRouter struct {
	Router *mux.Router
	Schema schemas.ReportSchema
}

func (rr *ReportRouter) MakeRouter() {
	subrouter := rr.Router.PathPrefix("/reports").Subrouter()
	// TODO: option route
	subrouter.HandleFunc("", rr.GetAll).Methods("GET")
	subrouter.HandleFunc("", rr.Create).Methods("POST")
	subrouter.HandleFunc("/{id}", rr.GetById).Methods("GET")
	subrouter.HandleFunc("/{id}", rr.Update).Methods("PUT", "PATCH")
}

func (rr *ReportRouter) GetAll(w http.ResponseWriter, r *http.Request) {
	reports, err := rr.Schema.GetAll(util.GenerateQueryFromURLQuery(r.URL.Query(), schemas.Report{}))

	if err != nil {
		http.Error(w, err.Error(), 400)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data, err := json.Marshal(reports)

	if err != nil {
		http.Error(w, err.Error(), 500)
		panic(err)
	}

	w.Write(data)
}

func (rr *ReportRouter) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	report, err := rr.Schema.GetById(params["id"])

	if err != nil {
		http.Error(w, err.Error(), 400)
		// TODO: advanced logging
		fmt.Println("ReportRouter.GetById " + err.Error() + " " + params["id"])
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data, err := json.Marshal(report)

	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Println("ReportRouter.GetById " + err.Error() + " " + params["id"])
		return
	}

	w.Write(data)
}

func (rr *ReportRouter) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body schemas.Report
	err := decoder.Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("ReportRouter.Create " + err.Error())
		return
	}

	err = rr.Schema.Create(&body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("ReportRouter.Create " + err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data, err := json.Marshal(body)

	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Println("ReportRouter.Create " + err.Error())
		return
	}

	w.Write(data)
}

func (rr *ReportRouter) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var body schemas.Report
	err := decoder.Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("ReportRouter.Update " + err.Error() + " " + params["id"])
		return
	}

	err = rr.Schema.Update(params["id"], &body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("ReportRouter.Update " + err.Error() + " " + params["id"])
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data, err := json.Marshal(body)

	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Println("ReportRouter.Update " + err.Error())
		return
	}

	w.Write(data)
}
