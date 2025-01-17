// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// PluginsNextboxUISavedtopologiesListReader is a Reader for the PluginsNextboxUISavedtopologiesList structure.
type PluginsNextboxUISavedtopologiesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsNextboxUISavedtopologiesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsNextboxUISavedtopologiesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsNextboxUISavedtopologiesListOK creates a PluginsNextboxUISavedtopologiesListOK with default headers values
func NewPluginsNextboxUISavedtopologiesListOK() *PluginsNextboxUISavedtopologiesListOK {
	return &PluginsNextboxUISavedtopologiesListOK{}
}

/* PluginsNextboxUISavedtopologiesListOK describes a response with status code 200, with default header values.

PluginsNextboxUISavedtopologiesListOK plugins nextbox Ui savedtopologies list o k
*/
type PluginsNextboxUISavedtopologiesListOK struct {
	Payload *PluginsNextboxUISavedtopologiesListOKBody
}

func (o *PluginsNextboxUISavedtopologiesListOK) Error() string {
	return fmt.Sprintf("[GET /plugins/nextbox-ui/savedtopologies/][%d] pluginsNextboxUiSavedtopologiesListOK  %+v", 200, o.Payload)
}
func (o *PluginsNextboxUISavedtopologiesListOK) GetPayload() *PluginsNextboxUISavedtopologiesListOKBody {
	return o.Payload
}

func (o *PluginsNextboxUISavedtopologiesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PluginsNextboxUISavedtopologiesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PluginsNextboxUISavedtopologiesListOKBody plugins nextbox UI savedtopologies list o k body
swagger:model PluginsNextboxUISavedtopologiesListOKBody
*/
type PluginsNextboxUISavedtopologiesListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.SavedTopology `json:"results"`
}

// Validate validates this plugins nextbox UI savedtopologies list o k body
func (o *PluginsNextboxUISavedtopologiesListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PluginsNextboxUISavedtopologiesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("pluginsNextboxUiSavedtopologiesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *PluginsNextboxUISavedtopologiesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("pluginsNextboxUiSavedtopologiesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PluginsNextboxUISavedtopologiesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("pluginsNextboxUiSavedtopologiesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PluginsNextboxUISavedtopologiesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("pluginsNextboxUiSavedtopologiesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pluginsNextboxUiSavedtopologiesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this plugins nextbox UI savedtopologies list o k body based on the context it is used
func (o *PluginsNextboxUISavedtopologiesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PluginsNextboxUISavedtopologiesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pluginsNextboxUiSavedtopologiesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PluginsNextboxUISavedtopologiesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PluginsNextboxUISavedtopologiesListOKBody) UnmarshalBinary(b []byte) error {
	var res PluginsNextboxUISavedtopologiesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
