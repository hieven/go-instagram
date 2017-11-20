package instagram

import (
	"context"
	"encoding/json"

	"github.com/stretchr/testify/mock"

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

			mockFeedResp *protos.InboxFeedResponse
			mockBody     string

			resp *protos.InboxFeedResponse
			err  error
		)

		BeforeEach(func() {
			req = &InboxFeedRequest{}

			mockFeedResp = &protos.InboxFeedResponse{
				Inbox: &protos.Inbox{},
			}

			mockBodyBytes, _ := json.Marshal(mockFeedResp)
			mockBody = string(mockBodyBytes)
		})

		JustBeforeEach(func() {
			mockRequestManager.On("Get", mock.Anything, mock.Anything).Return(nil, mockBody, nil)

			resp, err = client.Feed(ctx, req)
		})

		Context("when success", func() {
			It("should return response", func() {
				Expect(err).To(BeNil())
				Expect(resp).NotTo(BeNil())
			})
		})
	})
})
