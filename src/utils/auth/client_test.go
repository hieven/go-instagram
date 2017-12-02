package auth

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"

	"github.com/satori/go.uuid"

	"github.com/hieven/go-instagram/src/constants"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("client", func() {
	var (
		client *authManager
	)

	BeforeEach(func() {
		client = &authManager{}
	})

	Describe(".New", func() {
		var (
			client AuthManager
			err    error
		)

		BeforeEach(func() {
			client = nil
		})

		JustBeforeEach(func() {
			client, err = New()
		})

		Context("when success", func() {
			It("should return client", func() {
				Expect(client).NotTo(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#GenerateSignature", func() {
		var (
			payload      *SignaturePayload
			payloadBytes []byte

			sigVersion string
			signedBody string
			err        error
		)

		BeforeEach(func() {
			payload = &SignaturePayload{}
		})

		JustBeforeEach(func() {
			payloadBytes, _ = json.Marshal(payload)

			sigVersion, signedBody, err = client.GenerateSignature(payload)
		})

		Context("when success", func() {
			It("should return data", func() {
				Expect(sigVersion).To(Equal(constants.SigVersion))
				Expect(strings.Contains(signedBody, string(payloadBytes))).To(BeTrue())
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#GenerateUUID", func() {
		var (
			uuidStr  string
			uuidStru uuid.UUID
			err      error
		)

		JustBeforeEach(func() {
			uuidStr = client.GenerateUUID()
			uuidStru, err = uuid.FromString(uuidStr)
		})

		Context("when success", func() {
			It("should return UUID string", func() {
				Expect(uuidStr).NotTo(BeNil())
				Expect(uuidStru).NotTo(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#GenerateRankToken", func() {
		var (
			userID    int64
			userIDStr string

			token string
		)

		BeforeEach(func() {
			userID = rand.Int63()

			token = ""
		})

		JustBeforeEach(func() {
			userIDStr = strconv.FormatInt(userID, 10)

			token = client.GenerateRankToken(userID)
		})

		Context("when success", func() {
			It("should return token string", func() {
				Expect(token).NotTo(BeEmpty())
			})

			It("should contain user id", func() {
				splits := strings.Split(token, "_")
				Expect(len(splits)).To(Equal(2))
				Expect(splits[0]).To(Equal(userIDStr))
			})

			It("should contain UUID", func() {
				splits := strings.Split(token, "_")
				uuidStr, err := uuid.FromString(splits[1])
				Expect(uuidStr).NotTo(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
