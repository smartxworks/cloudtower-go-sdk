package integration

import (
	"testing"

	apiclient "github.com/Sczlog/cloudtower-go-sdk/client"
	"github.com/Sczlog/cloudtower-go-sdk/integration/fixture"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var Client *apiclient.Cloudtower = nil

var _ = BeforeSuite(func() {
	Client = fixture.GetClient()
})
