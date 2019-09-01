package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	u "rest-api/utils"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	UserID           uint64  `json:"user_id"`
	Operation        string  `json:"operation"`
	NumberCoin       int     `json:"quantity_coin"`
	DateTransaction  string  `json:"operation_date"`
	ValueTransaction float64 `json:"price"`
}

type CoinMarketCap struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
	} `json:"status"`
	Data struct {
		ID          int       `json:"id"`
		Symbol      string    `json:"symbol"`
		Name        string    `json:"name"`
		Amount      int       `json:"amount"`
		LastUpdated time.Time `json:"last_updated"`
		Quote       struct {
			BRL struct {
				Price       float64   `json:"price"`
				LastUpdated time.Time `json:"last_updated"`
			} `json:"BRL"`
		} `json:"quote"`
	} `json:"data"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (operation *Transaction) Validate() (map[string]interface{}, bool) {

	if operation.UserID <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	if operation.Operation == "" {
		return u.Message(false, "Operation type should be on the payload"), false
	} else if operation.Operation != "buy" && operation.Operation != "sell" {
		return u.Message(false, "Choose the action: buy or sell"), false
	}

	if operation.NumberCoin <= 0 {
		return u.Message(false, "Choose the quantity of coins"), false
	}

	//All the required parameters are present
	return u.Message(false, "Requirement passed"), true
}

func (operation *Transaction) Create() map[string]interface{} {

	if resp, ok := operation.Validate(); !ok {
		return resp
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/tools/price-conversion", nil)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	number := fmt.Sprintf("%d", operation.NumberCoin)

	q := url.Values{}
	q.Add("id", "1")
	q.Add("amount", number)
	q.Add("convert", "BRL")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "8df08593-5094-448e-8512-cff4b8d273ac")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	var d CoinMarketCap
	err = json.Unmarshal([]byte(respBody), &d)
	if err != nil {
		panic(err)
	}

	operation.ValueTransaction = d.Data.Quote.BRL.Price

	datetime := time.Now()
	date := fmt.Sprintf(datetime.Format("2006-01-02"))
	operation.DateTransaction = date

	GetDB().Create(operation)

	if operation.ID <= 0 {
		return u.Message(false, "Failed to create operation, connection error.")
	}

	respo := u.Message(true, "success")
	respo["operation"] = operation
	return respo
}

func GetTransactionUser(id uint64) []*Transaction {

	user := make([]*Transaction, 0)
	err := GetDB().Table("transactions").Where("user_id = ?", id).Find(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return user
}

func GetDateTransaction(operation_date string) []*Transaction {

	//operationDate, _ := time.Parse("2006-01-02", operation_date)
	//fmt.Print(operationDate)
	date := make([]*Transaction, 0)
	err := GetDB().Table("transactions").Where("operation_date = ?", operation_date).Find(date).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return date
}
