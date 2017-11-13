package instagram

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/hieven/go-instagram/src/config"
	"github.com/hieven/go-instagram/src/utils/auth"

	"github.com/hieven/go-instagram/src/constants"
	"github.com/hieven/go-instagram/src/protos"
	"github.com/hieven/go-instagram/src/utils/request"
)

type media struct {
	config *config.Config

	requestManager request.RequestManger
	authManager    auth.AuthManager
}

func (media *media) Like(ctx context.Context, req *protos.MediaLikeRequest) (*protos.MediaLikeResponse, error) {
	if req == nil {
		return nil, ErrRequestRequired
	}

	urlStru, _ := url.Parse(fmt.Sprintf(constants.MediaLikeEndpoint, req.MediaID)) // TODO: handle error

	sigPayload := &auth.SignaturePayload{
		Csrftoken:         constants.SigCsrfToken,
		DeviceID:          constants.SigDeviceID,
		UUID:              media.authManager.GenerateUUID(),
		UserName:          media.config.Username,
		Password:          media.config.Password,
		LoginAttemptCount: 0,
	}
	igSigKeyVersion, signedBody, _ := media.authManager.GenerateSignature(sigPayload) // TODO: handle error

	req.Src = "profile"
	req.IgSigKeyVersion = igSigKeyVersion
	req.SignedBody = signedBody

	_, body, _ := media.requestManager.Post(ctx, urlStru.String(), req) // TODO: handle error

	result := &protos.MediaLikeResponse{}
	json.Unmarshal([]byte(body), result) // TODO: handle error
	return result, nil
}
