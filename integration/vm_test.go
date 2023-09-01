package integration

import (
	"context"
	"fmt"
	"sync"
	"time"

	. "go.openly.dev/pointy"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_disk"
	"github.com/smartxworks/cloudtower-go-sdk/v2/integration/fixture"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	taskutil "github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vm Api", Ordered, func() {
	var cluster *models.Cluster = nil
	var vlan *models.Vlan = nil
	BeforeAll(func() {
		cluster = fixture.GetDefaultCluster(Client, "xiaojun-502-ELF-nestedX20220426111859")
		vlan = fixture.GetDefaultVlan(Client, cluster.ID)
	})
	It("should get vms", func() {
		params := vm.NewGetVmsParams()
		params.RequestBody = &models.GetVmsRequestBody{}
		connectionParams := vm.NewGetVmsConnectionParams()
		connectionParams.RequestBody = &models.GetVmsConnectionRequestBody{}
		res, err := Client.VM.GetVms(params)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		connectionRes, err := Client.VM.GetVmsConnection(connectionParams)
		Expect(err).To(BeNil())
		Expect(connectionRes).ToNot(BeNil())
	})
	It("should get vm connections", func() {
		params := vm.NewGetVmsConnectionParams()
		params.RequestBody = &models.GetVmsConnectionRequestBody{}
		res, err := Client.VM.GetVmsConnection(params)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
	})
	It("should create, update and delete vm", func() {
		createParams := vm.NewCreateVMParams()
		createParams.RequestBody = []*models.VMCreationParams{
			{
				ClusterID:  cluster.ID,
				Name:       String(fmt.Sprintf("tower-go-sdk-test-create-vm-%d", time.Now().Unix())),
				Ha:         Bool(true),
				CPUCores:   Int32(1),
				CPUSockets: Int32(1),
				Memory:     Int64(1 * 1024 * 1024 * 1024),
				Vcpu:       Int32(1),
				Status:     models.VMStatusSTOPPED.Pointer(),
				Firmware:   models.VMFirmwareBIOS.Pointer(),
				VMNics: []*models.VMNicParams{
					{
						ConnectVlanID: vlan.ID,
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
		createRes, err := Client.VM.CreateVM(createParams)
		Expect(err).To(BeNil())
		Expect(createRes).ToNot(BeNil())
		Expect(createRes.Payload).To(HaveLen(1))
		_vm := createRes.Payload[0]
		err = taskutil.WaitTask(context.TODO(), Client, _vm.TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
		updateParams := vm.NewUpdateVMParams()
		updateParams.RequestBody = &models.VMUpdateParams{
			Where: &models.VMWhereInput{
				ID: _vm.Data.ID,
			},
			Data: &models.VMUpdateParamsData{
				Description: String("tower-go-sdk-updated-description"),
			},
		}
		updatedRes, err := Client.VM.UpdateVM(updateParams)
		Expect(err).To(BeNil())
		Expect(updatedRes).ToNot(BeNil())
		Expect(updatedRes.Payload).To(HaveLen(1))
		err = taskutil.WaitTask(context.TODO(), Client, updatedRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
		deleteParams := vm.NewDeleteVMParams()
		deleteParams.RequestBody = &models.VMDeleteParams{
			Where: &models.VMWhereInput{
				ID: _vm.Data.ID,
			},
		}
		deleteRes, err := Client.VM.DeleteVM(deleteParams)
		Expect(err).To(BeNil())
		Expect(deleteRes).ToNot(BeNil())
		Expect(deleteRes.Payload).To(HaveLen(1))
		err = taskutil.WaitTask(context.TODO(), Client, deleteRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
	})
	It("should clone vm", func() {
		vmId, dispose := fixture.CreateVm(Client, cluster.ID, vlan.ID, false)
		cloneParams := vm.NewCloneVMParams()
		cloneParams.RequestBody = []*models.VMCloneParams{
			{
				SrcVMID:   vmId,
				ClusterID: cluster.ID,
				Name:      String(fmt.Sprintf("tower-go-sdk-cloned-vm-%d", time.Now().Unix())),
			},
		}
		res, err := Client.VM.CloneVM(cloneParams)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		taskutil.WaitTask(context.TODO(), Client, res.Payload[0].TaskID, 5*time.Minute)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			dispose()
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			deleteParams := vm.NewDeleteVMParams()
			deleteParams.RequestBody = &models.VMDeleteParams{
				Where: &models.VMWhereInput{
					ID: res.Payload[0].Data.ID,
				},
			}
			res, _ := Client.VM.DeleteVM(deleteParams)
			err = taskutil.WaitTask(context.TODO(), Client, res.Payload[0].TaskID, 5*time.Minute)
			Expect(err).To(BeNil())
		}()
		wg.Wait()
	})
	It("should start vm", func() {
		vmId, dispose := fixture.CreateVm(Client, cluster.ID, vlan.ID, false)
		startParams := vm.NewStartVMParams()
		startParams.RequestBody = &models.VMStartParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
		}
		res, err := Client.VM.StartVM(startParams)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, res.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
		dispose()
	})

	It("should migrate vm", func() {
		vmId, _ := fixture.CreateVm(Client, cluster.ID, vlan.ID, false)
		migrateParams := vm.NewMigrateVMParams()
		migrateParams.RequestBody = &models.VMMigrateParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
		}
		res, err := Client.VM.MigrateVM(migrateParams)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, res.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
		deleteParams := vm.NewDeleteVMParams()
		deleteParams.RequestBody = &models.VMDeleteParams{
			Where: &models.VMWhereInput{
				ID: res.Payload[0].Data.ID,
			},
		}
		deleteRes, err := Client.VM.DeleteVM(deleteParams)
		Expect(err).To(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, deleteRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
	})
	XIt("should restart vm", func() {
		Skip("we can't restart a vm without installing os")
		vmId, dispose := fixture.CreateVm(Client, cluster.ID, vlan.ID, true)
		defer dispose()
		restartParams := vm.NewRestartVMParams()
		restartParams.RequestBody = &models.VMOperateParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
		}
		res, err := Client.VM.RestartVM(restartParams)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, res.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
	})
	It("should force restart vm", func() {
		vmId, dispose := fixture.CreateVm(Client, cluster.ID, vlan.ID, true)
		defer dispose()
		restartParams := vm.NewForceRestartVMParams()
		restartParams.RequestBody = &models.VMOperateParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
		}
		res, err := Client.VM.ForceRestartVM(restartParams)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, res.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
	})
	It("should suspend and resume vm", func() {
		vmId, dispose := fixture.CreateVm(Client, cluster.ID, vlan.ID, true)
		defer dispose()
		suspendParams := vm.NewSuspendVMParams()
		suspendParams.RequestBody = &models.VMOperateParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
		}
		suspendRes, err := Client.VM.SuspendVM(suspendParams)
		Expect(err).To(BeNil())
		Expect(suspendRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, suspendRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
		resumeParams := vm.NewResumeVMParams()
		resumeParams.RequestBody = &models.VMOperateParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
		}
		resumeRes, err := Client.VM.ResumeVM(resumeParams)
		Expect(err).To(BeNil())
		Expect(resumeRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, resumeRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
	})

	XIt("should shutdown vm", func() {
		Skip("we can't shutdown a vm without installing os")
		vmId, dispose := fixture.CreateVm(Client, cluster.ID, vlan.ID, true)
		defer dispose()
		shutdownParams := vm.NewShutDownVMParams()
		shutdownParams.RequestBody = &models.VMOperateParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
		}
		res, err := Client.VM.ShutDownVM(shutdownParams)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, res.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
	})
	It("should force shutdown vm", func() {
		vmId, dispose := fixture.CreateVm(Client, cluster.ID, vlan.ID, true)
		defer dispose()
		shutdownParams := vm.NewPoweroffVMParams()
		shutdownParams.RequestBody = &models.VMOperateParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
		}
		res, err := Client.VM.PoweroffVM(shutdownParams)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, res.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())

	})

	It("should add and remove vm cd-rom", func() {
		vmId, dispose := fixture.CreateVm(Client, cluster.ID, vlan.ID, false)
		defer dispose()
		addParams := vm.NewAddVMCdRomParams()
		addParams.RequestBody = &models.VMAddCdRomParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
			Data: &models.VMAddCdRomParamsData{
				VMCdRoms: []*models.VMCdRomParams{
					{
						Boot:  Int32(2),
						Index: Int32(2),
					},
				},
			},
		}
		addRes, err := Client.VM.AddVMCdRom(addParams)
		Expect(err).To(BeNil())
		Expect(addRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, addRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
		getParams := vm_disk.NewGetVMDisksParams()
		getParams.RequestBody = &models.GetVMDisksRequestBody{
			Where: &models.VMDiskWhereInput{
				VM: &models.VMWhereInput{
					ID: vmId,
				},
				Boot: Int32(2),
			},
		}
		diskRes, err := Client.VMDisk.GetVMDisks(getParams)
		Expect(err).To(BeNil())
		deleteParams := vm.NewRemoveVMCdRomParams()
		deleteParams.RequestBody = &models.VMRemoveCdRomParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
			Data: &models.VMRemoveCdRomParamsData{
				CdRomIds: []string{*diskRes.Payload[0].ID},
			},
		}
		deleteRes, err := Client.VM.RemoveVMCdRom(deleteParams)
		Expect(err).To(BeNil())
		Expect(deleteRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, deleteRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
	})

	It("should add, update and remove vm disk", func() {
		vmId, dispose := fixture.CreateVm(Client, cluster.ID, vlan.ID, false)
		defer dispose()
		addParams := vm.NewAddVMDiskParams()
		addParams.RequestBody = &models.VMAddDiskParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
			Data: &models.VMAddDiskParamsData{
				VMDisks: &models.VMAddDiskParamsDataVMDisks{
					MountNewCreateDisks: []*models.MountNewCreateDisksParams{
						{
							Bus:  models.BusIDE.Pointer(),
							Boot: Int32(2),
							VMVolume: &models.MountNewCreateDisksParamsVMVolume{
								ElfStoragePolicy: models.VMVolumeElfStoragePolicyTypeREPLICA2THINPROVISION.Pointer(),
								Name:             String(fmt.Sprintf("tower-go-sdk-added-vm-volume-%d", time.Now().Unix())),
								Size:             Int64(1 * 1024 * 1024 * 1024),
							},
						},
					},
				},
			},
		}
		addRes, err := Client.VM.AddVMDisk(addParams)
		Expect(err).To(BeNil())
		Expect(addRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, addRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
		updateParams := vm.NewUpdateVMDiskParams()
		updateParams.RequestBody = &models.VMUpdateDiskParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
			Data: &models.VMUpdateDiskParamsData{
				Bus:      models.BusSCSI.Pointer(),
				VMDiskID: addRes.Payload[0].Data.VMDisks[0].ID,
			},
		}
		updateRes, err := Client.VM.UpdateVMDisk(updateParams)
		Expect(err).To(BeNil())
		Expect(updateRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, updateRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
		removeParams := vm.NewRemoveVMDiskParams()
		removeParams.RequestBody = &models.VMRemoveDiskParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
			Data: &models.VMRemoveDiskParamsData{
				DiskIds: []string{
					*addRes.Payload[0].Data.VMDisks[0].ID,
				},
			},
		}
		removeRes, err := Client.VM.RemoveVMDisk(removeParams)
		Expect(err).To(BeNil())
		Expect(removeRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, removeRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
	})

	It("should add, update and remove vm nic", func() {
		vmId, dispose := fixture.CreateVm(Client, cluster.ID, vlan.ID, false)
		defer dispose()
		addParams := vm.NewAddVMNicParams()
		addParams.RequestBody = &models.VMAddNicParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
			Data: &models.VMAddNicParamsData{
				VMNics: []*models.VMNicParams{
					{
						ConnectVlanID: vlan.ID,
					},
				},
			},
		}
		addRes, err := Client.VM.AddVMNic(addParams)
		Expect(err).To(BeNil())
		Expect(addRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, addRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
		updateParams := vm.NewUpdateVMNicParams()
		updateParams.RequestBody = &models.VMUpdateNicParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
			Data: &models.VMUpdateNicParamsData{
				NicIndex: Int32(0),
				Enabled:  Bool(false),
			},
		}
		updateRes, err := Client.VM.UpdateVMNic(updateParams)
		Expect(err).To(BeNil())
		Expect(updateRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, updateRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
		removeParams := vm.NewRemoveVMNicParams()
		removeParams.RequestBody = &models.VMRemoveNicParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
			Data: &models.VMRemoveNicParamsData{
				NicIndex: []int32{0},
			},
		}
		removeRes, err := Client.VM.RemoveVMNic(removeParams)
		Expect(err).To(BeNil())
		Expect(removeRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, removeRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
	})
	It("should move to and recover from recycle bin", func() {
		vmId, dispose := fixture.CreateVm(Client, cluster.ID, vlan.ID, false)
		disposeSetting := fixture.OpenRecycleBinIfClosed(Client, cluster.ID)
		defer func() {
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				dispose()
			}()
			wg.Add(1)
			go func() {
				defer wg.Done()
				disposeSetting()
			}()
			wg.Wait()
		}()
		moveParam := vm.NewMoveVMToRecycleBinParams()
		moveParam.RequestBody = &models.VMOperateParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
		}
		moveRes, err := Client.VM.MoveVMToRecycleBin(moveParam)
		Expect(err).To(BeNil())
		Expect(moveRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, moveRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
		recoverParams := vm.NewRecoverVMFromRecycleBinParams()
		recoverParams.RequestBody = &models.VMOperateParams{
			Where: &models.VMWhereInput{
				ID: vmId,
			},
		}
		recoverRes, err := Client.VM.RecoverVMFromRecycleBin(recoverParams)
		Expect(err).To(BeNil())
		Expect(recoverRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, recoverRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
	})
})
