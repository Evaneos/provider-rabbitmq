// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	providerconfig "github.com/evaneos/provider-rabbitmq/internal/controller/namespaced/providerconfig"
	binding "github.com/evaneos/provider-rabbitmq/internal/controller/namespaced/rabbitmq/binding"
	exchange "github.com/evaneos/provider-rabbitmq/internal/controller/namespaced/rabbitmq/exchange"
	federationupstream "github.com/evaneos/provider-rabbitmq/internal/controller/namespaced/rabbitmq/federationupstream"
	operatorpolicy "github.com/evaneos/provider-rabbitmq/internal/controller/namespaced/rabbitmq/operatorpolicy"
	permissions "github.com/evaneos/provider-rabbitmq/internal/controller/namespaced/rabbitmq/permissions"
	policy "github.com/evaneos/provider-rabbitmq/internal/controller/namespaced/rabbitmq/policy"
	queue "github.com/evaneos/provider-rabbitmq/internal/controller/namespaced/rabbitmq/queue"
	shovel "github.com/evaneos/provider-rabbitmq/internal/controller/namespaced/rabbitmq/shovel"
	topicpermissions "github.com/evaneos/provider-rabbitmq/internal/controller/namespaced/rabbitmq/topicpermissions"
	user "github.com/evaneos/provider-rabbitmq/internal/controller/namespaced/rabbitmq/user"
	vhost "github.com/evaneos/provider-rabbitmq/internal/controller/namespaced/rabbitmq/vhost"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		binding.Setup,
		exchange.Setup,
		federationupstream.Setup,
		operatorpolicy.Setup,
		permissions.Setup,
		policy.Setup,
		queue.Setup,
		shovel.Setup,
		topicpermissions.Setup,
		user.Setup,
		vhost.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.SetupGated,
		binding.SetupGated,
		exchange.SetupGated,
		federationupstream.SetupGated,
		operatorpolicy.SetupGated,
		permissions.SetupGated,
		policy.SetupGated,
		queue.SetupGated,
		shovel.SetupGated,
		topicpermissions.SetupGated,
		user.SetupGated,
		vhost.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
