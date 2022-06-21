package rest

import (
	"bytes"
	"encoding/json"
	"github.com/aftaab60/store_oauth-api/src/domain/user"
	"github.com/aftaab60/store_oauth-api/src/utils/errors"
	"io/ioutil"
	"net/http"
)

const (
	userLoginUrl = "http://localhost:8080/users/login"
)

type UserRepository interface {
	Login(loginRequest user.LoginRequest) (*user.User, *errors.RestErr)
}

type userRepository struct{}

func NewRepository() UserRepository {
	return &userRepository{}
}

func (u *userRepository) Login(loginRequest user.LoginRequest) (*user.User, *errors.RestErr) {
	requestBody, err := json.Marshal(loginRequest)
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}
	response, err := http.Post(userLoginUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		return nil, errors.NewInternalServerError(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	var user user.User
	if err := json.Unmarshal(responseData, &user); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &user, nil
}
