package models

import (
	"strings"
	"time"

	u "rest-api/utils"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	gorm.Model
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birthday"`
	Token     string    `json:"-"`
}

//Validate incoming user details...
func (account *Account) Validate() (map[string]interface{}, bool) {

	if account.Name == "" {
		return u.Message(false, "User name should be on the payload"), false
	}

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password has to have at 6 digits"), false
	}

	// if !account.BirthDate.Format("02-01-2006") {
	// 	return u.Message(false, "Birth date is required"), false
	// }

	//Email must be unique
	temp := &Account{}

	//check for errors and duplicate emails
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in suse by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (account *Account) Create() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	account.BirthDate.Format("2006-01-02")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	account.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(email, password string) map[string]interface{} {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	account.Password = ""

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}
