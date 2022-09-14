package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"finalProject/account"
	"finalProject/constants"
	"finalProject/wallet"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateHttpErrorMessage(message string) []byte {
	resp := make(map[string]string)
	resp["error"] = message
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Print(err)
	}

	return jsonResp
}

func SendHttpError(httpErrorCode int, w http.ResponseWriter, err error) {
	log.Print(err)
	w.WriteHeader(httpErrorCode)
	_, err = w.Write(CreateHttpErrorMessage(err.Error()))
	if err != nil {
		log.Printf("%s : %s", err.Error(), err)
	}
}

func CheckHttpMethod(method string, w http.ResponseWriter, r *http.Request) error {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write(CreateHttpErrorMessage("Only POST method is allowed"))
		if err != nil {
			log.Printf("Unable to write to http response : %s", err)
		}

		return errors.New(constants.HTTP_METHODE_NOT_ALLOWED)
	}

	return nil
}

func CheckData(account account.Account, w http.ResponseWriter) error {
	log.Print(account)
	err := checkUserName(account.UserName)
	if err != nil {
		SendHttpError(http.StatusBadRequest, w, err)
		return err
	}

	err = CheckPassword(account.Password)
	if err != nil {
		SendHttpError(http.StatusBadRequest, w, err)
		return err
	}

	err = checkPinCode(account.PinCode)
	if err != nil {
		SendHttpError(http.StatusBadRequest, w, err)
		return err
	}
	return nil
}

func checkUserName(userName string) error {
	// Username naming policy
	if !constants.RegexUsername.MatchString(userName) {
		return errors.New(constants.USERNAME_FORMAT_ERROR)
	}
	return nil
}

func CheckPassword(password string) error {
	// Password policy
	if !constants.RegexPassword.MatchString(password) {
		return errors.New(constants.PASSWORD_FORMAT_ERROR)
	}
	return nil
}

func checkPinCode(pinCode string) error {
	// Pincode policy
	if !constants.RegexPinCode.MatchString(pinCode) {
		return errors.New(constants.PIN_CODE_FORMAT_ERROR)
	}
	return nil
}

func CreateWallet(account account.Account, w http.ResponseWriter) (wallet.MockWalletReturn, error) {
	// Prepare external service call
	walletMock := wallet.MockWalletCall{
		Pincode:    account.PinCode,
		Blockchain: "ETH",
	}

	// Serialize external call arguments
	data, err := json.Marshal(walletMock)
	if err != nil {
		SendHttpError(http.StatusInternalServerError, w, err)
		return wallet.MockWalletReturn{}, err
	}

	// Send request to external service
	response, err := http.Post(fmt.Sprintf("http://%s:%s/wallets/create", constants.MOCK_SERVER_ADDRESS, constants.MOCK_SERVER_PORT), "application/json", bytes.NewBuffer(data))
	if err != nil {
		SendHttpError(http.StatusInternalServerError, w, err)
		return wallet.MockWalletReturn{}, err
	}

	// Read body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		SendHttpError(http.StatusInternalServerError, w, err)
		return wallet.MockWalletReturn{}, err
	}

	// Check error
	if response.StatusCode != 200 {
		SendHttpError(http.StatusInternalServerError, w, errors.New(string(responseBody)))
		return wallet.MockWalletReturn{}, err
	}

	// Handle body closure
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("IO error, Unable to close http request body : %s", err)
		}
	}(response.Body)

	// Map body to struct
	mockWallet := wallet.MockWalletReturn{}
	err = json.Unmarshal(responseBody, &mockWallet)
	if err != nil {
		SendHttpError(http.StatusInternalServerError, w, err)
		return wallet.MockWalletReturn{}, err
	}

	return mockWallet, nil
}
