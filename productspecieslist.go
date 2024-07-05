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

// ProductSpeciesListService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductSpeciesListService] method instead.
type ProductSpeciesListService struct {
	Options []option.RequestOption
}

// NewProductSpeciesListService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProductSpeciesListService(opts ...option.RequestOption) (r *ProductSpeciesListService) {
	r = &ProductSpeciesListService{}
	r.Options = opts
	return
}

// Get a list of species codes ever seen in a region, in taxonomic order (species
// taxa only)
//
// #### Notes The results are usually updated every 10 seconds for locations, every day for larger regions.
func (r *ProductSpeciesListService) List(ctx context.Context, regionCode string, opts ...option.RequestOption) (res *[]string, err error) {
	opts = append(r.Options[:], opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("product/spplist/%s", regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
