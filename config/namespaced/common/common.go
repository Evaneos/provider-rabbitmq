package common

import (
	"strings"

	"github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	jresource "github.com/crossplane/upjet/v2/pkg/resource"
)

// ExtractResourceName extracts the value of `spec.atProvider.id`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractResourceName() reference.ExtractValueFn {
	return func(mr resource.Managed) string {
		tr, ok := mr.(jresource.Terraformed)
		if !ok {
			return ""
		}
		sp := strings.Split(tr.GetID(), "@")
		return sp[0]
	}
}
