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

// ProductChecklistService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductChecklistService] method instead.
type ProductChecklistService struct {
	Options []option.RequestOption
	Summary *ProductChecklistSummaryService
}

// NewProductChecklistService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProductChecklistService(opts ...option.RequestOption) (r *ProductChecklistService) {
	r = &ProductChecklistService{}
	r.Options = opts
	r.Summary = NewProductChecklistSummaryService(opts...)
	return
}

// Get information on the checklists submitted on a given date for a country or
// region.
func (r *ProductChecklistService) Get(ctx context.Context, regionCode string, y int64, m int64, d int64, query ProductChecklistGetParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("product/lists/%s/%v/%v/%v", regionCode, y, m, d)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type ProductChecklistGetParams struct {
	// Only fetch this number of checklists.
	MaxResults param.Field[int64] `query:"maxResults"`
	// Order the results by the date of the checklist or by the date it was submitted.
	SortKey param.Field[ProductChecklistGetParamsSortKey] `query:"sortKey"`
}

// URLQuery serializes [ProductChecklistGetParams]'s query parameters as
// `url.Values`.
func (r ProductChecklistGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Order the results by the date of the checklist or by the date it was submitted.
type ProductChecklistGetParamsSortKey string

const (
	ProductChecklistGetParamsSortKeyObsDt      ProductChecklistGetParamsSortKey = "obs_dt"
	ProductChecklistGetParamsSortKeyCreationDt ProductChecklistGetParamsSortKey = "creation_dt"
)

func (r ProductChecklistGetParamsSortKey) IsKnown() bool {
	switch r {
	case ProductChecklistGetParamsSortKeyObsDt, ProductChecklistGetParamsSortKeyCreationDt:
		return true
	}
	return false
}
