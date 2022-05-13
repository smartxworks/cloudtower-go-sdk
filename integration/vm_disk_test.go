package integration

import (
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_disk"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vm disk api", Ordered, func() {
	It("should get vm disks", func() {
		params := vm_disk.NewGetVMDisksParams()
		params.RequestBody = &models.GetVMDisksRequestBody{}
		connectionParams := vm_disk.NewGetVMDisksConnectionParams()
		connectionParams.RequestBody = &models.GetVMDisksConnectionRequestBody{}
		res, err := Client.VMDisk.GetVMDisks(params)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		connectionRes, err := Client.VMDisk.GetVMDisksConnection(connectionParams)
		Expect(err).To(BeNil())
		Expect(connectionRes).ToNot(BeNil())
	})
})
