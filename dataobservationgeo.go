// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"github.com/phoebe-bird/phoebe-go/option"
)

// DataObservationGeoService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObservationGeoService] method instead.
type DataObservationGeoService struct {
	Options []option.RequestOption
	Recent  *DataObservationGeoRecentService
}

// NewDataObservationGeoService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDataObservationGeoService(opts ...option.RequestOption) (r *DataObservationGeoService) {
	r = &DataObservationGeoService{}
	r.Options = opts
	r.Recent = NewDataObservationGeoRecentService(opts...)
	return
}
