package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"rest-api/models"
	u "rest-api/utils"
)

var Transaction = func(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	user, _ := strconv.ParseUint(id, 10, 32)

	account := &models.Account{}
	err := models.GetDB().Table("accounts").Where("id = ?", user).First(account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		u.Respond(w, u.Message(false, "Connection error. Please retry"))
		return
	}
	if account.ID == 0 {
		u.Respond(w, u.Message(false, "User is not recognized"))
		return
	}

	operation := &models.Transaction{}
	err = json.NewDecoder(r.Body).Decode(operation)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

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

	if data == nil || len(data) == 0 {
		u.Respond(w, u.Message(false, "This date hasn't made any transactions"))
		return
	} else {
		resp := u.Message(true, "success")
		resp["data"] = data
		u.Respond(w, resp)
		return
	}
}
