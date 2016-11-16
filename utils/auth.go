package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/hieven/go-instagram/constants"
	UUID "github.com/satori/go.uuid"
)

func GenerateSignature(data []byte) (sigVersion string, signedBody string) {
	h := hmac.New(sha256.New, []byte(constants.SIG_KEY))
	h.Write(data)

	var b []byte
	hash := hex.EncodeToString(h.Sum(b))

	sigVersion = constants.SIG_VERSION
	signedBody = hash + "." + string(data)

	return sigVersion, signedBody
}

func GenerateUUID() (uuid string) {
	uuid = UUID.NewV4().String()

	return uuid
}

func GenerateRankToken(userID string) string {
	return userID + "_" + GenerateUUID()
}
