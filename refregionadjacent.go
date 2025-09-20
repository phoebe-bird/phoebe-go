// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/phoebe-bird/phoebe-go/internal/apijson"
	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
)

// RefRegionAdjacentService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefRegionAdjacentService] method instead.
type RefRegionAdjacentService struct {
	Options []option.RequestOption
}

// NewRefRegionAdjacentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRefRegionAdjacentService(opts ...option.RequestOption) (r *RefRegionAdjacentService) {
	r = &RefRegionAdjacentService{}
	r.Options = opts
	return
}

// Get the list of countries or regions that share a border with this one. ####
// Notes Only subnational2 codes in the United States, New Zealand, or Mexico are
// currently supported
func (r *RefRegionAdjacentService) List(ctx context.Context, regionCode string, opts ...option.RequestOption) (res *[]RefRegionAdjacentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("ref/adjacent/%s", regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RefRegionAdjacentListResponse struct {
	Code string                            `json:"code"`
	Name string                            `json:"name"`
	JSON refRegionAdjacentListResponseJSON `json:"-"`
}

// refRegionAdjacentListResponseJSON contains the JSON metadata for the struct
// [RefRegionAdjacentListResponse]
type refRegionAdjacentListResponseJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RefRegionAdjacentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refRegionAdjacentListResponseJSON) RawJSON() string {
	return r.raw
}
