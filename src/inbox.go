package instagram

import (
	"encoding/json"
	"net/url"

	"github.com/hieven/go-instagram/src/constants"
	"github.com/hieven/go-instagram/src/protos"
	"github.com/hieven/go-instagram/src/utils/request"
)

type inbox struct {
	requestManager request.RequestManger
}

func (inbox *inbox) Feed() (*protos.InboxFeedResponse, error) {
	urlStru, _ := url.Parse(constants.InboxEndpoint) // TODO: handle error

	_, body, _ := inbox.requestManager.Get(urlStru.String()) // TODO: handle error

	feedResponse := &protos.InboxFeedResponse{}
	json.Unmarshal([]byte(body), feedResponse) // TODO: handle error

	return feedResponse, nil
}
