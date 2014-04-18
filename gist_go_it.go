package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"os"
)

type EmailMessage struct {
	ToAddr   string
	FromAddr string
	Subject  string
	Text     string
}

func main() {
	sg_user := GetEnvOrExit("SG_USER")
	sg_key := GetEnvOrExit("SG_KEY")

	message := &EmailMessage{
		ToAddr:   "siyegen@gmail.com",
		FromAddr: "siyegen@gmail.com",
		Subject:  "Testing Testing",
		Text:     "Just some go",
	}

	SendEmail(message, sg_user, sg_key)
}

func GetEnvOrExit(name string) string {
	log.Printf("Accessing %s", name)
	value := os.Getenv(name)

	if value == "" {
		log.Fatalf("%s was empty!", name)
	}
	return value
}
func SendEmail(sgmessage *EmailMessage, sg_user, sg_key string) {
	sg := sendgrid.NewSendGridClient(sg_user, sg_key)
	message := sendgrid.NewMail()
	message.AddTo(sgmessage.ToAddr)
	message.AddSubject(sgmessage.Subject)
	message.AddText(sgmessage.Text)
	message.AddFrom(sgmessage.FromAddr)

	if r := sg.Send(message); r == nil {
		fmt.Println("Email sent!")
	} else {
		fmt.Println(r)
	}
}
