/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package provider

import (
	"testing"

	"github.com/birdycloud/fabric-sdk-go/pkg/fabsdk"
	"github.com/birdycloud/fabric-sdk-go/test/integration"
	"github.com/birdycloud/fabric-sdk-go/test/integration/util/runner"
)

const (
	org1Name      = "Org1"
	org1User      = "User1"
	org1AdminUser = "Admin"
)

var mainSDK *fabsdk.FabricSDK
var mainTestSetup *integration.BaseSetupImpl

func TestMain(m *testing.M) {
	r := runner.New()
	r.Initialize()
	mainSDK = r.SDK()
	mainTestSetup = r.TestSetup()

	r.Run(m)
}
