// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	group "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/ecs/group"
	providerconfig "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/providerconfig"
	vpc "github.com/crossplane-contrib/provider-upjet-alibabacloud/internal/controller/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		group.Setup,
		providerconfig.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
