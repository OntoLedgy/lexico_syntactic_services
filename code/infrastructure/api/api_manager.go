package api

import (
	"encoding/json"
	"fmt"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/issues"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//TODO add code here for api for exposing services

type JsonResponse struct {
	Type    string          `json:"type"`
	Data    []issues.Issues `json:"data"`
	Message string          `json:"message"`
}

func RunApi() {

	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	router.HandleFunc(
		"/syntacticChecker/",
		RunSyntacticCheck).
		Methods("GET")

	// serve the app
	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func RunSyntacticCheck(
	w http.ResponseWriter,
	r *http.Request) {

	printMessage("Running Syntactic Check...")

	issue := issues.Issues{}
	issue.Issue_type.Issue_type_name = "testing"
	issue.Issue_type.Issue_check_type = "regex"

	issues := []issues.Issues{}

	issues = append(issues, issue)

	var response = JsonResponse{
		Type: "success",
		Data: issues}

	json.NewEncoder(w).Encode(response)
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}
