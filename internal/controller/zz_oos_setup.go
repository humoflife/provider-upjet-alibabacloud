// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	application "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/oos/application"
	applicationgroup "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/oos/applicationgroup"
	defaultpatchbaseline "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/oos/defaultpatchbaseline"
	execution "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/oos/execution"
	parameter "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/oos/parameter"
	patchbaseline "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/oos/patchbaseline"
	secretparameter "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/oos/secretparameter"
	servicesetting "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/oos/servicesetting"
	stateconfiguration "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/oos/stateconfiguration"
	template "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/oos/template"
)

// Setup_oos creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_oos(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		applicationgroup.Setup,
		defaultpatchbaseline.Setup,
		execution.Setup,
		parameter.Setup,
		patchbaseline.Setup,
		secretparameter.Setup,
		servicesetting.Setup,
		stateconfiguration.Setup,
		template.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
