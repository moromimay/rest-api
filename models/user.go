package models

import (
	"fmt"
	"time"

	u "github.com/moromimay/rest-api/utils"
)

type User struct {
	UserId    int
	Name      string
	Email     string
	Password  string
	BirthDate time.Time
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (user *User) Validate() (map[string]interface{}, bool) {

	if user.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	if user.Name == "" {
		return u.Message(false, "User name should be on the payload"), false
	}

	if user.Email == "" {
		return u.Message(false, "Email number should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (user *User) Create() map[string]interface{} {

	if resp, ok := user.Validate(); !ok {
		return resp
	}

	GetDB().Create(user)

	resp := u.Message(true, "success")
	resp["user"] = user
	return resp
}

func GetUser(id uint) *User {

	user := &User{}
	err := GetDB().Table("users").Where("id = ?", id).First(user).Error
	if err != nil {
		return nil
	}
	return user
}

func GetUsers(user uint) []*User {

	users := make([]*User, 0)
	err := GetDB().Table("users").Where("user_id = ?", user).Find(&users).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return users
}
