package models

import (
	"encoding/json"
	"errors"

	"github.com/hieven/go-instagram/constants"
	"github.com/hieven/go-instagram/utils"
)

type Thread struct {
	ID             string         `json:"thread_id"`
	Users          []*UserSchema  `json:"users"`
	Items          []*ThreadItem  `json:"items"`
	ImageVersions2 ImageVersions2 `json:"image_versions2"`
	HasNewer       bool           `json:"has_newer"`
	Instagram      *Instagram     `json:"-"`
}

type UserSchema struct {
	Username string `json:"username"`
	Pk       int    `json:"pk"`
}

type broadcastTextRequest struct {
	UUID          string `json:"_uuid"`
	ThreadIds     string `json:"thread_ids"`
	ClientContext string `json:"client_context"`
	Text          string `json:"text"`
}

type broadcastTextResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type showResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Thread  Thread `json:"thread"`
}

func (thread *Thread) BroadcastText(text string) error {
	payload := broadcastTextRequest{
		UUID:          utils.GenerateUUID(),
		ThreadIds:     "[" + string(thread.ID) + "]",
		ClientContext: utils.GenerateUUID(),
		Text:          text,
	}

	jsonData, _ := json.Marshal(payload)

	agent := thread.Instagram.AgentPool.Get()
	defer thread.Instagram.AgentPool.Put(agent)

	_, body, _ := thread.Instagram.SendRequest(
		agent.Post(constants.ROUTES.ThreadsBroadcastText).
			Type("multipart").
			Send(string(jsonData)))

	var resp broadcastTextResponse
	json.Unmarshal([]byte(body), &resp)

	if resp.Status == "fail" {
		return errors.New(resp.Message)
	}

	return nil
}

func (thread *Thread) Show() (*Thread, error) {
	agent := thread.Instagram.AgentPool.Get()
	defer thread.Instagram.AgentPool.Put(agent)

	_, body, _ := thread.Instagram.SendRequest(
		agent.Get(constants.ROUTES.ThreadsShow + thread.ID + "/"))

	var resp showResponse
	json.Unmarshal([]byte(body), &resp)

	if resp.Status == "fail" {
		return nil, errors.New(resp.Message)
	}

	for _, item := range resp.Thread.Items {
		// if item.Location == nil {
		// 	continue
		// }

		item.Location.Instagram = thread.Instagram
	}

	thread.Items = resp.Thread.Items

	return thread, nil
}
