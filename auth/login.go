package auth

import (
	"context"
	"log"

	"github.com/Nerzal/gocloak/v6"
)

// Permissions for the user
func Permissions() {
	hostname := "localhost"
	clientid := ""
	clientSecret := ""
	realm := ""
	client := gocloak.NewClient(hostname)
	ctx := context.Background()
	token, err := client.LoginClient(ctx, clientid, clientSecret, realm)
	if err != nil {
		panic("Login failed:" + err.Error())
	}

	rptResult, err := client.RetrospectToken(ctx, token.AccessToken, clientid, clientSecret, realm)
	if err != nil {
		panic("Inspection failed:" + err.Error())
	}

	if *rptResult.Active == false {
		panic("Token is not active")
	}

	permissions := rptResult.Permissions
	log.Print("Getting user persmissions: ", permissions)
}
