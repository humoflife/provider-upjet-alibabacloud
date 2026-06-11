// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package cr

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_cr_chain", func(r *config.Resource) {
		r.ShortGroup = "cr"
	})

	p.AddResourceConfigurator("alicloud_cr_chart_namespace", func(r *config.Resource) {
		r.ShortGroup = "cr"
	})

	p.AddResourceConfigurator("alicloud_cr_chart_repository", func(r *config.Resource) {
		r.ShortGroup = "cr"
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
		// instance_id is populated via a cross-resource reference
		// (instanceIdSelector). It is Required in the Terraform schema, which
		// makes upjet emit a required-parameter CEL rule that is not
		// reference-aware, so a selector cannot satisfy it at admission. Mark it
		// Optional so the reference path works, consistent with the other cr
		// resources whose instance_id is reference-populated.
		if s, ok := r.TerraformResource.Schema["instance_id"]; ok {
			s.Required = false
			s.Optional = true
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
