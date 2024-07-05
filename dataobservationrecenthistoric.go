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

// DataObservationRecentHistoricService contains methods and other services that
// help with interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataObservationRecentHistoricService] method instead.
type DataObservationRecentHistoricService struct {
	Options []option.RequestOption
}

// NewDataObservationRecentHistoricService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDataObservationRecentHistoricService(opts ...option.RequestOption) (r *DataObservationRecentHistoricService) {
	r = &DataObservationRecentHistoricService{}
	r.Options = opts
	return
}

// Get a list of all taxa seen in a country, region or location on a specific date,
// with the specific observations determined by the "rank" parameter (defaults to
// latest observation on the date).
//
// #### Notes Responses may be cached for 30 minutes
func (r *DataObservationRecentHistoricService) List(ctx context.Context, regionCode string, y int64, m int64, d int64, query DataObservationRecentHistoricListParams, opts ...option.RequestOption) (res *[]Observation, err error) {
	opts = append(r.Options[:], opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("data/obs/%s/historic/%v/%v/%v", regionCode, y, m, d)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DataObservationRecentHistoricListParams struct {
	// Only fetch observations from these taxonomic categories
	Cat param.Field[DataObservationRecentHistoricListParamsCat] `query:"cat"`
	// Include a subset (simple), or all (full), of the fields available.
	Detail param.Field[DataObservationRecentHistoricListParamsDetail] `query:"detail"`
	// Only fetch observations from hotspots
	Hotspot param.Field[bool] `query:"hotspot"`
	// Include observations which have not yet been reviewed.
	IncludeProvisional param.Field[bool] `query:"includeProvisional"`
	// Only fetch this number of observations
	MaxResults param.Field[int64] `query:"maxResults"`
	// Fetch observations from up to 50 locations
	R param.Field[[]string] `query:"r"`
	// Include latest observation of the day, or the first added
	Rank param.Field[DataObservationRecentHistoricListParamsRank] `query:"rank"`
	// Use this language for species common names
	SppLocale param.Field[string] `query:"sppLocale"`
}

// URLQuery serializes [DataObservationRecentHistoricListParams]'s query parameters
// as `url.Values`.
func (r DataObservationRecentHistoricListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only fetch observations from these taxonomic categories
type DataObservationRecentHistoricListParamsCat string

const (
	DataObservationRecentHistoricListParamsCatSpecies    DataObservationRecentHistoricListParamsCat = "species"
	DataObservationRecentHistoricListParamsCatSlash      DataObservationRecentHistoricListParamsCat = "slash"
	DataObservationRecentHistoricListParamsCatIssf       DataObservationRecentHistoricListParamsCat = "issf"
	DataObservationRecentHistoricListParamsCatSpuh       DataObservationRecentHistoricListParamsCat = "spuh"
	DataObservationRecentHistoricListParamsCatHybrid     DataObservationRecentHistoricListParamsCat = "hybrid"
	DataObservationRecentHistoricListParamsCatDomestic   DataObservationRecentHistoricListParamsCat = "domestic"
	DataObservationRecentHistoricListParamsCatForm       DataObservationRecentHistoricListParamsCat = "form"
	DataObservationRecentHistoricListParamsCatIntergrade DataObservationRecentHistoricListParamsCat = "intergrade"
)

func (r DataObservationRecentHistoricListParamsCat) IsKnown() bool {
	switch r {
	case DataObservationRecentHistoricListParamsCatSpecies, DataObservationRecentHistoricListParamsCatSlash, DataObservationRecentHistoricListParamsCatIssf, DataObservationRecentHistoricListParamsCatSpuh, DataObservationRecentHistoricListParamsCatHybrid, DataObservationRecentHistoricListParamsCatDomestic, DataObservationRecentHistoricListParamsCatForm, DataObservationRecentHistoricListParamsCatIntergrade:
		return true
	}
	return false
}

// Include a subset (simple), or all (full), of the fields available.
type DataObservationRecentHistoricListParamsDetail string

const (
	DataObservationRecentHistoricListParamsDetailSimple DataObservationRecentHistoricListParamsDetail = "simple"
	DataObservationRecentHistoricListParamsDetailFull   DataObservationRecentHistoricListParamsDetail = "full"
)

func (r DataObservationRecentHistoricListParamsDetail) IsKnown() bool {
	switch r {
	case DataObservationRecentHistoricListParamsDetailSimple, DataObservationRecentHistoricListParamsDetailFull:
		return true
	}
	return false
}

// Include latest observation of the day, or the first added
type DataObservationRecentHistoricListParamsRank string

const (
	DataObservationRecentHistoricListParamsRankMrec   DataObservationRecentHistoricListParamsRank = "mrec"
	DataObservationRecentHistoricListParamsRankCreate DataObservationRecentHistoricListParamsRank = "create"
)

func (r DataObservationRecentHistoricListParamsRank) IsKnown() bool {
	switch r {
	case DataObservationRecentHistoricListParamsRankMrec, DataObservationRecentHistoricListParamsRankCreate:
		return true
	}
	return false
}
