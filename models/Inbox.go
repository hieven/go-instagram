package models

import (
	"encoding/json"

	"github.com/hieven/go-instagram/constants"
	"github.com/hieven/go-instagram/utils"
	"github.com/parnurzeal/gorequest"
)

type Inbox struct {
	Threads []*Thread `json:"threads"`
	Request *gorequest.SuperAgent
}

// Parsing instagram response
type feedResponse struct {
	Inbox Inbox `json:inbox`
}

func (inbox *Inbox) GetFeed() (threads []*Thread) {
	_, body, _ := utils.WrapRequest(inbox.Request.Get(constants.ROUTES.Inbox))

	var resp feedResponse
	json.Unmarshal([]byte(body), &resp)

	// fmt.Println(body)
	inbox.Threads = resp.Inbox.Threads

	for _, thread := range inbox.Threads {
		thread.Request = inbox.Request

		for _, item := range thread.Items {
			item.Location.Request = inbox.Request
		}
	}

	return inbox.Threads
}
