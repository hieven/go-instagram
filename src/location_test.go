package instagram

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"

	"github.com/satori/go.uuid"

	"github.com/stretchr/testify/mock"

	"github.com/hieven/go-instagram/src/constants"
	"github.com/hieven/go-instagram/src/protos"
	authMocks "github.com/hieven/go-instagram/src/utils/auth/mocks"
	requestMocks "github.com/hieven/go-instagram/src/utils/request/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("location", func() {
	var (
		mockAuthManager    *authMocks.AuthManager
		mockRequestManager *requestMocks.RequestManger

		ctx    context.Context
		client *location
	)

	BeforeEach(func() {
		mockAuthManager = &authMocks.AuthManager{}
		mockRequestManager = &requestMocks.RequestManger{}

		ctx = context.Background()

		client = &location{
			authManager:    mockAuthManager,
			requestManager: mockRequestManager,
		}
	})

	Describe("Feed", func() {
		var (
			req *LocationFeedRequest

			mockGenerateUUIDResp string
			mockResp             *protos.LocationFeedResponse
			mockBody             string

			resp *protos.LocationFeedResponse
			err  error

			expURLStru  *url.URL
			expURLQuery url.Values
			expURLStr   string
		)

		BeforeEach(func() {
			req = &LocationFeedRequest{
				Pk: rand.Int63(),
			}

			mockGenerateUUIDResp = uuid.NewV4().String()

			mockResp = &protos.LocationFeedResponse{}
			mockBodyBytes, _ := json.Marshal(mockResp)
			mockBody = string(mockBodyBytes)

			expURLStru, _ = url.Parse(fmt.Sprintf(constants.LocationFeedEndpoint, req.Pk))
			expURLQuery = expURLStru.Query()
			expURLQuery.Set("rank_token", mockGenerateUUIDResp)
		})

		JustBeforeEach(func() {
			mockAuthManager.On("GenerateUUID").Return(mockGenerateUUIDResp)
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

			It("should call authManager.GenerateUUID", func() {
				mockAuthManager.AssertNumberOfCalls(GinkgoT(), "GenerateUUID", 1)
				mockAuthManager.AssertCalled(GinkgoT(), "GenerateUUID")
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
	})

	Describe("Section", func() {
		var (
			req *LocationSectionRequest

			mockGenerateUUIDResp string
			mockResp             *protos.LocationSectionResponse
			mockBody             string

			resp *protos.LocationSectionResponse
			err  error

			expURLStru *url.URL
			expURLStr  string
		)

		BeforeEach(func() {
			req = &LocationSectionRequest{
				Pk:  rand.Int63(),
				Tab: LocationSectionTabRanked,
			}

			mockGenerateUUIDResp = uuid.NewV4().String()

			mockResp = &protos.LocationSectionResponse{}
			mockBodyBytes, _ := json.Marshal(mockResp)
			mockBody = string(mockBodyBytes)

			expURLStru, _ = url.Parse(fmt.Sprintf(constants.LocationSectionEndpoint, req.Pk))
		})

		JustBeforeEach(func() {
			mockAuthManager.On("GenerateUUID").Return(mockGenerateUUIDResp)
			mockRequestManager.On("Post", mock.Anything, mock.Anything, mock.Anything).Return(nil, mockBody, nil)

			resp, err = client.Section(ctx, req)

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
	})
})
