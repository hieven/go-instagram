package instagram

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/url"

	"github.com/stretchr/testify/mock"

	"github.com/hieven/go-instagram/src/constants"
	"github.com/hieven/go-instagram/src/protos"
	authMocks "github.com/hieven/go-instagram/src/utils/auth/mocks"
	requestMocks "github.com/hieven/go-instagram/src/utils/request/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("timeline", func() {
	var (
		mockAuthManager    *authMocks.AuthManager
		mockRequestManager *requestMocks.RequestManger

		ctx    context.Context
		client *timeline
	)

	BeforeEach(func() {
		mockAuthManager = &authMocks.AuthManager{}
		mockRequestManager = &requestMocks.RequestManger{}

		ctx = context.Background()

		client = &timeline{
			authManager:    mockAuthManager,
			requestManager: mockRequestManager,
		}
	})

	Describe("Feed", func() {
		var (
			req *TimelineFeedRequest

			mockGenerateRankTokenResp string
			mockResp                  *protos.TimelineFeedResponse
			mockBody                  string

			resp *protos.TimelineFeedResponse
			err  error

			expURLStru  *url.URL
			expURLQuery url.Values
			expURLStr   string
		)

		BeforeEach(func() {
			req = &TimelineFeedRequest{
				UserID: rand.Int63(),
			}

			mockGenerateRankTokenResp = "rank token"

			mockResp = &protos.TimelineFeedResponse{}
			mockBodyBytes, _ := json.Marshal(mockResp)
			mockBody = string(mockBodyBytes)

			expURLStru, _ = url.Parse(constants.TimelineFeedEndpoint)
			expURLQuery = expURLStru.Query()
		})

		JustBeforeEach(func() {
			mockAuthManager.On("GenerateRankToken", mock.Anything).Return(mockGenerateRankTokenResp)
			mockRequestManager.On("Post", mock.Anything, mock.Anything, mock.Anything).Return(nil, mockBody, nil)

			resp, err = client.Feed(ctx, req)

			expURLStru.RawQuery = expURLQuery.Encode()
			expURLStr = expURLStru.String()
		})

		Context("when success", func() {
			It("should return response", func() {
				Expect(err).To(BeNil())
				Expect(resp).NotTo(BeNil())
				Expect(resp).To(Equal(mockResp))
			})

			It("should call authManager.GenerateRankToken", func() {
				mockAuthManager.AssertNumberOfCalls(GinkgoT(), "GenerateRankToken", 1)
				mockAuthManager.AssertCalled(GinkgoT(), "GenerateRankToken", req.UserID)
			})

			It("should call requestManager.Post", func() {
				mockRequestManager.AssertNumberOfCalls(GinkgoT(), "Post", 1)
				mockRequestManager.AssertCalled(GinkgoT(), "Post", mock.Anything, expURLStr, mock.Anything)
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

		Context("when MaxID is provided", func() {
			BeforeEach(func() {
				req.MaxID = "max id"
			})

			It("should add it to request body", func() {
				mockRequestManager.AssertCalled(GinkgoT(), "Post",
					mock.Anything,
					mock.Anything,
					mock.MatchedBy(func(internalReq *protos.TimelineFeedRequest) bool {
						Expect(internalReq.MaxID).To(Equal(req.MaxID))

						return true
					}),
				)
			})
		})
	})
})
