// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"github.com/phoebe-bird/phoebe-go/option"
)

// RefRegionService contains methods and other services that help with interacting
// with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefRegionService] method instead.
type RefRegionService struct {
	Options []option.RequestOption
	// With the ref/geo end-point you can find a country's or region's neighbours.
	Adjacent *RefRegionAdjacentService
	// The ref/region end-points return information on regions.
	Info *RefRegionInfoService
	// The ref/region end-points return information on regions.
	List *RefRegionListService
}

// NewRefRegionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRefRegionService(opts ...option.RequestOption) (r *RefRegionService) {
	r = &RefRegionService{}
	r.Options = opts
	r.Adjacent = NewRefRegionAdjacentService(opts...)
	r.Info = NewRefRegionInfoService(opts...)
	r.List = NewRefRegionListService(opts...)
	return
}
