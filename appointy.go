package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Participant struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	RSVP  string `json:"RSVP"`
}

//defining the attributes
type meetings struct {
	ID               string        `json:"ID"`
	Title            string        `json:"Title"`
	Participant      []Participant `json:"Participants"`
	StartTime        string        `json:"StartTime"`
	EndTime          string        `json:"EndTime"`
	Creationtimestap string        `json:"CreateTime"`
}

type allEvents []meetings

var events = allEvents{}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

//function for finding schedulingmeeting

func ScheduleAmeeting(w http.ResponseWriter, r *http.Request) {
	var newEvent meetings
	// Convert r.Body into a readable formart
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event id, title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)

	events = append(events, newEvent)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

//function to get meetings

func getameeting(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func listtimefra(w http.ResponseWriter, r *http.Request) {
	sttime := mux.Vars(r)["id1"]
	entime := mux.Vars(r)["id2"]
	var eve = allEvents{}
	for _, singleEvent := range events {
		if singleEvent.StartTime >= sttime && singleEvent.EndTime <= entime {
			eve = append(eve, singleEvent)
		}
	}
	json.NewEncoder(w).Encode(eve)
}

func listemail(w http.ResponseWriter, r *http.Request) {
	fiemail := mux.Vars(r)["id"]
	var sar = allEvents{}
	for _, singleEvent := range events {
		if singleEvent.ID == fiemail {
			sar = append(sar, singleEvent)
		}
	}
	json.NewEncoder(w).Encode(sar)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/meetings", ScheduleAmeeting).Methods("POST")
	router.HandleFunc("/meetings", getAllEvents).Methods("GET")
	router.HandleFunc("/meetings/{id}", getameeting).Methods("GET")
	router.HandleFunc("/meetings?start=<id1>&end=<id2>", listtimefra).Methods("GET")
	router.HandleFunc("/meetings?participant=<id>", listemail).Methods("GET")
	log.Fatal(http.ListenAndServe(":8072", router))
}