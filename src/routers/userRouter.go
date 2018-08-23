package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../schemas"
	"../util"
	"github.com/gorilla/mux"
)

type UserRouter struct {
	Router *mux.Router
	Schema schemas.UserSchema
}

func (ur *UserRouter) MakeRouter() {
	subrouter := ur.Router.PathPrefix("/users").Subrouter()
	// TODO: option route
	subrouter.HandleFunc("", ur.GetAll).Methods("GET")
	subrouter.HandleFunc("", ur.Create).Methods("POST")
	subrouter.HandleFunc("/{id}", ur.GetById).Methods("GET")
	subrouter.HandleFunc("/{id}", ur.Update).Methods("PUT", "PATCH")
}

func (ur *UserRouter) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := ur.Schema.GetAll(util.GenerateQueryFromURLQuery(r.URL.Query(), schemas.Algorithm{}))

	if err != nil {
		http.Error(w, err.Error(), 400)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data, err := json.Marshal(users)

	if err != nil {
		http.Error(w, err.Error(), 500)
		panic(err)
	}

	w.Write(data)
}

func (ur *UserRouter) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := ur.Schema.GetById(params["id"])

	if err != nil {
		http.Error(w, err.Error(), 400)
		// TODO: advanced logging
		fmt.Println("UserRouter.GetById " + err.Error() + " " + params["id"])
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Println("UserRouter.GetById " + err.Error() + " " + params["id"])
		return
	}

	w.Write(data)
}

func (ur *UserRouter) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body schemas.User
	err := decoder.Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("UserRouter.Create " + err.Error())
		return
	}

	err = ur.Schema.Create(&body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("UserRouter.Create " + err.Error())
		return
	}
}

func (ur *UserRouter) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var body schemas.User
	err := decoder.Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("UserRouter.Update " + err.Error() + " " + params["id"])
		return
	}

	err = ur.Schema.Update(params["id"], &body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("UserRouter.Update " + err.Error() + " " + params["id"])
		return
	}
}
