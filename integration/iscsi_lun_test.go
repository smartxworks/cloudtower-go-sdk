package integration

import (
	"fmt"
	"time"

	"github.com/smartxworks/cloudtower-go-sdk/v2/client/iscsi_lun"
	"github.com/smartxworks/cloudtower-go-sdk/v2/integration/fixture"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	taskutil "github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/openlyinc/pointy"
)

var _ = Describe("Iscsi lun api", Ordered, func() {
	var cluster *models.Cluster = nil
	BeforeAll(func() {
		cluster = fixture.GetDefaultCluster(Client, "xiaojun-nested-X20211110124941")
	})
	It("should get iscsi lun", func() {
		params := iscsi_lun.NewGetIscsiLunsParams()
		params.RequestBody = &models.GetIscsiLunsRequestBody{}
		res, err := Client.IscsiLun.GetIscsiLuns(params)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		connectionParams := iscsi_lun.NewGetIscsiLunsConnectionParams()
		connectionParams.RequestBody = &models.GetIscsiLunsConnectionRequestBody{}
		connectionRes, err := Client.IscsiLun.GetIscsiLunsConnection(connectionParams)
		Expect(err).To(BeNil())
		Expect(connectionRes).ToNot(BeNil())
	})

	It("should create, update and delete iscsi lun", func() {
		targetId, dispose := fixture.CreateIscsiTarget(Client, cluster.ID)
		defer dispose()
		createParams := iscsi_lun.NewCreateIscsiLunParams()
		createParams.RequestBody = []*models.IscsiLunCreationParams{
			{
				IscsiTargetID: targetId,
				Name:          String(fmt.Sprintf("tower-go-sdk-iscsi-lun-%d", time.Now().Unix())),
				ReplicaNum:    Int32(2),
				AssignedSize:  Int64(30 * 1024 * 1024 * 1024),
			},
		}
		createRes, err := Client.IscsiLun.CreateIscsiLun(createParams)
		Expect(err).To(BeNil())
		Expect(createRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, createRes.Payload[0].TaskID)
		Expect(err).To(BeNil())

		updateParams := iscsi_lun.NewUpdateIscsiLunParams()
		updateParams.RequestBody = &models.IscsiLunUpdationParams{
			Where: &models.IscsiLunWhereInput{
				ID: createRes.Payload[0].Data.ID,
			},
			Data: &models.IscsiLunUpdationParamsData{
				IscsiLunCommonParams: models.IscsiLunCommonParams{
					Iops: Int64(1024),
				},
			},
		}
		updateRes, err := Client.IscsiLun.UpdateIscsiLun(updateParams)
		Expect(err).To(BeNil())
		Expect(updateRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, updateRes.Payload[0].TaskID)
		Expect(err).To(BeNil())

		deleteParams := iscsi_lun.NewDeleteIscsiLunParams()
		deleteParams.RequestBody = &models.IscsiLunDeletionParams{
			Where: &models.IscsiLunWhereInput{
				ID: createRes.Payload[0].Data.ID,
			},
		}
		deleteRes, err := Client.IscsiLun.DeleteIscsiLun(deleteParams)
		Expect(err).To(BeNil())
		Expect(deleteRes).ToNot(BeNil())
		err = taskutil.WaitTask(Client, deleteRes.Payload[0].TaskID)
		Expect(err).To(BeNil())
	})
})
