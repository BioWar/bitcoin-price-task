package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "encoding/json"
    "fmt"
    // "net/url"
    "io/ioutil"
	"errors"

	"example/csv_utils"
	"example/mail_utils"
)

var binanceURL string = "https://api.binance.com/api/v3/avgPrice?symbol=BTCUAH"
var emailFile string = "example.csv"

type BinanceResponse struct {
	Mins  int    `json:"mins"`
	Price string `json:"price"`
}

type EmailSubscribeRequest struct {
	Email string `json:"email"`
}

type Rate struct {
	BTCUAH string `json:"rate"`
}

func main() {
    router := gin.Default()
    router.GET("/rate", getRate)
    router.POST("/subscribe", postSubscribe)
    router.POST("/sendEmails", postSendEmails)

    router.Run("0.0.0.0:12321")
}


func sendRequestForPrice() (string, error){
    // Get request
    resp, err := http.Get(binanceURL)
    if err != nil {
        fmt.Println("No response from request")
		return "", errors.New("No response from request")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body) // response body is []byte

    var result BinanceResponse
    if err := json.Unmarshal(body, &result); err != nil {  // Parse []byte to the go struct pointer
        fmt.Println("Can not unmarshal JSON")
		return "", errors.New("Can not unmarshal JSON")
    }

    // fmt.Println(PrettyPrint(result))

    // Loop through the data node for the FirstName
    // fmt.Println(result.Price)
	return result.Price, nil
}

func getRate(c *gin.Context) {
	rate, err := sendRequestForPrice()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "")
	}
    c.IndentedJSON(http.StatusOK, rate)
}

func postSubscribe(c *gin.Context) {
	var newEmail EmailSubscribeRequest

	if err := c.BindJSON(&newEmail); err != nil {
		c.IndentedJSON(http.StatusConflict, "")
		return
	}

	returnValue := csv_utils.WriteEmailRecordToFile(emailFile, newEmail.Email)

	if returnValue != 0{
		c.IndentedJSON(http.StatusConflict, "")
		return
	}

    c.IndentedJSON(http.StatusOK, "")
}

func postSendEmails(c *gin.Context){
	emailRecords := csv_utils.ReadCSV(emailFile)
	rate, err := sendRequestForPrice()
	if err != nil {
		return
	}

	for _, emailRecord := range emailRecords {
		mail_utils.SendMail(emailRecord.Email, rate)
		fmt.Println("Email sent to " + emailRecord.Email)
	}

    c.IndentedJSON(http.StatusOK, "")
}
