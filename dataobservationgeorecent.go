// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/phoebe-bird/phoebe-go/internal/apiquery"
	"github.com/phoebe-bird/phoebe-go/internal/param"
	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
)

// DataObservationGeoRecentService contains methods and other services that help
// with interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObservationGeoRecentService] method instead.
type DataObservationGeoRecentService struct {
	Options []option.RequestOption
	Species *DataObservationGeoRecentSpecieService
	Notable *DataObservationGeoRecentNotableService
}

// NewDataObservationGeoRecentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDataObservationGeoRecentService(opts ...option.RequestOption) (r *DataObservationGeoRecentService) {
	r = &DataObservationGeoRecentService{}
	r.Options = opts
	r.Species = NewDataObservationGeoRecentSpecieService(opts...)
	r.Notable = NewDataObservationGeoRecentNotableService(opts...)
	return
}

// Get the list of recent observations (up to 30 days ago) of birds seen at
// locations within a radius of up to 50 kilometers, from a given set of
// coordinates. Results include only the most recent observation for each species
// in the region specified.
func (r *DataObservationGeoRecentService) List(ctx context.Context, query DataObservationGeoRecentListParams, opts ...option.RequestOption) (res *[]Observation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "data/obs/geo/recent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DataObservationGeoRecentListParams struct {
	Lat param.Field[float64] `query:"lat" api:"required"`
	Lng param.Field[float64] `query:"lng" api:"required"`
	// The number of days back to fetch observations.
	Back param.Field[int64] `query:"back"`
	// Only fetch observations from these taxonomic categories
	Cat param.Field[DataObservationGeoRecentListParamsCat] `query:"cat"`
	// The search radius from the given position, in kilometers.
	Dist param.Field[int64] `query:"dist"`
	// Only fetch observations from hotspots
	Hotspot param.Field[bool] `query:"hotspot"`
	// Include observations which have not yet been reviewed.
	IncludeProvisional param.Field[bool] `query:"includeProvisional"`
	// Only fetch this number of observations
	MaxResults param.Field[int64] `query:"maxResults"`
	// Sort observations by taxonomy or by date, most recent first.
	Sort param.Field[DataObservationGeoRecentListParamsSort] `query:"sort"`
	// Use this language for species common names
	SppLocale param.Field[string] `query:"sppLocale"`
}

// URLQuery serializes [DataObservationGeoRecentListParams]'s query parameters as
// `url.Values`.
func (r DataObservationGeoRecentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only fetch observations from these taxonomic categories
type DataObservationGeoRecentListParamsCat string

const (
	DataObservationGeoRecentListParamsCatSpecies    DataObservationGeoRecentListParamsCat = "species"
	DataObservationGeoRecentListParamsCatSlash      DataObservationGeoRecentListParamsCat = "slash"
	DataObservationGeoRecentListParamsCatIssf       DataObservationGeoRecentListParamsCat = "issf"
	DataObservationGeoRecentListParamsCatSpuh       DataObservationGeoRecentListParamsCat = "spuh"
	DataObservationGeoRecentListParamsCatHybrid     DataObservationGeoRecentListParamsCat = "hybrid"
	DataObservationGeoRecentListParamsCatDomestic   DataObservationGeoRecentListParamsCat = "domestic"
	DataObservationGeoRecentListParamsCatForm       DataObservationGeoRecentListParamsCat = "form"
	DataObservationGeoRecentListParamsCatIntergrade DataObservationGeoRecentListParamsCat = "intergrade"
)

func (r DataObservationGeoRecentListParamsCat) IsKnown() bool {
	switch r {
	case DataObservationGeoRecentListParamsCatSpecies, DataObservationGeoRecentListParamsCatSlash, DataObservationGeoRecentListParamsCatIssf, DataObservationGeoRecentListParamsCatSpuh, DataObservationGeoRecentListParamsCatHybrid, DataObservationGeoRecentListParamsCatDomestic, DataObservationGeoRecentListParamsCatForm, DataObservationGeoRecentListParamsCatIntergrade:
		return true
	}
	return false
}

// Sort observations by taxonomy or by date, most recent first.
type DataObservationGeoRecentListParamsSort string

const (
	DataObservationGeoRecentListParamsSortDate    DataObservationGeoRecentListParamsSort = "date"
	DataObservationGeoRecentListParamsSortSpecies DataObservationGeoRecentListParamsSort = "species"
)

func (r DataObservationGeoRecentListParamsSort) IsKnown() bool {
	switch r {
	case DataObservationGeoRecentListParamsSortDate, DataObservationGeoRecentListParamsSortSpecies:
		return true
	}
	return false
}
