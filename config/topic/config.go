package topic

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rabbitmq_topic_permissions", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "rabbitmq"
		r.Kind = "TopicPermissions"
		r.References["vhost"] = config.Reference{
			TerraformName: "rabbitmq_vhost",
		}
	})
}
