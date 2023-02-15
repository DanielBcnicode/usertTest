package user

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"usertest.com/broker"
	"usertest.com/event"
)

func UpdateUserController(uRe UserRepo, br broker.Broker) func(w http.ResponseWriter, r *http.Request) {
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

		jsonPayload, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error updating Marshalling user: %s\n", err)
			return
		}

		_, err = w.Write(jsonPayload)
		if err != nil {
			log.Printf("Error UpdateUser encoding: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		de := event.DomainEvent{
			Type:        "UserUpdated",
			Version:     "1.0",
			AggregateID: data.ID.String(),
			Payload:     string(jsonPayload),
		}

		br.PublishDomainEvent(&de)
	}
}
