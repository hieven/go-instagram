package instagram

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/hieven/go-instagram/src/config"
	"github.com/hieven/go-instagram/src/constants"
	"github.com/hieven/go-instagram/src/protos"
	"github.com/hieven/go-instagram/src/utils/auth"
	"github.com/hieven/go-instagram/src/utils/request"
	"github.com/hieven/go-instagram/src/utils/session"
	"github.com/hieven/go-instagram/src/utils/text"
)

type instagram struct {
	config *config.Config

	// utils
	authManager    auth.AuthManager
	requestManager request.RequestManger
	sessionManager session.SessionManager

	// API
	inbox    Inbox
	timeline Timeline
	thread   Thread
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
	textManager, _ := text.New()

	timeline := &timeline{}
	inbox := &inbox{requestManager: requestManager}
	thread := &thread{requestManager: requestManager, authManager: authManager, textManager: textManager}

	ig := &instagram{
		config: cnf,

		timeline: timeline,
		inbox:    inbox,
		thread:   thread,

		// utils
		authManager:    authManager,
		requestManager: requestManager,
		sessionManager: sessionManager,
	}

	return ig, nil
}

func (ig *instagram) Login(ctx context.Context) error {
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

	resp, body, _ := ig.requestManager.Post(ctx, constants.LoginEndpoint, req)

	loginResp := &protos.LoginResponse{}
	json.Unmarshal([]byte(body), loginResp) // TODO: handle error

	if loginResp.Status == constants.InstagramStatusFail {
		return errors.New(loginResp.Message)
	}

	ig.sessionManager.SetCookies(resp.Cookies())

	return nil
}

func (ig *instagram) Timeline() Timeline {
	return ig.timeline
}

func (ig *instagram) Inbox() Inbox {
	return ig.inbox
}

func (ig *instagram) Thread() Thread {
	return ig.thread
}
