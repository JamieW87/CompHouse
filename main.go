package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thedevsaddam/gojsonq"
)

//Open the json file using gojsonq package and assign it to jq
var jq = gojsonq.New().File("./data/games.json")

func getReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	report := jq.From("games").Only("title", "likes")
	json.NewEncoder(w).Encode(&report)

}

//Function that returns game specified by the parameter
func getGame(w http.ResponseWriter, r *http.Request) {
	//Set header type to json
	w.Header().Set("Content-Type", "application/json")
	//Pull the parameter from the url then convert it to an int
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	//Query the json using the parameter.
	game := jq.Copy().From("games").WhereEqual("id", id).Get()
	json.NewEncoder(w).Encode(&game)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Simple get request.
	games := jq.Get()
	json.NewEncoder(w).Encode(&games)
}

func main() {

	//Initialises basic router and endpoints
	r := mux.NewRouter()
	r.HandleFunc("/", getAll).Methods("GET")
	r.HandleFunc("/games/{id:[0-9]+}", getGame).Methods("GET")
	r.HandleFunc("/games/report", getReport).Methods("GET")

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)

}

/*
func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Open jsonFile and handle the error
	jsonFile, err := os.Open("./data/games.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened games.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	var games models.Games
	json.NewDecoder(jsonFile).Decode(&games)
	json.NewEncoder(w).Encode(games)
}
*/
