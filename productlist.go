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

// ProductListService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductListService] method instead.
type ProductListService struct {
	Options    []option.RequestOption
	Historical *ProductListHistoricalService
}

// NewProductListService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProductListService(opts ...option.RequestOption) (r *ProductListService) {
	r = &ProductListService{}
	r.Options = opts
	r.Historical = NewProductListHistoricalService(opts...)
	return
}

// Get information on the most recently submitted checklists for a region.
func (r *ProductListService) Get(ctx context.Context, regionCode string, query ProductListGetParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("product/lists/%s", regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type ProductListGetParams struct {
	// Only fetch this number of checklists.
	MaxResults param.Field[int64] `query:"maxResults"`
}

// URLQuery serializes [ProductListGetParams]'s query parameters as `url.Values`.
func (r ProductListGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
