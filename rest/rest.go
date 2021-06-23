package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/chiwon99881/chyocoin/blockchain"
	"github.com/chiwon99881/chyocoin/utils"
	"github.com/gorilla/mux"
)

var port string

// URL of custom type
type url string

type errorResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// AddBlockBody of body in post request
type addBlockBody struct {
	Message string `json:"message"`
}

// MarshalText of URL receiver
func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

// URLDescription struct
type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "See all blocks",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{height}"),
			Method:      "GET",
			Description: "Get a single block",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	// b, err := json.Marshal(data)
	// utils.HandleError(err)
	// fmt.Fprintf(rw, "%s\n", b)
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleError(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	height, err := strconv.Atoi(vars["height"])
	utils.HandleError(err)
	block, err := blockchain.GetBlockchain().Getblock(height)
	if err != nil {
		json.NewEncoder(rw).Encode(errorResponse{ErrorMessage: fmt.Sprint(err)})
	} else {
		json.NewEncoder(rw).Encode(block)
	}
}

// Start of rest.go
func Start(aPort int) {
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", aPort)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")
	fmt.Printf("Server listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
