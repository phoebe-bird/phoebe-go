// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/phoebe-bird/phoebe-go/internal/apiquery"
	"github.com/phoebe-bird/phoebe-go/internal/param"
	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
)

// DataObservationNearestGeoSpecieService contains methods and other services that
// help with interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObservationNearestGeoSpecieService] method instead.
type DataObservationNearestGeoSpecieService struct {
	Options []option.RequestOption
}

// NewDataObservationNearestGeoSpecieService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDataObservationNearestGeoSpecieService(opts ...option.RequestOption) (r *DataObservationNearestGeoSpecieService) {
	r = &DataObservationNearestGeoSpecieService{}
	r.Options = opts
	return
}

// Find the nearest locations where a species has been seen recently. #### Notes
// The species code is typically a 6-letter code, e.g. barswa for Barn Swallow. You
// can get complete set of species code from the GET eBird Taxonomy end-point.
func (r *DataObservationNearestGeoSpecieService) List(ctx context.Context, speciesCode string, query DataObservationNearestGeoSpecieListParams, opts ...option.RequestOption) (res *[]Observation, err error) {
	opts = slices.Concat(r.Options, opts)
	if speciesCode == "" {
		err = errors.New("missing required speciesCode parameter")
		return
	}
	path := fmt.Sprintf("data/nearest/geo/recent/%s", speciesCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DataObservationNearestGeoSpecieListParams struct {
	Lat param.Field[float64] `query:"lat,required"`
	Lng param.Field[float64] `query:"lng,required"`
	// The number of days back to fetch observations.
	Back param.Field[int64] `query:"back"`
	// Only fetch observations within this distance of the provided lat/lng
	Dist param.Field[int64] `query:"dist"`
	// Only fetch observations from hotspots
	Hotspot param.Field[bool] `query:"hotspot"`
	// Include observations which have not yet been reviewed.
	IncludeProvisional param.Field[bool] `query:"includeProvisional"`
	// Only fetch up to this number of observations
	MaxResults param.Field[int64] `query:"maxResults"`
	// Use this language for species common names
	SppLocale param.Field[string] `query:"sppLocale"`
}

// URLQuery serializes [DataObservationNearestGeoSpecieListParams]'s query
// parameters as `url.Values`.
func (r DataObservationNearestGeoSpecieListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
