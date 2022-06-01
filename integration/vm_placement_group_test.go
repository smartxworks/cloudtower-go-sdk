package integration

import (
	"fmt"
	"time"

	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_placement_group"
	"github.com/smartxworks/cloudtower-go-sdk/v2/integration/fixture"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	taskutil "github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/openlyinc/pointy"
)

var _ = Describe("Vm placment group api", Ordered, func() {
	var cluster *models.Cluster = nil
	BeforeAll(func() {
		cluster = fixture.GetDefaultCluster(Client, "xiaojun-502-ELF-nestedX20220426111859")
	})

	It("should get vm placement group", func() {
		params := vm_placement_group.NewGetVMPlacementGroupsParams()
		params.RequestBody = &models.GetVMPlacementGroupsRequestBody{}
		connectionParams := vm_placement_group.NewGetVMPlacementGroupsConnectionParams()
		connectionParams.RequestBody = &models.GetVMPlacementGroupsConnectionRequestBody{}
		res, err := Client.VMPlacementGroup.GetVMPlacementGroups(params)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		connectionRes, err := Client.VMPlacementGroup.GetVMPlacementGroupsConnection(connectionParams)
		Expect(err).To(BeNil())
		Expect(connectionRes).ToNot(BeNil())
	})

	It("should create, update and delete vm placement group", func() {
		createParams := vm_placement_group.NewCreateVMPlacementGroupParams()
		createParams.RequestBody = []*models.VMPlacementGroupCreationParams{
			{
				ClusterID:           cluster.ID,
				Name:                String(fmt.Sprintf("tower-go-sdk-vm-placement-group-%d", time.Now().Unix())),
				VMVMPolicy:          models.VMVMPolicyPREFERDIFFERENT.Pointer(),
				VMVMPolicyEnabled:   Bool(true),
				VMHostPreferPolicy:  Bool(true),
				VMHostPreferEnabled: Bool(false),
				VMHostMustPolicy:    Bool(false),
				VMHostMustEnabled:   Bool(false),
				Enabled:             Bool(true),
			},
		}
		createRes, err := Client.VMPlacementGroup.CreateVMPlacementGroup(createParams)
		Expect(err).To(BeNil())
		Expect(createRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, createRes.Payload[0].TaskID)
		Expect(err).To(BeNil())

		updateParams := vm_placement_group.NewUpdateVMPlacementGroupParams()
		updateParams.RequestBody = &models.VMPlacementGroupUpdationParams{
			Where: &models.VMPlacementGroupWhereInput{
				ID: createRes.Payload[0].Data.ID,
			},
			Data: &models.VMPlacementGroupUpdationParamsData{
				Description: String("Updated placement group"),
			},
		}
		updateRes, err := Client.VMPlacementGroup.UpdateVMPlacementGroup(updateParams)
		Expect(err).To(BeNil())
		Expect(updateRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, updateRes.Payload[0].TaskID)
		Expect(err).To(BeNil())

		deleteParams := vm_placement_group.NewDeleteVMPlacementGroupParams()
		deleteParams.RequestBody = &models.VMPlacementGroupDeletionParams{
			Where: &models.VMPlacementGroupWhereInput{
				ID: createRes.Payload[0].Data.ID,
			},
		}
		deleteRes, err := Client.VMPlacementGroup.DeleteVMPlacementGroup(deleteParams)
		Expect(err).To(BeNil())
		Expect(deleteRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, deleteRes.Payload[0].TaskID)
		Expect(err).To(BeNil())
	})

})
