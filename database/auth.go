package database

import (
	"github.com/coffemanfp/docucentertest/auth"
	"github.com/coffemanfp/docucentertest/client"
)

// AUTH_REPOSITORY is the key to be used when creating the repositories hashmap.
const AUTH_REPOSITORY RepositoryID = "AUTH"

// AuthRepository defines the behaviors to be used by a AuthRepository implementation.
type AuthRepository interface {
	// GetIdAndHashedPassword retrieves the user ID and hashed password for the given authentication data.
	GetIdAndHashedPassword(auth auth.Auth) (id int, hash string, err error)

	// Register registers a new client with authentication and returns the assigned ID.
	Register(client client.Client) (id int, err error)
}
