package main

import (
	"flag"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	kms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/kms/v20190118"
)

func main() {
	var ak string
	var sk string
	var region string
	var endpoint string
	flag.StringVar(&ak, "ak", "", "")
	flag.StringVar(&sk, "sk", "", "")
	flag.StringVar(&region, "region", "", "")
	flag.StringVar(&endpoint, "endpoint", "", "optional")
	flag.Parse()
	cred := common.NewCredential(ak, sk)
	pf := profile.NewClientProfile()
	if endpoint != "" {
		pf.HttpProfile.Endpoint = endpoint
	}
	if region == "" {
		region = "ap-guangzhou"
	}
	kmsCli, err := kms.NewClient(cred, region, pf)
	if err != nil {
		panic(err)
	}
	resp, err := kmsCli.GetServiceStatus(kms.NewGetServiceStatusRequest())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp.Response)
}
