package user

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"usertest.com/user/common"
)

func AddNewUserController(uRe UserRepo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("AddNewUser end-point called")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		log.Printf("Payload: %s\n", string(d))
		data := NewUser()
		json.Unmarshal(d, &data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)

		err = uRe.Save(context.TODO(), &data)
		if err != nil {
			log.Printf("Error creating row: %s\n", err)
		}

	}
}

func UpdateUserController(uRe UserRepo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("UpdateUser end-point called")
		id, ok := mux.Vars(r)["id"]
		if !ok {
			log.Printf("ERROR: UpdateUser, Id not allowed \n")
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		uId, err := uuid.Parse(id)
		if err != nil {
			log.Printf("ERROR: UpdateUser, Id not valid: %s \n", id)
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		d, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		log.Printf("Payload: %s\n", string(d))
		data := NewUser()
		json.Unmarshal(d, &data)
		data.UpdateDate()
		data.ID = uId

		err = uRe.Update(context.TODO(), &data)
		if err != nil {
			log.Printf("ERROR UpdateUser: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Printf("Error UpdateUser encoding: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
}

func DeleteUserController(uRe UserRepo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("DeleteUser end-point called")
		w.Header().Set("Content-Type", "application/json")
		id, ok := mux.Vars(r)["id"]
		if !ok {
			log.Printf("ERROR: DeleteUser, Id not allowed \n")
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		uId, err := uuid.Parse(id)
		if err != nil {
			log.Printf("ERROR: DeleteUser, Id not valid: %s \n", id)
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		err = uRe.Delete(context.TODO(), uId)
		if err != nil {
			log.Printf("ERROR DeleteUser: %s\n", err)
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func ListUserController(userRepository UserRepo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("ListUser end-point called")
		w.Header().Set("Content-Type", "application/json")
		
		paginator := paginationData(r)
		filter := filterData(r)
		d, err := userRepository.FindByFilter(context.TODO(), filter, &paginator)
		if err != nil {
			log.Printf("Error FindByFilter: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(d)
		if err != nil {
			log.Printf("Error FindByFilter encoding: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// Get filters from url. Only get variables defined in repository.FilterFields map
func filterData(r *http.Request) RepositoryFilter {
	f := make(map[string]string)

	for i, _ := range common.FilterFields {
		va := r.URL.Query().Get(i)
		if va != "" {
			f[i] = va
		}
	}
	return RepositoryFilter{Filters: f}
}

// Get the paginator variables from the URL
// Variable p contains the page wanted, 0 starting
// Variable ps contains the items per page
func paginationData(r *http.Request) Paginator {
	p, _ := strconv.Atoi(r.URL.Query().Get("p"))
	ps, _ := strconv.Atoi(r.URL.Query().Get("ps"))

	return Paginator{CurrentPage: p, PagSize: ps}
}
