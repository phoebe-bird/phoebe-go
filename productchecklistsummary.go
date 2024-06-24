// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/phoebe-go/internal/requestconfig"
	"github.com/stainless-sdks/phoebe-go/option"
)

// ProductChecklistSummaryService contains methods and other services that help
// with interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductChecklistSummaryService] method instead.
type ProductChecklistSummaryService struct {
	Options []option.RequestOption
}

// NewProductChecklistSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProductChecklistSummaryService(opts ...option.RequestOption) (r *ProductChecklistSummaryService) {
	r = &ProductChecklistSummaryService{}
	r.Options = opts
	return
}

// Get a summary of the number of checklist submitted, species seen and
// contributors on a given date for a country or region.
//
// #### Notes The results are updated every 15 minutes.
func (r *ProductChecklistSummaryService) Get(ctx context.Context, regionCode string, y int64, m int64, d int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("product/stats/%s/%v/%v/%v", regionCode, y, m, d)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}
