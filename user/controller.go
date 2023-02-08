package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func AddNewUserController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("AddNewUser end-point called")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		log.Printf("Payload: %s\n", string(d))
		data := User{}
		json.Unmarshal(d, &data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func UpdateUserController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("UpdateUser end-point called")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteUserController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("DeleteUser end-point called")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func ListUserController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("ListUser end-point called")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
