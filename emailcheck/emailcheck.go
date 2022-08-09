package emailcheck

import (
	"encoding/base64"
	"fmt"
	"github.com/lazyguru/emailcheck/emailcheck/gmailclient"
	"google.golang.org/api/gmail/v1"
	"log"
	"os"
	"strings"
)

var srv *gmail.Service

func Initialize() error {
	service, err := gmailclient.GetService()
	srv = service
	return err
}

func checkMessages(data *CheckData) error {
	user := "me"
	r, err := srv.Users.Messages.List(user).Q(data.Filter).Do()
	if err != nil {
		return err
	}

	unreadCount := len(r.Messages)
	data.UpdateUnread(unreadCount)
	return nil
}

func sendNotice(data *CheckData) error {
	var message gmail.Message
	emailTo := "To: " + os.Getenv("EMAILTO") + "\r\n"
	subject := "Subject: " + data.Company + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + mime + "\n" + fmt.Sprintf(data.Message, data.Count))

	message.Raw = base64.URLEncoding.EncodeToString(msg)

	// Send the message
	_, err := srv.Users.Messages.Send("me", &message).Do()
	return err
}

func getDataFiles() ([]CheckData, error) {
	var dir, err = os.ReadDir("checks")
	if err != nil {
		return nil, err
	}
	checkData := []CheckData{}
	for _, file := range dir {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		data := &CheckData{}
		err := data.Load(file.Name())
		if err != nil {
			return nil, err
		}
		checkData = append(checkData, *data)
	}
	return checkData, err
}

func Run() {
	searchData, err := getDataFiles()
	if err != nil {
		log.Fatalf("Error loading files: %v", err)
	}

	for _, data := range searchData {
		err := checkMessages(&data)
		if err != nil {
			log.Fatalf("Search failed: %v", err)
		}
		if data.IsModified() {
			data.Save()
		}

		if data.ShouldNotify() {
			log.Printf("Sending email notice about %s\n", data.Company)
			err = sendNotice(&data)
			if err != nil {
				log.Fatalf("Failed sending message: %v", err)
			}
		} else {
			log.Printf("No new unread messages from %s\n", data.Company)
		}
	}
}
