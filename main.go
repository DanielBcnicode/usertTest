package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"usertest.com/config"
	"usertest.com/user"
	"usertest.com/user/repository"
)

var (
	ListUserController   func(http.ResponseWriter, *http.Request)
	AddNewUserController func(http.ResponseWriter, *http.Request)
	DeleteUserController func(http.ResponseWriter, *http.Request)
	UpdateUserController func(http.ResponseWriter, *http.Request)
)

func main() {
	config := config.GetConfig()
	db, err := initDatabase(config)
	if err != nil {
		log.Fatalf("ERROR: can't initialize db: %s", err)
	}

	defer db.Close()

	initializeControllers(db)

	log.Printf("Configuration %#v\n", config)
	log.Print("Server started at port 8088")
	err = http.ListenAndServe(":8088", Server())
	if err != nil {
		log.Fatalf("ERROR: can't initialize the server: %s", err)
	}
}

func Server() *mux.Router {
	s := mux.NewRouter()
	s.HandleFunc("/health_check", health_check)

	server := s.PathPrefix("/api/v1/").Subrouter()
	server.HandleFunc("/user", ListUserController).Methods("GET")
	//server.HandleFunc("/user/{id}", user.ListUserController()).Methods("GET")
	server.HandleFunc("/user", AddNewUserController).Methods("POST")
	server.HandleFunc("/user/{id}", DeleteUserController).Methods("DELETE")
	server.HandleFunc("/user/{id}", UpdateUserController).Methods("PUT")

	return s
}

func initializeControllers(db *repository.PostgresConn) error {
	userRepository := repository.NewUserPostgresRepository(db)

	AddNewUserController = user.AddNewUserController(&userRepository)
	ListUserController = user.ListUserController(&userRepository)
	UpdateUserController = user.UpdateUserController(&userRepository)
	DeleteUserController = user.DeleteUserController(&userRepository)

	return nil
}

func health_check(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func initDatabase(conf *config.Config) (*repository.PostgresConn, error) {
	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		conf.Db.User,
		conf.Db.Password,
		conf.Db.Host,
		conf.Db.Port,
		conf.Db.Database,
	)

	log.Printf("Connection URL: %s\n", connString)

	return repository.NewPostgresConn(connString)
}
