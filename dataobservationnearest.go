// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"github.com/phoebe-bird/phoebe-go/option"
)

// DataObservationNearestService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObservationNearestService] method instead.
type DataObservationNearestService struct {
	Options    []option.RequestOption
	GeoSpecies *DataObservationNearestGeoSpecieService
}

// NewDataObservationNearestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDataObservationNearestService(opts ...option.RequestOption) (r *DataObservationNearestService) {
	r = &DataObservationNearestService{}
	r.Options = opts
	r.GeoSpecies = NewDataObservationNearestGeoSpecieService(opts...)
	return
}
