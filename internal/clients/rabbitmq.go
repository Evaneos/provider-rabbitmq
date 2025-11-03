/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/v2/pkg/terraform"

	"github.com/evaneos/provider-rabbitmq/apis/cluster/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal rabbitmq credentials as JSON"
	errRequiredProperty     = "missing %s required property"

	endpoint = "endpoint"
	username = "username"
	password = "password"
	insecure = "insecure"
	proxy    = "proxy"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		return extractConfigurationFromCreds(ps, creds)
	}
}

func extractConfigurationFromCreds(ps terraform.Setup, creds map[string]string) (terraform.Setup, error) {
	// Set credentials in Terraform provider configuration.
	ps.Configuration = map[string]any{}
	if v, ok := creds[endpoint]; ok {
		ps.Configuration[endpoint] = v
	} else {
		return ps, fmt.Errorf(errRequiredProperty, endpoint)
	}
	if v, ok := creds[username]; ok {
		ps.Configuration[username] = v
	} else {
		return ps, fmt.Errorf(errRequiredProperty, username)
	}
	if v, ok := creds[password]; ok {
		ps.Configuration[password] = v
	}
	if v, ok := creds[insecure]; ok {
		ps.Configuration[insecure] = v
	}
	if v, ok := creds[proxy]; ok {
		ps.Configuration[proxy] = v
	}

	return ps, nil
}
