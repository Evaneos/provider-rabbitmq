/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/evaneos/provider-rabbitmq/config/binding"
	"github.com/evaneos/provider-rabbitmq/config/exchange"
	"github.com/evaneos/provider-rabbitmq/config/federation"
	"github.com/evaneos/provider-rabbitmq/config/operator"
	"github.com/evaneos/provider-rabbitmq/config/permissions"
	"github.com/evaneos/provider-rabbitmq/config/policy"
	"github.com/evaneos/provider-rabbitmq/config/queue"
	"github.com/evaneos/provider-rabbitmq/config/shovel"
	"github.com/evaneos/provider-rabbitmq/config/topic"
	"github.com/evaneos/provider-rabbitmq/config/user"
	"github.com/evaneos/provider-rabbitmq/config/vhost"
)

const (
	resourcePrefix = "rabbitmq"
	modulePath     = "github.com/evaneos/provider-rabbitmq"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("evaneos.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		binding.Configure,
		exchange.Configure,
		federation.Configure,
		operator.Configure,
		permissions.Configure,
		policy.Configure,
		queue.Configure,
		shovel.Configure,
		topic.Configure,
		user.Configure,
		vhost.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
