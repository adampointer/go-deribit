package deribit

import (
	"fmt"
	"github.com/adampointer/go-deribit/models/public"
	"time"
)

// Renew login 10 minutes before we have to
const renewBefore  = 600

func (e *Exchange) Authenticate(key, secret string) error {
	auth, err := e.PublicAuth(&public.AuthRequest{ClientID: key, ClientSecret: secret, GrantType: "client_credentials"})
	if err != nil {
		return fmt.Errorf("error authenticating: %s", err)
	}
	e.auth = auth
	d, err := time.ParseDuration(fmt.Sprintf("%ds", auth.ExpiresIn-renewBefore))
	if err != nil {
		return fmt.Errorf("unable to parse %ds as a duration: %s", auth.ExpiresIn-renewBefore, err)
	}
	go func() {
		time.Sleep(d)
		if err := e.Authenticate(key, secret);err != nil {
			e.errors <-err
		}
	}()
	return nil
}
