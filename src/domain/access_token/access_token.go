package access_token

import (
	"github.com/aftaab60/store_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"accessToken"`
	UserId      int64  `json:"userId"`
	ClientId    int64  `json:"clientId"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	if at == nil || strings.TrimSpace(at.AccessToken) == "" {
		return errors.NewBadRequestError("invalid token in request")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid userId in request")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid clientId in request")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time in request")
	}
	return nil
}

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(time.Hour * expirationTime).Unix(),
	}
}

func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) GetAccessTokenById() {

}
