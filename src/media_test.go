package instagram

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/hieven/go-instagram/src/config"
	"github.com/hieven/go-instagram/src/utils/auth"

	"github.com/satori/go.uuid"

	"github.com/stretchr/testify/mock"

	"github.com/hieven/go-instagram/src/constants"
	"github.com/hieven/go-instagram/src/protos"
	authMocks "github.com/hieven/go-instagram/src/utils/auth/mocks"
	requestMocks "github.com/hieven/go-instagram/src/utils/request/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("media", func() {
	var (
		mockAuthManager    *authMocks.AuthManager
		mockRequestManager *requestMocks.RequestManger

		ctx    context.Context
		client *media
	)

	BeforeEach(func() {
		mockAuthManager = &authMocks.AuthManager{}
		mockRequestManager = &requestMocks.RequestManger{}

		ctx = context.Background()

		client = &media{
			config: &config.Config{
				Username: "username",
				Password: "password",
			},
			authManager:    mockAuthManager,
			requestManager: mockRequestManager,
		}
	})

	Describe("Like", func() {
		var (
			req *MediaLikeRequest

			mockGenerateUUIDResp            string
			mockGenerateSignatureKeyVersion string
			mockGenerateSignatureBody       string
			mockResp                        *protos.MediaLikeResponse
			mockBody                        string

			resp *protos.MediaLikeResponse
			err  error

			expURLStru                  *url.URL
			expURLQuery                 url.Values
			expURLStr                   string
			expGenerateSignaturePayload *auth.SignaturePayload
			expInternalReq              *protos.MediaLikeRequest
		)

		BeforeEach(func() {
			req = &MediaLikeRequest{
				MediaID: "media id",
			}

			mockGenerateUUIDResp = uuid.NewV4().String()

			mockGenerateSignatureKeyVersion = "key version"
			mockGenerateSignatureBody = "sig body"

			mockResp = &protos.MediaLikeResponse{}
			mockBodyBytes, _ := json.Marshal(mockResp)
			mockBody = string(mockBodyBytes)

			expURLStru, _ = url.Parse(fmt.Sprintf(constants.MediaLikeEndpoint, req.MediaID))
			expURLQuery = expURLStru.Query()

			expGenerateSignaturePayload = &auth.SignaturePayload{
				Csrftoken:         constants.SigCsrfToken,
				DeviceID:          constants.SigDeviceID,
				UUID:              mockGenerateUUIDResp,
				UserName:          client.config.Username,
				Password:          client.config.Password,
				LoginAttemptCount: 0,
			}

			expInternalReq = &protos.MediaLikeRequest{
				MediaID: req.MediaID,
				Src:     "profile",
				LoginRequest: protos.LoginRequest{
					IgSigKeyVersion: mockGenerateSignatureKeyVersion,
					SignedBody:      mockGenerateSignatureBody,
				},
			}
		})

		JustBeforeEach(func() {
			mockAuthManager.On("GenerateUUID").Return(mockGenerateUUIDResp)
			mockAuthManager.On("GenerateSignature", mock.Anything).Return(mockGenerateSignatureKeyVersion, mockGenerateSignatureBody, nil)
			mockRequestManager.On("Post", mock.Anything, mock.Anything, mock.Anything).Return(nil, mockBody, nil)

			resp, err = client.Like(ctx, req)

			expURLStru.RawQuery = expURLQuery.Encode()
			expURLStr = expURLStru.String()
		})

		Context("when success", func() {
			It("should return response", func() {
				Expect(err).To(BeNil())
				Expect(resp).NotTo(BeNil())
			})

			It("should call authManager.GenerateUUID", func() {
				mockAuthManager.AssertNumberOfCalls(GinkgoT(), "GenerateUUID", 1)
				mockAuthManager.AssertCalled(GinkgoT(), "GenerateUUID")
			})

			It("should call authManager.GenerateSignature", func() {
				mockAuthManager.AssertNumberOfCalls(GinkgoT(), "GenerateSignature", 1)
				mockAuthManager.AssertCalled(GinkgoT(), "GenerateSignature", expGenerateSignaturePayload)
			})

			It("should call requestManager.Post", func() {
				mockRequestManager.AssertNumberOfCalls(GinkgoT(), "Post", 1)
				mockRequestManager.AssertCalled(GinkgoT(), "Post", mock.Anything, expURLStr, expInternalReq)
			})
		})

		Context("when req isn't provided", func() {
			BeforeEach(func() {
				req = nil
			})

			It("should return error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(ErrRequestRequired.Error()))
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Unlike", func() {
		var (
			req *MediaUnlikeRequest

			mockGenerateUUIDResp            string
			mockGenerateSignatureKeyVersion string
			mockGenerateSignatureBody       string
			mockResp                        *protos.MediaUnlikeResponse
			mockBody                        string

			resp *protos.MediaUnlikeResponse
			err  error

			expURLStru                  *url.URL
			expURLQuery                 url.Values
			expURLStr                   string
			expGenerateSignaturePayload *auth.SignaturePayload
			expInternalReq              *protos.MediaUnlikeRequest
		)

		BeforeEach(func() {
			req = &MediaUnlikeRequest{
				MediaID: "media id",
			}

			mockGenerateUUIDResp = uuid.NewV4().String()

			mockGenerateSignatureKeyVersion = "key version"
			mockGenerateSignatureBody = "sig body"

			mockResp = &protos.MediaUnlikeResponse{}
			mockBodyBytes, _ := json.Marshal(mockResp)
			mockBody = string(mockBodyBytes)

			expURLStru, _ = url.Parse(fmt.Sprintf(constants.MediaUnlikeEndpoint, req.MediaID))
			expURLQuery = expURLStru.Query()

			expGenerateSignaturePayload = &auth.SignaturePayload{
				Csrftoken:         constants.SigCsrfToken,
				DeviceID:          constants.SigDeviceID,
				UUID:              mockGenerateUUIDResp,
				UserName:          client.config.Username,
				Password:          client.config.Password,
				LoginAttemptCount: 0,
			}

			expInternalReq = &protos.MediaUnlikeRequest{
				MediaID: req.MediaID,
				Src:     "profile",
				LoginRequest: protos.LoginRequest{
					IgSigKeyVersion: mockGenerateSignatureKeyVersion,
					SignedBody:      mockGenerateSignatureBody,
				},
			}
		})

		JustBeforeEach(func() {
			mockAuthManager.On("GenerateUUID").Return(mockGenerateUUIDResp)
			mockAuthManager.On("GenerateSignature", mock.Anything).Return(mockGenerateSignatureKeyVersion, mockGenerateSignatureBody, nil)
			mockRequestManager.On("Post", mock.Anything, mock.Anything, mock.Anything).Return(nil, mockBody, nil)

			resp, err = client.Unlike(ctx, req)

			expURLStru.RawQuery = expURLQuery.Encode()
			expURLStr = expURLStru.String()
		})

		Context("when success", func() {
			It("should return response", func() {
				Expect(err).To(BeNil())
				Expect(resp).NotTo(BeNil())
			})

			It("should call authManager.GenerateUUID", func() {
				mockAuthManager.AssertNumberOfCalls(GinkgoT(), "GenerateUUID", 1)
				mockAuthManager.AssertCalled(GinkgoT(), "GenerateUUID")
			})

			It("should call authManager.GenerateSignature", func() {
				mockAuthManager.AssertNumberOfCalls(GinkgoT(), "GenerateSignature", 1)
				mockAuthManager.AssertCalled(GinkgoT(), "GenerateSignature", expGenerateSignaturePayload)
			})

			It("should call requestManager.Post", func() {
				mockRequestManager.AssertNumberOfCalls(GinkgoT(), "Post", 1)
				mockRequestManager.AssertCalled(GinkgoT(), "Post", mock.Anything, expURLStr, expInternalReq)
			})
		})

		Context("when req isn't provided", func() {
			BeforeEach(func() {
				req = nil
			})

			It("should return error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(ErrRequestRequired.Error()))
				Expect(resp).To(BeNil())
			})
		})
	})
})
