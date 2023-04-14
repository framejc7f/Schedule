package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	wf "github.com/framejc7f/Schedule/WorkingWithFiles"
	"github.com/framejc7f/Schedule/structs"
	"github.com/gorilla/mux"
)

var weeks = wf.ReadFile("ИИТ-22-о.xlsx")

func getWeek(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range weeks {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&structs.Week{})
}

func getDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	d, _ := strconv.Atoi(params["day"])
	sb, _ := strconv.Atoi(params["sb"])
	if d > 6 {
		json.NewEncoder(w).Encode(&structs.Day{})
		return
	}
	day := wf.ReadDay("ИИТ-22-о.xlsx", weeks[1].Name, d, sb)
	json.NewEncoder(w).Encode(day)
}

// func getDay(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for _, item := range weeks {
// 		if item.Id == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&structs.Week{})
// }

func Server() {
	// timeNow := time.Now().Hour()
	// fmt.Println("Hour now - ", timeNow)
	// if (timeNow == 6) || (timeNow == 23) {
	// 	// wf.Parsing()
	// }

	// for i := 0; i <= 6; i++ {

	// }
	// var week time.Weekday = 1
	// fmt.Println(week.String())
	r := mux.NewRouter()
	r.HandleFunc("/home/{id}", getWeek).Methods("GET")
	r.HandleFunc("/home/{id}/{day}/{sb}", getDay).Methods("GET")
	http.Handle("/", r)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)

}
