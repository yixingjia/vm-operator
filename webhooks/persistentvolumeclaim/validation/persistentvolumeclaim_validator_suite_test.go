// Copyright (c) 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package validation_test

import (
	"testing"

	. "github.com/onsi/ginkgo"

	"github.com/vmware-tanzu/vm-operator/test/builder"
	"github.com/vmware-tanzu/vm-operator/webhooks/persistentvolumeclaim/validation"
)

// suite is used for unit and integration testing this webhook.
var suite = builder.NewTestSuiteForValidatingWebhook(
	validation.AddToManager,
	validation.NewValidator,
	"default.validating.persistentvolumeclaim.vmoperator.vmware.com")

func TestWebhook(t *testing.T) {
	suite.Register(t, "Validation webhook suite", intgTests, unitTests)
}

var _ = BeforeSuite(suite.BeforeSuite)

var _ = AfterSuite(suite.AfterSuite)
