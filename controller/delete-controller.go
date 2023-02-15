package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"usertest.com/broker"
	"usertest.com/event"
	"usertest.com/user"
)

func DeleteUserController(uRe user.UserRepo, br broker.Broker) func(w http.ResponseWriter, r *http.Request) {
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

		de := event.DomainEvent{
			Type:        "UserDeleted",
			Version:     "1.0",
			AggregateID: uId.String(),
			Payload:     "",
		}

		br.PublishDomainEvent(&de)
	}
}
