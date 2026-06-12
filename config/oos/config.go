package oos

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_oos_application", func(r *config.Resource) {
		r.ShortGroup = "oos"
	})
	p.AddResourceConfigurator("alicloud_oos_parameter", func(r *config.Resource) {
		r.ShortGroup = "oos"
	})
	p.AddResourceConfigurator("alicloud_oos_template", func(r *config.Resource) {
		r.ShortGroup = "oos"
	})
}
