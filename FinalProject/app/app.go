package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"finalProject/account"
	"finalProject/constants"
	"finalProject/utils"
	"fmt"
	"io"
	"log"
	"net/http"
)

var accountRepository *account.PostgreSqlRepository

func HandleCreateAccount(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("pgx", "host=localhost port=5432 user=postgres password=password dbname=database sslmode=disable")

	if err != nil {
		log.Fatal(err)
		fmt.Println("Open db fail")
	}
	defer db.Close()

	accountRepository = account.NewPostgreSqlRepository(db)
	if err := accountRepository.Migrate(); err != nil {
		log.Fatal(err.Error())
	}

	// Check method
	err = utils.CheckHttpMethod("POST", w, r)
	if err != nil {
		log.Print(err.Error())
		return
	}

	// Read body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.SendHttpError(http.StatusInternalServerError, w, err)
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
	account := account.Account{}
	err = json.Unmarshal(body, &account)
	if err != nil {
		utils.SendHttpError(http.StatusInternalServerError, w, err)
		log.Print("Read body fail")
		return
	}

	utils.CheckData(account, w)

	wallet, err := utils.CreateWallet(account, w)
	if err != nil {
		utils.SendHttpError(http.StatusInternalServerError, w, err)
	}

	account, err = accountRepository.Create(account, wallet)
	if errors.Is(err, constants.ErrDuplicate) {
		fmt.Printf("record: %+v already exists\n", account)
	} else if err != nil {
		log.Fatal(err.Error())
	}

	accountJson, err := json.Marshal(account)
	if err != nil {
		utils.SendHttpError(http.StatusInternalServerError, w, err)
		return
	}

	// Send successful response
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(accountJson)
	if err != nil {
		log.Printf("Unable to write response : %s", err)
	}
}
