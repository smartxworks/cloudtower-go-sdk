package fixture

import (
	"context"
	"fmt"
	"sync"
	"time"

	. "go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/iscsi_lun"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/iscsi_lun_snapshot"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/iscsi_target"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	taskutil "github.com/smartxworks/cloudtower-go-sdk/v2/utils"
)

func CreateIscsiTarget(client *apiclient.Cloudtower, clusterId *string) (*string, func()) {
	createParams := iscsi_target.NewCreateIscsiTargetParams()
	var res *iscsi_target.CreateIscsiTargetOK = nil
	var dispose = func() {
		if res == nil {
			return
		}
		deleteParams := iscsi_target.NewDeleteIscsiTargetParams()
		deleteParams.RequestBody = &models.IscsiTargetDeletionParams{
			Where: &models.IscsiTargetWhereInput{
				ID: res.Payload[0].Data.ID,
			},
		}
		res, err := client.IscsiTarget.DeleteIscsiTarget(deleteParams)
		if err != nil {
			panic(err.Error())
		}
		taskutil.WaitTask(context.TODO(), client, res.Payload[0].TaskID, 5*time.Minute)
	}
	defer func() {
		if err := recover(); err != nil {
			dispose()
		}
	}()
	createParams.RequestBody = []*models.IscsiTargetCreationParams{
		{
			ClusterID:     clusterId,
			ThinProvision: Bool(true),
			ReplicaNum:    Int32(2),
			StripeNum:     Int32(4),
			StripeSize:    Int64(256 * 1024),
			Name:          String(fmt.Sprintf("tower-go-sdk-iscsi-target-%d", time.Now().Unix())),
		},
	}
	res, err := client.IscsiTarget.CreateIscsiTarget(createParams)
	if err != nil {
		panic(err.Error())
	}
	taskutil.WaitTask(context.TODO(), client, res.Payload[0].TaskID, 5*time.Minute)
	return res.Payload[0].Data.ID, dispose
}

func CreateIscsiLun(client *apiclient.Cloudtower, clusterId *string) (*string, *string, func()) {
	targetId, disposeTarget := CreateIscsiTarget(client, clusterId)
	createParams := iscsi_lun.NewCreateIscsiLunParams()
	var res *iscsi_lun.CreateIscsiLunOK = nil
	var dispose = func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			disposeTarget()
		}()
		if res != nil {
			wg.Add(1)
			go func() {
				defer wg.Done()
				deleteParams := iscsi_lun.NewDeleteIscsiLunParams()
				deleteParams.RequestBody = &models.IscsiLunDeletionParams{
					Where: &models.IscsiLunWhereInput{
						ID: res.Payload[0].Data.ID,
					},
				}
				res, err := client.IscsiLun.DeleteIscsiLun(deleteParams)
				if err != nil {
					panic(err.Error())
				}
				taskutil.WaitTask(context.TODO(), client, res.Payload[0].TaskID, 5*time.Minute)
			}()
		}
		wg.Wait()
	}
	defer func() {
		if err := recover(); err != nil {
			dispose()
		}
	}()
	createParams.RequestBody = []*models.IscsiLunCreationParams{
		{
			IscsiTargetID: targetId,
			Name:          String(fmt.Sprintf("tower-go-sdk-iscsi-lun-%d", time.Now().Unix())),
			ReplicaNum:    Int32(2),
			AssignedSize:  Int64(30 * 1024 * 1024 * 1024),
		},
	}
	res, err := client.IscsiLun.CreateIscsiLun(createParams)
	if err != nil {
		panic(err.Error())
	}
	taskutil.WaitTask(context.TODO(), client, res.Payload[0].TaskID, 5*time.Minute)
	return res.Payload[0].Data.ID, targetId, dispose
}

func CreateIscsiLunSnapshot(client *apiclient.Cloudtower, clusterId *string) (*string, *string, *string, func()) {
	lunId, targetId, disposeLun := CreateIscsiLun(client, clusterId)
	var res *iscsi_lun_snapshot.CreateIscsiLunSnapshotOK = nil
	var dispose = func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			disposeLun()
		}()
		if res != nil {
			wg.Add(1)
			go func() {
				defer wg.Done()
				deleteParams := iscsi_lun_snapshot.NewDeleteIscsiLunSnapshotParams()
				deleteParams.RequestBody = &models.IscsiLunSnapshotDeletionParams{
					Where: &models.IscsiLunSnapshotWhereInput{
						ID: res.Payload[0].Data.ID,
					},
				}
				res, err := client.IscsiLunSnapshot.DeleteIscsiLunSnapshot(deleteParams)
				if err != nil {
					panic(err.Error())
				}
				taskutil.WaitTask(context.TODO(), client, res.Payload[0].TaskID, 5*time.Minute)
			}()
		}
		wg.Wait()
	}
	defer func() {
		if err := recover(); err != nil {
			dispose()
		}
	}()
	createParams := iscsi_lun_snapshot.NewCreateIscsiLunSnapshotParams()
	createParams.RequestBody = []*models.IscsiLunSnapshotCreationParams{
		{
			IscsiLunID:    lunId,
			IscsiTargetID: targetId,
			Name:          String(fmt.Sprintf("tower-go-sdk-iscsi-lun-snapshot-%d", time.Now().Unix())),
		},
	}
	res, err := client.IscsiLunSnapshot.CreateIscsiLunSnapshot(createParams)
	if err != nil {
		panic(err.Error())
	}
	return res.Payload[0].Data.ID, lunId, targetId, dispose
}
