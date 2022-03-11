package integration

import (
	"github.com/smartxworks/cloudtower-go-sdk/client/vm_entity_filter_result"
	"github.com/smartxworks/cloudtower-go-sdk/models"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vm entity filter api", Ordered, func() {
	It("should get vm entity filter result", func() {
		params := vm_entity_filter_result.NewGetVMEntityFilterResultsParams()
		params.RequestBody = &models.GetVMEntityFilterResultsRequestBody{}
		connectionParams := vm_entity_filter_result.NewGetVMEntityFilterResultsConnectionParams()
		connectionParams.RequestBody = &models.GetVMEntityFilterResultsConnectionRequestBody{}
		res, err := Client.VMEntityFilterResult.GetVMEntityFilterResults(params)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		connectionRes, err := Client.VMEntityFilterResult.GetVMEntityFilterResultsConnection(connectionParams)
		Expect(err).To(BeNil())
		Expect(connectionRes).ToNot(BeNil())
	})
})
