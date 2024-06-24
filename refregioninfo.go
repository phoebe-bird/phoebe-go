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

// RefRegionInfoService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefRegionInfoService] method instead.
type RefRegionInfoService struct {
	Options []option.RequestOption
}

// NewRefRegionInfoService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRefRegionInfoService(opts ...option.RequestOption) (r *RefRegionInfoService) {
	r = &RefRegionInfoService{}
	r.Options = opts
	return
}

// Get information on the name and geographical area covered by a region.
//
// #### Notes
//
// Taking Madison County, New York, USA (location code US-NY-053) as an example the
// various values for the regionNameFormat query parameter work as follows:
//
// | Value          | Description                                | Result                           |
// | -------------- | ------------------------------------------ | -------------------------------- |
// | detailed       | return a detailed description              | Madison County, New York, US     |
// | detailednoqual | return the name to the subnational1 level  | Madison, New York                |
// | full           | return the full description                | Madison, New York, United States |
// | namequal       | return the qualified name                  | Madison County                   |
// | nameonly       | return only the name of the region         | Madison                          |
// | revdetailed    | return the detailed description in reverse | US, New York, Madison County     |
func (r *RefRegionInfoService) Get(ctx context.Context, regionCode string, query RefRegionInfoGetParams, opts ...option.RequestOption) (res *RefRegionInfoGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("ref/region/info/%s", regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RefRegionInfoGetResponse struct {
	Bounds RefRegionInfoGetResponseBounds `json:"bounds"`
	Result string                         `json:"result"`
	JSON   refRegionInfoGetResponseJSON   `json:"-"`
}

// refRegionInfoGetResponseJSON contains the JSON metadata for the struct
// [RefRegionInfoGetResponse]
type refRegionInfoGetResponseJSON struct {
	Bounds      apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RefRegionInfoGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refRegionInfoGetResponseJSON) RawJSON() string {
	return r.raw
}

type RefRegionInfoGetResponseBounds struct {
	MaxX float64                            `json:"maxX"`
	MaxY float64                            `json:"maxY"`
	MinX float64                            `json:"minX"`
	MinY float64                            `json:"minY"`
	JSON refRegionInfoGetResponseBoundsJSON `json:"-"`
}

// refRegionInfoGetResponseBoundsJSON contains the JSON metadata for the struct
// [RefRegionInfoGetResponseBounds]
type refRegionInfoGetResponseBoundsJSON struct {
	MaxX        apijson.Field
	MaxY        apijson.Field
	MinX        apijson.Field
	MinY        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RefRegionInfoGetResponseBounds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refRegionInfoGetResponseBoundsJSON) RawJSON() string {
	return r.raw
}

type RefRegionInfoGetParams struct {
	// The characters used to separate elements in the name.
	Delim param.Field[string] `query:"delim"`
	// Control how the name is displayed.
	RegionNameFormat param.Field[RefRegionInfoGetParamsRegionNameFormat] `query:"regionNameFormat"`
}

// URLQuery serializes [RefRegionInfoGetParams]'s query parameters as `url.Values`.
func (r RefRegionInfoGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Control how the name is displayed.
type RefRegionInfoGetParamsRegionNameFormat string

const (
	RefRegionInfoGetParamsRegionNameFormatDetailed       RefRegionInfoGetParamsRegionNameFormat = "detailed"
	RefRegionInfoGetParamsRegionNameFormatDetailednoqual RefRegionInfoGetParamsRegionNameFormat = "detailednoqual"
	RefRegionInfoGetParamsRegionNameFormatFull           RefRegionInfoGetParamsRegionNameFormat = "full"
	RefRegionInfoGetParamsRegionNameFormatNamequal       RefRegionInfoGetParamsRegionNameFormat = "namequal"
	RefRegionInfoGetParamsRegionNameFormatNameonly       RefRegionInfoGetParamsRegionNameFormat = "nameonly"
	RefRegionInfoGetParamsRegionNameFormatRevdetailed    RefRegionInfoGetParamsRegionNameFormat = "revdetailed"
)

func (r RefRegionInfoGetParamsRegionNameFormat) IsKnown() bool {
	switch r {
	case RefRegionInfoGetParamsRegionNameFormatDetailed, RefRegionInfoGetParamsRegionNameFormatDetailednoqual, RefRegionInfoGetParamsRegionNameFormatFull, RefRegionInfoGetParamsRegionNameFormatNamequal, RefRegionInfoGetParamsRegionNameFormatNameonly, RefRegionInfoGetParamsRegionNameFormatRevdetailed:
		return true
	}
	return false
}
