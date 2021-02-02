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

// Backend Backend
//
// HAProxy backend configuration
//
// swagger:model backend
type Backend struct {

	// abortonclose
	// Enum: [enabled disabled]
	Abortonclose string `json:"abortonclose,omitempty"`

	// adv check
	// Enum: [ssl-hello-chk smtpchk ldap-check mysql-check pgsql-check tcp-check redis-check]
	AdvCheck string `json:"adv_check,omitempty"`

	// allbackups
	// Enum: [enabled disabled]
	Allbackups string `json:"allbackups,omitempty"`

	// balance
	Balance *Balance `json:"balance,omitempty"`

	// bind process
	// Pattern: ^[^\s]+$
	BindProcess string `json:"bind_process,omitempty"`

	// check timeout
	CheckTimeout *int64 `json:"check_timeout,omitempty"`

	// connect timeout
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`

	// client timeout
	ClientTimeout *int64 `json:"client_timeout,omitempty"`

	// cookie
	Cookie *Cookie `json:"cookie,omitempty"`

	// default server
	DefaultServer *DefaultServer `json:"default_server,omitempty"`

	// external check
	// Enum: [enabled disabled]
	ExternalCheck string `json:"external_check,omitempty"`

	// external check command
	// Pattern: ^[^\s]+$
	ExternalCheckCommand string `json:"external_check_command,omitempty"`

	// external check path
	// Pattern: ^[^\s]+$
	ExternalCheckPath string `json:"external_check_path,omitempty"`

	// forwardfor
	Forwardfor *Forwardfor `json:"forwardfor,omitempty"`

	// hash type
	HashType *BackendHashType `json:"hash_type,omitempty"`

	// http buffer request
	// Enum: [enabled disabled]
	HTTPBufferRequest string `json:"http-buffer-request,omitempty"`

	// http check
	HTTPCheck *HTTPCheck `json:"http-check,omitempty"`

	// http use htx
	// Pattern: ^[^\s]+$
	// Enum: [enabled disabled]
	HTTPUseHtx string `json:"http-use-htx,omitempty"`

	// http connection mode
	// Enum: [httpclose http-server-close http-keep-alive]
	HTTPConnectionMode string `json:"http_connection_mode,omitempty"`

	// http keep alive timeout
	HTTPKeepAliveTimeout *int64 `json:"http_keep_alive_timeout,omitempty"`

	// http pretend keepalive
	// Enum: [enabled disabled]
	HTTPPretendKeepalive string `json:"http_pretend_keepalive,omitempty"`

	// http request timeout
	HTTPRequestTimeout *int64 `json:"http_request_timeout,omitempty"`

	// http reuse
	// Enum: [aggressive always never safe]
	HTTPReuse string `json:"http_reuse,omitempty"`

	// httpchk
	Httpchk *Httpchk `json:"httpchk,omitempty"`

	// log tag
	// Pattern: ^[^\s]+$
	LogTag string `json:"log_tag,omitempty"`

	// mode
	// Enum: [http tcp]
	Mode string `json:"mode,omitempty"`

	// name
	// Required: true
	// Pattern: ^[A-Za-z0-9-_.:]+$
	Name string `json:"name"`

	// queue timeout
	QueueTimeout *int64 `json:"queue_timeout,omitempty"`

	// redispatch
	Redispatch *Redispatch `json:"redispatch,omitempty"`

	// retries
	Retries *int64 `json:"retries,omitempty"`

	// server timeout
	ServerTimeout *int64 `json:"server_timeout,omitempty"`

	// stats options
	StatsOptions *StatsOptions `json:"stats_options,omitempty"`

	// stick table
	StickTable *BackendStickTable `json:"stick_table,omitempty"`

	// tunnel timeout
	TunnelTimeout *int64 `json:"tunnel_timeout,omitempty"`
}

// Validate validates this backend
func (m *Backend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbortonclose(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdvCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllbackups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBindProcess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCookie(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalCheckCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalCheckPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForwardfor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHashType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPBufferRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPUseHtx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPConnectionMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPPretendKeepalive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPReuse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHttpchk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedispatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatsOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStickTable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var backendTypeAbortonclosePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeAbortonclosePropEnum = append(backendTypeAbortonclosePropEnum, v)
	}
}

const (

	// BackendAbortoncloseEnabled captures enum value "enabled"
	BackendAbortoncloseEnabled string = "enabled"

	// BackendAbortoncloseDisabled captures enum value "disabled"
	BackendAbortoncloseDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateAbortoncloseEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeAbortonclosePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateAbortonclose(formats strfmt.Registry) error {

	if swag.IsZero(m.Abortonclose) { // not required
		return nil
	}

	// value enum
	if err := m.validateAbortoncloseEnum("abortonclose", "body", m.Abortonclose); err != nil {
		return err
	}

	return nil
}

var backendTypeAdvCheckPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ssl-hello-chk","smtpchk","ldap-check","mysql-check","pgsql-check","tcp-check","redis-check"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeAdvCheckPropEnum = append(backendTypeAdvCheckPropEnum, v)
	}
}

const (

	// BackendAdvCheckSslHelloChk captures enum value "ssl-hello-chk"
	BackendAdvCheckSslHelloChk string = "ssl-hello-chk"

	// BackendAdvCheckSmtpchk captures enum value "smtpchk"
	BackendAdvCheckSmtpchk string = "smtpchk"

	// BackendAdvCheckLdapCheck captures enum value "ldap-check"
	BackendAdvCheckLdapCheck string = "ldap-check"

	// BackendAdvCheckMysqlCheck captures enum value "mysql-check"
	BackendAdvCheckMysqlCheck string = "mysql-check"

	// BackendAdvCheckPgsqlCheck captures enum value "pgsql-check"
	BackendAdvCheckPgsqlCheck string = "pgsql-check"

	// BackendAdvCheckTCPCheck captures enum value "tcp-check"
	BackendAdvCheckTCPCheck string = "tcp-check"

	// BackendAdvCheckRedisCheck captures enum value "redis-check"
	BackendAdvCheckRedisCheck string = "redis-check"
)

// prop value enum
func (m *Backend) validateAdvCheckEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeAdvCheckPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateAdvCheck(formats strfmt.Registry) error {

	if swag.IsZero(m.AdvCheck) { // not required
		return nil
	}

	// value enum
	if err := m.validateAdvCheckEnum("adv_check", "body", m.AdvCheck); err != nil {
		return err
	}

	return nil
}

var backendTypeAllbackupsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeAllbackupsPropEnum = append(backendTypeAllbackupsPropEnum, v)
	}
}

const (

	// BackendAllbackupsEnabled captures enum value "enabled"
	BackendAllbackupsEnabled string = "enabled"

	// BackendAllbackupsDisabled captures enum value "disabled"
	BackendAllbackupsDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateAllbackupsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeAllbackupsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateAllbackups(formats strfmt.Registry) error {

	if swag.IsZero(m.Allbackups) { // not required
		return nil
	}

	// value enum
	if err := m.validateAllbackupsEnum("allbackups", "body", m.Allbackups); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateBalance(formats strfmt.Registry) error {

	if swag.IsZero(m.Balance) { // not required
		return nil
	}

	if m.Balance != nil {
		if err := m.Balance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balance")
			}
			return err
		}
	}

	return nil
}

func (m *Backend) validateBindProcess(formats strfmt.Registry) error {

	if swag.IsZero(m.BindProcess) { // not required
		return nil
	}

	if err := validate.Pattern("bind_process", "body", string(m.BindProcess), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateCookie(formats strfmt.Registry) error {

	if swag.IsZero(m.Cookie) { // not required
		return nil
	}

	if m.Cookie != nil {
		if err := m.Cookie.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cookie")
			}
			return err
		}
	}

	return nil
}

func (m *Backend) validateDefaultServer(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultServer) { // not required
		return nil
	}

	if m.DefaultServer != nil {
		if err := m.DefaultServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_server")
			}
			return err
		}
	}

	return nil
}

var backendTypeExternalCheckPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeExternalCheckPropEnum = append(backendTypeExternalCheckPropEnum, v)
	}
}

const (

	// BackendExternalCheckEnabled captures enum value "enabled"
	BackendExternalCheckEnabled string = "enabled"

	// BackendExternalCheckDisabled captures enum value "disabled"
	BackendExternalCheckDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateExternalCheckEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeExternalCheckPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateExternalCheck(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalCheck) { // not required
		return nil
	}

	// value enum
	if err := m.validateExternalCheckEnum("external_check", "body", m.ExternalCheck); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateExternalCheckCommand(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalCheckCommand) { // not required
		return nil
	}

	if err := validate.Pattern("external_check_command", "body", string(m.ExternalCheckCommand), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateExternalCheckPath(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalCheckPath) { // not required
		return nil
	}

	if err := validate.Pattern("external_check_path", "body", string(m.ExternalCheckPath), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateForwardfor(formats strfmt.Registry) error {

	if swag.IsZero(m.Forwardfor) { // not required
		return nil
	}

	if m.Forwardfor != nil {
		if err := m.Forwardfor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forwardfor")
			}
			return err
		}
	}

	return nil
}

func (m *Backend) validateHashType(formats strfmt.Registry) error {

	if swag.IsZero(m.HashType) { // not required
		return nil
	}

	if m.HashType != nil {
		if err := m.HashType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hash_type")
			}
			return err
		}
	}

	return nil
}

var backendTypeHTTPBufferRequestPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHTTPBufferRequestPropEnum = append(backendTypeHTTPBufferRequestPropEnum, v)
	}
}

const (

	// BackendHTTPBufferRequestEnabled captures enum value "enabled"
	BackendHTTPBufferRequestEnabled string = "enabled"

	// BackendHTTPBufferRequestDisabled captures enum value "disabled"
	BackendHTTPBufferRequestDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateHTTPBufferRequestEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeHTTPBufferRequestPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateHTTPBufferRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPBufferRequest) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPBufferRequestEnum("http-buffer-request", "body", m.HTTPBufferRequest); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateHTTPCheck(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPCheck) { // not required
		return nil
	}

	if m.HTTPCheck != nil {
		if err := m.HTTPCheck.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http-check")
			}
			return err
		}
	}

	return nil
}

var backendTypeHTTPUseHtxPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHTTPUseHtxPropEnum = append(backendTypeHTTPUseHtxPropEnum, v)
	}
}

const (

	// BackendHTTPUseHtxEnabled captures enum value "enabled"
	BackendHTTPUseHtxEnabled string = "enabled"

	// BackendHTTPUseHtxDisabled captures enum value "disabled"
	BackendHTTPUseHtxDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateHTTPUseHtxEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeHTTPUseHtxPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateHTTPUseHtx(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPUseHtx) { // not required
		return nil
	}

	if err := validate.Pattern("http-use-htx", "body", string(m.HTTPUseHtx), `^[^\s]+$`); err != nil {
		return err
	}

	// value enum
	if err := m.validateHTTPUseHtxEnum("http-use-htx", "body", m.HTTPUseHtx); err != nil {
		return err
	}

	return nil
}

var backendTypeHTTPConnectionModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["httpclose","http-server-close","http-keep-alive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHTTPConnectionModePropEnum = append(backendTypeHTTPConnectionModePropEnum, v)
	}
}

const (

	// BackendHTTPConnectionModeHttpclose captures enum value "httpclose"
	BackendHTTPConnectionModeHttpclose string = "httpclose"

	// BackendHTTPConnectionModeHTTPServerClose captures enum value "http-server-close"
	BackendHTTPConnectionModeHTTPServerClose string = "http-server-close"

	// BackendHTTPConnectionModeHTTPKeepAlive captures enum value "http-keep-alive"
	BackendHTTPConnectionModeHTTPKeepAlive string = "http-keep-alive"
)

// prop value enum
func (m *Backend) validateHTTPConnectionModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeHTTPConnectionModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateHTTPConnectionMode(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPConnectionMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPConnectionModeEnum("http_connection_mode", "body", m.HTTPConnectionMode); err != nil {
		return err
	}

	return nil
}

var backendTypeHTTPPretendKeepalivePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHTTPPretendKeepalivePropEnum = append(backendTypeHTTPPretendKeepalivePropEnum, v)
	}
}

const (

	// BackendHTTPPretendKeepaliveEnabled captures enum value "enabled"
	BackendHTTPPretendKeepaliveEnabled string = "enabled"

	// BackendHTTPPretendKeepaliveDisabled captures enum value "disabled"
	BackendHTTPPretendKeepaliveDisabled string = "disabled"
)

// prop value enum
func (m *Backend) validateHTTPPretendKeepaliveEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeHTTPPretendKeepalivePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateHTTPPretendKeepalive(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPPretendKeepalive) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPPretendKeepaliveEnum("http_pretend_keepalive", "body", m.HTTPPretendKeepalive); err != nil {
		return err
	}

	return nil
}

var backendTypeHTTPReusePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aggressive","always","never","safe"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeHTTPReusePropEnum = append(backendTypeHTTPReusePropEnum, v)
	}
}

const (

	// BackendHTTPReuseAggressive captures enum value "aggressive"
	BackendHTTPReuseAggressive string = "aggressive"

	// BackendHTTPReuseAlways captures enum value "always"
	BackendHTTPReuseAlways string = "always"

	// BackendHTTPReuseNever captures enum value "never"
	BackendHTTPReuseNever string = "never"

	// BackendHTTPReuseSafe captures enum value "safe"
	BackendHTTPReuseSafe string = "safe"
)

// prop value enum
func (m *Backend) validateHTTPReuseEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeHTTPReusePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateHTTPReuse(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPReuse) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPReuseEnum("http_reuse", "body", m.HTTPReuse); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateHttpchk(formats strfmt.Registry) error {

	if swag.IsZero(m.Httpchk) { // not required
		return nil
	}

	if m.Httpchk != nil {
		if err := m.Httpchk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpchk")
			}
			return err
		}
	}

	return nil
}

func (m *Backend) validateLogTag(formats strfmt.Registry) error {

	if swag.IsZero(m.LogTag) { // not required
		return nil
	}

	if err := validate.Pattern("log_tag", "body", string(m.LogTag), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var backendTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","tcp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendTypeModePropEnum = append(backendTypeModePropEnum, v)
	}
}

const (

	// BackendModeHTTP captures enum value "http"
	BackendModeHTTP string = "http"

	// BackendModeTCP captures enum value "tcp"
	BackendModeTCP string = "tcp"
)

// prop value enum
func (m *Backend) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Backend) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Backend) validateRedispatch(formats strfmt.Registry) error {

	if swag.IsZero(m.Redispatch) { // not required
		return nil
	}

	if m.Redispatch != nil {
		if err := m.Redispatch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redispatch")
			}
			return err
		}
	}

	return nil
}

func (m *Backend) validateStatsOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.StatsOptions) { // not required
		return nil
	}

	if m.StatsOptions != nil {
		if err := m.StatsOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats_options")
			}
			return err
		}
	}

	return nil
}

func (m *Backend) validateStickTable(formats strfmt.Registry) error {

	if swag.IsZero(m.StickTable) { // not required
		return nil
	}

	if m.StickTable != nil {
		if err := m.StickTable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stick_table")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Backend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Backend) UnmarshalBinary(b []byte) error {
	var res Backend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BackendHashType backend hash type
//
// swagger:model BackendHashType
type BackendHashType struct {

	// function
	// Enum: [sdbm djb2 wt6 crc32]
	Function string `json:"function,omitempty"`

	// method
	// Enum: [map-based consistent]
	Method string `json:"method,omitempty"`

	// modifier
	// Enum: [avalanche]
	Modifier string `json:"modifier,omitempty"`
}

// Validate validates this backend hash type
func (m *BackendHashType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFunction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var backendHashTypeTypeFunctionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sdbm","djb2","wt6","crc32"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendHashTypeTypeFunctionPropEnum = append(backendHashTypeTypeFunctionPropEnum, v)
	}
}

const (

	// BackendHashTypeFunctionSdbm captures enum value "sdbm"
	BackendHashTypeFunctionSdbm string = "sdbm"

	// BackendHashTypeFunctionDjb2 captures enum value "djb2"
	BackendHashTypeFunctionDjb2 string = "djb2"

	// BackendHashTypeFunctionWt6 captures enum value "wt6"
	BackendHashTypeFunctionWt6 string = "wt6"

	// BackendHashTypeFunctionCrc32 captures enum value "crc32"
	BackendHashTypeFunctionCrc32 string = "crc32"
)

// prop value enum
func (m *BackendHashType) validateFunctionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendHashTypeTypeFunctionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BackendHashType) validateFunction(formats strfmt.Registry) error {

	if swag.IsZero(m.Function) { // not required
		return nil
	}

	// value enum
	if err := m.validateFunctionEnum("hash_type"+"."+"function", "body", m.Function); err != nil {
		return err
	}

	return nil
}

var backendHashTypeTypeMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["map-based","consistent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendHashTypeTypeMethodPropEnum = append(backendHashTypeTypeMethodPropEnum, v)
	}
}

const (

	// BackendHashTypeMethodMapBased captures enum value "map-based"
	BackendHashTypeMethodMapBased string = "map-based"

	// BackendHashTypeMethodConsistent captures enum value "consistent"
	BackendHashTypeMethodConsistent string = "consistent"
)

// prop value enum
func (m *BackendHashType) validateMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendHashTypeTypeMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BackendHashType) validateMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.Method) { // not required
		return nil
	}

	// value enum
	if err := m.validateMethodEnum("hash_type"+"."+"method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

var backendHashTypeTypeModifierPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["avalanche"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendHashTypeTypeModifierPropEnum = append(backendHashTypeTypeModifierPropEnum, v)
	}
}

const (

	// BackendHashTypeModifierAvalanche captures enum value "avalanche"
	BackendHashTypeModifierAvalanche string = "avalanche"
)

// prop value enum
func (m *BackendHashType) validateModifierEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendHashTypeTypeModifierPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BackendHashType) validateModifier(formats strfmt.Registry) error {

	if swag.IsZero(m.Modifier) { // not required
		return nil
	}

	// value enum
	if err := m.validateModifierEnum("hash_type"+"."+"modifier", "body", m.Modifier); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackendHashType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackendHashType) UnmarshalBinary(b []byte) error {
	var res BackendHashType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BackendStickTable backend stick table
//
// swagger:model BackendStickTable
type BackendStickTable struct {

	// expire
	Expire *int64 `json:"expire,omitempty"`

	// keylen
	Keylen *int64 `json:"keylen,omitempty"`

	// nopurge
	Nopurge bool `json:"nopurge,omitempty"`

	// peers
	// Pattern: ^[^\s]+$
	Peers string `json:"peers,omitempty"`

	// size
	Size *int64 `json:"size,omitempty"`

	// store
	// Pattern: ^[^\s]+$
	Store string `json:"store,omitempty"`

	// type
	// Enum: [ip ipv6 integer string binary]
	Type string `json:"type,omitempty"`
}

// Validate validates this backend stick table
func (m *BackendStickTable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStore(formats); err != nil {
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

func (m *BackendStickTable) validatePeers(formats strfmt.Registry) error {

	if swag.IsZero(m.Peers) { // not required
		return nil
	}

	if err := validate.Pattern("stick_table"+"."+"peers", "body", string(m.Peers), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *BackendStickTable) validateStore(formats strfmt.Registry) error {

	if swag.IsZero(m.Store) { // not required
		return nil
	}

	if err := validate.Pattern("stick_table"+"."+"store", "body", string(m.Store), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var backendStickTableTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ip","ipv6","integer","string","binary"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendStickTableTypeTypePropEnum = append(backendStickTableTypeTypePropEnum, v)
	}
}

const (

	// BackendStickTableTypeIP captures enum value "ip"
	BackendStickTableTypeIP string = "ip"

	// BackendStickTableTypeIPV6 captures enum value "ipv6"
	BackendStickTableTypeIPV6 string = "ipv6"

	// BackendStickTableTypeInteger captures enum value "integer"
	BackendStickTableTypeInteger string = "integer"

	// BackendStickTableTypeString captures enum value "string"
	BackendStickTableTypeString string = "string"

	// BackendStickTableTypeBinary captures enum value "binary"
	BackendStickTableTypeBinary string = "binary"
)

// prop value enum
func (m *BackendStickTable) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendStickTableTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BackendStickTable) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("stick_table"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackendStickTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackendStickTable) UnmarshalBinary(b []byte) error {
	var res BackendStickTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
