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

type GHResponse struct {
	Files GistFiles
}

func main() {
	sgUser := GetEnvOrExit("SG_USER")
	sgKey := GetEnvOrExit("SG_KEY")
	ghGist := GetEnvOrExit("GH_GIST")
	ghKey := GetEnvOrExit("GH_KEY")

	url := fmt.Sprintf("https://api.github.com/gists/%s", ghGist)
	gist := GetGist(url, ghKey)

	message := &EmailMessage{
		ToAddr:   "siyegen@gmail.com",
		FromAddr: "siyegen@gmail.com",
		Subject:  gist.Filename,
		Text:     gist.Content,
	}

	SendEmail(message, sgUser, sgKey)
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

	var gist_resp GHResponse
	err = json.Unmarshal(body, &gist_resp)
	if err != nil {
		log.Fatalf("Error with get %s", err)
	}

	return gist_resp.Files["remind_todo"]
}

func SendEmail(sgmessage *EmailMessage, sgUser, sgKey string) {
	sg := sendgrid.NewSendGridClient(sgUser, sgKey)
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
