package auth

type AuthManager interface {
	GenerateSignature(payload *SignaturePayload) (sigVersion string, signedBody string, err error)
	GenerateUUID() (uuid string)
	GenerateRankToken(userID string) string
}

type SignaturePayload struct {
	Csrftoken         string `json:"_csrftoken"`
	DeviceID          string `json:"device_id"`
	UUID              string `json:"_uuid"`
	UserName          string `json:"username"`
	Password          string `json:"password"`
	LoginAttemptCount int    `json:"login_attempt_count"`
}
