package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"os"
)

func main() {
	sg_user := GetEnvOrExit("SG_USER")
	sg_key := GetEnvOrExit("SG_KEY")

	fmt.Println("Hello", sg_user, sg_key)
	sg := sendgrid.NewSendGridClient(sg_user, sg_key)
	message := sendgrid.NewMail()
	message.AddTo("siyegen@gmail.com")
	message.AddToName("Moo Moo")
	message.AddSubject("SendGrid Testing")
	message.AddText("WIN")
	message.AddFrom("siyegen@gmail.com")

	if r := sg.Send(message); r == nil {
		fmt.Println("Email sent!")
	} else {
		fmt.Println(r)
	}
}

func GetEnvOrExit(name string) string {
	log.Printf("Accessing %s", name)
	value := os.Getenv(name)

	if value == "" {
		log.Fatalf("%s was empty!", name)
	}
	return value
}
