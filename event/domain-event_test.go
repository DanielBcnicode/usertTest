package event

import "testing"

func TestDomainEvent_Serialize(t *testing.T) {
	type fields struct {
		Type        string
		Version     string
		AggregateID string
		Payload     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Basic test",
			fields: fields{
				Type: "Basic",
				Version: "Basic",
				AggregateID: "1111-1111-111111-1111",
				Payload: `{"payload":"payload}`,
			},
			want: `{"type":"Basic","version":"Basic","aggregate_id":"1111-1111-111111-1111","payload":"{\"payload\":\"payload}"}`,
			wantErr: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DomainEvent{
				Type:        tt.fields.Type,
				Version:     tt.fields.Version,
				AggregateID: tt.fields.AggregateID,
				Payload:     tt.fields.Payload,
			}
			got, err := d.Serialize()
			if (err != nil) != tt.wantErr {
				t.Errorf("DomainEvent.Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DomainEvent.Serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
