package integration

import (
	"context"
	"fmt"
	"time"

	"github.com/smartxworks/cloudtower-go-sdk/v2/client/iscsi_target"
	"github.com/smartxworks/cloudtower-go-sdk/v2/integration/fixture"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	taskutil "github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "go.openly.dev/pointy"
)

var _ = Describe("Iscsi target api", Ordered, func() {
	var cluster *models.Cluster = nil
	BeforeAll(func() {
		cluster = fixture.GetDefaultCluster(Client, "xiaojun-502-ELF-nestedX20220426111859")
	})
	It("should get iscsi target", func() {
		params := iscsi_target.NewGetIscsiTargetsParams()
		params.RequestBody = &models.GetIscsiTargetsRequestBody{}
		res, err := Client.IscsiTarget.GetIscsiTargets(params)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		connectionParams := iscsi_target.NewGetIscsiTargetsConnectionParams()
		connectionParams.RequestBody = &models.GetIscsiTargetsConnectionRequestBody{}
		connectionRes, err := Client.IscsiTarget.GetIscsiTargetsConnection(connectionParams)
		Expect(err).To(BeNil())
		Expect(connectionRes).ToNot(BeNil())
	})

	It("should create, update and delete iscsi target", func() {
		createParams := iscsi_target.NewCreateIscsiTargetParams()
		createParams.RequestBody = []*models.IscsiTargetCreationParams{
			{
				ClusterID:     cluster.ID,
				ThinProvision: Bool(true),
				ReplicaNum:    Int32(2),
				StripeNum:     Int32(4),
				StripeSize:    Int64(256 * 1024),
				Name:          String(fmt.Sprintf("tower-go-sdk-iscsi-target-%d", time.Now().Unix())),
			},
		}
		createRes, err := Client.IscsiTarget.CreateIscsiTarget(createParams)
		Expect(err).To(BeNil())
		Expect(createRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, createRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())

		updateParams := iscsi_target.NewUpdateIscsiTargetParams()
		updateParams.RequestBody = &models.IscsiTargetUpdationParams{
			Where: &models.IscsiTargetWhereInput{
				ID: createRes.Payload[0].Data.ID,
			},
			Data: &models.IscsiTargetCommonParams{
				Iops: Int64(1024),
			},
		}
		updateRes, err := Client.IscsiTarget.UpdateIscsiTarget(updateParams)
		Expect(err).To(BeNil())
		Expect(updateRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, updateRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())

		deleteParams := iscsi_target.NewDeleteIscsiTargetParams()
		deleteParams.RequestBody = &models.IscsiTargetDeletionParams{
			Where: &models.IscsiTargetWhereInput{
				ID: createRes.Payload[0].Data.ID,
			},
		}
		deleteRes, err := Client.IscsiTarget.DeleteIscsiTarget(deleteParams)
		Expect(err).To(BeNil())
		Expect(deleteRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, deleteRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
	})
})
