package models

import (
	"encoding/json"
	"log"

	"github.com/hieven/go-instagram/constants"
	"github.com/hieven/go-instagram/utils"
	"github.com/parnurzeal/gorequest"
)

type Thread struct {
	ID       string                `json:"thread_id"`
	Users    []UserSchema          `json:"users"`
	Items    []ItemSchema          `json:"items"`
	HasNewer bool                  `json:"has_newer"`
	Request  *gorequest.SuperAgent `json:"request"`
}

type UserSchema struct {
	Username string `json:"username"`
	Pk       int    `json:"pk"`
}

type ItemSchema struct {
	ID        string `json:"item_id"`
	UserID    int    `json:"user_id"`
	ItemType  string `json:"item_type"`
	Timestamp int    `json:"timestamp"`

	// depends on ItemType
	Placeholder placeholderSchema `json:"placeholder"`
	Text        string            `json:"text"`
	MediaShare  mediaShareSchema  `json:"media_share"`
	Location    locationSchema    `json:"location"`
}

type broadcastTextSchema struct {
	UUID          string `json:"_uuid"`
	ThreadIds     string `json:"thread_ids"`
	ClientContext string `json:"client_context"`
	Text          string `json:"text"`
}

type placeholderSchema struct {
	IsLinked bool   `json:"is_linked"`
	Message  string `json:"message"`
	Title    string `json:"title"`
}

type mediaShareSchema struct {
	ImageVersion2 imageVersion2Schema `json:"image_versions2"`
	Location      locationSchema      `json:"location"`
	Lat           float32             `json:"lat"`
	Lng           float32             `json:"lng"`
}

type imageVersion2Schema struct {
	Candidates []candidateSchema `json:"candidates"`
}

type candidateSchema struct {
	url    string
	width  int
	height int
}

type locationSchema struct {
	ExternalSource   string  `json:"external_source"`
	City             string  `json:"city"`
	Name             string  `json:"name"`
	FacebookPlacesID int     `json:"facebook_places_id"`
	Address          string  `json:"address"`
	Lat              float32 `json:"lat"`
	Lng              float32 `json:"lng"`
	Pk               int     `json:"pk"`
}

func (thread Thread) BroadcastText(text string) (body string) {
	log.Println("Thread---------------->")
	log.Println("Method: BroadcastText ", text)

	payload := broadcastTextSchema{
		UUID:          utils.GenerateUUID(),
		ThreadIds:     "[" + string(thread.ID) + "]",
		ClientContext: utils.GenerateUUID(),
		Text:          text,
	}

	jsonData, _ := json.Marshal(payload)

	_, body, _ = utils.WrapRequest(
		thread.Request.Post(constants.ROUTES.ThreadsBroadcastText).
			Type("multipart").
			Send(string(jsonData)))

	return body
}
