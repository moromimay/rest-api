package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/moromimay/rest-api/models"
	u "github.com/moromimay/rest-api/utils"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(int) //Grab the id of the user that send the request
	contact := &models.User{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetUser(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
