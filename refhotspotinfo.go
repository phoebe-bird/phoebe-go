// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebebird

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/phoebe-bird/phoebe-go/internal/apijson"
	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
)

// RefHotspotInfoService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefHotspotInfoService] method instead.
type RefHotspotInfoService struct {
	Options []option.RequestOption
}

// NewRefHotspotInfoService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRefHotspotInfoService(opts ...option.RequestOption) (r *RefHotspotInfoService) {
	r = &RefHotspotInfoService{}
	r.Options = opts
	return
}

// Get information on the location of a hotspot. #### Notes This API call only
// works for hotspots. If you pass the location code for a private location or an
// invalid location code then an HTTP 410 (Gone) error is returned.
func (r *RefHotspotInfoService) Get(ctx context.Context, locID string, opts ...option.RequestOption) (res *RefHotspotInfoGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if locID == "" {
		err = errors.New("missing required locId parameter")
		return
	}
	path := fmt.Sprintf("ref/hotspot/info/%s", locID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RefHotspotInfoGetResponse struct {
	CountryCode      string                        `json:"countryCode"`
	CountryName      string                        `json:"countryName"`
	HierarchicalName string                        `json:"hierarchicalName"`
	IsHotspot        bool                          `json:"isHotspot"`
	Lat              float64                       `json:"lat"`
	Latitude         float64                       `json:"latitude"`
	Lng              float64                       `json:"lng"`
	LocID            string                        `json:"locId"`
	LocID            string                        `json:"locID"`
	LocName          string                        `json:"locName"`
	Longitude        float64                       `json:"longitude"`
	Name             string                        `json:"name"`
	Subnational1Code string                        `json:"subnational1Code"`
	Subnational1Name string                        `json:"subnational1Name"`
	JSON             refHotspotInfoGetResponseJSON `json:"-"`
}

// refHotspotInfoGetResponseJSON contains the JSON metadata for the struct
// [RefHotspotInfoGetResponse]
type refHotspotInfoGetResponseJSON struct {
	CountryCode      apijson.Field
	CountryName      apijson.Field
	HierarchicalName apijson.Field
	IsHotspot        apijson.Field
	Lat              apijson.Field
	Latitude         apijson.Field
	Lng              apijson.Field
	LocID            apijson.Field
	LocID            apijson.Field
	LocName          apijson.Field
	Longitude        apijson.Field
	Name             apijson.Field
	Subnational1Code apijson.Field
	Subnational1Name apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RefHotspotInfoGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refHotspotInfoGetResponseJSON) RawJSON() string {
	return r.raw
}
