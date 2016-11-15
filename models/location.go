package models

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/hieven/go-instagram/constants"
	"github.com/hieven/go-instagram/utils"
)

type Location struct {
	ExternalSource   string     `json:"external_source"`
	City             string     `json:"city"`
	Name             string     `json:"name"`
	FacebookPlacesID int64      `json:"facebook_places_id"`
	Address          string     `json:"address"`
	Lat              float64    `json:"lat"`
	Lng              float64    `json:"lng"`
	Pk               int64      `json:"pk"`
	Instagram        *Instagram `json:"-"`
}

type mediaResponse struct {
	Status      string
	Message     string
	RankedItems []*Media `json:"ranked_items"`
	Items       []*Media `json:"items"`
}

func (location Location) GetRankedMedias() ([]*Media, error) {
	url := constants.ROUTES.LocationFeed + strconv.FormatInt(location.Pk, 10) + "/"

	agent := location.Instagram.AgentPool.Get()
	defer location.Instagram.AgentPool.Put(agent)

	_, body, _ := location.Instagram.SendRequest(
		agent.Get(url).Query("rank_token=" + utils.GenerateUUID()))

	var resp mediaResponse
	json.Unmarshal([]byte(body), &resp)

	if resp.Status == "fail" {
		return nil, errors.New(resp.Message)
	}

	return resp.RankedItems, nil
}

func (location Location) GetRecentMedias() ([]*Media, error) {
	url := constants.ROUTES.LocationFeed + strconv.FormatInt(location.Pk, 10) + "/"

	agent := location.Instagram.AgentPool.Get()
	defer location.Instagram.AgentPool.Put(agent)

	_, body, _ := location.Instagram.SendRequest(
		agent.Get(url).Query("rank_token=" + utils.GenerateUUID()))

	var resp mediaResponse
	json.Unmarshal([]byte(body), &resp)

	if resp.Status == "fail" {
		return nil, errors.New(resp.Message)
	}

	return resp.Items, nil
}
