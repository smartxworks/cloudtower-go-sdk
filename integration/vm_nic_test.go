package integration

import (
	"github.com/smartxworks/cloudtower-go-sdk/client/vm_nic"
	"github.com/smartxworks/cloudtower-go-sdk/models"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vm nic api", Ordered, func() {
	It("should get vm nics", func() {
		params := vm_nic.NewGetVMNicsParams()
		params.RequestBody = &models.GetVMNicsRequestBody{}
		connectionParams := vm_nic.NewGetVMNicsConnectionParams()
		connectionParams.RequestBody = &models.GetVMNicsConnectionRequestBody{}
		res, err := Client.VMNic.GetVMNics(params)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		connectionRes, err := Client.VMNic.GetVMNicsConnection(connectionParams)
		Expect(err).To(BeNil())
		Expect(connectionRes).ToNot(BeNil())
	})
})
