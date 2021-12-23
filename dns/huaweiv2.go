package dns

import (
	"fmt"
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack"
	"github.com/sirupsen/logrus"
)

func dns() {

	tokenOpts := golangsdk.AuthOptions{
		IdentityEndpoint: "https://openstack.example.com:5000/v2.0",
		Username:         "{username}",
		Password:         "{password}",
	}

	//初始化provider client
	providerClient, providerErr := openstack.AuthenticatedClient(tokenOpts)
	if providerErr != nil {
		fmt.Println("init provider client error:", providerErr)
		panic(providerErr)
	}

	dnsClient, err := openstack.NewDNSV2(providerClient, golangsdk.EndpointOpts{
		Type:         "",
		Name:         "",
		Region:       "",
		Availability: "",
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("init dns client fail")
	}
	dnsClient.Token()
}
