package request

import (
	"context"
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/mock"

	requestMocks "github.com/hieven/go-instagram/src/utils/request/mocks"
	"github.com/hieven/go-instagram/src/utils/session"
	sessionMocks "github.com/hieven/go-instagram/src/utils/session/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("client", func() {
	var (
		mockSessionManager *sessionMocks.SessionManager
		mockCommon         *requestMocks.Common

		manager *requestManager
	)

	BeforeEach(func() {
		mockSessionManager = &sessionMocks.SessionManager{}
		mockCommon = &requestMocks.Common{}

		manager = &requestManager{
			sessionManager: mockSessionManager,
		}
	})

	Describe(".New", func() {
		var (
			manager RequestManger
			err     error
		)

		JustBeforeEach(func() {
			manager, err = New(mockSessionManager)
		})

		Context("when success", func() {
			It("should return manager", func() {
				Expect(manager).NotTo(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#Get", func() {
		var (
			ctx    context.Context
			urlStr string

			resp *http.Response
			body string
			err  error

			expMethod string

			oriWithDefaultHeader func(sessionManager session.SessionManager, req *gorequest.SuperAgent) *gorequest.SuperAgent
		)

		BeforeEach(func() {
			ctx = context.Background()

			tsHandler = func(w http.ResponseWriter, r *http.Request) {
				expMethod = r.Method
				fmt.Fprintln(w, "Hello, client")
			}
			urlStr = ts.URL

			mockCommon.On("WithDefaultHeader", mock.Anything, mock.Anything).Return(nil)
			oriWithDefaultHeader = withDefaultHeader
			withDefaultHeader = mockCommon.WithDefaultHeader

			expMethod = ""
		})

		JustBeforeEach(func() {
			resp, body, err = manager.Get(ctx, urlStr)
		})

		AfterEach(func() {
			withDefaultHeader = oriWithDefaultHeader
		})

		Context("when success", func() {
			It("should return result", func() {
				Expect(resp).NotTo(BeNil())
				Expect(body).NotTo(BeEmpty())
				Expect(err).To(BeNil())
				Expect(expMethod).To(Equal(http.MethodGet))
			})

			It("should call withDefaultHeader", func() {
				mockCommon.AssertNumberOfCalls(GinkgoT(), "WithDefaultHeader", 1)
			})
		})
	})

	Describe("#Post", func() {
		var (
			ctx    context.Context
			urlStr string

			resp *http.Response
			body string
			err  error

			expMethod string

			oriWithDefaultHeader func(sessionManager session.SessionManager, req *gorequest.SuperAgent) *gorequest.SuperAgent
		)

		BeforeEach(func() {
			ctx = context.Background()

			tsHandler = func(w http.ResponseWriter, r *http.Request) {
				expMethod = r.Method
				fmt.Fprintln(w, "Hello, client")
			}
			urlStr = ts.URL

			mockCommon.On("WithDefaultHeader", mock.Anything, mock.Anything).Return(nil)
			oriWithDefaultHeader = withDefaultHeader
			withDefaultHeader = mockCommon.WithDefaultHeader

			expMethod = ""
		})

		JustBeforeEach(func() {
			resp, body, err = manager.Post(ctx, urlStr, nil)
		})

		AfterEach(func() {
			withDefaultHeader = oriWithDefaultHeader
		})

		Context("when success", func() {
			It("should return result", func() {
				Expect(resp).NotTo(BeNil())
				Expect(body).NotTo(BeEmpty())
				Expect(err).To(BeNil())
				Expect(expMethod).To(Equal(http.MethodPost))
			})

			It("should call withDefaultHeader", func() {
				mockCommon.AssertNumberOfCalls(GinkgoT(), "WithDefaultHeader", 1)
			})
		})
	})

})
