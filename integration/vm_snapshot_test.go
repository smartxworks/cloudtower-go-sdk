package integration

import (
	"fmt"
	"time"

	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_snapshot"
	"github.com/smartxworks/cloudtower-go-sdk/v2/integration/fixture"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	taskutil "github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/openlyinc/pointy"
)

var _ = Describe("Vm snapshot api", Ordered, func() {
	var cluster *models.Cluster = nil
	var vlan *models.Vlan = nil
	BeforeAll(func() {
		cluster = fixture.GetDefaultCluster(Client, "xiaojun-nested-X20211110124941")
		vlan = fixture.GetDefaultVlan(Client, cluster.ID)
	})

	It("should get vm snapshot", func() {
		params := vm_snapshot.NewGetVMSnapshotsParams()
		params.RequestBody = &models.GetVMSnapshotsRequestBody{}
		connectionParams := vm_snapshot.NewGetVMSnapshotsConnectionParams()
		connectionParams.RequestBody = &models.GetVMSnapshotsConnectionRequestBody{}
		res, err := Client.VMSnapshot.GetVMSnapshots(params)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		connectionRes, err := Client.VMSnapshot.GetVMSnapshotsConnection(connectionParams)
		Expect(err).To(BeNil())
		Expect(connectionRes).ToNot(BeNil())
	})

	It("should create and delete vm snapshot", func() {
		vmId, disposeVm := fixture.CreateVm(Client, cluster.ID, vlan.ID, false)
		defer disposeVm()
		createParams := vm_snapshot.NewCreateVMSnapshotParams()
		createParams.RequestBody = &models.VMSnapshotCreationParams{
			Data: []*models.VMSnapshotCreationParamsData{
				{
					VMID: vmId,
					Name: String(fmt.Sprintf("tower-go-sdk-vm-snapshot-%d", time.Now().Unix())),
				},
			},
		}
		createRes, err := Client.VMSnapshot.CreateVMSnapshot(createParams)
		Expect(err).To(BeNil())
		Expect(createRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, createRes.Payload[0].TaskID)
		Expect(err).To(BeNil())

		deleteParams := vm_snapshot.NewDeleteVMSnapshotParams()
		deleteParams.RequestBody = &models.VMSnapshotDeletionParams{
			Where: &models.VMSnapshotWhereInput{
				ID: createRes.Payload[0].Data.ID,
			},
		}
		deleteRes, err := Client.VMSnapshot.DeleteVMSnapshot(deleteParams)
		Expect(err).To(BeNil())
		Expect(deleteRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, deleteRes.Payload[0].TaskID)
		Expect(err).To(BeNil())
	})

	It("should rollback vm", func() {
		snapShotId, vmId, dispose := fixture.CreateSnapshot(Client, cluster.ID, vlan.ID)
		defer dispose()
		rollbackParams := vm.NewRollbackVMParams()
		rollbackParams.RequestBody = &models.VMRollbackParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
			Data: &models.VMRollbackParamsData{
				SnapshotID: snapShotId,
			},
		}
		rollbackRes, err := Client.VM.RollbackVM(rollbackParams)
		Expect(err).To(BeNil())
		Expect(rollbackRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, rollbackRes.Payload[0].TaskID)
		Expect(err).To(BeNil())
	})

	It("should rebuild vm", func() {
		snapShotId, _, dispose := fixture.CreateSnapshot(Client, cluster.ID, vlan.ID)
		defer dispose()
		rebuildParams := vm.NewRebuildVMParams()
		rebuildParams.RequestBody = []*models.VMRebuildParams{
			{
				ClusterID:             cluster.ID,
				RebuildFromSnapshotID: snapShotId,
				Name:                  String(fmt.Sprintf("tower-go-sdk-vm-rebuild-%d", time.Now().Unix())),
			},
		}
		rebuildRes, err := Client.VM.RebuildVM(rebuildParams)
		Expect(err).To(BeNil())
		Expect(rebuildRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, rebuildRes.Payload[0].TaskID)
		Expect(err).To(BeNil())
		deleteParams := vm.NewDeleteVMParams()
		deleteParams.RequestBody = &models.VMOperateParams{
			Where: &models.VMWhereInput{
				ID: rebuildRes.Payload[0].Data.ID,
			},
		}
	})
})
