package deribit

import (
	"fmt"
	"time"

	"github.com/adampointer/go-deribit/v3/client/operations"
)

// Renew login 10 minutes before we have to
const renewBefore int64 = 600

// Keep clientID and clientSecret for re-connecting automatically
var (
	clientID     string
	clientSecret string
)

func authArgsHelper(keys []string) error {
	if len(keys) == 2 {
		clientID = keys[0]
		clientSecret = keys[1]
	} else {
		if clientID == "" || clientSecret == "" {
			return fmt.Errorf("API Key and Secret must be provided")
		}
	}
	return nil
}

// Authenticate : 1st param: API Key, 2nd param: API Secret
func (e *Exchange) Authenticate(keys ...string) error {
	if err := authArgsHelper(keys); err != nil {
		return err
	}
	auth, err := e.Client().GetPublicAuth(&operations.GetPublicAuthParams{ClientID: clientID, ClientSecret: clientSecret, GrantType: "client_credentials"})
	if err != nil {
		return fmt.Errorf("error authenticating: %s", err)
	}
	e.auth = auth.Payload
	d, err := time.ParseDuration(fmt.Sprintf("%ds", *(e.auth.Result.ExpiresIn)-renewBefore))
	if err != nil {
		return fmt.Errorf("unable to parse %ds as a duration: %s", *(e.auth.Result.ExpiresIn)-renewBefore, err)
	}
	go e.refreshAuth(d)
	return nil
}

func (e *Exchange) refreshAuth(d time.Duration) {
	time.Sleep(d)
	auth, err := e.Client().GetPublicAuth(&operations.GetPublicAuthParams{RefreshToken: *e.auth.Result.RefreshToken, GrantType: "refresh_token"})
	if err != nil {
		e.errors <- fmt.Errorf("error authenticating: %s", err)
	}
	e.auth = auth.Payload
}
