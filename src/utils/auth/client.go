package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/hieven/go-instagram/src/constants"

	uuid "github.com/satori/go.uuid"
)

type authManager struct {
}

func New() (AuthManager, error) {
	auth := &authManager{}
	return auth, nil
}

func (auth *authManager) GenerateSignature(payload *SignaturePayload) (string, string, error) {
	payloadBytes, _ := json.Marshal(payload)

	h := hmac.New(sha256.New, []byte(constants.SigKey))
	h.Write(payloadBytes)

	var b []byte
	hash := hex.EncodeToString(h.Sum(b))

	sigVersion := constants.SigVersion
	signedBody := hash + "." + string(payloadBytes)

	return sigVersion, signedBody, nil
}

func (auth *authManager) GenerateUUID() string {
	return uuid.NewV4().String()
}

func (auth *authManager) GenerateRankToken(userID string) string {
	return userID + "_" + auth.GenerateUUID()
}
