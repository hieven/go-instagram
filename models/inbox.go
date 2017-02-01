package models

import (
	"encoding/json"
	"errors"

	"github.com/hieven/go-instagram/constants"
	"github.com/hieven/go-instagram/utils"
)

// Inbox type
type Inbox struct {
	Threads      []*Thread  `json:"threads"`
	HasOlder     bool       `json:"has_older"`
	OldestCursor string     `json:"oldest_cursor"`
	Instagram    *Instagram `json:"-"`
}

// Parsing instagram response
type inboxFeedResponse struct {
	Status  string
	Message string
	Inbox   Inbox `json:"inbox"`
}

type approveAllThreadRequest struct {
	UUID string `json:"_uuid"`
}

type approveAllThreadResponse struct {
	Status  string
	Message string
}

func (inbox *Inbox) GetCursor() string {
	return inbox.OldestCursor
}

func (inbox *Inbox) SetCursor(cursor string) {
	inbox.OldestCursor = cursor
}

func (inbox *Inbox) IsMoreAvailalbe() bool {
	return inbox.HasOlder
}

// GetFeed returns you inbox feed
func (inbox *Inbox) GetFeed() ([]*Thread, error) {
	url := constants.GetURL("Inbox", struct {
		Cursor string
	}{
		Cursor: inbox.GetCursor(),
	})

	agent := inbox.Instagram.AgentPool.Get()
	defer inbox.Instagram.AgentPool.Put(agent)

	_, body, _ := inbox.Instagram.SendRequest(agent.Get(url))

	var resp inboxFeedResponse
	json.Unmarshal([]byte(body), &resp)

	if resp.Status == "fail" {
		return nil, errors.New(resp.Message)
	}

	inbox.HasOlder = resp.Inbox.HasOlder
	inbox.Threads = resp.Inbox.Threads

	if inbox.HasOlder {
		inbox.SetCursor(resp.Inbox.OldestCursor)
	}

	for _, thread := range inbox.Threads {
		thread.Instagram = inbox.Instagram

		for _, item := range thread.Items {
			// if item.Location == {} {
			// 	continue
			// }

			item.Location.Instagram = inbox.Instagram
		}
	}

	return inbox.Threads, nil
}

// ApproveAllThreads will approve all pending message requests
func (inbox *Inbox) ApproveAllThreads() error {
	payload := approveAllThreadRequest{
		UUID: utils.GenerateUUID(),
	}

	jsonData, _ := json.Marshal(payload)

	agent := inbox.Instagram.AgentPool.Get()
	defer inbox.Instagram.AgentPool.Put(agent)

	_, body, _ := inbox.Instagram.SendRequest(
		agent.Post(constants.ROUTES.ThreadsApproveAll).
			Type("multipart").
			Send(string(jsonData)))

	var resp approveAllThreadResponse
	json.Unmarshal([]byte(body), &resp)

	if resp.Status == "fail" {
		return errors.New(resp.Message)
	}

	return nil
}
