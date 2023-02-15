package user

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"usertest.com/broker"
	"usertest.com/event"
)

func AddNewUserController(uRe UserRepo, br broker.Broker) func(w http.ResponseWriter, r *http.Request) {
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

		jsonPayload, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error creating Marshalling user: %s\n", err)
			return
		}

		_, err = w.Write(jsonPayload)
		if err != nil {
			log.Printf("Error writing AddNewUser response: %s\n", err)
			return
		}

		err = uRe.Save(context.TODO(), &data)
		if err != nil {
			log.Printf("Error creating row: %s\n", err)
			return
		}

		de := event.DomainEvent{
			Type:        "UserAdded",
			Version:     "1.0",
			AggregateID: data.ID.String(),
			Payload:     string(jsonPayload),
		}

		br.PublishDomainEvent(&de)
	}
}
