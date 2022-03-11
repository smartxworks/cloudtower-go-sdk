package fixture

import (
	"fmt"
	"sync"
	"time"

	. "github.com/openlyinc/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/client"
	"github.com/smartxworks/cloudtower-go-sdk/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/client/vm_folder"
	"github.com/smartxworks/cloudtower-go-sdk/client/vm_snapshot"
	"github.com/smartxworks/cloudtower-go-sdk/client/vm_template"
	"github.com/smartxworks/cloudtower-go-sdk/models"
	taskutil "github.com/smartxworks/cloudtower-go-sdk/utils"
)

func CreateVm(client *apiclient.Cloudtower, clusterId *string, vlanId *string, isRunning bool) (*string, func()) {
	createParams := vm.NewCreateVMParams()
	var res *vm.CreateVMOK = nil
	var dispose = func() {
		if res == nil {
			return
		}
		deleteParams := vm.NewDeleteVMParams()
		deleteParams.RequestBody = &models.VMOperateParams{
			Where: &models.VMWhereInput{
				ID: res.Payload[0].Data.ID,
			},
		}
		res, err := client.VM.DeleteVM(deleteParams)
		if err != nil {
			panic(err.Error())
		}
		taskutil.WaitTask(client, res.Payload[0].TaskID)
	}
	defer func() {
		if err := recover(); err != nil {
			dispose()
		}
	}()
	var status *models.VMStatus = nil
	if isRunning {
		status = models.VMStatusRUNNING.Pointer()
	} else {
		status = models.VMStatusSTOPPED.Pointer()
	}
	createParams.RequestBody = []*models.VMCreationParams{
		{
			ClusterID:  clusterId,
			Name:       String(fmt.Sprintf("tower-go-sdk-test-create-vm-%d", time.Now().Unix())),
			Ha:         Bool(true),
			CPUCores:   Int32(1),
			CPUSockets: Int32(1),
			Memory:     Float64(1 * 1024 * 1024 * 1024),
			Vcpu:       Int32(1),
			Status:     status,
			Firmware:   models.VMFirmwareBIOS.Pointer(),
			VMNics: []*models.VMNicParams{
				{
					ConnectVlanID: vlanId,
				},
			},
			VMDisks: &models.VMDiskParams{
				MountCdRoms: []*models.VMCdRomParams{
					{
						Index: Int32(1),
						Boot:  Int32(1),
					},
				},
			},
		},
	}
	res, err := client.VM.CreateVM(createParams)
	if err != nil {
		panic(err.Error())
	}
	taskutil.WaitTask(client, res.Payload[0].TaskID)
	return res.Payload[0].Data.ID, dispose
}

func CreateVmFolder(client *apiclient.Cloudtower, clusterId *string) (*string, func()) {
	createParams := vm_folder.NewCreateVMFolderParams()
	var res *vm_folder.CreateVMFolderOK = nil
	var dispose = func() {
		deleteParams := vm_folder.NewDeleteVMFolderParams()
		deleteParams.RequestBody = &models.VMFolderDeletionParams{
			Where: &models.VMFolderWhereInput{
				ID: res.Payload[0].Data.ID,
			},
		}
		res, err := client.VMFolder.DeleteVMFolder(deleteParams)
		if err != nil {
			panic(err.Error())
		}
		taskutil.WaitTask(client, res.Payload[0].TaskID)
	}
	defer func() {
		if err := recover(); err != nil {
			dispose()
		}
	}()
	createParams.RequestBody = []*models.VMFolderCreationParams{
		{
			ClusterID: clusterId,
			Name:      String(fmt.Sprintf("tower-go-sdk-vm-folder-%d", time.Now().Unix())),
		},
	}
	res, err := client.VMFolder.CreateVMFolder(createParams)
	if err != nil {
		panic(err.Error())
	}
	taskutil.WaitTask(client, res.Payload[0].TaskID)
	return res.Payload[0].Data.ID, dispose
}

func CreateSnapshot(client *apiclient.Cloudtower, clusterId *string, vlanId *string) (*string, *string, func()) {
	vmId, disposeVm := CreateVm(client, clusterId, vlanId, false)
	createParams := vm_snapshot.NewCreateVMSnapshotParams()
	var res *vm_snapshot.CreateVMSnapshotOK = nil
	var dispose = func() {
		var wg sync.WaitGroup
		if res != nil {
			wg.Add(1)
			go func() {
				defer wg.Done()
				deleteParams := vm_snapshot.NewDeleteVMSnapshotParams()
				deleteParams.RequestBody = &models.VMSnapshotDeletionParams{
					Where: &models.VMSnapshotWhereInput{
						ID: res.Payload[0].Data.ID,
					},
				}
				res, err := client.VMSnapshot.DeleteVMSnapshot(deleteParams)
				if err != nil {
					panic(err.Error())
				}
				taskutil.WaitTask(client, res.Payload[0].TaskID)
			}()
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			disposeVm()
		}()
		wg.Wait()
	}
	defer func() {
		if err := recover(); err != nil {
			dispose()
		}
	}()
	createParams.RequestBody = &models.VMSnapshotCreationParams{
		Data: []*models.VMSnapshotCreationParamsDataItems0{
			{
				VMID: vmId,
				Name: String(fmt.Sprintf("tower-go-sdk-vm-snapshot-%d", time.Now().Unix())),
			},
		},
	}
	res, err := client.VMSnapshot.CreateVMSnapshot(createParams)
	if err != nil {
		panic(err.Error())
	}
	taskutil.WaitTask(client, res.Payload[0].TaskID)
	return res.Payload[0].Data.ID, vmId, dispose
}

func CreateTemplate(client *apiclient.Cloudtower, clusterId *string, vlanId *string) (*string, func()) {
	vmId, disposeVm := CreateVm(client, clusterId, vlanId, false)
	createParams := vm_template.NewConvertVMTemplateFromVMParams()
	var res *vm_template.ConvertVMTemplateFromVMOK = nil
	var dispose = func() {
		if res != nil {
			deleteParams := vm_template.NewDeleteVMTemplateParams()
			deleteParams.RequestBody = &models.VMTemplateDeletionParams{
				Where: &models.VMTemplateWhereInput{
					ID: res.Payload[0].Data.ID,
				},
			}
			res, err := client.VMTemplate.DeleteVMTemplate(deleteParams)
			if err != nil {
				panic(err.Error())
			}
			taskutil.WaitTask(client, res.Payload[0].TaskID)
		} else {
			disposeVm()
		}
	}
	defer func() {
		if err := recover(); err != nil {
			dispose()
		}
	}()
	createParams.RequestBody = []*models.VMTemplateCreationParams{
		{
			Name:               String(fmt.Sprintf("tower-go-sdk-vm-template-clone-%d", time.Now().Unix())),
			VMID:               vmId,
			ClusterID:          clusterId,
			CloudInitSupported: Bool(false),
		},
	}
	res, err := client.VMTemplate.ConvertVMTemplateFromVM(createParams)
	if err != nil {
		panic(err.Error())
	}
	taskutil.WaitTask(client, res.Payload[0].TaskID)
	return res.Payload[0].Data.ID, dispose
}
