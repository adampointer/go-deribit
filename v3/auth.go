package deribit

import (
	"fmt"
	"time"

	"github.com/adampointer/go-deribit/v3/client/operations"
)

// Renew login 10 minutes before we have to
const renewBefore int64 = 600

func (e *Exchange) Authenticate(key, secret string) error {
	client := e.Client()
	auth, err := client.GetPublicAuth(&operations.GetPublicAuthParams{ClientID: key, ClientSecret: secret, GrantType: "client_credentials"})
	if err != nil {
		return fmt.Errorf("error authenticating: %s", err)
	}
	e.auth = auth.Payload
	d, err := time.ParseDuration(fmt.Sprintf("%ds", *(e.auth.Result.ExpiresIn)-renewBefore))
	if err != nil {
		return fmt.Errorf("unable to parse %ds as a duration: %s", *(e.auth.Result.ExpiresIn)-renewBefore, err)
	}
	go func() {
		time.Sleep(d)
		auth, err := client.GetPublicAuth(&operations.GetPublicAuthParams{RefreshToken: *e.auth.Result.RefreshToken, GrantType: "refresh_token"})
		if err != nil {
			e.errors <- fmt.Errorf("error authenticating: %s", err)
		}
		e.auth = auth.Payload
	}()
	return nil
}
