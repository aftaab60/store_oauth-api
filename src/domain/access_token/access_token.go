package access_token

import "time"

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string
	UserId      int64
	ClientId    int64
	Expires     int64
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
