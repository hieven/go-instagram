package models

import (
	"encoding/json"

	"errors"

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

type loginRequestSchema struct {
	SignedBody      string `json:"signed_body"`
	IgSigKeyVersion string `json:"ig_sig_key_version"`
}

type loginResponseSchema struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (login Login) CreateSignature() (sigVersion string, signedBody string) {
	jsonData, _ := json.Marshal(login)

	return utils.GenerateSignature(jsonData)
}

func (login Login) Login() error {
	igSigKeyVersion, signedBody := login.CreateSignature()

	payload := loginRequestSchema{
		IgSigKeyVersion: igSigKeyVersion,
		SignedBody:      signedBody,
	}

	jsonData, _ := json.Marshal(payload)

	_, body, err := utils.WrapRequest(
		login.Agent.Post(constants.ROUTES.Login).
			Type("multipart").
			Send(string(jsonData)))

	if len(err) > 0 {
		return err[0]
	}

	var resp loginResponseSchema
	json.Unmarshal([]byte(body), &resp)

	if resp.Status == "fail" {

		return errors.New(resp.Message)
	}

	return nil
}
