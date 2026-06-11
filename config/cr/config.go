// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package cr

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_cr_chain", func(r *config.Resource) {
		r.ShortGroup = "cr"
		// Auto-detection mis-resolves instance_id to an EeNamespace; it is the
		// Enterprise Edition instance ID.
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_cr_ee_instance",
			Extractor:     common.PathIdExtractor,
		}
	})

	p.AddResourceConfigurator("alicloud_cr_chart_namespace", func(r *config.Resource) {
		r.ShortGroup = "cr"
	})

	p.AddResourceConfigurator("alicloud_cr_chart_repository", func(r *config.Resource) {
		r.ShortGroup = "cr"
		// Auto-detection mis-resolves instance_id to a ChartNamespace; it is the
		// Enterprise Edition instance ID.
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_cr_ee_instance",
			Extractor:     common.PathIdExtractor,
		}
	})

	p.AddResourceConfigurator("alicloud_cr_ee_instance", func(r *config.Resource) {
		r.ShortGroup = "cr"
	})

	p.AddResourceConfigurator("alicloud_cr_ee_namespace", func(r *config.Resource) {
		r.ShortGroup = "cr"
	})

	p.AddResourceConfigurator("alicloud_cr_ee_repo", func(r *config.Resource) {
		r.ShortGroup = "cr"
	})

	p.AddResourceConfigurator("alicloud_cr_ee_sync_rule", func(r *config.Resource) {
		r.ShortGroup = "cr"
	})

	p.AddResourceConfigurator("alicloud_cr_endpoint_acl_policy", func(r *config.Resource) {
		r.ShortGroup = "cr"
		// Auto-detection generated no instance_id reference for this resource,
		// so instanceIdSelector could not resolve and the (non reference-aware)
		// required-parameter CEL rejected the manifest at admission. Configure
		// the reference explicitly; this both enables the selector and excludes
		// instance_id from the required-parameter CEL, matching the other cr
		// resources.
		r.References["instance_id"] = config.Reference{
			TerraformName: "alicloud_cr_ee_instance",
			Extractor:     common.PathIdExtractor,
		}
	})

	p.AddResourceConfigurator("alicloud_cr_namespace", func(r *config.Resource) {
		r.ShortGroup = "cr"
		// The default kind "Namespace" yields the plural "namespaces", which
		// collides with the Kubernetes API server's reserved namespace-scoping
		// path grammar (/apis/<g>/<v>/namespaces/<name>/...), making individual
		// objects unaddressable by name (GET/PATCH/DELETE 404). Rename the kind
		// so the plural avoids that reserved segment.
		r.Kind = "RegistryNamespace"
	})

	p.AddResourceConfigurator("alicloud_cr_repo", func(r *config.Resource) {
		r.ShortGroup = "cr"
	})

	p.AddResourceConfigurator("alicloud_cr_scan_rule", func(r *config.Resource) {
		r.ShortGroup = "cr"
	})

	p.AddResourceConfigurator("alicloud_cr_storage_domain_routing_rule", func(r *config.Resource) {
		r.ShortGroup = "cr"
	})

	p.AddResourceConfigurator("alicloud_cr_vpc_endpoint_linked_vpc", func(r *config.Resource) {
		r.ShortGroup = "cr"
	})
}
