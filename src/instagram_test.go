package instagram

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hieven/go-instagram/src/constants"
	"github.com/hieven/go-instagram/src/utils/auth"

	"github.com/satori/go.uuid"

	"github.com/stretchr/testify/mock"

	"github.com/hieven/go-instagram/src/config"
	"github.com/hieven/go-instagram/src/protos"
	authMocks "github.com/hieven/go-instagram/src/utils/auth/mocks"
	requestMocks "github.com/hieven/go-instagram/src/utils/request/mocks"
	sessionMocks "github.com/hieven/go-instagram/src/utils/session/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("instagram", func() {
	var (
		mockAuthManager    *authMocks.AuthManager
		mockSessionManager *sessionMocks.SessionManager
		mockRequestManager *requestMocks.RequestManger

		mockTimeline *timeline
		mockInbox    *inbox
		mockThread   *thread
		mockMedia    *media
		mockLocation *location

		cnf *config.Config
		ig  *instagram
	)

	BeforeEach(func() {
		mockAuthManager = &authMocks.AuthManager{}
		mockSessionManager = &sessionMocks.SessionManager{}
		mockRequestManager = &requestMocks.RequestManger{}

		mockTimeline = &timeline{}
		mockInbox = &inbox{}
		mockThread = &thread{}
		mockMedia = &media{}
		mockLocation = &location{}

		ig = &instagram{
			config: &config.Config{
				Username: "Johnny",
				Password: "123456",
			},

			authManager:    mockAuthManager,
			sessionManager: mockSessionManager,
			requestManager: mockRequestManager,

			timeline: mockTimeline,
			inbox:    mockInbox,
			thread:   mockThread,
			media:    mockMedia,
			location: mockLocation,
		}
	})

	Describe(".New", func() {
		var (
			ig  Instagram
			err error
		)

		BeforeEach(func() {
			cnf = &config.Config{
				Username: "Johnny",
				Password: "123456",
			}
		})

		JustBeforeEach(func() {
			ig, err = New(cnf)
		})

		Context("when success", func() {
			It("should return Instagram client", func() {
				Expect(err).To(BeNil())
				Expect(ig).NotTo(BeNil())
			})
		})

		tests := []struct {
			desc        string
			beforeFunc  func()
			expectedErr error
		}{
			{
				desc:        "when config is missing",
				beforeFunc:  func() { cnf = nil },
				expectedErr: ErrConfigRequired,
			},
			{
				desc:        "when username is missing",
				beforeFunc:  func() { cnf.Username = "" },
				expectedErr: ErrUsernameRequired,
			},
			{
				desc:        "when password is missing",
				beforeFunc:  func() { cnf.Password = "" },
				expectedErr: ErrPasswordRequired,
			},
		}

		for _, test := range tests {
			t := test
			Context(t.desc, func() {
				BeforeEach(t.beforeFunc)

				It("should return error", func() {
					Expect(ig).To(BeNil())
					Expect(err).NotTo(BeNil())
					Expect(err).To(Equal(t.expectedErr))
				})
			})
		}
	})

	Describe("#Login", func() {
		var (
			ctx context.Context

			mockGenerateUUIDResp            string
			mockGenerateSignatureKeyVersion string
			mockGenerateSignatureBody       string
			mockLoginResp                   protos.LoginResponse
			mockResp                        *http.Response
			mockBody                        string

			err error

			expectedGenerateSignatureParam *auth.SignaturePayload
			expectedPostReq                *protos.LoginRequest
			expectedSetCookiesParam        []*http.Cookie
		)

		BeforeEach(func() {
			ctx = context.Background()

			mockGenerateUUIDResp = uuid.NewV4().String()
			mockGenerateSignatureKeyVersion = "key version"
			mockGenerateSignatureBody = "sig body"

			cookie := &http.Cookie{}
			mockResp = &http.Response{
				Header: http.Header{
					"Set-Cookie": []string{cookie.String()},
				},
			}

			mockLoginResp = protos.LoginResponse{}
			mockLoginRespbytes, _ := json.Marshal(mockLoginResp)
			mockBody = string(mockLoginRespbytes)

			expectedGenerateSignatureParam = &auth.SignaturePayload{
				Csrftoken:         constants.SigCsrfToken,
				DeviceID:          constants.SigDeviceID,
				UUID:              mockGenerateUUIDResp,
				UserName:          ig.config.Username,
				Password:          ig.config.Password,
				LoginAttemptCount: 0,
			}

			expectedPostReq = &protos.LoginRequest{
				IgSigKeyVersion: mockGenerateSignatureKeyVersion,
				SignedBody:      mockGenerateSignatureBody,
			}

			expectedSetCookiesParam = mockResp.Cookies()
		})

		JustBeforeEach(func() {
			mockAuthManager.On("GenerateUUID").Return(mockGenerateUUIDResp)
			mockAuthManager.On("GenerateSignature", mock.Anything).Return(mockGenerateSignatureKeyVersion, mockGenerateSignatureBody, nil)
			mockRequestManager.On("Post", mock.Anything, mock.Anything, mock.Anything).Return(mockResp, mockBody, nil)
			mockSessionManager.On("SetCookies", mockResp.Cookies()).Return(nil)

			err = ig.Login(ctx)
		})

		Context("when success", func() {
			It("should return no error", func() {
				Expect(err).To(BeNil())
			})

			It("should call GenerateUUID", func() {
				mockAuthManager.AssertNumberOfCalls(GinkgoT(), "GenerateUUID", 1)
				mockAuthManager.AssertCalled(GinkgoT(), "GenerateUUID")
			})

			It("should call GenerateSignature", func() {
				mockAuthManager.AssertNumberOfCalls(GinkgoT(), "GenerateSignature", 1)
				mockAuthManager.AssertCalled(GinkgoT(), "GenerateSignature", expectedGenerateSignatureParam)
			})

			It("should call Post", func() {
				mockRequestManager.AssertNumberOfCalls(GinkgoT(), "Post", 1)
				mockRequestManager.AssertCalled(GinkgoT(), "Post", mock.Anything, constants.LoginEndpoint, expectedPostReq)
			})

			It("should call SetCookies", func() {
				mockSessionManager.AssertNumberOfCalls(GinkgoT(), "SetCookies", 1)
				mockSessionManager.AssertCalled(GinkgoT(), "SetCookies", expectedSetCookiesParam)
			})
		})

		Context("when return status is fail", func() {
			BeforeEach(func() {
				mockLoginResp.Status = constants.InstagramStatusFail
				mockLoginResp.Message = "oops"
				bytes, _ := json.Marshal(mockLoginResp)
				mockBody = string(bytes)
			})

			It("should return error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(mockLoginResp.Message))
			})

			It("should not call SetCookies", func() {
				mockSessionManager.AssertNumberOfCalls(GinkgoT(), "SetCookies", 0)
			})
		})
	})

	Describe("#Timeline", func() {
		var (
			client Timeline
		)

		JustBeforeEach(func() {
			client = ig.Timeline()
		})

		Context("when success", func() {
			It("should return client", func() {
				Expect(client).NotTo(BeNil())
			})
		})
	})

	Describe("#Inbox", func() {
		var (
			client Inbox
		)

		JustBeforeEach(func() {
			client = ig.Inbox()
		})

		Context("when success", func() {
			It("should return client", func() {
				Expect(client).NotTo(BeNil())
			})
		})
	})

	Describe("#Thread", func() {
		var (
			client Thread
		)

		JustBeforeEach(func() {
			client = ig.Thread()
		})

		Context("when success", func() {
			It("should return client", func() {
				Expect(client).NotTo(BeNil())
			})
		})
	})

	Describe("#Media", func() {
		var (
			client Media
		)

		JustBeforeEach(func() {
			client = ig.Media()
		})

		Context("when success", func() {
			It("should return client", func() {
				Expect(client).NotTo(BeNil())
			})
		})
	})

	Describe("#Location", func() {
		var (
			client Location
		)

		JustBeforeEach(func() {
			client = ig.Location()
		})

		Context("when success", func() {
			It("should return client", func() {
				Expect(client).NotTo(BeNil())
			})
		})
	})
})
