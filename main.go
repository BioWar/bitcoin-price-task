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
    // router.POST("/albums", postAlbums)

    router.Run("localhost:12321")
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

// // postAlbums adds an album from JSON received in the request body.
// func postAlbums(c *gin.Context) {
//     var newAlbum album

//     // Call BindJSON to bind the received JSON to
//     // newAlbum.
//     if err := c.BindJSON(&newAlbum); err != nil {
//         return
//     }

//     // Add the new album to the slice.
//     albums = append(albums, newAlbum)
//     c.IndentedJSON(http.StatusCreated, newAlbum)
// }

// // getAlbumByID locates the album whose ID value matches the id
// // parameter sent by the client, then returns that album as a response.
// func getAlbumByID(c *gin.Context) {
//     id := c.Param("id")

//     // Loop through the list of albums, looking for
//     // an album whose ID value matches the parameter.
//     for _, a := range albums {
//         if a.ID == id {
//             c.IndentedJSON(http.StatusOK, a)
//             return
//         }
//     }
//     c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }