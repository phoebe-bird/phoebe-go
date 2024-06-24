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

// RefGeoAdjacentService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefGeoAdjacentService] method instead.
type RefGeoAdjacentService struct {
	Options []option.RequestOption
}

// NewRefGeoAdjacentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRefGeoAdjacentService(opts ...option.RequestOption) (r *RefGeoAdjacentService) {
	r = &RefGeoAdjacentService{}
	r.Options = opts
	return
}

// Get the list of countries or regions that share a border with this one. ####
// Notes Only subnational2 codes in the United States, New Zealand, or Mexico are
// currently supported
func (r *RefGeoAdjacentService) Get(ctx context.Context, regionCode string, opts ...option.RequestOption) (res *RefGeoAdjacentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("ref/adjacent/%s", regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RefGeoAdjacentGetResponse struct {
	AdjacentRegions []RefGeoAdjacentGetResponseAdjacentRegion `json:"adjacentRegions"`
	JSON            refGeoAdjacentGetResponseJSON             `json:"-"`
}

// refGeoAdjacentGetResponseJSON contains the JSON metadata for the struct
// [RefGeoAdjacentGetResponse]
type refGeoAdjacentGetResponseJSON struct {
	AdjacentRegions apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RefGeoAdjacentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refGeoAdjacentGetResponseJSON) RawJSON() string {
	return r.raw
}

type RefGeoAdjacentGetResponseAdjacentRegion struct {
	Code string                                      `json:"code"`
	Name string                                      `json:"name"`
	JSON refGeoAdjacentGetResponseAdjacentRegionJSON `json:"-"`
}

// refGeoAdjacentGetResponseAdjacentRegionJSON contains the JSON metadata for the
// struct [RefGeoAdjacentGetResponseAdjacentRegion]
type refGeoAdjacentGetResponseAdjacentRegionJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RefGeoAdjacentGetResponseAdjacentRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refGeoAdjacentGetResponseAdjacentRegionJSON) RawJSON() string {
	return r.raw
}
