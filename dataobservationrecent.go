// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/phoebe-go/internal/apiquery"
	"github.com/stainless-sdks/phoebe-go/internal/param"
	"github.com/stainless-sdks/phoebe-go/internal/requestconfig"
	"github.com/stainless-sdks/phoebe-go/option"
)

// DataObservationRecentService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObservationRecentService] method instead.
type DataObservationRecentService struct {
	Options  []option.RequestOption
	Notable  *DataObservationRecentNotableService
	Species  *DataObservationRecentSpecieService
	Historic *DataObservationRecentHistoricService
}

// NewDataObservationRecentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDataObservationRecentService(opts ...option.RequestOption) (r *DataObservationRecentService) {
	r = &DataObservationRecentService{}
	r.Options = opts
	r.Notable = NewDataObservationRecentNotableService(opts...)
	r.Species = NewDataObservationRecentSpecieService(opts...)
	r.Historic = NewDataObservationRecentHistoricService(opts...)
	return
}

// Get the list of recent observations (up to 30 days ago) of birds seen in a
// country, state, county, or location. Results include only the most recent
// observation for each species in the region specified.
func (r *DataObservationRecentService) List(ctx context.Context, regionCode string, query DataObservationRecentListParams, opts ...option.RequestOption) (res *[]Observation, err error) {
	opts = append(r.Options[:], opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("data/obs/%s/recent", regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DataObservationRecentListParams struct {
	// The number of days back to fetch observations.
	Back param.Field[int64] `query:"back"`
	// Only fetch observations from these taxonomic categories
	Cat param.Field[DataObservationRecentListParamsCat] `query:"cat"`
	// Only fetch observations from hotspots
	Hotspot param.Field[bool] `query:"hotspot"`
	// Include observations which have not yet been reviewed
	IncludeProvisional param.Field[bool] `query:"includeProvisional"`
	// Only fetch this number of observations
	MaxResults param.Field[int64] `query:"maxResults"`
	// Fetch observations from up to 10 locations
	R param.Field[[]string] `query:"r"`
	// Use this language for species common names
	SppLocale param.Field[string] `query:"sppLocale"`
}

// URLQuery serializes [DataObservationRecentListParams]'s query parameters as
// `url.Values`.
func (r DataObservationRecentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only fetch observations from these taxonomic categories
type DataObservationRecentListParamsCat string

const (
	DataObservationRecentListParamsCatSpecies    DataObservationRecentListParamsCat = "species"
	DataObservationRecentListParamsCatSlash      DataObservationRecentListParamsCat = "slash"
	DataObservationRecentListParamsCatIssf       DataObservationRecentListParamsCat = "issf"
	DataObservationRecentListParamsCatSpuh       DataObservationRecentListParamsCat = "spuh"
	DataObservationRecentListParamsCatHybrid     DataObservationRecentListParamsCat = "hybrid"
	DataObservationRecentListParamsCatDomestic   DataObservationRecentListParamsCat = "domestic"
	DataObservationRecentListParamsCatForm       DataObservationRecentListParamsCat = "form"
	DataObservationRecentListParamsCatIntergrade DataObservationRecentListParamsCat = "intergrade"
)

func (r DataObservationRecentListParamsCat) IsKnown() bool {
	switch r {
	case DataObservationRecentListParamsCatSpecies, DataObservationRecentListParamsCatSlash, DataObservationRecentListParamsCatIssf, DataObservationRecentListParamsCatSpuh, DataObservationRecentListParamsCatHybrid, DataObservationRecentListParamsCatDomestic, DataObservationRecentListParamsCatForm, DataObservationRecentListParamsCatIntergrade:
		return true
	}
	return false
}
