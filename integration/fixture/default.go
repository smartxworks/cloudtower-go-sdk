package fixture

import (
	"os"

	apiclient "github.com/Sczlog/cloudtower-go-sdk/client"
	"github.com/Sczlog/cloudtower-go-sdk/client/cluster"
	"github.com/Sczlog/cloudtower-go-sdk/client/user"
	"github.com/Sczlog/cloudtower-go-sdk/client/vlan"
	"github.com/Sczlog/cloudtower-go-sdk/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	. "github.com/openlyinc/pointy"
)

var client *apiclient.Cloudtower = nil

func GetClient() *apiclient.Cloudtower {
	if client != nil {
		return client
	}
	transport := httptransport.New(os.Getenv("CLOUDTOWER_SDK_ENDPOINT"), "v2/api", []string{"http"})

	client := apiclient.New(transport, strfmt.Default)
	token := os.Getenv("CLOUDTOWER_SDK_TOKEN")
	if len(token) == 0 {
		loginParams := user.NewLoginParams()
		loginParams.RequestBody = &models.LoginInput{
			Username: String(os.Getenv("CLOUDTOWER_SDK_USERNAME")),
			Password: String(os.Getenv("CLOUDTOWER_SDK_PASSWORD")),
			Source:   models.NewUserSource(models.UserSource(os.Getenv("CLOUDTOWER_SDK_USERSOURCE"))),
		}
		loginRes, err := client.User.Login(loginParams)
		if err != nil {
			panic(err.Error())
		}
		token = *loginRes.Payload.Data.Token
	}

	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", token)
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
		panic("Deafult clusteer not found")
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
