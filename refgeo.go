// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"github.com/stainless-sdks/phoebe-go/option"
)

// RefGeoService contains methods and other services that help with interacting
// with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefGeoService] method instead.
type RefGeoService struct {
	Options  []option.RequestOption
	Adjacent *RefGeoAdjacentService
}

// NewRefGeoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRefGeoService(opts ...option.RequestOption) (r *RefGeoService) {
	r = &RefGeoService{}
	r.Options = opts
	r.Adjacent = NewRefGeoAdjacentService(opts...)
	return
}
