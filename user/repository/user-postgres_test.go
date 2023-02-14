package repository

import (
	"testing"

	"usertest.com/user"
)

func Test_generateUserFilterQuery(t *testing.T) {
	type args struct {
		filter user.RepositoryFilter
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "no filter added",
			args: args{filter: user.RepositoryFilter{Filters: map[string]string{}}},
			want: "",
			wantErr: false,
		},
		{
			name: "filter not allowed",
			args: args{filter: user.RepositoryFilter{Filters: map[string]string{"test":"wrong"}}},
			want: "",
			wantErr: true,
		},
		{
			name: "one filter added",
			args: args{filter: user.RepositoryFilter{Filters: map[string]string{"first_name":"Joe"}}},
			want: ` WHERE "first_name" = "Joe" `,
			wantErr: false,
		},
		{
			name: "two filters added",
			args: args{filter: user.RepositoryFilter{Filters: map[string]string{"first_name":"Joe","last_name":"Greenfield"}}},
			want: ` WHERE "first_name" = "Joe" AND "last_name" = "Greenfield" `,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateUserFilterQuery(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateUserFilterQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("generateUserFilterQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
