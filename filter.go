// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Filter Filter
//
// HAProxy filters
//
// swagger:model filter
type Filter struct {

	// cache name
	// Pattern: ^[^\s]+$
	CacheName string `json:"cache_name,omitempty"`

	// index
	// Required: true
	Index *int64 `json:"index"`

	// spoe config
	// Pattern: ^[^\s]+$
	SpoeConfig string `json:"spoe_config,omitempty"`

	// spoe engine
	// Pattern: ^[^\s]+$
	SpoeEngine string `json:"spoe_engine,omitempty"`

	// trace hexdump
	TraceHexdump bool `json:"trace_hexdump,omitempty"`

	// trace name
	// Pattern: ^[^\s]+$
	TraceName string `json:"trace_name,omitempty"`

	// trace rnd forwarding
	TraceRndForwarding bool `json:"trace_rnd_forwarding,omitempty"`

	// trace rnd parsing
	TraceRndParsing bool `json:"trace_rnd_parsing,omitempty"`

	// type
	// Required: true
	// Enum: [trace compression spoe cache]
	Type string `json:"type"`
}

// Validate validates this filter
func (m *Filter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCacheName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpoeConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpoeEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Filter) validateCacheName(formats strfmt.Registry) error {

	if swag.IsZero(m.CacheName) { // not required
		return nil
	}

	if err := validate.Pattern("cache_name", "body", string(m.CacheName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Filter) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *Filter) validateSpoeConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.SpoeConfig) { // not required
		return nil
	}

	if err := validate.Pattern("spoe_config", "body", string(m.SpoeConfig), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Filter) validateSpoeEngine(formats strfmt.Registry) error {

	if swag.IsZero(m.SpoeEngine) { // not required
		return nil
	}

	if err := validate.Pattern("spoe_engine", "body", string(m.SpoeEngine), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Filter) validateTraceName(formats strfmt.Registry) error {

	if swag.IsZero(m.TraceName) { // not required
		return nil
	}

	if err := validate.Pattern("trace_name", "body", string(m.TraceName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var filterTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["trace","compression","spoe","cache"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		filterTypeTypePropEnum = append(filterTypeTypePropEnum, v)
	}
}

const (

	// FilterTypeTrace captures enum value "trace"
	FilterTypeTrace string = "trace"

	// FilterTypeCompression captures enum value "compression"
	FilterTypeCompression string = "compression"

	// FilterTypeSpoe captures enum value "spoe"
	FilterTypeSpoe string = "spoe"

	// FilterTypeCache captures enum value "cache"
	FilterTypeCache string = "cache"
)

// prop value enum
func (m *Filter) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, filterTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Filter) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Filter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Filter) UnmarshalBinary(b []byte) error {
	var res Filter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
