package main

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type EmailMessage struct {
	ToAddr   string
	FromAddr string
	Subject  string
	Text     string
}

type GistResponse struct {
	Filename string
	Content  string
}

type GistFiles map[string]*GistResponse

type XXResponse struct {
	Files GistFiles
}

func main() {
	// sg_user := GetEnvOrExit("SG_USER")
	// sg_key := GetEnvOrExit("SG_KEY")

	// message := &EmailMessage{
	// 	ToAddr:   "siyegen@gmail.com",
	// 	FromAddr: "siyegen@gmail.com",
	// 	Subject:  "Testing Testing",
	// 	Text:     "Just some go",
	// }

	// SendEmail(message, sg_user, sg_key)

	// gh_user := GetEnvOrExit("GH_USER")
	gh_gist := GetEnvOrExit("GH_GIST")
	gh_key := GetEnvOrExit("GH_KEY")

	url := fmt.Sprintf("https://api.github.com/gists/%s", gh_gist)
	fmt.Println(GetGist(url, gh_key))
}

func GetGist(url, key string) *GistResponse {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("authorization token", key)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Error with get %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error with get %s", err)
	}

	var gist_resp XXResponse
	err = json.Unmarshal(body, &gist_resp)
	if err != nil {
		log.Fatalf("Error with get %s", err)
	}

	return gist_resp.Files["remind_todo"]
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

func GetEnvOrExit(name string) string {
	log.Printf("Accessing %s", name)
	value := os.Getenv(name)

	if value == "" {
		log.Fatalf("%s was empty!", name)
	}
	return value
}
