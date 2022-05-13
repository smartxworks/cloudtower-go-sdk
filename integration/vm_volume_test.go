package integration

import (
	"fmt"
	"time"

	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_volume"
	"github.com/smartxworks/cloudtower-go-sdk/v2/integration/fixture"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	taskutil "github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/openlyinc/pointy"
)

var _ = Describe("Vm volume api", Ordered, func() {
	var cluster *models.Cluster = nil
	BeforeAll(func() {
		cluster = fixture.GetDefaultCluster(Client, "xiaojun-nested-X20211110124941")
	})

	It("should get vm volume", func() {
		params := vm_volume.NewGetVMVolumesParams()
		params.RequestBody = &models.GetVMVolumesRequestBody{}
		connectionParams := vm_volume.NewGetVMVolumesConnectionParams()
		connectionParams.RequestBody = &models.GetVMVolumesConnectionRequestBody{}
		res, err := Client.VMVolume.GetVMVolumes(params)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		connectionRes, err := Client.VMVolume.GetVMVolumesConnection(connectionParams)
		Expect(err).To(BeNil())
		Expect(connectionRes).ToNot(BeNil())
	})

	It("should create and delete vm volume", func() {
		createParams := vm_volume.NewCreateVMVolumeParams()
		createParams.RequestBody = []*models.VMVolumeCreationParams{
			{
				ClusterID:        cluster.ID,
				Name:             String(fmt.Sprintf("tower-go-sdk-vm-volume-%d", time.Now().Unix())),
				ElfStoragePolicy: models.VMVolumeElfStoragePolicyTypeREPLICA2THINPROVISION.Pointer(),
				Size:             Int64(1024 * 1024 * 1024),
				Sharing:          Bool(false),
			},
		}
		createRes, err := Client.VMVolume.CreateVMVolume(createParams)
		Expect(err).To(BeNil())
		Expect(createRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, createRes.Payload[0].TaskID)
		Expect(err).To(BeNil())

		deleteParams := vm_volume.NewDeleteVMVolumeFromVMParams()
		deleteParams.RequestBody = &models.VMVolumeDeletionParams{
			Where: &models.VMVolumeWhereInput{
				ID: createRes.Payload[0].Data.ID,
			},
		}
		deleteRes, err := Client.VMVolume.DeleteVMVolumeFromVM(deleteParams)
		Expect(err).To(BeNil())
		Expect(deleteRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, deleteRes.Payload[0].TaskID)
		Expect(err).To(BeNil())
	})

})
