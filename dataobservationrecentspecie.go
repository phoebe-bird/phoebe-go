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

// DataObservationRecentSpecieService contains methods and other services that help
// with interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObservationRecentSpecieService] method instead.
type DataObservationRecentSpecieService struct {
	Options []option.RequestOption
}

// NewDataObservationRecentSpecieService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDataObservationRecentSpecieService(opts ...option.RequestOption) (r *DataObservationRecentSpecieService) {
	r = &DataObservationRecentSpecieService{}
	r.Options = opts
	return
}

// Get the recent observations, up to 30 days ago, of a particular species in a
// country, region or location. Results include only the most recent observation
// from each location in the region specified.
//
// #### Notes
//
// The species code is typically a 6-letter code, e.g. cangoo for Canada Goose. You
// can get complete set of species code from the GET eBird Taxonomy end-point.
//
// When using the _r_ query parameter set the _regionCode_ URL parameter to an
// empty string.
func (r *DataObservationRecentSpecieService) Get(ctx context.Context, regionCode string, speciesCode string, query DataObservationRecentSpecieGetParams, opts ...option.RequestOption) (res *[]Observation, err error) {
	opts = append(r.Options[:], opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	if speciesCode == "" {
		err = errors.New("missing required speciesCode parameter")
		return
	}
	path := fmt.Sprintf("data/obs/%s/recent/%s", regionCode, speciesCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DataObservationRecentSpecieGetParams struct {
	// The number of days back to fetch observations.
	Back param.Field[int64] `query:"back"`
	// Only fetch observations from hotspots
	Hotspot param.Field[bool] `query:"hotspot"`
	// Include observations which have not yet been reviewed.
	IncludeProvisional param.Field[bool] `query:"includeProvisional"`
	// Only fetch this number of observations
	MaxResults param.Field[int64] `query:"maxResults"`
	// Fetch observations from up to 10 locations
	R param.Field[[]string] `query:"r"`
	// Use this language for species common names
	SppLocale param.Field[string] `query:"sppLocale"`
}

// URLQuery serializes [DataObservationRecentSpecieGetParams]'s query parameters as
// `url.Values`.
func (r DataObservationRecentSpecieGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
