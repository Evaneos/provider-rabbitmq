/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	bindingCluster "github.com/evaneos/provider-rabbitmq/config/cluster/binding"
	exchangeCluster "github.com/evaneos/provider-rabbitmq/config/cluster/exchange"
	federationCluster "github.com/evaneos/provider-rabbitmq/config/cluster/federation"
	operatorCluster "github.com/evaneos/provider-rabbitmq/config/cluster/operator"
	permissionsCluster "github.com/evaneos/provider-rabbitmq/config/cluster/permissions"
	policyCluster "github.com/evaneos/provider-rabbitmq/config/cluster/policy"
	queueCluster "github.com/evaneos/provider-rabbitmq/config/cluster/queue"
	shovelCluster "github.com/evaneos/provider-rabbitmq/config/cluster/shovel"
	topicCluster "github.com/evaneos/provider-rabbitmq/config/cluster/topic"
	userCluster "github.com/evaneos/provider-rabbitmq/config/cluster/user"
	vhostCluster "github.com/evaneos/provider-rabbitmq/config/cluster/vhost"

	bindingNamespaced "github.com/evaneos/provider-rabbitmq/config/namespaced/binding"
	exchangeNamespaced "github.com/evaneos/provider-rabbitmq/config/namespaced/exchange"
	federationNamespaced "github.com/evaneos/provider-rabbitmq/config/namespaced/federation"
	operatorNamespaced "github.com/evaneos/provider-rabbitmq/config/namespaced/operator"
	permissionsNamespaced "github.com/evaneos/provider-rabbitmq/config/namespaced/permissions"
	policyNamespaced "github.com/evaneos/provider-rabbitmq/config/namespaced/policy"
	queueNamespaced "github.com/evaneos/provider-rabbitmq/config/namespaced/queue"
	shovelNamespaced "github.com/evaneos/provider-rabbitmq/config/namespaced/shovel"
	topicNamespaced "github.com/evaneos/provider-rabbitmq/config/namespaced/topic"
	userNamespaced "github.com/evaneos/provider-rabbitmq/config/namespaced/user"
	vhostNamespaced "github.com/evaneos/provider-rabbitmq/config/namespaced/vhost"
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
		bindingCluster.Configure,
		exchangeCluster.Configure,
		federationCluster.Configure,
		operatorCluster.Configure,
		permissionsCluster.Configure,
		policyCluster.Configure,
		queueCluster.Configure,
		shovelCluster.Configure,
		topicCluster.Configure,
		userCluster.Configure,
		vhostCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("m.evaneos.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		bindingNamespaced.Configure,
		exchangeNamespaced.Configure,
		federationNamespaced.Configure,
		operatorNamespaced.Configure,
		permissionsNamespaced.Configure,
		policyNamespaced.Configure,
		queueNamespaced.Configure,
		shovelNamespaced.Configure,
		topicNamespaced.Configure,
		userNamespaced.Configure,
		vhostNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
