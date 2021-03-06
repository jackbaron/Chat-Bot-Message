package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	r := mux.NewRouter()
	port, _ := getPort()

	r.HandleFunc("/", indexHandler).Methods("GET")
	fmt.Printf("Server up and running . Running PORT: %s\n", port)

	r.HandleFunc("/webhook", webhookGetHandler).Methods("GET")
	r.HandleFunc("/webhook", webhookPostHandler).Methods("POST")
	r.HandleFunc("/add-message", addMessageHandler).Methods("GET")
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Error listening and server", err)
	}
}

func addMessageHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("add")
	t, _ = t.ParseFiles("tmpl/add.html")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Got my server up and running in Go.  Yay!")
}

func getPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		// run local
		// return ":3500", fmt.Errorf("$PORT not set")
		// deploy
		return "", fmt.Errorf("$PORT not set")
	}

	return ":" + port, nil
}
