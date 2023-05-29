// Code generated by go-swagger; DO NOT EDIT.

package sample_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetV1SchemaNameSampleCsvURL generates an URL for the get v1 schema name sample csv operation
type GetV1SchemaNameSampleCsvURL struct {
	SchemaName string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetV1SchemaNameSampleCsvURL) WithBasePath(bp string) *GetV1SchemaNameSampleCsvURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetV1SchemaNameSampleCsvURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetV1SchemaNameSampleCsvURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v1/{schemaName}/sample-csv"

	schemaName := o.SchemaName
	if schemaName != "" {
		_path = strings.Replace(_path, "{schemaName}", schemaName, -1)
	} else {
		return nil, errors.New("schemaName is required on GetV1SchemaNameSampleCsvURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetV1SchemaNameSampleCsvURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetV1SchemaNameSampleCsvURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetV1SchemaNameSampleCsvURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetV1SchemaNameSampleCsvURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetV1SchemaNameSampleCsvURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetV1SchemaNameSampleCsvURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
