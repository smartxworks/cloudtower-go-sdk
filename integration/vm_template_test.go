package integration

import (
	"fmt"
	"sync"
	"time"

	"github.com/Sczlog/cloudtower-go-sdk/client/vm"
	"github.com/Sczlog/cloudtower-go-sdk/client/vm_template"
	"github.com/Sczlog/cloudtower-go-sdk/integration/fixture"
	"github.com/Sczlog/cloudtower-go-sdk/models"
	taskutil "github.com/Sczlog/cloudtower-go-sdk/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/openlyinc/pointy"
)

var _ = Describe("Vm template api", Ordered, func() {
	var cluster *models.Cluster = nil
	var vlan *models.Vlan = nil
	BeforeAll(func() {
		cluster = fixture.GetDefaultCluster(Client, "xiaojun-nested-X20211110124941")
		vlan = fixture.GetDefaultVlan(Client, cluster.ID)
	})

	It("should get vm template", func() {
		params := vm_template.NewGetVMTemplatesParams()
		params.RequestBody = &models.GetVMTemplatesRequestBody{}
		connectionParams := vm_template.NewGetVMTemplatesConnectionParams()
		connectionParams.RequestBody = &models.GetVMTemplatesConnectionRequestBody{}
		res, err := Client.VMTemplate.GetVMTemplates(params)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		connectionRes, err := Client.VMTemplate.GetVMTemplatesConnection(connectionParams)
		Expect(err).To(BeNil())
		Expect(connectionRes).ToNot(BeNil())
	})

	It("should clone vm template from vm and delete", func() {
		vmId, disposeVm := fixture.CreateVm(Client, cluster.ID, vlan.ID, false)
		defer disposeVm()
		cloneParams := vm_template.NewCloneVMTemplateFromVMParams()
		cloneParams.RequestBody = []*models.VMTemplateCreationParams{
			{
				Name:               String(fmt.Sprintf("tower-go-sdk-vm-template-clone-%d", time.Now().Unix())),
				VMID:               vmId,
				ClusterID:          cluster.ID,
				CloudInitSupported: Bool(false),
			},
		}
		cloneRes, err := Client.VMTemplate.CloneVMTemplateFromVM(cloneParams)
		Expect(err).To(BeNil())
		Expect(cloneRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, cloneRes.Payload[0].TaskID)
		Expect(err).To(BeNil())

		deleteParams := vm_template.NewDeleteVMTemplateParams()
		deleteParams.RequestBody = &models.VMTemplateDeletionParams{
			Where: &models.VMTemplateWhereInput{
				ID: cloneRes.Payload[0].Data.ID,
			},
		}
		deleteRes, err := Client.VMTemplate.DeleteVMTemplate(deleteParams)
		Expect(err).To(BeNil())
		Expect(deleteRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, deleteRes.Payload[0].TaskID)
		Expect(err).To(BeNil())
	})

	It("should convert vm template from vm and delete", func() {
		vmId, dispose := fixture.CreateVm(Client, cluster.ID, vlan.ID, false)
		var cloneRes *vm_template.ConvertVMTemplateFromVMOK = nil
		defer func() {
			if err := recover(); cloneRes == nil && err != nil {
				dispose()
			}
		}()
		convertParams := vm_template.NewConvertVMTemplateFromVMParams()
		convertParams.RequestBody = []*models.VMTemplateCreationParams{
			{
				Name:               String(fmt.Sprintf("tower-go-sdk-vm-template-clone-%d", time.Now().Unix())),
				VMID:               vmId,
				ClusterID:          cluster.ID,
				CloudInitSupported: Bool(false),
			},
		}
		cloneRes, err := Client.VMTemplate.ConvertVMTemplateFromVM(convertParams)
		Expect(err).To(BeNil())
		Expect(cloneRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, cloneRes.Payload[0].TaskID)
		Expect(err).To(BeNil())

		deleteParams := vm_template.NewDeleteVMTemplateParams()
		deleteParams.RequestBody = &models.VMTemplateDeletionParams{
			Where: &models.VMTemplateWhereInput{
				ID: cloneRes.Payload[0].Data.ID,
			},
		}
		deleteRes, err := Client.VMTemplate.DeleteVMTemplate(deleteParams)
		Expect(err).To(BeNil())
		Expect(deleteRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, deleteRes.Payload[0].TaskID)
		Expect(err).To(BeNil())
	})

	It("should update vm template", func() {
		templateId, dispose := fixture.CreateTemplate(Client, cluster.ID, vlan.ID)
		defer dispose()
		updateParams := vm_template.NewUpdateVMTemplateParams()
		updateParams.RequestBody = &models.VMTemplateUpdationParams{
			Where: &models.VMTemplateWhereInput{
				ID: templateId,
			},
			Data: &models.VMTemplateUpdationParamsData{
				Description: String("Updated vm template"),
			},
		}
		updateRes, err := Client.VMTemplate.UpdateVMTemplate(updateParams)
		Expect(err).To(BeNil())
		Expect(updateRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, updateRes.Payload[0].TaskID)
		Expect(err).To(BeNil())
	})

	It("should convert vm template to vm", func() {
		templateId, dispose := fixture.CreateTemplate(Client, cluster.ID, vlan.ID)
		var convertRes *vm.ConvertVMTemplateToVMOK = nil
		defer func() {
			if err := recover(); err != nil && convertRes == nil {
				dispose()
			} else if convertRes != nil {
				deleteParams := vm.NewDeleteVMParams()
				deleteParams.RequestBody = &models.VMOperateParams{
					Where: &models.VMWhereInput{
						ID: convertRes.Payload[0].Data.ID,
					},
				}
				deleteRes, _ := Client.VM.DeleteVM(deleteParams)
				taskutil.WaitTask(Client, deleteRes.Payload[0].TaskID)
			}
		}()
		convertParams := vm.NewConvertVMTemplateToVMParams()
		convertParams.RequestBody = []*models.ConvertVMTemplateToVMParams{
			{
				ConvertedFromTemplateID: templateId,
				Name:                    String(fmt.Sprintf("tower-go-sdk-vm-convert-from-template-%d", time.Now().Unix())),
			},
		}
		convertRes, err := Client.VM.ConvertVMTemplateToVM(convertParams)
		Expect(err).To(BeNil())
		Expect(convertRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, convertRes.Payload[0].TaskID)
		Expect(err).To(BeNil())
	})

	It("should create vm from vm template", func() {
		templateId, dispose := fixture.CreateTemplate(Client, cluster.ID, vlan.ID)
		var createRes *vm.CreateVMFromTemplateOK = nil
		defer func() {
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				dispose()
			}()
			if createRes != nil {
				wg.Add(1)
				go func() {
					defer wg.Done()
					deleteParams := vm.NewDeleteVMParams()
					deleteParams.RequestBody = &models.VMOperateParams{
						Where: &models.VMWhereInput{
							ID: createRes.Payload[0].Data.ID,
						},
					}
					deleteRes, _ := Client.VM.DeleteVM(deleteParams)
					taskutil.WaitTask(Client, deleteRes.Payload[0].TaskID)
				}()
			}
			wg.Wait()
		}()
		createParams := vm.NewCreateVMFromTemplateParams()
		createParams.RequestBody = []*models.VMCreateVMFromTemplateParams{
			{
				TemplateID: templateId,
				Name:       String(fmt.Sprintf("tower-go-sdk-vm-create-from-template-%d", time.Now().Unix())),
				ClusterID:  cluster.ID,
				IsFullCopy: Bool(false),
			},
		}
		createRes, err := Client.VM.CreateVMFromTemplate(createParams)
		Expect(err).To(BeNil())
		Expect(createRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, createRes.Payload[0].TaskID)
		Expect(err).To(BeNil())
	})
})
