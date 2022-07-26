package mail_utils

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/mailgun/mailgun-go/v4"
)

var yourDomain string = "my-homework-vorotyntsev.co" 

var partOne string = "c0a1070116e12756a1e912"
var partTwo string = "399824ec81-835621cf-d"
var partThree string = "07e8349"
var competedK string = partOne + partTwo + partThree

func SendMail(receiver string, price string) int{

	mg := mailgun.NewMailgun(yourDomain, competedK)
	

	mg.SetAPIBase("https://api.eu.mailgun.net/v3")

	sender := "example@example.co"
	subject := "Here is your Bitcoin to UAH price sir!"
	body := "Bitcoin to UAH = " + price + " UAH"
    fmt.Println(body)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, receiver)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
        return 1
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
    return 0
}
