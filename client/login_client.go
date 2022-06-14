package client

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/user"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

type ClientConfig struct {
	Host     string
	BasePath string
	Schemes  []string
	formats  *strfmt.Registry
}

type UserConfig struct {
	Name     string
	Password string
	Source   models.UserSource
}

func NewWithUserConfig(clientConfig ClientConfig, userConfig UserConfig) (*Cloudtower, error) {
	transport := httptransport.New(clientConfig.Host, clientConfig.BasePath, clientConfig.Schemes)
	var client *Cloudtower
	if clientConfig.formats == nil {
		client = New(transport, strfmt.Default)
	} else {
		client = New(transport, *clientConfig.formats)
	}
	params := user.NewLoginParams()
	params.RequestBody = &models.LoginInput{
		Username: &userConfig.Name,
		Password: &userConfig.Password,
		Source:   userConfig.Source.Pointer(),
	}
	resp, err := client.User.Login(params)
	if err != nil {
		return nil, err
	}
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", *resp.Payload.Data.Token)
	return client, nil
}
