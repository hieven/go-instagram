package instagram

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/satori/go.uuid"

	"github.com/stretchr/testify/mock"

	"github.com/hieven/go-instagram/src/constants"
	"github.com/hieven/go-instagram/src/protos"
	authMocks "github.com/hieven/go-instagram/src/utils/auth/mocks"
	requestMocks "github.com/hieven/go-instagram/src/utils/request/mocks"
	textMocks "github.com/hieven/go-instagram/src/utils/text/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("thread", func() {
	var (
		mockAuthManager    *authMocks.AuthManager
		mockRequestManager *requestMocks.RequestManger
		mockTextManager    *textMocks.TextManager

		ctx    context.Context
		client *thread
	)

	BeforeEach(func() {
		mockAuthManager = &authMocks.AuthManager{}
		mockRequestManager = &requestMocks.RequestManger{}
		mockTextManager = &textMocks.TextManager{}

		ctx = context.Background()

		client = &thread{
			authManager:    mockAuthManager,
			requestManager: mockRequestManager,
			textManager:    mockTextManager,
		}
	})

	Describe("ApprovedAll", func() {
		var (
			req *ThreadApproveAllRequest

			mockGenerateUUIDResp string
			mockResp             *protos.ThreadApproveAllResponse
			mockBody             string

			resp *protos.ThreadApproveAllResponse
			err  error

			expURLStru     *url.URL
			expURLQuery    url.Values
			expURLStr      string
			expInternalReq *protos.ThreadApproveAllRequest
		)

		BeforeEach(func() {
			req = &ThreadApproveAllRequest{}

			mockGenerateUUIDResp = uuid.NewV4().String()

			mockResp = &protos.ThreadApproveAllResponse{}

			expURLStru, _ = url.Parse(constants.ThreadApproveAllEndpoint)
			expURLQuery = expURLStru.Query()

			expInternalReq = &protos.ThreadApproveAllRequest{
				UUID: mockGenerateUUIDResp,
			}
		})

		JustBeforeEach(func() {
			mockBodyBytes, _ := json.Marshal(mockResp)
			mockBody = string(mockBodyBytes)

			mockAuthManager.On("GenerateUUID").Return(mockGenerateUUIDResp)
			mockRequestManager.On("Post", mock.Anything, mock.Anything, mock.Anything).Return(nil, mockBody, nil)

			resp, err = client.ApproveAll(ctx, req)

			expURLStru.RawQuery = expURLQuery.Encode()
			expURLStr = expURLStru.String()
		})

		Context("when success", func() {
			It("should return response", func() {
				Expect(err).To(BeNil())
				Expect(resp).NotTo(BeNil())
				Expect(resp).To(Equal(mockResp))
			})

			It("should call authManager.GenerateUUID", func() {
				mockAuthManager.AssertNumberOfCalls(GinkgoT(), "GenerateUUID", 1)
				mockAuthManager.AssertCalled(GinkgoT(), "GenerateUUID")
			})

			It("should call requestManager.Get", func() {
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

		Context("when user didn't login", func() {
			BeforeEach(func() {
				mockResp = &protos.ThreadApproveAllResponse{}
				mockResp.Status = instaStatusFail
				mockResp.Message = instaMsgLoginRequired
			})

			It("should return login required error", func() {
				Expect(resp).NotTo(BeNil())
				Expect(resp.Status).To(Equal(mockResp.Status))
				Expect(resp.Message).To(Equal(mockResp.Message))

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(ErrLoginRequired.Error()))
			})
		})

		Context("when unknown error happens", func() {
			BeforeEach(func() {
				mockResp = &protos.ThreadApproveAllResponse{}
				mockResp.Status = instaStatusFail
				mockResp.Message = "unknown error"
			})

			It("should return error", func() {
				Expect(resp).NotTo(BeNil())
				Expect(resp.Status).To(Equal(mockResp.Status))
				Expect(resp.Message).To(Equal(mockResp.Message))

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(ErrUnknown.Error()))
			})
		})
	})

	Describe("BroadcastText", func() {
		var (
			req *ThreadBroadcastTextRequest

			mockGenerateUUIDResp string
			mockResp             *protos.ThreadBroadcastTextResponse
			mockBody             string

			resp *protos.ThreadBroadcastTextResponse
			err  error

			expURLStru     *url.URL
			expURLQuery    url.Values
			expURLStr      string
			expInternalReq *protos.ThreadBroadcastTextRequest
		)

		BeforeEach(func() {
			req = &ThreadBroadcastTextRequest{
				ThreadIDs: "thread id",
				Text:      "text",
			}

			mockGenerateUUIDResp = uuid.NewV4().String()

			mockResp = &protos.ThreadBroadcastTextResponse{
				Threads: []*protos.Thread{
					{ThreadID: "thread id"},
				},
			}

			expURLStru, _ = url.Parse(constants.ThreadBroadcastTextEndpoint)
			expURLQuery = expURLStru.Query()

			expInternalReq = &protos.ThreadBroadcastTextRequest{
				ThreadIDs:     "[" + req.ThreadIDs + "]",
				Text:          req.Text,
				UUID:          mockGenerateUUIDResp,
				ClientContext: mockGenerateUUIDResp,
			}
		})

		JustBeforeEach(func() {
			mockBodyBytes, _ := json.Marshal(mockResp)
			mockBody = string(mockBodyBytes)

			mockAuthManager.On("GenerateUUID").Return(mockGenerateUUIDResp)
			mockRequestManager.On("Post", mock.Anything, mock.Anything, mock.Anything).Return(nil, mockBody, nil)

			resp, err = client.BroadcastText(ctx, req)

			expURLStru.RawQuery = expURLQuery.Encode()
			expURLStr = expURLStru.String()
		})

		Context("when success", func() {
			It("should return response", func() {
				Expect(err).To(BeNil())
				Expect(resp).NotTo(BeNil())
				Expect(resp).To(Equal(mockResp))
			})

			It("should call authManager.GenerateUUID", func() {
				mockAuthManager.AssertNumberOfCalls(GinkgoT(), "GenerateUUID", 2)
				mockAuthManager.AssertCalled(GinkgoT(), "GenerateUUID")
			})

			It("should call requestManager.Get", func() {
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

		Context("when user didn't login", func() {
			BeforeEach(func() {
				mockResp = &protos.ThreadBroadcastTextResponse{}
				mockResp.Status = instaStatusFail
				mockResp.Message = instaMsgLoginRequired
			})

			It("should return login required error", func() {
				Expect(resp).NotTo(BeNil())
				Expect(resp.Status).To(Equal(mockResp.Status))
				Expect(resp.Message).To(Equal(mockResp.Message))

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(ErrLoginRequired.Error()))
			})
		})

		Context("when unknown error happens", func() {
			BeforeEach(func() {
				mockResp = &protos.ThreadBroadcastTextResponse{}
				mockResp.Status = instaStatusFail
				mockResp.Message = "unknown error"
			})

			It("should return error", func() {
				Expect(resp).NotTo(BeNil())
				Expect(resp.Status).To(Equal(mockResp.Status))
				Expect(resp.Message).To(Equal(mockResp.Message))

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(ErrUnknown.Error()))
			})
		})
	})

	Describe("BroadcastLink", func() {
		var (
			req *ThreadBroadcastLinkRequest

			mockGenerateUUIDResp string
			mockResp             *protos.ThreadBroadcastLinkResponse
			mockBody             string
			mockExtractURLResp   string

			resp *protos.ThreadBroadcastLinkResponse
			err  error

			expURLStru     *url.URL
			expURLQuery    url.Values
			expURLStr      string
			expInternalReq *protos.ThreadBroadcastLinkRequest
		)

		BeforeEach(func() {
			req = &ThreadBroadcastLinkRequest{
				ThreadIDs: "thread id",
				LinkText:  "https://www.google.com",
			}

			mockGenerateUUIDResp = uuid.NewV4().String()

			mockResp = &protos.ThreadBroadcastLinkResponse{
				Threads: []*protos.Thread{
					{ThreadID: "thread id"},
				},
			}

			mockExtractURLResp = "mock url string"

			expURLStru, _ = url.Parse(constants.ThreadBroadcastLinkEndpoint)
			expURLQuery = expURLStru.Query()

			expInternalReq = &protos.ThreadBroadcastLinkRequest{
				ThreadIDs:     "[" + req.ThreadIDs + "]",
				LinkText:      req.LinkText,
				LinkURLs:      `["` + mockExtractURLResp + `"]`,
				UUID:          mockGenerateUUIDResp,
				ClientContext: mockGenerateUUIDResp,
			}
		})

		JustBeforeEach(func() {
			mockBodyBytes, _ := json.Marshal(mockResp)
			mockBody = string(mockBodyBytes)

			mockTextManager.On("ExtractURL", mock.Anything).Return(mockExtractURLResp)
			mockAuthManager.On("GenerateUUID").Return(mockGenerateUUIDResp)
			mockRequestManager.On("Post", mock.Anything, mock.Anything, mock.Anything).Return(nil, mockBody, nil)

			resp, err = client.BroadcastLink(ctx, req)

			expURLStru.RawQuery = expURLQuery.Encode()
			expURLStr = expURLStru.String()
		})

		Context("when success", func() {
			It("should return response", func() {
				Expect(err).To(BeNil())
				Expect(resp).NotTo(BeNil())
				Expect(resp).To(Equal(mockResp))
			})

			It("should call textManager.ExtractURL", func() {
				mockTextManager.AssertNumberOfCalls(GinkgoT(), "ExtractURL", 1)
				mockTextManager.AssertCalled(GinkgoT(), "ExtractURL", req.LinkText)
			})

			It("should call authManager.GenerateUUID", func() {
				mockAuthManager.AssertNumberOfCalls(GinkgoT(), "GenerateUUID", 2)
				mockAuthManager.AssertCalled(GinkgoT(), "GenerateUUID")
			})

			It("should call requestManager.Get", func() {
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

		Context("when user didn't login", func() {
			BeforeEach(func() {
				mockResp = &protos.ThreadBroadcastLinkResponse{}
				mockResp.Status = instaStatusFail
				mockResp.Message = instaMsgLoginRequired
			})

			It("should return login required error", func() {
				Expect(resp).NotTo(BeNil())
				Expect(resp.Status).To(Equal(mockResp.Status))
				Expect(resp.Message).To(Equal(mockResp.Message))

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(ErrLoginRequired.Error()))
			})
		})

		Context("when unknown error happens", func() {
			BeforeEach(func() {
				mockResp = &protos.ThreadBroadcastLinkResponse{}
				mockResp.Status = instaStatusFail
				mockResp.Message = "unknown error"
			})

			It("should return error", func() {
				Expect(resp).NotTo(BeNil())
				Expect(resp.Status).To(Equal(mockResp.Status))
				Expect(resp.Message).To(Equal(mockResp.Message))

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(ErrUnknown.Error()))
			})
		})
	})

	Describe("BroadcastShare", func() {
		var (
			req *ThreadBroadcastShareRequest

			mockGenerateUUIDResp string
			mockResp             *protos.ThreadBroadcastShareResponse
			mockBody             string

			resp *protos.ThreadBroadcastShareResponse
			err  error

			expURLStru     *url.URL
			expURLQuery    url.Values
			expURLStr      string
			expInternalReq *protos.ThreadBroadcastShareRequest
		)

		BeforeEach(func() {
			req = &ThreadBroadcastShareRequest{
				ThreadIDs: "thread id",
				MediaID:   "media id",
				Text:      "text",
			}

			mockGenerateUUIDResp = uuid.NewV4().String()

			mockResp = &protos.ThreadBroadcastShareResponse{
				Threads: []*protos.Thread{
					{ThreadID: "thread id"},
				},
			}

			expURLStru, _ = url.Parse(constants.ThreadBroadcastShareEndpoint)
			expURLQuery = expURLStru.Query()

			expInternalReq = &protos.ThreadBroadcastShareRequest{
				ThreadIDs:     "[" + req.ThreadIDs + "]",
				Text:          req.Text,
				MediaID:       req.MediaID,
				ClientContext: mockGenerateUUIDResp,
			}
		})

		JustBeforeEach(func() {
			mockBodyBytes, _ := json.Marshal(mockResp)
			mockBody = string(mockBodyBytes)

			mockAuthManager.On("GenerateUUID").Return(mockGenerateUUIDResp)
			mockRequestManager.On("Post", mock.Anything, mock.Anything, mock.Anything).Return(nil, mockBody, nil)

			resp, err = client.BroadcastShare(ctx, req)

			expURLStru.RawQuery = expURLQuery.Encode()
			expURLStr = expURLStru.String()
		})

		Context("when success", func() {
			It("should return response", func() {
				Expect(err).To(BeNil())
				Expect(resp).NotTo(BeNil())
				Expect(resp).To(Equal(mockResp))
			})

			It("should call authManager.GenerateUUID", func() {
				mockAuthManager.AssertNumberOfCalls(GinkgoT(), "GenerateUUID", 1)
				mockAuthManager.AssertCalled(GinkgoT(), "GenerateUUID")
			})

			It("should call requestManager.Get", func() {
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

		Context("when user didn't login", func() {
			BeforeEach(func() {
				mockResp = &protos.ThreadBroadcastShareResponse{}
				mockResp.Status = instaStatusFail
				mockResp.Message = instaMsgLoginRequired
			})

			It("should return login required error", func() {
				Expect(resp).NotTo(BeNil())
				Expect(resp.Status).To(Equal(mockResp.Status))
				Expect(resp.Message).To(Equal(mockResp.Message))

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(ErrLoginRequired.Error()))
			})
		})

		Context("when unknown error happens", func() {
			BeforeEach(func() {
				mockResp = &protos.ThreadBroadcastShareResponse{}
				mockResp.Status = instaStatusFail
				mockResp.Message = "unknown error"
			})

			It("should return error", func() {
				Expect(resp).NotTo(BeNil())
				Expect(resp.Status).To(Equal(mockResp.Status))
				Expect(resp.Message).To(Equal(mockResp.Message))

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(ErrUnknown.Error()))
			})
		})
	})

	Describe("Show", func() {
		var (
			req *ThreadShowRequest

			mockResp *protos.ThreadShowResponse
			mockBody string

			resp *protos.ThreadShowResponse
			err  error

			expURLStru  *url.URL
			expURLQuery url.Values
			expURLStr   string
		)

		BeforeEach(func() {
			req = &ThreadShowRequest{
				ThreadID: "thread id",
			}

			mockResp = &protos.ThreadShowResponse{
				Thread: &protos.Thread{},
			}

			expURLStru, _ = url.Parse(fmt.Sprintf(constants.ThreadShowEndpoint, req.ThreadID))
			expURLQuery = expURLStru.Query()
		})

		JustBeforeEach(func() {
			mockBodyBytes, _ := json.Marshal(mockResp)
			mockBody = string(mockBodyBytes)

			mockRequestManager.On("Get", mock.Anything, mock.Anything).Return(nil, mockBody, nil)

			resp, err = client.Show(ctx, req)

			expURLStru.RawQuery = expURLQuery.Encode()
			expURLStr = expURLStru.String()
		})

		Context("when success", func() {
			It("should return response", func() {
				Expect(err).To(BeNil())
				Expect(resp).NotTo(BeNil())
				Expect(resp).To(Equal(mockResp))
				Expect(resp).To(Equal(mockResp))
			})

			It("should call requestManager.Get", func() {
				mockRequestManager.AssertNumberOfCalls(GinkgoT(), "Get", 1)
				mockRequestManager.AssertCalled(GinkgoT(), "Get", mock.Anything, expURLStr)
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

		Context("when user didn't login", func() {
			BeforeEach(func() {
				mockResp = &protos.ThreadShowResponse{}
				mockResp.Status = instaStatusFail
				mockResp.Message = instaMsgLoginRequired
			})

			It("should return login required error", func() {
				Expect(resp).NotTo(BeNil())
				Expect(resp.Status).To(Equal(mockResp.Status))
				Expect(resp.Message).To(Equal(mockResp.Message))

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(ErrLoginRequired.Error()))
			})
		})

		Context("when unknown error happens", func() {
			BeforeEach(func() {
				mockResp = &protos.ThreadShowResponse{}
				mockResp.Status = instaStatusFail
				mockResp.Message = "unknown error"
			})

			It("should return error", func() {
				Expect(resp).NotTo(BeNil())
				Expect(resp.Status).To(Equal(mockResp.Status))
				Expect(resp.Message).To(Equal(mockResp.Message))

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(ErrUnknown.Error()))
			})
		})

		Context("when Instagram is down", func() {
			BeforeEach(func() {
				mockResp = &protos.ThreadShowResponse{}
			})

			It("should return error", func() {
				Expect(resp).NotTo(BeNil())
				Expect(resp.Status).To(Equal(mockResp.Status))
				Expect(resp.Message).To(Equal(mockResp.Message))

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal(ErrUnknown.Error()))
			})
		})
	})
})
