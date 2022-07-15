package oauth

import (
	"microservices-ex-app/src/api/utils/errors"
)

const (
	queryGetUerByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User{
		"jhon": &User{
			Id:       123,
			Username: "Jhon",
		},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError("no user found with given parameters")
	}
	return user, nil
}
