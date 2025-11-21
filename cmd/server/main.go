package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	// On strip le "/static/" du path pour servir les fichiers statiques
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//http.HandleFunc("/", power4.HomeHandler)
	//http.HandleFunc("/game", power4.GameHandler)

	log.Println("Groupie tracker server started listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}
