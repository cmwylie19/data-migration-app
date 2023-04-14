package server

import (
	"auth/pkg/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type httpServer struct {
	server *http.Server
	router *mux.Router
}

func NewServer(port string, keycloak *keycloak) *httpServer {

	// create a root router
	router := mux.NewRouter()

	// enable CORS

	// add a subrouter based on matcher func
	// note, routers are processed one by one in order, so that if one of the routing matches other won't be processed
	noAuthRouter := router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return r.Header.Get("Authorization") == ""
	}).Subrouter()

	// add one more subrouter for the authenticated service methods
	authRouter := router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return true
	}).Subrouter()

	// instantiate a new controller which is supposed to serve our routes
	controller := newController(keycloak)

	// map url routes to controller's methods
	noAuthRouter.HandleFunc("/login/{username}/{password}", func(writer http.ResponseWriter, request *http.Request) {

		controller.login(writer, request)
	}).Methods("GET")

	authRouter.HandleFunc("/migration/egress", func(writer http.ResponseWriter, request *http.Request) {
		controller.getEgress(writer, request)
	}).Methods("GET")

	authRouter.HandleFunc("/migration/restricted", func(writer http.ResponseWriter, request *http.Request) {
		controller.getRestricted(writer, request)
	}).Methods("GET")

	authRouter.HandleFunc("/migration/mdm", func(writer http.ResponseWriter, request *http.Request) {
		controller.getMDM(writer, request)
	}).Methods("GET")

	// apply middleware
	mdw := newMiddleware(keycloak)
	authRouter.Use(mdw.verifyEgress)

	// create a server object
	s := &httpServer{
		server: &http.Server{
			Addr:         fmt.Sprintf("0.0.0.0:%s", port),
			Handler:      router,
			WriteTimeout: time.Hour,
			ReadTimeout:  time.Hour,
		},
		router: router,
	}

	return s
}

func (s *httpServer) Listen(log utils.Log) error {

	err := s.server.ListenAndServe()
	if err != nil {
		log.Printf("Error while serving: %s", err.Error())

	}
	return err
}
