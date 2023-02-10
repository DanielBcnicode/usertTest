package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"usertest.com/config"
	"usertest.com/user"
)

func main() {
	config := config.GetConfig()
	log.Print(config)
	log.Print("Server started at port 8088")
	http.ListenAndServe(":8088", Server())
}

func Server() *mux.Router {
	s := mux.NewRouter()
	s.HandleFunc("/health_check", health_check)

	server := s.PathPrefix("/api/v1/").Subrouter()
	server.HandleFunc("/user", user.ListUserController()).Methods("GET")
	//server.HandleFunc("/user/{id}", user.ListUserController()).Methods("GET")
	server.HandleFunc("/user", user.AddNewUserController()).Methods("POST")
	server.HandleFunc("/user/{id}", user.DeleteUserController()).Methods("DELETE")
	server.HandleFunc("/user/{id}", user.UpdateUserController()).Methods("PUT")

	return s
}

func health_check(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
