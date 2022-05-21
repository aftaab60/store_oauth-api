package access_token

import (
	"testing"
	"time"
)

func TestGetNewAccessToken(t *testing.T) {
	accessToken := GetNewAccessToken()
	if accessToken.IsExpired() {
		t.Error("new access token should not be expired")
	}
	if accessToken.AccessToken != "" {
		t.Error("new access should not have tokenId")
	}
	if accessToken.UserId != 0 {
		t.Error("new access should not have userId")
	}
}

func TestAccessToken_IsExpired(t *testing.T) {
	accessToken := &AccessToken{}
	if !accessToken.IsExpired() {
		t.Error("empty access token should be expired by default")
	}
	accessToken.Expires = time.Now().UTC().Add(time.Hour * 3).Unix()
	if accessToken.IsExpired() {
		t.Error("access token expiring 3 hours from now should not be expired now")
	}
}
