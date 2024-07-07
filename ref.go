// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"github.com/phoebe-bird/phoebe-go/option"
)

// RefService contains methods and other services that help with interacting with
// the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefService] method instead.
type RefService struct {
	Options  []option.RequestOption
	Region   *RefRegionService
	Hotspot  *RefHotspotService
	Taxonomy *RefTaxonomyService
}

// NewRefService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRefService(opts ...option.RequestOption) (r *RefService) {
	r = &RefService{}
	r.Options = opts
	r.Region = NewRefRegionService(opts...)
	r.Hotspot = NewRefHotspotService(opts...)
	r.Taxonomy = NewRefTaxonomyService(opts...)
	return
}
