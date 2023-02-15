package common

// FilterFields has the possible url filter fields and the representation as fields in the queries
var FilterFields = map[string]string{
	"first_name": `"first_name"`,
	"last_name":  `"last_name"`,
	"nickname":   `"nickname"`,
	"email":      `"email"`,
	"country":    `"country"`,
}