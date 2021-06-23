package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// URLDescription struct
type URLDescription struct {
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

const port string = ":4000"

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         "/blocks",
			Method:      "POST",
			Description: "Add a block",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	// b, err := json.Marshal(data)
	// utils.HandleError(err)
	// fmt.Fprintf(rw, "%s\n", b)
	json.NewEncoder(rw).Encode(data)
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Server listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
