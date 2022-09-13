package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"wallet_api/server"
	"wallet_api/utils"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/wallets/create", mockCreateWallet)
	err := http.ListenAndServe(fmt.Sprintf(":%s", "5002"), nil)
	if err != nil {
		fmt.Print(err)
	}
}

/*
Mocks a blockchain service to create a wallet
*/
func mockCreateWallet(w http.ResponseWriter, r *http.Request) {
	// Check method
	err := utils.CheckHttpMethod("POST", w, r)
	if err != nil {
		log.Print(err)
		return
	}

	// Read body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.SendHttpError(http.StatusInternalServerError, "Unable to read body", w, err)
		return
	}

	// Handle body closure
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("IO error, Unable to close http request body : %s", err)
		}
	}(r.Body)

	// Map body to struct
	mockServiceExpectation := server.MockServiceExpectation{}
	err = json.Unmarshal(body, &mockServiceExpectation)
	if err != nil {
		utils.SendHttpError(http.StatusInternalServerError, "Unable to deserialize body", w, err)
		return
	}

	// The service should do something here with the provided pincode and blockchain to create the wallet, once it's done, the new wallet address is sent back
	mockServiceAnswer := server.MockServiceAnswer{
		WalletAddress:   uuid.NewString(),
		CurrencyCode:    "ETH",
		CurrencyBalance: "0",
	}

	// Serialize response to json
	mockServiceAnswerJson, err := json.Marshal(mockServiceAnswer)
	if err != nil {
		utils.SendHttpError(http.StatusInternalServerError, "Unable to deserialize response", w, err)
		return
	}

	// Send back the response
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(mockServiceAnswerJson)
	if err != nil {
		log.Printf("Unable to write response : %s", err)
	}
}
