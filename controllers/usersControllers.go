package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"rest-api/models"
	u "rest-api/utils"
)

var Transaction = func(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	operation := &models.Transaction{}
	err := json.NewDecoder(r.Body).Decode(operation)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	user, _ := strconv.ParseUint(id, 10, 32)
	operation.UserID = user
	resp := operation.Create()
	u.Respond(w, resp)
}

var GetUser = func(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	idUser, _ := strconv.ParseUint(id, 10, 32)
	transactions := models.GetTransactionUser(idUser)

	if len(transactions) == 0 {
		u.Respond(w, u.Message(false, "This User hasn't made any transactions"))
		return
	} else {
		resp := u.Message(true, "success")
		resp["transactions"] = transactions
		u.Respond(w, resp)
		return
	}
}

var GetDate = func(w http.ResponseWriter, r *http.Request) {

	date := mux.Vars(r)["operation_date"]
	data := models.GetDateTransaction(date)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
