package server

import (
	"auth/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func EnableCors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		f(w, r)
	}
}

type keyCloakMiddleware struct {
	keycloak *keycloak
}

func newMiddleware(keycloak *keycloak) *keyCloakMiddleware {
	return &keyCloakMiddleware{keycloak: keycloak}
}

func (auth *keyCloakMiddleware) extractBearerToken(token string) string {
	return strings.Replace(token, "Bearer ", "", 1)
}

func ElementExists(arr []string, element string) bool {
	for _, v := range arr {
		fmt.Println(v)
		if v == element {
			return true
		}
	}
	return false
}

func (auth *keyCloakMiddleware) verifyEgress(next http.Handler) http.Handler {

	f := func(w http.ResponseWriter, r *http.Request) {

		// try to extract Authorization parameter from the HTTP header
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// extract Bearer token
		token = auth.extractBearerToken(token)

		if token == "" {
			http.Error(w, "Bearer Token missing", http.StatusUnauthorized)
			return
		}

		//// call Keycloak API to verify the access token
		result, err := auth.keycloak.gocloak.RetrospectToken(context.Background(), token, auth.keycloak.clientId, auth.keycloak.clientSecret, auth.keycloak.realm)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		jwt, claims, err := auth.keycloak.gocloak.DecodeAccessToken(context.Background(), token, auth.keycloak.realm)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		// Check if egress role is assocaited with the token

		realm_access, _ := (*claims)["realm_access"].(map[string]interface{})
		roles := utils.ExtractRoles(realm_access)

		if !ElementExists(roles, "egress") {
			http.Error(w, "Invalid role or unsufficient permissions: need egress role", http.StatusUnauthorized)
			return
		}

		// arr := realm_access["roles"]
		// fmt.Printf("\nType of arr: %T\n", arr)
		// for _, v := range realm_access["roles"] {
		// 	fmt.Println(v)
		// }
		// fmt.Printf("realm_access %+v", realm_access["roles"])
		// roles, _ := realm_access["roles"].([]string)
		// for _, role := range roles {
		// 	fmt.Println(role)
		// }
		// if !ElementExists(arr, "egress") {
		// 	http.Error(w, "Invalid role or unsufficient permissions: need egress role", http.StatusUnauthorized)
		// 	return
		// }

		familyName, _ := (*claims)["family_name"].(string)

		fmt.Printf("\nFamily Name: %v\n", familyName)
		jwtj, _ := json.Marshal(jwt)
		fmt.Printf("token: %v\n", string(jwtj))

		// check if the token isn't expired and valid
		if !*result.Active {
			http.Error(w, "Invalid or expired Token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}

func (auth *keyCloakMiddleware) verifyMDM(next http.Handler) http.Handler {

	f := func(w http.ResponseWriter, r *http.Request) {

		// try to extract Authorization parameter from the HTTP header
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// extract Bearer token
		token = auth.extractBearerToken(token)

		if token == "" {
			http.Error(w, "Bearer Token missing", http.StatusUnauthorized)
			return
		}

		//// call Keycloak API to verify the access token
		result, err := auth.keycloak.gocloak.RetrospectToken(context.Background(), token, auth.keycloak.clientId, auth.keycloak.clientSecret, auth.keycloak.realm)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		jwt, claims, err := auth.keycloak.gocloak.DecodeAccessToken(context.Background(), token, auth.keycloak.realm)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}
		familyName, _ := (*claims)["family_name"].(string)

		fmt.Printf("\nFamily Name: %v\n", familyName)
		jwtj, _ := json.Marshal(jwt)
		fmt.Printf("token: %v\n", string(jwtj))

		// check if the token isn't expired and valid
		if !*result.Active {
			http.Error(w, "Invalid or expired Token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}
