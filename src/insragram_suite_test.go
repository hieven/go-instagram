package instagram_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInstagram(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Instagram Suite")
}
