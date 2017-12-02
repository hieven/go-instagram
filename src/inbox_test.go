package instagram

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/stretchr/testify/mock"

	"github.com/hieven/go-instagram/src/constants"
	"github.com/hieven/go-instagram/src/protos"
	requestMocks "github.com/hieven/go-instagram/src/utils/request/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("inbox", func() {
	var (
		mockRequestManager *requestMocks.RequestManger

		ctx    context.Context
		client *inbox
	)

	BeforeEach(func() {
		mockRequestManager = &requestMocks.RequestManger{}

		ctx = context.Background()

		client = &inbox{
			requestManager: mockRequestManager,
		}
	})

	Describe("Feed", func() {
		var (
			req *InboxFeedRequest

			mockResp *protos.InboxFeedResponse
			mockBody string

			resp *protos.InboxFeedResponse
			err  error

			expURLStru  *url.URL
			expURLQuery url.Values
			expURLStr   string
		)

		BeforeEach(func() {
			req = &InboxFeedRequest{}

			mockResp = &protos.InboxFeedResponse{
				Inbox: &protos.Inbox{},
			}
			mockBodyBytes, _ := json.Marshal(mockResp)
			mockBody = string(mockBodyBytes)

			expURLStru, _ = url.Parse(constants.InboxEndpoint)
			expURLQuery = expURLStru.Query()
		})

		JustBeforeEach(func() {
			mockRequestManager.On("Get", mock.Anything, mock.Anything).Return(nil, mockBody, nil)

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

		Context("when cursor is provided", func() {
			BeforeEach(func() {
				req.Cursor = "hello"
				expURLQuery.Set("cursor", req.Cursor)
			})

			It("should be added to query string", func() {
				mockRequestManager.AssertCalled(GinkgoT(), "Get", mock.Anything, expURLStr)
			})
		})
	})
})
