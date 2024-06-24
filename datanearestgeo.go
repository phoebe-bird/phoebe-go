// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"github.com/stainless-sdks/phoebe-go/option"
)

// DataNearestGeoService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataNearestGeoService] method instead.
type DataNearestGeoService struct {
	Options []option.RequestOption
	Recent  *DataNearestGeoRecentService
}

// NewDataNearestGeoService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDataNearestGeoService(opts ...option.RequestOption) (r *DataNearestGeoService) {
	r = &DataNearestGeoService{}
	r.Options = opts
	r.Recent = NewDataNearestGeoRecentService(opts...)
	return
}
