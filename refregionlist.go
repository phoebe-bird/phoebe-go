// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/phoebe-bird/phoebe-go/internal/apijson"
	"github.com/phoebe-bird/phoebe-go/internal/apiquery"
	"github.com/phoebe-bird/phoebe-go/internal/param"
	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
)

// RefRegionListService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefRegionListService] method instead.
type RefRegionListService struct {
	Options []option.RequestOption
}

// NewRefRegionListService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRefRegionListService(opts ...option.RequestOption) (r *RefRegionListService) {
	r = &RefRegionListService{}
	r.Options = opts
	return
}

// Get the list of sub-regions for a given country or region. #### Notes Not all
// combinations of region type and region code are valid. You can fetch all the
// subnational1 or subnational2 regions for a country however you can only specify
// a region type of 'country' when using 'world' as a region code.
func (r *RefRegionListService) List(ctx context.Context, regionType string, parentRegionCode string, query RefRegionListListParams, opts ...option.RequestOption) (res *[]RefRegionListListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if regionType == "" {
		err = errors.New("missing required regionType parameter")
		return
	}
	if parentRegionCode == "" {
		err = errors.New("missing required parentRegionCode parameter")
		return
	}
	path := fmt.Sprintf("ref/region/list/%s/%s", regionType, parentRegionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RefRegionListListResponse struct {
	Code string                        `json:"code"`
	Name string                        `json:"name"`
	JSON refRegionListListResponseJSON `json:"-"`
}

// refRegionListListResponseJSON contains the JSON metadata for the struct
// [RefRegionListListResponse]
type refRegionListListResponseJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RefRegionListListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refRegionListListResponseJSON) RawJSON() string {
	return r.raw
}

type RefRegionListListParams struct {
	// Fetch the records in CSV or JSON format.
	Fmt param.Field[RefRegionListListParamsFmt] `query:"fmt"`
}

// URLQuery serializes [RefRegionListListParams]'s query parameters as
// `url.Values`.
func (r RefRegionListListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Fetch the records in CSV or JSON format.
type RefRegionListListParamsFmt string

const (
	RefRegionListListParamsFmtCsv  RefRegionListListParamsFmt = "csv"
	RefRegionListListParamsFmtJson RefRegionListListParamsFmt = "json"
)

func (r RefRegionListListParamsFmt) IsKnown() bool {
	switch r {
	case RefRegionListListParamsFmtCsv, RefRegionListListParamsFmtJson:
		return true
	}
	return false
}
