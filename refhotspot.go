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

// RefHotspotService contains methods and other services that help with interacting
// with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefHotspotService] method instead.
type RefHotspotService struct {
	Options []option.RequestOption
	Geo     *RefHotspotGeoService
	Info    *RefHotspotInfoService
}

// NewRefHotspotService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRefHotspotService(opts ...option.RequestOption) (r *RefHotspotService) {
	r = &RefHotspotService{}
	r.Options = opts
	r.Geo = NewRefHotspotGeoService(opts...)
	r.Info = NewRefHotspotInfoService(opts...)
	return
}

// Hotspots in a region
func (r *RefHotspotService) List(ctx context.Context, regionCode string, query RefHotspotListParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("ref/hotspot/%s", regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type RefHotspotListParams struct {
	// The number of days back to fetch hotspots.
	Back param.Field[int64] `query:"back"`
	// Fetch the records in CSV or JSON format.
	Fmt param.Field[RefHotspotListParamsFmt] `query:"fmt"`
}

// URLQuery serializes [RefHotspotListParams]'s query parameters as `url.Values`.
func (r RefHotspotListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Fetch the records in CSV or JSON format.
type RefHotspotListParamsFmt string

const (
	RefHotspotListParamsFmtCsv  RefHotspotListParamsFmt = "csv"
	RefHotspotListParamsFmtJson RefHotspotListParamsFmt = "json"
)

func (r RefHotspotListParamsFmt) IsKnown() bool {
	switch r {
	case RefHotspotListParamsFmtCsv, RefHotspotListParamsFmtJson:
		return true
	}
	return false
}
