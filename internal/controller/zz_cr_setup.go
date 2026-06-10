// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	chain "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/chain"
	chartnamespace "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/chartnamespace"
	chartrepository "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/chartrepository"
	eeinstance "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/eeinstance"
	eenamespace "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/eenamespace"
	eerepo "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/eerepo"
	eesyncrule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/eesyncrule"
	endpointaclpolicy "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/endpointaclpolicy"
	registrynamespace "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/registrynamespace"
	repo "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/repo"
	scanrule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/scanrule"
	storagedomainroutingrule "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/storagedomainroutingrule"
	vpcendpointlinkedvpc "github.com/crossplane-contrib/provider-alibabacloud/internal/controller/cr/vpcendpointlinkedvpc"
)

// Setup_cr creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cr(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		chain.Setup,
		chartnamespace.Setup,
		chartrepository.Setup,
		eeinstance.Setup,
		eenamespace.Setup,
		eerepo.Setup,
		eesyncrule.Setup,
		endpointaclpolicy.Setup,
		registrynamespace.Setup,
		repo.Setup,
		scanrule.Setup,
		storagedomainroutingrule.Setup,
		vpcendpointlinkedvpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
