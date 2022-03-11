package integration

import (
	"github.com/smartxworks/cloudtower-go-sdk/client/iscsi_connection"
	"github.com/smartxworks/cloudtower-go-sdk/models"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Iscsi api", Ordered, func() {
	It("should get iscsi connection", func() {
		params := iscsi_connection.NewGetIscsiConnectionsParams()
		params.RequestBody = &models.GetIscsiConnectionsRequestBody{}
		res, err := Client.IscsiConnection.GetIscsiConnections(params)
		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
	})
})
