package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/diegoahg/blockchain/app"
	"github.com/julienschmidt/httprouter"
)

func Init(ai *app.App) {
	router := ai.Router
	router.GET("/api/blocks", GetBlockHandler)
	router.POST("/api/blocks", PostBlockHandler)
	router.PUT("/api/hack", HackBlockHandler)

	log.Println("Initialized api")
}

// GetBlockHandler write blockchain when we receive an http request
func GetBlockHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bytes, err := json.MarshalIndent(app.Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !app.IsChainValid(app.Blockchain) {
		io.WriteString(w, "Chain is not valid")
		return
	}
	io.WriteString(w, string(bytes))
}

// PostBlockHandler takes JSON payload as an input for heart rate (Car)
func PostBlockHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var input app.CarInput

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	newBlock, err := app.GenerateBlock(input.LicensePlate, input.Owner)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, input)
		return
	}

	if app.IsBlockValid(newBlock, app.Blockchain[len(app.Blockchain)-1]) {
		newBlockchain := append(app.Blockchain, newBlock)
		app.ReplaceChain(newBlockchain)
		spew.Dump(app.Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

// PostBlockHandler takes JSON payload as an input for heart rate (Car)
func HackBlockHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var input app.HackInput

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	hackBlock := app.HackBlock(input.Index, input.Hash, input.Owner)
	respondWithJSON(w, r, http.StatusOK, hackBlock)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
