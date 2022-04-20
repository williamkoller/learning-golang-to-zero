package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/williamkoller/7.1-crud/database"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Fail the write body of request"))
		return
	}

	var user user

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		w.Write([]byte("Error the converter user for struct"))
	}

	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Fail the connect database"))
		return
	}

	statement, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")

	if err != nil {
		w.Write([]byte("Fail in create statement"))
		return
	}

	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email)

	if err != nil {
		w.Write([]byte("Fail in exec statement"))
		return
	}

	idInsert, err := insert.LastInsertId()

	if err != nil {
		w.Write([]byte("Failed to get id entered"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User entered with successfully! Id: %d", idInsert)))

}
