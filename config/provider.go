/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/ackone"
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/alidns"
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/cdn"
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/cloudmonitorservice"
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/ecs"
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/kms"
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/messageservice"
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/oss"
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/polardb"
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/ram"
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/tair"
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/vpc"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "alicloud"
	modulePath     = "github.com/crossplane-contrib/provider-upjet-alibabacloud"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("alibabacloud.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		ackone.Configure,
		alidns.Configure,
		cdn.Configure,
		cloudmonitorservice.Configure,
		ecs.Configure,
		kms.Configure,
		messageservice.Configure,
		oss.Configure,
		polardb.Configure,
		ram.Configure,
		tair.Configure,
		vpc.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
