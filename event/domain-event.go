package event

import "encoding/json"

// DomainEvent basic domain event structure
type DomainEvent struct {
	Type        string `json:"type"`
	Version     string `json:"version"`
	AggregateID string `json:"aggregate_id"`
	Payload     string `json:"payload"`
}

// Serialize the DomainEvent to send as string
func (d *DomainEvent) Serialize() (string, error) {
	s, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	return string(s), nil
}
