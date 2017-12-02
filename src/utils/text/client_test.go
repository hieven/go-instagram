package text

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("client", func() {
	var (
		manager *textManager
	)

	BeforeEach(func() {
		manager = &textManager{}
	})

	Describe(".New", func() {
		var (
			manager TextManager
			err     error
		)

		JustBeforeEach(func() {
			manager, err = New()
		})

		Context("when success", func() {
			It("should return manager", func() {
				Expect(manager).NotTo(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#ExtractURL", func() {
		var (
			text string

			result string

			expURL string
		)

		BeforeEach(func() {
			expURL = "https://www.google.com.tw"

			text = "hello " + expURL
		})

		JustBeforeEach(func() {
			result = manager.ExtractURL(text)
		})

		Context("when success", func() {
			It("should return url", func() {
				Expect(result).To(Equal(expURL))
			})
		})
	})
})
