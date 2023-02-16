package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"usertest.com/user"
	"usertest.com/user/common"
)

// ListUserController is the http handle to list users using filters and pagination
func ListUserController(userRepository user.UserRepo) func(w http.ResponseWriter, r *http.Request) {
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
func filterData(r *http.Request) user.RepositoryFilter {
	f := make(map[string]string)

	for i, _ := range common.FilterFields {
		va := r.URL.Query().Get(i)
		if va != "" {
			f[i] = va
		}
	}
	return user.RepositoryFilter{Filters: f}
}

// Get the paginator variables from the URL
// Variable p contains the page wanted, 0 starting
// Variable ps contains the items per page
func paginationData(r *http.Request) user.Paginator {
	p, _ := strconv.Atoi(r.URL.Query().Get("p"))
	ps, _ := strconv.Atoi(r.URL.Query().Get("ps"))

	return user.Paginator{CurrentPage: p, PagSize: ps}
}
