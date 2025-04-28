// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	providerconfig "github.com/evaneos/provider-rabbitmq/internal/controller/providerconfig"
	binding "github.com/evaneos/provider-rabbitmq/internal/controller/rabbitmq/binding"
	exchange "github.com/evaneos/provider-rabbitmq/internal/controller/rabbitmq/exchange"
	federationupstream "github.com/evaneos/provider-rabbitmq/internal/controller/rabbitmq/federationupstream"
	operatorpolicy "github.com/evaneos/provider-rabbitmq/internal/controller/rabbitmq/operatorpolicy"
	permissions "github.com/evaneos/provider-rabbitmq/internal/controller/rabbitmq/permissions"
	policy "github.com/evaneos/provider-rabbitmq/internal/controller/rabbitmq/policy"
	queue "github.com/evaneos/provider-rabbitmq/internal/controller/rabbitmq/queue"
	shovel "github.com/evaneos/provider-rabbitmq/internal/controller/rabbitmq/shovel"
	topicpermissions "github.com/evaneos/provider-rabbitmq/internal/controller/rabbitmq/topicpermissions"
	user "github.com/evaneos/provider-rabbitmq/internal/controller/rabbitmq/user"
	vhost "github.com/evaneos/provider-rabbitmq/internal/controller/rabbitmq/vhost"
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
