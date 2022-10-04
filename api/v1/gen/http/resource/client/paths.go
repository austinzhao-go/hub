// Code generated by goa v3.9.1, DO NOT EDIT.
//
// HTTP request path constructors for the resource service.
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package client

import (
	"fmt"
)

// QueryResourcePath returns the URL path to the resource service Query HTTP endpoint.
func QueryResourcePath() string {
	return "/v1/query"
}

// ListResourcePath returns the URL path to the resource service List HTTP endpoint.
func ListResourcePath() string {
	return "/v1/resources"
}

// VersionsByIDResourcePath returns the URL path to the resource service VersionsByID HTTP endpoint.
func VersionsByIDResourcePath(id uint) string {
	return fmt.Sprintf("/v1/resource/%v/versions", id)
}

// ByCatalogKindNameVersionResourcePath returns the URL path to the resource service ByCatalogKindNameVersion HTTP endpoint.
func ByCatalogKindNameVersionResourcePath(catalog string, kind string, name string, version string) string {
	return fmt.Sprintf("/v1/resource/%v/%v/%v/%v", catalog, kind, name, version)
}

// ByCatalogKindNameVersionReadmeResourcePath returns the URL path to the resource service ByCatalogKindNameVersionReadme HTTP endpoint.
func ByCatalogKindNameVersionReadmeResourcePath(catalog string, kind string, name string, version string) string {
	return fmt.Sprintf("/v1/resource/%v/%v/%v/%v/readme", catalog, kind, name, version)
}

// ByCatalogKindNameVersionYamlResourcePath returns the URL path to the resource service ByCatalogKindNameVersionYaml HTTP endpoint.
func ByCatalogKindNameVersionYamlResourcePath(catalog string, kind string, name string, version string) string {
	return fmt.Sprintf("/v1/resource/%v/%v/%v/%v/yaml", catalog, kind, name, version)
}

// ByVersionIDResourcePath returns the URL path to the resource service ByVersionId HTTP endpoint.
func ByVersionIDResourcePath(versionID uint) string {
	return fmt.Sprintf("/v1/resource/version/%v", versionID)
}

// ByCatalogKindNameResourcePath returns the URL path to the resource service ByCatalogKindName HTTP endpoint.
func ByCatalogKindNameResourcePath(catalog string, kind string, name string) string {
	return fmt.Sprintf("/v1/resource/%v/%v/%v", catalog, kind, name)
}

// ByIDResourcePath returns the URL path to the resource service ById HTTP endpoint.
func ByIDResourcePath(id uint) string {
	return fmt.Sprintf("/v1/resource/%v", id)
}

// GetRawYamlByCatalogKindNameVersionResourcePath returns the URL path to the resource service GetRawYamlByCatalogKindNameVersion HTTP endpoint.
func GetRawYamlByCatalogKindNameVersionResourcePath(catalog string, kind string, name string, version string) string {
	return fmt.Sprintf("/v1/resource/%v/%v/%v/%v/raw", catalog, kind, name, version)
}
