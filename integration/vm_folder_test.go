package integration

import (
	"fmt"
	"sync"
	"time"

	"github.com/Sczlog/cloudtower-go-sdk/client/vm"
	"github.com/Sczlog/cloudtower-go-sdk/client/vm_folder"
	"github.com/Sczlog/cloudtower-go-sdk/integration/fixture"
	"github.com/Sczlog/cloudtower-go-sdk/models"
	taskutil "github.com/Sczlog/cloudtower-go-sdk/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/openlyinc/pointy"
)

var _ = Describe("Vm folder api", Ordered, func() {
	var cluster *models.Cluster = nil
	var vlan *models.Vlan = nil
	BeforeAll(func() {
		cluster = fixture.GetDefaultCluster(Client, "xiaojun-nested-X20211110124941")
		vlan = fixture.GetDefaultVlan(Client, cluster.ID)
	})
	It("should get vm folders", func() {
		params := vm_folder.NewGetVMFoldersParams()
		params.RequestBody = &models.GetVMFoldersRequestBody{}
		connectionParams := vm_folder.NewGetVMFoldersConnectionParams()
		connectionParams.RequestBody = &models.GetVMFoldersConnectionRequestBody{}
		res, err := Client.VMFolder.GetVMFolders(params)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		connectionRes, err := Client.VMFolder.GetVMFoldersConnection(connectionParams)
		Expect(err).To(BeNil())
		Expect(connectionRes).ToNot(BeNil())
	})

	It("should create, update and delete vm folder", func() {
		createParams := vm_folder.NewCreateVMFolderParams()
		createParams.RequestBody = []*models.VMFolderCreationParams{
			{
				ClusterID: cluster.ID,
				Name:      String(fmt.Sprintf("tower-go-sdk-vm-folder-%d", time.Now().Unix())),
			},
		}
		createRes, err := Client.VMFolder.CreateVMFolder(createParams)
		Expect(err).To(BeNil())
		Expect(createRes).ToNot(BeNil())

		updateParams := vm_folder.NewUpdateVMFolderParams()
		updateParams.RequestBody = &models.VMFolderUpdationParams{
			Where: &models.VMFolderWhereInput{
				ID: createRes.Payload[0].Data.ID,
			},
			Data: &models.VMFolderUpdationParamsData{
				Name: String(fmt.Sprintf("tower-go-sdk-vm-folder-updated-%d", time.Now().Unix())),
			},
		}
		updateRes, err := Client.VMFolder.UpdateVMFolder(updateParams)
		Expect(err).To(BeNil())
		Expect(updateRes).ToNot(BeNil())

		deleteParam := vm_folder.NewDeleteVMFolderParams()
		deleteParam.RequestBody = &models.VMFolderDeletionParams{
			Where: &models.VMFolderWhereInput{
				ID: createRes.Payload[0].Data.ID,
			},
		}
		deleteRes, err := Client.VMFolder.DeleteVMFolder(deleteParam)
		Expect(err).To(BeNil())
		Expect(deleteRes).ToNot(BeNil())
	})

	It("should move in and out vm from vm folder", func() {
		folderId, disposeFolder := fixture.CreateVmFolder(Client, cluster.ID)
		vmId, disposeVm := fixture.CreateVm(Client, cluster.ID, vlan.ID, false)
		defer func() {
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				disposeFolder()
			}()
			wg.Add(1)
			go func() {
				defer wg.Done()
				disposeVm()
			}()
		}()
		addParams := vm.NewAddVMToFolderParams()
		addParams.RequestBody = &models.VMAddFolderParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
			Data: &models.VMAddFolderParamsData{
				FolderID: folderId,
			},
		}
		addRes, err := Client.VM.AddVMToFolder(addParams)
		Expect(err).To(BeNil())
		Expect(addRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, addRes.Payload[0].TaskID)
		Expect(err).To(BeNil())
		removeParams := vm.NewRemoveVMToFolderParams()
		removeParams.RequestBody = &models.VMOperateParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
		}
		removeRes, err := Client.VM.RemoveVMToFolder(removeParams)
		Expect(err).To(BeNil())
		Expect(removeRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, removeRes.Payload[0].TaskID)
		Expect(err).To(BeNil())
	})
})
