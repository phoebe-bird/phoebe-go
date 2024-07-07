// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/phoebe-bird/phoebe-go/internal/apiquery"
	"github.com/phoebe-bird/phoebe-go/internal/param"
	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
)

// DataObservationGeoRecentSpecieService contains methods and other services that
// help with interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObservationGeoRecentSpecieService] method instead.
type DataObservationGeoRecentSpecieService struct {
	Options []option.RequestOption
}

// NewDataObservationGeoRecentSpecieService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDataObservationGeoRecentSpecieService(opts ...option.RequestOption) (r *DataObservationGeoRecentSpecieService) {
	r = &DataObservationGeoRecentSpecieService{}
	r.Options = opts
	return
}

// Get all observations of a species, seen up to 30 days ago, at any location
// within a radius of up to 50 kilometers, from a given set of coordinates. Results
// include only the most recent observation from each location in the region
// specified.
//
// #### URL parameters
//
// | Name        | Description             |
// | ----------- | ----------------------- |
// | speciesCode | The eBird species code. |
//
// #### Notes
//
// The species code is typically a 6-letter code, e.g. horlar for Horned Lark. You
// can get complete set of species code from the GET eBird Taxonomy end-point.
func (r *DataObservationGeoRecentSpecieService) List(ctx context.Context, speciesCode string, query DataObservationGeoRecentSpecieListParams, opts ...option.RequestOption) (res *[]Observation, err error) {
	opts = append(r.Options[:], opts...)
	if speciesCode == "" {
		err = errors.New("missing required speciesCode parameter")
		return
	}
	path := fmt.Sprintf("data/obs/geo/recent/%s", speciesCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DataObservationGeoRecentSpecieListParams struct {
	Lat param.Field[float64] `query:"lat,required"`
	Lng param.Field[float64] `query:"lng,required"`
	// The number of days back to fetch observations.
	Back param.Field[int64] `query:"back"`
	// The search radius from the given position, in kilometers.
	Dist param.Field[int64] `query:"dist"`
	// Only fetch observations from hotspots
	Hotspot param.Field[bool] `query:"hotspot"`
	// Include observations which have not yet been reviewed.
	IncludeProvisional param.Field[bool] `query:"includeProvisional"`
	// Only fetch this number of observations
	MaxResults param.Field[int64] `query:"maxResults"`
	// Use this language for species common names
	SppLocale param.Field[string] `query:"sppLocale"`
}

// URLQuery serializes [DataObservationGeoRecentSpecieListParams]'s query
// parameters as `url.Values`.
func (r DataObservationGeoRecentSpecieListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
