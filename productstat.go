// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/phoebe-go/internal/apijson"
	"github.com/stainless-sdks/phoebe-go/internal/requestconfig"
	"github.com/stainless-sdks/phoebe-go/option"
)

// ProductStatService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductStatService] method instead.
type ProductStatService struct {
	Options []option.RequestOption
}

// NewProductStatService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProductStatService(opts ...option.RequestOption) (r *ProductStatService) {
	r = &ProductStatService{}
	r.Options = opts
	return
}

// Get a summary of the number of checklist submitted, species seen and
// contributors on a given date for a country or region.
//
// #### Notes The results are updated every 15 minutes.
func (r *ProductStatService) Get(ctx context.Context, regionCode string, y int64, m int64, d int64, opts ...option.RequestOption) (res *ProductStatGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("product/stats/%s/%v/%v/%v", regionCode, y, m, d)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ProductStatGetResponse struct {
	NumChecklists   int64                      `json:"numChecklists"`
	NumContributors int64                      `json:"numContributors"`
	NumSpecies      int64                      `json:"numSpecies"`
	JSON            productStatGetResponseJSON `json:"-"`
}

// productStatGetResponseJSON contains the JSON metadata for the struct
// [ProductStatGetResponse]
type productStatGetResponseJSON struct {
	NumChecklists   apijson.Field
	NumContributors apijson.Field
	NumSpecies      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProductStatGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productStatGetResponseJSON) RawJSON() string {
	return r.raw
}
