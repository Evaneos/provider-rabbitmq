package binding

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rabbitmq_binding", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "rabbitmq"
		r.References["vhost"] = config.Reference{
			TerraformName: "rabbitmq_vhost",
		}
		r.References["source"] = config.Reference{
			TerraformName: "rabbitmq_exchange",
			Extractor:     "github.com/evaneos/provider-rabbitmq/config/common.ExtractResourceName()",
		}
	})
}
