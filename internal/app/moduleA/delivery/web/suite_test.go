package web_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSigningDeliveryWebHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	suiteConfig, reporterConfig := GinkgoConfiguration()
	reporterConfig.NoColor = true // Prevents unsupported ANSI color characters.
	RunSpecs(t, "Module A Web Handler Suite", suiteConfig, reporterConfig)
}
