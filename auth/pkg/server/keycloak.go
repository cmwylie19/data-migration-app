package server

import (
	"os"

	"github.com/Nerzal/gocloak/v13"
)

type keycloak struct {
	gocloak      gocloak.GoCloak // keycloak client
	clientId     string          // clientId specified in Keycloak
	clientSecret string          // client secret specified in Keycloak
	realm        string          // realm specified in Keycloak
}

// gocloak.NewClient("http://localhost:8081")
func NewKeycloak() *keycloak {
	return &keycloak{
		gocloak:      *gocloak.NewClient(os.Getenv("KC_URL")),
		clientId:     os.Getenv("KC_CLIENT_ID"),
		clientSecret: os.Getenv("KC_CLIENT_SECRET"),
		realm:        os.Getenv("KC_REALM"),
	}
}
