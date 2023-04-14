package server

import (
	"auth/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Members struct {
	Members []Person `json:"members"`
}
type Person struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type Token struct {
	RealmAccess struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Email      string `json:"email"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	ExpiresIn    int      `json:"expiresIn"`
	Roles        []string `json:"roles"`
	GivenName    string   `json:"given_name"`
	FamilyName   string   `json:"family_name"`
	Email        string   `json:"email"`
}

type controller struct {
	keycloak *keycloak
}

func newController(keycloak *keycloak) *controller {
	return &controller{
		keycloak: keycloak,
	}
}

func (c *controller) login(w http.ResponseWriter, r *http.Request) {

	rq := &loginRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(rq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jwt, err := c.keycloak.gocloak.Login(context.Background(),
		c.keycloak.clientId,
		c.keycloak.clientSecret,
		c.keycloak.realm,
		rq.Username,
		rq.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	// var token Token
	_, claims, err := c.keycloak.gocloak.DecodeAccessToken(context.Background(), jwt.AccessToken, c.keycloak.realm)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
		return
	}
	// arr := strings.Split(jwt.AccessToken, ".")
	// fmt.Printf("jwt: %+v\n", arr[1])
	// uDec, _ := b64.URLEncoding.DecodeString(arr[1])
	// uDec = append(uDec, '}')
	// fmt.Println("uDec: ", string(uDec))
	// unmarshall_err := json.Unmarshal(uDec, &token)
	// if unmarshall_err != nil {
	// 	log.Fatal("Unmarshall error: ", unmarshall_err.Error())
	// }
	fmt.Printf("\nCLIAMS %+v\n", claims)
	realm_access, _ := (*claims)["realm_access"].(map[string]interface{})
	rs := &loginResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
		Roles:        utils.ExtractRoles(realm_access),
		GivenName:    (*claims)["given_name"].(string),
		FamilyName:   (*claims)["family_name"].(string),
		Email:        (*claims)["email"].(string),
	}

	rsJs, _ := json.Marshal(rs)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rsJs)
}

const URL = "http://localhost:8084/members/"
const (
	MDM        = URL + "mdm"
	EGRESS     = URL + "egress"
	RESTRICTED = URL + "restricted"
)

func API(url string) Members {
	var members Members
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &members)

	// fmt.Println(members[0].First)
	return members
}
func (c *controller) getEgress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	members, _ := json.Marshal(API(EGRESS))
	_, _ = w.Write(members)
}
func (c *controller) getMDM(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	members, _ := json.Marshal(API(MDM))
	_, _ = w.Write(members)
}
func (c *controller) getRestricted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	members, _ := json.Marshal(API(RESTRICTED))
	_, _ = w.Write(members)
}
