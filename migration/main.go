package main

import (
	"encoding/json"
	"fmt"
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

func mdmHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MDM Handler")
	w.Header().Set("Content-Type", "application/json")
	mdmMembers := Members{
		Members: []Person{
			{
				First: "John",
				Last:  "Doe",
			},
			{
				First: "Jane",
				Last:  "Doe",
			},
			{
				First: "Harry",
				Last:  "Potter",
			},
		},
	}
	membersJs, _ := json.Marshal(mdmMembers)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(membersJs)

}

func egressHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Egress Handler")
	w.Header().Set("Content-Type", "application/json")
	egressMembers := Members{
		Members: []Person{
			{
				First: "Rubeus",
				Last:  "Hagrid",
			},
			{
				First: "Hermoine",
				Last:  "Granger",
			},
			{
				First: "Ron",
				Last:  "Weisley",
			},
		},
	}
	membersJs, _ := json.Marshal(egressMembers)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(membersJs)
}

func restrictedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restrictedMemebers := Members{
		Members: []Person{
			{
				First: "Santa",
				Last:  "Claus",
			},
			{
				First: "Easter",
				Last:  "Bunny",
			},
			{
				First: "Bad",
				Last:  "Bunny",
			},
		},
	}
	membersJs, _ := json.Marshal(restrictedMemebers)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(membersJs)
}
func EnableCors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		f(w, r)
	}
}
func handleRequests() {
	http.HandleFunc("/members/mdm", EnableCors(mdmHandler))
	http.HandleFunc("/members/egress", egressHandler)
	http.HandleFunc("/members/restricted", restrictedHandler)
	log.Fatal(http.ListenAndServe(":8084", nil))
}

func main() {
	handleRequests()
}
