package instagram

import (
	"encoding/json"
	"errors"

	"github.com/hieven/go-instagram/src/config"
	"github.com/hieven/go-instagram/src/constants"
	"github.com/hieven/go-instagram/src/protos"
	"github.com/hieven/go-instagram/src/utils/auth"
	"github.com/hieven/go-instagram/src/utils/request"
	"github.com/hieven/go-instagram/src/utils/session"
)

type instagram struct {
	config *config.Config

	// utils
	authManager    auth.AuthManager
	requestManager request.RequestManger
	sessionManager session.SessionManager

	// API
	inbox    *inbox
	timeline *timeline
}

func New(cnf *config.Config) (Instagram, error) {
	if cnf == nil {
		return nil, ErrConfigRequired
	}

	if cnf.Username == "" {
		return nil, ErrUsernameRequired
	}

	if cnf.Password == "" {
		return nil, ErrPasswordRequired
	}

	// utils
	authManager, _ := auth.New()
	sessionManager, _ := session.NewSession(cnf)
	requestManager, _ := request.New(sessionManager)

	inbox := &inbox{
		requestManager: requestManager,
	}
	timeline := &timeline{}

	ig := &instagram{
		config: cnf,

		inbox:    inbox,
		timeline: timeline,

		// utils
		authManager:    authManager,
		requestManager: requestManager,
		sessionManager: sessionManager,
	}

	return ig, nil
}

func (ig *instagram) Login() error {
	sigPayload := &auth.SignaturePayload{
		Csrftoken:         constants.SigCsrfToken,
		DeviceID:          constants.SigDeviceID,
		UUID:              ig.authManager.GenerateUUID(),
		UserName:          ig.config.Username,
		Password:          ig.config.Password,
		LoginAttemptCount: 0,
	}
	igSigKeyVersion, signedBody, _ := ig.authManager.GenerateSignature(sigPayload) // TODO: handle error

	req := protos.LoginRequest{
		IgSigKeyVersion: igSigKeyVersion,
		SignedBody:      signedBody,
	}

	resp, body, _ := ig.requestManager.Post(constants.LoginEndpoint, req)

	loginResp := &protos.LoginResponse{}
	json.Unmarshal([]byte(body), loginResp) // TODO: handle error

	if loginResp.Status == constants.InstagramStatusFail {
		return errors.New(loginResp.Message)
	}

	ig.sessionManager.SetCookies(resp.Cookies())

	return nil
}

func (ig *instagram) Inbox() Inbox {
	return ig.inbox
}

func (ig *instagram) Timeline() Timeline {
	return ig.timeline
}
