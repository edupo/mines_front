package main

import (
	"encoding/json"
	"log"
	"mines"
	"net/http"
)

var game *mines.Game
var port string = "8080"

func handleMinesGame(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := json.NewEncoder(w).Encode(game.State()); err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else if r.Method == "POST" {
		var command mines.Command
		if err := json.NewDecoder(r.Body).Decode(&command); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func main() {

	game = mines.NewGame(10, 10, 10)
	game.Seed(42)
	game.InitializeClassic()
	game.Uncover(8)
	game.Print()
	http.HandleFunc("/mines", handleMinesGame)

	log.Printf("Listenting on port http://localhost:%v/mines ...", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
