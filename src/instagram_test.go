package instagram_test

import (
	. "github.com/hieven/go-instagram/src"
	"github.com/hieven/go-instagram/src/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("instagram", func() {
	var (
		cnf *config.Config
		ig  Instagram
	)

	BeforeEach(func() {
		cnf = &config.Config{
			Username: "Johnny",
			Password: "123456",
		}

		ig, _ = New(cnf)
	})

	Describe(".New", func() {
		var (
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

	Describe("#Login", func() {})

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
})
