package config

import (
	"github.com/caesar-rocks/core"
	"github.com/caesar-rocks/orm"
)

// EnvironmentVariables is a struct that holds all the environment variables that need to be validated.
// For full reference, see: https://github.com/go-playground/validator.
type EnvironmentVariables struct {
	// APP_KEY is the key used for encryption and decryption.
	APP_KEY string `validate:"required" json:"app_key,omitempty"`

	// Addr is the address to listen on for incoming requests.
	ADDR string `validate:"required" json:"addr,omitempty"`

	// DBMS is the database management system to use ("postgres", "mysql", "sqlite").
	DBMS orm.DBMS `validate:"oneof=postgres mysql sqlite" json:"dbms,omitempty"`

	// DSN is the data source name, which is a connection string for the database.
	DSN string `validate:"required" json:"dsn,omitempty"`
}

func ProvideEnvironmentVariables() *EnvironmentVariables {
	return core.ValidateEnvironmentVariables[EnvironmentVariables]()
}
