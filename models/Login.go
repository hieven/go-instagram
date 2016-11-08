package models

import (
	"encoding/json"

	"github.com/hieven/go-instagram/constants"
	"github.com/hieven/go-instagram/utils"
	"github.com/parnurzeal/gorequest"
)

type Login struct {
	Csrftoken         string                `json:"_csrftoken"`
	DeviceID          string                `json:"device_id"`
	UUID              string                `json:"_uuid"`
	UserName          string                `json:"username"`
	Password          string                `json:"password"`
	LoginAttemptCount int                   `json:"login_attempt_count"`
	Agent             *gorequest.SuperAgent `json:"-"`
}

type LoginRequestSchema struct {
	SignedBody      string `json:"signed_body"`
	IgSigKeyVersion string `json:"ig_sig_key_version"`
}

func (login Login) CreateSignature() (sigVersion string, signedBody string) {
	jsonData, _ := json.Marshal(login)

	return utils.GenerateSignature(jsonData)
}

func (login Login) Login() (body string) {
	igSigKeyVersion, signedBody := login.CreateSignature()

	payload := LoginRequestSchema{
		IgSigKeyVersion: igSigKeyVersion,
		SignedBody:      signedBody,
	}

	jsonData, _ := json.Marshal(payload)

	_, body, _ = utils.WrapRequest(
		login.Agent.Post(constants.ROUTES.Login).
			Type("multipart").
			Send(string(jsonData)))

	return body
}
