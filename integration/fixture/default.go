package fixture

import (
	"os"

	. "github.com/openlyinc/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/cluster"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vlan"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

var client *apiclient.Cloudtower = nil

func GetClient() *apiclient.Cloudtower {
	if client != nil {
		return client
	}
	client, err := apiclient.NewWithUserConfig(apiclient.ClientConfig{
		Host:     os.Getenv("CLOUDTOWER_SDK_ENDPOINT"),
		BasePath: "v2/api",
		Schemes:  []string{"http"},
	}, apiclient.UserConfig{
		Name:     os.Getenv("CLOUDTOWER_SDK_USERNAME"),
		Password: os.Getenv("CLOUDTOWER_SDK_PASSWORD"),
		Source:   models.UserSource(os.Getenv("CLOUDTOWER_SDK_USERSOURCE")),
	})
	if err != nil {
		panic(err.Error())
	}
	return client
}

func GetDefaultCluster(client *apiclient.Cloudtower, name string) *models.Cluster {
	params := cluster.NewGetClustersParams()
	params.RequestBody = &models.GetClustersRequestBody{
		Where: &models.ClusterWhereInput{
			Name: String(name),
		},
		First: Int32(1),
	}
	res, err := client.Cluster.GetClusters(params)
	if err != nil {
		panic(err.Error())
	}
	if len(res.Payload) == 0 {
		panic("Deafult cluster not found")
	}
	return res.Payload[0]
}

func GetDefaultVlan(client *apiclient.Cloudtower, clusterId *string) *models.Vlan {
	params := vlan.NewGetVlansParams()
	params.RequestBody = &models.GetVlansRequestBody{
		Where: &models.VlanWhereInput{
			Name: String("default"),
			Vds: &models.VdsWhereInput{
				Cluster: &models.ClusterWhereInput{
					ID: clusterId,
				},
				Internal: Bool(false),
			},
		},
		First: Int32(1),
	}
	res, err := client.Vlan.GetVlans(params)
	if err != nil {
		panic(err.Error())
	}
	if len(res.Payload) == 0 {
		panic("Deafult vlan not found")
	}
	return res.Payload[0]
}
