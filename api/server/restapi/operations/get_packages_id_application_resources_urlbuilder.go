// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetPackagesIDApplicationResourcesURL generates an URL for the get packages ID application resources operation
type GetPackagesIDApplicationResourcesURL struct {
	ID string

	Page                                        int64
	PageSize                                    int64
	ReportingSBOMAnalyzersContainElements       []string
	ReportingSBOMAnalyzersDoesntContainElements []string
	ResourceHashContains                        []string
	ResourceHashEnd                             *string
	ResourceHashIsNot                           []string
	ResourceHashIs                              []string
	ResourceHashStart                           *string
	ResourceNameContains                        []string
	ResourceNameEnd                             *string
	ResourceNameIsNot                           []string
	ResourceNameIs                              []string
	ResourceNameStart                           *string
	SortDir                                     *string
	SortKey                                     string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetPackagesIDApplicationResourcesURL) WithBasePath(bp string) *GetPackagesIDApplicationResourcesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetPackagesIDApplicationResourcesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetPackagesIDApplicationResourcesURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/packages/{id}/applicationResources"

	id := o.ID
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("id is required on GetPackagesIDApplicationResourcesURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	pageQ := swag.FormatInt64(o.Page)
	if pageQ != "" {
		qs.Set("page", pageQ)
	}

	pageSizeQ := swag.FormatInt64(o.PageSize)
	if pageSizeQ != "" {
		qs.Set("pageSize", pageSizeQ)
	}

	var reportingSBOMAnalyzersContainElementsIR []string
	for _, reportingSBOMAnalyzersContainElementsI := range o.ReportingSBOMAnalyzersContainElements {
		reportingSBOMAnalyzersContainElementsIS := reportingSBOMAnalyzersContainElementsI
		if reportingSBOMAnalyzersContainElementsIS != "" {
			reportingSBOMAnalyzersContainElementsIR = append(reportingSBOMAnalyzersContainElementsIR, reportingSBOMAnalyzersContainElementsIS)
		}
	}

	reportingSBOMAnalyzersContainElements := swag.JoinByFormat(reportingSBOMAnalyzersContainElementsIR, "")

	if len(reportingSBOMAnalyzersContainElements) > 0 {
		qsv := reportingSBOMAnalyzersContainElements[0]
		if qsv != "" {
			qs.Set("reportingSBOMAnalyzers[containElements]", qsv)
		}
	}

	var reportingSBOMAnalyzersDoesntContainElementsIR []string
	for _, reportingSBOMAnalyzersDoesntContainElementsI := range o.ReportingSBOMAnalyzersDoesntContainElements {
		reportingSBOMAnalyzersDoesntContainElementsIS := reportingSBOMAnalyzersDoesntContainElementsI
		if reportingSBOMAnalyzersDoesntContainElementsIS != "" {
			reportingSBOMAnalyzersDoesntContainElementsIR = append(reportingSBOMAnalyzersDoesntContainElementsIR, reportingSBOMAnalyzersDoesntContainElementsIS)
		}
	}

	reportingSBOMAnalyzersDoesntContainElements := swag.JoinByFormat(reportingSBOMAnalyzersDoesntContainElementsIR, "")

	if len(reportingSBOMAnalyzersDoesntContainElements) > 0 {
		qsv := reportingSBOMAnalyzersDoesntContainElements[0]
		if qsv != "" {
			qs.Set("reportingSBOMAnalyzers[doesntContainElements]", qsv)
		}
	}

	var resourceHashContainsIR []string
	for _, resourceHashContainsI := range o.ResourceHashContains {
		resourceHashContainsIS := resourceHashContainsI
		if resourceHashContainsIS != "" {
			resourceHashContainsIR = append(resourceHashContainsIR, resourceHashContainsIS)
		}
	}

	resourceHashContains := swag.JoinByFormat(resourceHashContainsIR, "")

	if len(resourceHashContains) > 0 {
		qsv := resourceHashContains[0]
		if qsv != "" {
			qs.Set("resourceHash[contains]", qsv)
		}
	}

	var resourceHashEndQ string
	if o.ResourceHashEnd != nil {
		resourceHashEndQ = *o.ResourceHashEnd
	}
	if resourceHashEndQ != "" {
		qs.Set("resourceHash[end]", resourceHashEndQ)
	}

	var resourceHashIsNotIR []string
	for _, resourceHashIsNotI := range o.ResourceHashIsNot {
		resourceHashIsNotIS := resourceHashIsNotI
		if resourceHashIsNotIS != "" {
			resourceHashIsNotIR = append(resourceHashIsNotIR, resourceHashIsNotIS)
		}
	}

	resourceHashIsNot := swag.JoinByFormat(resourceHashIsNotIR, "")

	if len(resourceHashIsNot) > 0 {
		qsv := resourceHashIsNot[0]
		if qsv != "" {
			qs.Set("resourceHash[isNot]", qsv)
		}
	}

	var resourceHashIsIR []string
	for _, resourceHashIsI := range o.ResourceHashIs {
		resourceHashIsIS := resourceHashIsI
		if resourceHashIsIS != "" {
			resourceHashIsIR = append(resourceHashIsIR, resourceHashIsIS)
		}
	}

	resourceHashIs := swag.JoinByFormat(resourceHashIsIR, "")

	if len(resourceHashIs) > 0 {
		qsv := resourceHashIs[0]
		if qsv != "" {
			qs.Set("resourceHash[is]", qsv)
		}
	}

	var resourceHashStartQ string
	if o.ResourceHashStart != nil {
		resourceHashStartQ = *o.ResourceHashStart
	}
	if resourceHashStartQ != "" {
		qs.Set("resourceHash[start]", resourceHashStartQ)
	}

	var resourceNameContainsIR []string
	for _, resourceNameContainsI := range o.ResourceNameContains {
		resourceNameContainsIS := resourceNameContainsI
		if resourceNameContainsIS != "" {
			resourceNameContainsIR = append(resourceNameContainsIR, resourceNameContainsIS)
		}
	}

	resourceNameContains := swag.JoinByFormat(resourceNameContainsIR, "")

	if len(resourceNameContains) > 0 {
		qsv := resourceNameContains[0]
		if qsv != "" {
			qs.Set("resourceName[contains]", qsv)
		}
	}

	var resourceNameEndQ string
	if o.ResourceNameEnd != nil {
		resourceNameEndQ = *o.ResourceNameEnd
	}
	if resourceNameEndQ != "" {
		qs.Set("resourceName[end]", resourceNameEndQ)
	}

	var resourceNameIsNotIR []string
	for _, resourceNameIsNotI := range o.ResourceNameIsNot {
		resourceNameIsNotIS := resourceNameIsNotI
		if resourceNameIsNotIS != "" {
			resourceNameIsNotIR = append(resourceNameIsNotIR, resourceNameIsNotIS)
		}
	}

	resourceNameIsNot := swag.JoinByFormat(resourceNameIsNotIR, "")

	if len(resourceNameIsNot) > 0 {
		qsv := resourceNameIsNot[0]
		if qsv != "" {
			qs.Set("resourceName[isNot]", qsv)
		}
	}

	var resourceNameIsIR []string
	for _, resourceNameIsI := range o.ResourceNameIs {
		resourceNameIsIS := resourceNameIsI
		if resourceNameIsIS != "" {
			resourceNameIsIR = append(resourceNameIsIR, resourceNameIsIS)
		}
	}

	resourceNameIs := swag.JoinByFormat(resourceNameIsIR, "")

	if len(resourceNameIs) > 0 {
		qsv := resourceNameIs[0]
		if qsv != "" {
			qs.Set("resourceName[is]", qsv)
		}
	}

	var resourceNameStartQ string
	if o.ResourceNameStart != nil {
		resourceNameStartQ = *o.ResourceNameStart
	}
	if resourceNameStartQ != "" {
		qs.Set("resourceName[start]", resourceNameStartQ)
	}

	var sortDirQ string
	if o.SortDir != nil {
		sortDirQ = *o.SortDir
	}
	if sortDirQ != "" {
		qs.Set("sortDir", sortDirQ)
	}

	sortKeyQ := o.SortKey
	if sortKeyQ != "" {
		qs.Set("sortKey", sortKeyQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetPackagesIDApplicationResourcesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetPackagesIDApplicationResourcesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetPackagesIDApplicationResourcesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetPackagesIDApplicationResourcesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetPackagesIDApplicationResourcesURL")
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
func (o *GetPackagesIDApplicationResourcesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
