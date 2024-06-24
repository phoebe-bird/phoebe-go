// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"github.com/stainless-sdks/phoebe-go/internal/apijson"
	"github.com/stainless-sdks/phoebe-go/option"
)

// DataObservationService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObservationService] method instead.
type DataObservationService struct {
	Options []option.RequestOption
	Recent  *DataObservationRecentService
	Geo     *DataObservationGeoService
	Nearest *DataObservationNearestService
}

// NewDataObservationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDataObservationService(opts ...option.RequestOption) (r *DataObservationService) {
	r = &DataObservationService{}
	r.Options = opts
	r.Recent = NewDataObservationRecentService(opts...)
	r.Geo = NewDataObservationGeoService(opts...)
	r.Nearest = NewDataObservationNearestService(opts...)
	return
}

type Observation struct {
	ID              int64           `json:"id"`
	ComName         string          `json:"comName"`
	Firstname       string          `json:"firstname"`
	HowMany         int64           `json:"howMany"`
	Lastname        string          `json:"lastname"`
	Lat             float64         `json:"lat"`
	Lng             float64         `json:"lng"`
	LocationPrivate bool            `json:"locationPrivate"`
	LocID           string          `json:"locId"`
	LocName         string          `json:"locName"`
	ObsDt           string          `json:"obsDt"`
	ObsReviewed     bool            `json:"obsReviewed"`
	ObsValid        bool            `json:"obsValid"`
	SciName         string          `json:"sciName"`
	SpeciesCode     string          `json:"speciesCode"`
	SubID           string          `json:"subId"`
	JSON            observationJSON `json:"-"`
}

// observationJSON contains the JSON metadata for the struct [Observation]
type observationJSON struct {
	ID              apijson.Field
	ComName         apijson.Field
	Firstname       apijson.Field
	HowMany         apijson.Field
	Lastname        apijson.Field
	Lat             apijson.Field
	Lng             apijson.Field
	LocationPrivate apijson.Field
	LocID           apijson.Field
	LocName         apijson.Field
	ObsDt           apijson.Field
	ObsReviewed     apijson.Field
	ObsValid        apijson.Field
	SciName         apijson.Field
	SpeciesCode     apijson.Field
	SubID           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Observation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observationJSON) RawJSON() string {
	return r.raw
}
