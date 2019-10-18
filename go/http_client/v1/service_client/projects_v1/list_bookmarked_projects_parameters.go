// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package projects_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListBookmarkedProjectsParams creates a new ListBookmarkedProjectsParams object
// with the default values initialized.
func NewListBookmarkedProjectsParams() *ListBookmarkedProjectsParams {
	var ()
	return &ListBookmarkedProjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListBookmarkedProjectsParamsWithTimeout creates a new ListBookmarkedProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListBookmarkedProjectsParamsWithTimeout(timeout time.Duration) *ListBookmarkedProjectsParams {
	var ()
	return &ListBookmarkedProjectsParams{

		timeout: timeout,
	}
}

// NewListBookmarkedProjectsParamsWithContext creates a new ListBookmarkedProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListBookmarkedProjectsParamsWithContext(ctx context.Context) *ListBookmarkedProjectsParams {
	var ()
	return &ListBookmarkedProjectsParams{

		Context: ctx,
	}
}

// NewListBookmarkedProjectsParamsWithHTTPClient creates a new ListBookmarkedProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListBookmarkedProjectsParamsWithHTTPClient(client *http.Client) *ListBookmarkedProjectsParams {
	var ()
	return &ListBookmarkedProjectsParams{
		HTTPClient: client,
	}
}

/*ListBookmarkedProjectsParams contains all the parameters to send to the API endpoint
for the list bookmarked projects operation typically these are written to a http.Request
*/
type ListBookmarkedProjectsParams struct {

	/*Limit
	  Limit size.

	*/
	Limit *string
	/*Page
	  Pagination.

	*/
	Page *string
	/*Query
	  Query filter the search search.

	*/
	Query *string
	/*Sort
	  Sort to order the search.

	*/
	Sort *string
	/*User
	  User

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) WithTimeout(timeout time.Duration) *ListBookmarkedProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) WithContext(ctx context.Context) *ListBookmarkedProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) WithHTTPClient(client *http.Client) *ListBookmarkedProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) WithLimit(limit *string) *ListBookmarkedProjectsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithPage adds the page to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) WithPage(page *string) *ListBookmarkedProjectsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) SetPage(page *string) {
	o.Page = page
}

// WithQuery adds the query to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) WithQuery(query *string) *ListBookmarkedProjectsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) WithSort(sort *string) *ListBookmarkedProjectsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithUser adds the user to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) WithUser(user string) *ListBookmarkedProjectsParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *ListBookmarkedProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit string
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage string
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := qrPage
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Query != nil {

		// query param query
		var qrQuery string
		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {
			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}