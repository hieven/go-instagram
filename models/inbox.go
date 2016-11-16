package models

import (
	"encoding/json"
	"errors"

	"github.com/hieven/go-instagram/constants"
	"github.com/hieven/go-instagram/utils"
)

// Inbox type
type Inbox struct {
	Threads   []*Thread `json:"threads"`
	Instagram *Instagram
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

// GetFeed returns you inbox feed
func (inbox *Inbox) GetFeed() ([]*Thread, error) {
	agent := inbox.Instagram.AgentPool.Get()
	defer inbox.Instagram.AgentPool.Put(agent)

	_, body, _ := inbox.Instagram.SendRequest(agent.Get(constants.ROUTES.Inbox))

	var resp inboxFeedResponse
	json.Unmarshal([]byte(body), &resp)

	if resp.Status == "fail" {
		return nil, errors.New(resp.Message)
	}

	inbox.Threads = resp.Inbox.Threads

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
