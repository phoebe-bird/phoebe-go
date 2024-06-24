// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/phoebe-go/internal/apijson"
	"github.com/stainless-sdks/phoebe-go/internal/apiquery"
	"github.com/stainless-sdks/phoebe-go/internal/param"
	"github.com/stainless-sdks/phoebe-go/internal/requestconfig"
	"github.com/stainless-sdks/phoebe-go/option"
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
func (r *RefRegionListService) Get(ctx context.Context, regionType string, parentRegionCode string, query RefRegionListGetParams, opts ...option.RequestOption) (res *[]RefRegionListGetResponse, err error) {
	opts = append(r.Options[:], opts...)
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

type RefRegionListGetResponse struct {
	Code string                       `json:"code"`
	Name string                       `json:"name"`
	JSON refRegionListGetResponseJSON `json:"-"`
}

// refRegionListGetResponseJSON contains the JSON metadata for the struct
// [RefRegionListGetResponse]
type refRegionListGetResponseJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RefRegionListGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refRegionListGetResponseJSON) RawJSON() string {
	return r.raw
}

type RefRegionListGetParams struct {
	// Fetch the records in CSV or JSON format.
	Fmt param.Field[RefRegionListGetParamsFmt] `query:"fmt"`
}

// URLQuery serializes [RefRegionListGetParams]'s query parameters as `url.Values`.
func (r RefRegionListGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Fetch the records in CSV or JSON format.
type RefRegionListGetParamsFmt string

const (
	RefRegionListGetParamsFmtCsv  RefRegionListGetParamsFmt = "csv"
	RefRegionListGetParamsFmtJson RefRegionListGetParamsFmt = "json"
)

func (r RefRegionListGetParamsFmt) IsKnown() bool {
	switch r {
	case RefRegionListGetParamsFmtCsv, RefRegionListGetParamsFmtJson:
		return true
	}
	return false
}
