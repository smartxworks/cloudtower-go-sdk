package fixture

import (
	"sync"

	. "go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/cluster_settings"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/global_settings"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

func OpenRecycleBinIfClosed(client *apiclient.Cloudtower, clusterId *string) func() {
	return switchRecycleBin(client, clusterId, true)
}
func CloseRecycleBinIfOpen(client *apiclient.Cloudtower, clusterId *string) func() {
	return switchRecycleBin(client, clusterId, false)
}
func switchRecycleBin(client *apiclient.Cloudtower, clusterId *string, targetStatus bool) func() {
	var wg sync.WaitGroup
	var originGlobalSetting *models.GlobalSettings = nil
	var originClusterSetting *models.ClusterSettings = nil

	globalSettingParams := global_settings.NewGetGlobalSettingsesParams()
	globalSettingParams.RequestBody = &models.GetGlobalSettingsesRequestBody{
		First: Int32(1),
	}
	globalSettingRes, err := client.GlobalSettings.GetGlobalSettingses(globalSettingParams)
	if err != nil {
		panic(err.Error())
	}
	if *globalSettingRes.Payload[0].VMRecycleBin.Enabled != targetStatus {
		originGlobalSetting = globalSettingRes.Payload[0]
		wg.Add(1)
		go func() {
			defer wg.Done()
			updateParams := global_settings.NewUpdateGlobalRecycleBinSettingParams()
			updateParams.RequestBody = &models.GlobalRecycleBinUpdationParams{
				Enabled: Bool(targetStatus),
				Retain:  originGlobalSetting.VMRecycleBin.Retain,
			}
			_, err := client.GlobalSettings.UpdateGlobalRecycleBinSetting(updateParams)
			if err != nil {
				panic(err.Error())
			}
		}()
	}
	if clusterId != nil {
		clusterSettingParams := cluster_settings.NewGetClusterSettingsesParams()
		clusterSettingParams.RequestBody = &models.GetClusterSettingsesRequestBody{
			Where: &models.ClusterSettingsWhereInput{
				Cluster: &models.ClusterWhereInput{
					ID: clusterId,
				},
			},
		}
		clusterSettingRes, err := client.ClusterSettings.GetClusterSettingses(clusterSettingParams)
		if err != nil {
			panic(err.Error())
		}
		if len(clusterSettingRes.Payload) > 0 && *clusterSettingRes.Payload[0].VMRecycleBin.Enabled != targetStatus {
			originClusterSetting = clusterSettingRes.Payload[0]
			wg.Add(1)
			go func() {
				defer wg.Done()
				updateParams := global_settings.NewUpdateClusterRecycleBinSettingParams()
				updateParams.RequestBody = &models.ClusterRecycleBinUpdationParams{
					Where: &models.ClusterWhereInput{
						ID: clusterId,
					},
					Data: &models.ClusterRecycleBinUpdationParamsData{
						Enabled: Bool(targetStatus),
						Retain:  originClusterSetting.VMRecycleBin.Retain,
					},
				}
				_, err := client.GlobalSettings.UpdateClusterRecycleBinSetting(updateParams)
				if err != nil {
					panic(err.Error())
				}
			}()
		}
	}
	wg.Wait()
	return func() {
		var wg sync.WaitGroup
		if originGlobalSetting != nil {
			wg.Add(1)
			go func() {
				defer wg.Done()
				updateParams := global_settings.NewUpdateGlobalRecycleBinSettingParams()
				updateParams.RequestBody = &models.GlobalRecycleBinUpdationParams{
					Enabled: originGlobalSetting.VMRecycleBin.Enabled,
					Retain:  originGlobalSetting.VMRecycleBin.Retain,
				}
				_, err := client.GlobalSettings.UpdateGlobalRecycleBinSetting(updateParams)
				if err != nil {
					panic(err.Error())
				}
			}()
		}
		if originClusterSetting != nil {
			wg.Add(1)
			go func() {
				defer wg.Done()
				updateParams := global_settings.NewUpdateClusterRecycleBinSettingParams()
				updateParams.RequestBody = &models.ClusterRecycleBinUpdationParams{
					Where: &models.ClusterWhereInput{
						ID: clusterId,
					},
					Data: &models.ClusterRecycleBinUpdationParamsData{
						Enabled: originClusterSetting.VMRecycleBin.Enabled,
						Retain:  originClusterSetting.VMRecycleBin.Retain,
					},
				}
				_, err := client.GlobalSettings.UpdateClusterRecycleBinSetting(updateParams)
				if err != nil {
					panic(err.Error())
				}
			}()
		}
		wg.Wait()
	}
}
