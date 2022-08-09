package emailcheck

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type CheckData struct {
	filename string
	notify   bool
	modified bool
	Count    int    `json:"count"`
	Filter   string `json:"filter"`
	Message  string `json:"message,omitempty"`
	Company  string `json:"company,omitempty"`
}

func (data *CheckData) Save() {
	fmt.Printf("Saving CheckData to file: %s\n", data.filename)
	f, err := os.OpenFile("checks/"+data.filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Unable to save checkdata to file: %v", err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(data)
	if err != nil {
		log.Fatalf("Unable to encode checkdata: %v", err)
	}
}

func (data *CheckData) Load(filename string) error {
	f, err := os.Open("checks/" + filename)
	if err != nil {
		return err
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(data)
	if err != nil {
		return err
	}
	data.filename = filename
	data.notify = false
	data.modified = false
	return nil
}

func (data *CheckData) IsModified() bool {
	return data.modified
}

func (data *CheckData) ShouldNotify() bool {
	return data.notify
}

func (data *CheckData) UpdateUnread(unreadCount int) {
	data.notify = (unreadCount > data.Count)
	data.modified = (unreadCount != data.Count)
	data.Count = unreadCount
}
