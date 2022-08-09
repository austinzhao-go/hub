// Code generated by goa v3.8.2, DO NOT EDIT.
//
// catalog HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	catalog "github.com/tektoncd/hub/api/gen/catalog"
	catalogviews "github.com/tektoncd/hub/api/gen/catalog/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildRefreshRequest instantiates a HTTP request object with method and path
// set to call the "catalog" service "Refresh" endpoint
func (c *Client) BuildRefreshRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		catalogName string
	)
	{
		p, ok := v.(*catalog.RefreshPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("catalog", "Refresh", "*catalog.RefreshPayload", v)
		}
		catalogName = p.CatalogName
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RefreshCatalogPath(catalogName)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("catalog", "Refresh", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRefreshRequest returns an encoder for requests sent to the catalog
// Refresh server.
func EncodeRefreshRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*catalog.RefreshPayload)
		if !ok {
			return goahttp.ErrInvalidType("catalog", "Refresh", "*catalog.RefreshPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeRefreshResponse returns a decoder for responses returned by the
// catalog Refresh endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeRefreshResponse may return the following errors:
//	- "not-found" (type *goa.ServiceError): http.StatusNotFound
//	- "internal-error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeRefreshResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body RefreshResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("catalog", "Refresh", err)
			}
			p := NewRefreshJobOK(&body)
			view := "default"
			vres := &catalogviews.Job{Projected: p, View: view}
			if err = catalogviews.ValidateJob(vres); err != nil {
				return nil, goahttp.ErrValidationError("catalog", "Refresh", err)
			}
			res := catalog.NewJob(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body RefreshNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("catalog", "Refresh", err)
			}
			err = ValidateRefreshNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("catalog", "Refresh", err)
			}
			return nil, NewRefreshNotFound(&body)
		case http.StatusInternalServerError:
			var (
				body RefreshInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("catalog", "Refresh", err)
			}
			err = ValidateRefreshInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("catalog", "Refresh", err)
			}
			return nil, NewRefreshInternalError(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("catalog", "Refresh", resp.StatusCode, string(body))
		}
	}
}

// BuildRefreshAllRequest instantiates a HTTP request object with method and
// path set to call the "catalog" service "RefreshAll" endpoint
func (c *Client) BuildRefreshAllRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RefreshAllCatalogPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("catalog", "RefreshAll", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRefreshAllRequest returns an encoder for requests sent to the catalog
// RefreshAll server.
func EncodeRefreshAllRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*catalog.RefreshAllPayload)
		if !ok {
			return goahttp.ErrInvalidType("catalog", "RefreshAll", "*catalog.RefreshAllPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeRefreshAllResponse returns a decoder for responses returned by the
// catalog RefreshAll endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeRefreshAllResponse may return the following errors:
//	- "internal-error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeRefreshAllResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body RefreshAllResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("catalog", "RefreshAll", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateJobResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("catalog", "RefreshAll", err)
			}
			res := NewRefreshAllJobOK(body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body RefreshAllInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("catalog", "RefreshAll", err)
			}
			err = ValidateRefreshAllInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("catalog", "RefreshAll", err)
			}
			return nil, NewRefreshAllInternalError(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("catalog", "RefreshAll", resp.StatusCode, string(body))
		}
	}
}

// BuildCatalogErrorRequest instantiates a HTTP request object with method and
// path set to call the "catalog" service "CatalogError" endpoint
func (c *Client) BuildCatalogErrorRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		catalogName string
	)
	{
		p, ok := v.(*catalog.CatalogErrorPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("catalog", "CatalogError", "*catalog.CatalogErrorPayload", v)
		}
		catalogName = p.CatalogName
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CatalogErrorCatalogPath(catalogName)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("catalog", "CatalogError", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCatalogErrorRequest returns an encoder for requests sent to the
// catalog CatalogError server.
func EncodeCatalogErrorRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*catalog.CatalogErrorPayload)
		if !ok {
			return goahttp.ErrInvalidType("catalog", "CatalogError", "*catalog.CatalogErrorPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeCatalogErrorResponse returns a decoder for responses returned by the
// catalog CatalogError endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeCatalogErrorResponse may return the following errors:
//	- "internal-error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeCatalogErrorResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CatalogErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("catalog", "CatalogError", err)
			}
			err = ValidateCatalogErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("catalog", "CatalogError", err)
			}
			res := NewCatalogErrorResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body CatalogErrorInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("catalog", "CatalogError", err)
			}
			err = ValidateCatalogErrorInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("catalog", "CatalogError", err)
			}
			return nil, NewCatalogErrorInternalError(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("catalog", "CatalogError", resp.StatusCode, string(body))
		}
	}
}

// unmarshalJobResponseToCatalogJob builds a value of type *catalog.Job from a
// value of type *JobResponse.
func unmarshalJobResponseToCatalogJob(v *JobResponse) *catalog.Job {
	res := &catalog.Job{
		ID:          *v.ID,
		CatalogName: *v.CatalogName,
		Status:      *v.Status,
	}

	return res
}

// unmarshalCatalogErrorsResponseBodyToCatalogCatalogErrors builds a value of
// type *catalog.CatalogErrors from a value of type *CatalogErrorsResponseBody.
func unmarshalCatalogErrorsResponseBodyToCatalogCatalogErrors(v *CatalogErrorsResponseBody) *catalog.CatalogErrors {
	res := &catalog.CatalogErrors{
		Type: *v.Type,
	}
	res.Errors = make([]string, len(v.Errors))
	for i, val := range v.Errors {
		res.Errors[i] = val
	}

	return res
}
