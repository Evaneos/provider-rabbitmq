package exchange

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rabbitmq_exchange", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "rabbitmq"
		r.References["vhost"] = config.Reference{
			TerraformName: "rabbitmq_vhost",
		}
	})
}
