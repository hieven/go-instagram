package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hieven/go-instagram/constants"
	"github.com/hieven/go-instagram/utils"
)

type Thread struct {
	ID             string                `json:"thread_id"`
	Users          []*UserSchema         `json:"users"`
	Items          []*ThreadItem         `json:"items"`
	ImageVersions2 ImageVersions2        `json:"image_versions2"`
	HasNewer       bool                  `json:"has_newer"`
	AgentPool      *utils.SuperAgentPool `json:"-"`
}

type UserSchema struct {
	Username string `json:"username"`
	Pk       int    `json:"pk"`
}

type broadcastTextSchema struct {
	UUID          string `json:"_uuid"`
	ThreadIds     string `json:"thread_ids"`
	ClientContext string `json:"client_context"`
	Text          string `json:"text"`
}

type showResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Thread  Thread `json:"thread"`
}

func (thread *Thread) BroadcastText(text string) (body string) {
	log.Println("Method: BroadcastText ", text)

	payload := broadcastTextSchema{
		UUID:          utils.GenerateUUID(),
		ThreadIds:     "[" + string(thread.ID) + "]",
		ClientContext: utils.GenerateUUID(),
		Text:          text,
	}

	jsonData, _ := json.Marshal(payload)

	agent := thread.AgentPool.Get()
	defer thread.AgentPool.Put(agent)

	_, body, _ = utils.WrapRequest(
		agent.Post(constants.ROUTES.ThreadsBroadcastText).
			Type("multipart").
			Send(string(jsonData)))

	return body
}

func (thread *Thread) Show() *Thread {
	agent := thread.AgentPool.Get()
	defer thread.AgentPool.Put(agent)

	_, body, _ := utils.WrapRequest(agent.Get(constants.ROUTES.ThreadsShow + thread.ID + "/"))

	var resp showResponse
	json.Unmarshal([]byte(body), &resp)

	if resp.Status == "fail" {
		fmt.Println(resp.Message)
		return thread
	}

	for _, item := range resp.Thread.Items {
		// if item.Location == nil {
		// 	continue
		// }

		item.Location.AgentPool = thread.AgentPool
	}

	thread.Items = resp.Thread.Items

	return thread
}
