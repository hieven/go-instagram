package instagram

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

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
	idToCodeMap    []string
}

func newMediaClient(
	config *config.Config,

	requestManager request.RequestManger,
	authManager auth.AuthManager,
) (*media, error) {
	m := &media{
		config:         config,
		requestManager: requestManager,
		authManager:    authManager,
		idToCodeMap:    []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-", "_"},
	}

	return m, nil
}

// GetShortCodeByMediaID convert media id to corresponding short code
func (media *media) GetShortCodeByMediaID(ctx context.Context, mediaID string) string {
	// media_id from Instagram is actually composed of {media_id}_{user_id}
	parts := strings.Split(mediaID, "_")
	realMediaID, err := strconv.Atoi(parts[0])
	if err != nil {
		return ""
	}

	code := ""

	for realMediaID > 0 {
		remainder := realMediaID % 64
		realMediaID = (realMediaID - remainder) / 64
		code = media.idToCodeMap[remainder] + code
	}

	return code
}

func (media *media) Info(ctx context.Context, req *MediaInfoRequest) (*protos.MediaInfoResponse, error) {
	if req == nil {
		return nil, ErrRequestRequired
	}

	urlStru, _ := url.Parse(fmt.Sprintf(constants.MediaInfoEndpoint, req.MediaID)) // TODO: handle error

	_, body, _ := media.requestManager.Get(ctx, urlStru.String()) // TODO: handle error

	result := &protos.MediaInfoResponse{}
	json.Unmarshal([]byte(body), result) // TODO: handle error
	return result, nil
}

func (media *media) Like(ctx context.Context, req *MediaLikeRequest) (*protos.MediaLikeResponse, error) {
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

	internalReq := &protos.MediaLikeRequest{
		MediaID: req.MediaID,
		Src:     "profile",
		LoginRequest: protos.LoginRequest{
			IgSigKeyVersion: igSigKeyVersion,
			SignedBody:      signedBody,
		},
	}

	_, body, _ := media.requestManager.Post(ctx, urlStru.String(), internalReq) // TODO: handle error

	result := &protos.MediaLikeResponse{}
	json.Unmarshal([]byte(body), result) // TODO: handle error
	return result, nil
}

func (media *media) Unlike(ctx context.Context, req *MediaUnlikeRequest) (*protos.MediaUnlikeResponse, error) {
	if req == nil {
		return nil, ErrRequestRequired
	}

	urlStru, _ := url.Parse(fmt.Sprintf(constants.MediaUnlikeEndpoint, req.MediaID)) // TODO: handle error

	sigPayload := &auth.SignaturePayload{
		Csrftoken:         constants.SigCsrfToken,
		DeviceID:          constants.SigDeviceID,
		UUID:              media.authManager.GenerateUUID(),
		UserName:          media.config.Username,
		Password:          media.config.Password,
		LoginAttemptCount: 0,
	}
	igSigKeyVersion, signedBody, _ := media.authManager.GenerateSignature(sigPayload) // TODO: handle error

	internalReq := &protos.MediaUnlikeRequest{
		MediaID: req.MediaID,
		Src:     "profile",
		LoginRequest: protos.LoginRequest{
			IgSigKeyVersion: igSigKeyVersion,
			SignedBody:      signedBody,
		},
	}

	_, body, _ := media.requestManager.Post(ctx, urlStru.String(), internalReq) // TODO: handle error

	result := &protos.MediaUnlikeResponse{}
	json.Unmarshal([]byte(body), result) // TODO: handle error
	return result, nil
}
