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

// ProductSpecieService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductSpecieService] method instead.
type ProductSpecieService struct {
	Options []option.RequestOption
}

// NewProductSpecieService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProductSpecieService(opts ...option.RequestOption) (r *ProductSpecieService) {
	r = &ProductSpecieService{}
	r.Options = opts
	return
}

// Get a list of species codes ever seen in a region, in taxonomic order (species
// taxa only)
//
// #### Notes The results are usually updated every 10 seconds for locations, every day for larger regions.
func (r *ProductSpecieService) List(ctx context.Context, regionCode string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("product/spplist/%s", regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}
