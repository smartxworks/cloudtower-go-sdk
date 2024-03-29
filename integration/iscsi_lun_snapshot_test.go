package integration

import (
	"context"
	"fmt"
	"sync"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/openlyinc/pointy"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/iscsi_lun"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/iscsi_lun_snapshot"
	"github.com/smartxworks/cloudtower-go-sdk/v2/integration/fixture"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	taskutil "github.com/smartxworks/cloudtower-go-sdk/v2/utils"
)

var _ = Describe("Iscsi lun snapshot api", Ordered, func() {
	var cluster *models.Cluster = nil
	BeforeAll(func() {
		cluster = fixture.GetDefaultCluster(Client, "xiaojun-502-ELF-nestedX20220426111859")
	})
	It("should get iscsi lun snapshot", func() {
		params := iscsi_lun_snapshot.NewGetIscsiLunSnapshotsParams()
		params.RequestBody = &models.GetIscsiLunSnapshotsRequestBody{}
		res, err := Client.IscsiLunSnapshot.GetIscsiLunSnapshots(params)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
		connectionParams := iscsi_lun_snapshot.NewGetIscsiLunSnapshotsConnectionParams()
		connectionParams.RequestBody = &models.GetIscsiLunSnapshotsConnectionRequestBody{}
		connectionRes, err := Client.IscsiLunSnapshot.GetIscsiLunSnapshotsConnection(connectionParams)
		Expect(err).To(BeNil())
		Expect(connectionRes).ToNot(BeNil())
	})

	It("should create and delete iscsi lun snapshot", func() {
		lunId, targetId, dispose := fixture.CreateIscsiLun(Client, cluster.ID)
		defer dispose()
		createParams := iscsi_lun_snapshot.NewCreateIscsiLunSnapshotParams()
		createParams.RequestBody = []*models.IscsiLunSnapshotCreationParams{
			{
				IscsiLunID:    lunId,
				IscsiTargetID: targetId,
				Name:          String(fmt.Sprintf("tower-go-sdk-iscsi-lun-snapshot-%d", time.Now().Unix())),
			},
		}
		createRes, err := Client.IscsiLunSnapshot.CreateIscsiLunSnapshot(createParams)
		Expect(err).To(BeNil())
		Expect(createRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, createRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())

		deleteParams := iscsi_lun_snapshot.NewDeleteIscsiLunSnapshotParams()
		deleteParams.RequestBody = &models.IscsiLunSnapshotDeletionParams{
			Where: &models.IscsiLunSnapshotWhereInput{
				ID: createRes.Payload[0].Data.ID,
			},
		}
		deleteRes, err := Client.IscsiLunSnapshot.DeleteIscsiLunSnapshot(deleteParams)
		Expect(err).To(BeNil())
		Expect(deleteRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, deleteRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
	})

	It("should clone iscsi lun from snapshot", func() {
		snapshotId, _, targetId, dispose := fixture.CreateIscsiLunSnapshot(Client, cluster.ID)
		var cloneRes *iscsi_lun.CloneIscsiLunFromSnapshotOK = nil
		defer func() {
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				dispose()
			}()
			if cloneRes != nil {
				wg.Add(1)
				go func() {
					defer wg.Done()
					deleteParams := iscsi_lun_snapshot.NewDeleteIscsiLunSnapshotParams()
					deleteParams.RequestBody = &models.IscsiLunSnapshotDeletionParams{
						Where: &models.IscsiLunSnapshotWhereInput{
							ID: cloneRes.Payload[0].Data.ID,
						},
					}
					deleteRes, err := Client.IscsiLunSnapshot.DeleteIscsiLunSnapshot(deleteParams)
					if err != nil {
						panic(err.Error())
					}
					taskutil.WaitTask(context.TODO(), Client, deleteRes.Payload[0].TaskID, 5*time.Minute)
				}()
			}
		}()
		cloneParams := iscsi_lun.NewCloneIscsiLunFromSnapshotParams()
		cloneParams.RequestBody = []*models.IscsiLunCloneParams{
			{
				SnapshotID:    snapshotId,
				IscsiTargetID: targetId,
			},
		}
		cloneRes, err := Client.IscsiLun.CloneIscsiLunFromSnapshot(cloneParams)
		Expect(err).To(BeNil())
		Expect(cloneRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, cloneRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
	})

	It("should rollback iscsi lun from snapshot", func() {
		snapshotId, lunId, _, dispose := fixture.CreateIscsiLunSnapshot(Client, cluster.ID)
		defer dispose()
		rollbackParams := iscsi_lun.NewRollbackIscsiLunFromSnapshotParams()
		rollbackParams.RequestBody = []*models.IscsiLunRollbackParams{
			{
				LunID:      lunId,
				SnapshotID: snapshotId,
			},
		}
		rollackRes, err := Client.IscsiLun.RollbackIscsiLunFromSnapshot(rollbackParams)
		Expect(err).To(BeNil())
		Expect(rollackRes).ToNot(BeNil())
		err = taskutil.WaitTask(context.TODO(), Client, rollackRes.Payload[0].TaskID, 5*time.Minute)
		Expect(err).To(BeNil())
	})
})
