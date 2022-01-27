package main

import (
	apiclient "github.com/Sczlog/cloudtower-go-sdk/client"
	"github.com/Sczlog/cloudtower-go-sdk/client/user"
	"github.com/Sczlog/cloudtower-go-sdk/client/vm"
	"github.com/Sczlog/cloudtower-go-sdk/models"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func strPtr(s string) *string {
	return &s
}

func int32Ptr(val int32) *int32 {
	return &val
}

func main() {
	transport := httptransport.New("tower.smartx.com", "/v2/api", []string{"http"})
	c := apiclient.New(transport, strfmt.Default)

	loginParams := user.NewLoginParams()
	loginParams.RequestBody = &models.LoginInput{
		Username: strPtr("Username"),
		Password: strPtr("Password"),
		Source:   models.NewUserSource(models.UserSourceLDAP),
	}

	logResponse, err := c.User.Login(loginParams)

	if err != nil {
		panic(err.Error())
	}

	httptransport.APIKeyAuth("Authorization", "header", *logResponse.Payload.Data.Token)

	getVmsParams := vm.NewGetVmsParams()

	getVmsParams.RequestBody = &models.GetVmsRequestBody{
		First: int32Ptr(32),
	}

	response, err := c.VM.GetVms(getVmsParams, nil)
	if err != nil {
		panic(err.Error())
	}
	vm := *response.GetPayload()[0]

	println(*vm.CPU.Cores)
}
