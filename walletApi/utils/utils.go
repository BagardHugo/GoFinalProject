package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func SendHttpError(httpErrorCode int, message string, w http.ResponseWriter, err error) {
	log.Print(err)
	w.WriteHeader(httpErrorCode)
	_, err = w.Write(BuildHttpErrorMessage(message))
	if err != nil {
		log.Printf("%s : %s", message, err)
	}
}

func CheckHttpMethod(method string, w http.ResponseWriter, r *http.Request) error {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write(BuildHttpErrorMessage("Only POST method is allowed"))
		if err != nil {
			log.Printf("Unable to write to http response : %s", err)
		}

		return errors.New("Method not allowed")
	}

	return nil
}

func BuildHttpErrorMessage(message string) []byte {
	resp := make(map[string]string)
	resp["error"] = message
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Print(err)
	}

	return jsonResp
}
