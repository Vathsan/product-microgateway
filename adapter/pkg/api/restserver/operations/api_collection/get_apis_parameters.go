// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package api_collection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetApisParams creates a new GetApisParams object
//
// There are no default values defined in the spec.
func NewGetApisParams() GetApisParams {

	return GetApisParams{}
}

// GetApisParams contains all the bound params for the get apis operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetApis
type GetApisParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Optional - Condition to filter APIs. Currently only filtering
	by API type (HTTP or WebSocket) is supported.
	"http" for HTTP type
	"ws" for WebSocket type

	  In: query
	*/
	APIType *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetApisParams() beforehand.
func (o *GetApisParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAPIType, qhkAPIType, _ := qs.GetOK("apiType")
	if err := o.bindAPIType(qAPIType, qhkAPIType, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAPIType binds and validates parameter APIType from query.
func (o *GetApisParams) bindAPIType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.APIType = &raw

	return nil
}
