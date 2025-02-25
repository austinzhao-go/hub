// Copyright © 2021 The Tekton Authors.
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

package resource

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/hub/api/pkg/testutils"
	"github.com/tektoncd/hub/api/v1/gen/resource"
)

func TestQuery_ByTags(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.QueryPayload{Name: "", Kinds: []string{}, Tags: []string{"atag"}, Limit: 100}
	all, err := resourceSvc.Query(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(all.Data))
}

func TestQuery_ByPlatforms(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.QueryPayload{Name: "", Kinds: []string{}, Platforms: []string{"linux/amd64"}, Limit: 100}
	all, err := resourceSvc.Query(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(all.Data))
}

func TestQuery_ByNameAndKind(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.QueryPayload{Name: "build", Kinds: []string{"pipeline"}, Limit: 100}
	all, err := resourceSvc.Query(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(all.Data))
	assert.Equal(t, "build-pipeline", all.Data[0].Name)
}

func TestQuery_ByCategories(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.QueryPayload{Name: "", Kinds: []string{}, Categories: []string{"abc"}, Limit: 100}
	all, err := resourceSvc.Query(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(all.Data))
	assert.Equal(t, "build-pipeline", all.Data[0].Name)
}

func TestQuery_NotFoundError(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.QueryPayload{Name: "foo", Kinds: []string{}, Limit: 100}
	_, err := resourceSvc.Query(context.Background(), payload)
	assert.Error(t, err)
	assert.EqualError(t, err, "resource not found")
}

func TestList_ByLimit(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ListPayload{Limit: 3}
	all, err := resourceSvc.List(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(all.Data))
	assert.Equal(t, "tekton", all.Data[0].Name)
}

func TestVersionsByID(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.VersionsByIDPayload{ID: 1}
	res, err := resourceSvc.VersionsByID(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(res.Data.Versions))
	assert.Equal(t, "0.2", res.Data.Latest.Version)
}

func TestVersionsByID_NotFoundError(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.VersionsByIDPayload{ID: 11}
	_, err := resourceSvc.VersionsByID(context.Background(), payload)
	assert.Error(t, err)
	assert.EqualError(t, err, "resource not found")
}

func TestByCatalogKindNameVersion(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByCatalogKindNameVersionPayload{Catalog: "catalog-official", Kind: "task", Name: "tkn", Version: "0.1"}
	res, err := resourceSvc.ByCatalogKindNameVersion(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, "0.1", res.Data.Version)
}

func TestByCatalogKindNameVersionReadme(t *testing.T) {
	os.Setenv("CLONE_BASE_PATH", "testdata/catalog")
	defer os.Unsetenv("CLONE_BASE_PATH")

	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByCatalogKindNameVersionReadmePayload{Catalog: "catalog-official", Kind: "task", Name: "tkn", Version: "0.1"}
	res, err := resourceSvc.ByCatalogKindNameVersionReadme(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, *res.Data.Readme, "# This works\n")
}

func TestByCatalogKindNameVersionYaml(t *testing.T) {
	os.Setenv("CLONE_BASE_PATH", "testdata/catalog")
	defer os.Unsetenv("CLONE_BASE_PATH")

	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByCatalogKindNameVersionYamlPayload{Catalog: "catalog-official", Kind: "task", Name: "tkn", Version: "0.1"}
	res, err := resourceSvc.ByCatalogKindNameVersionYaml(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, *res.Data.Yaml, "Hub: works\n")
}

func TestByCatalogKindNameVersion_NoResourceWithName(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByCatalogKindNameVersionPayload{Catalog: "catalog-official", Kind: "task", Name: "foo", Version: "0.1"}
	_, err := resourceSvc.ByCatalogKindNameVersion(context.Background(), payload)
	assert.Error(t, err)
	assert.EqualError(t, err, "resource not found")
}

func TestByVersionID(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByVersionIDPayload{VersionID: 6}
	res, err := resourceSvc.ByVersionID(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, "0.1.1", res.Data.Version)
}

func TestByVersionID_NotFoundError(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByVersionIDPayload{VersionID: 111}
	_, err := resourceSvc.ByVersionID(context.Background(), payload)
	assert.Error(t, err)
	assert.EqualError(t, err, "resource not found")
}

func TestByCatalogKindName(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByCatalogKindNamePayload{Catalog: "catalog-community", Kind: "task", Name: "img"}
	res, err := resourceSvc.ByCatalogKindName(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, "img", res.Data.Name)
}

func TestByEnterpriseCatalogKindName(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByCatalogKindNamePayload{Catalog: "catalog-enterprise", Kind: "task", Name: "tkn-enterprise"}
	res, err := resourceSvc.ByCatalogKindName(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, "tkn-enterprise", res.Data.Name)
}

func TestByCatalogKindNameIfCompatible(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	version := "0.12.3"
	payload := &resource.ByCatalogKindNamePayload{Catalog: "catalog-official", Kind: "task", Name: "tekton", Pipelinesversion: &version}
	res, err := resourceSvc.ByCatalogKindName(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, "tekton", res.Data.Name)
	assert.Equal(t, 1, len(res.Data.Versions))
	assert.Equal(t, "0.1", res.Data.Versions[0].Version)
}

func TestByCatalogKindName_CompatibleVersionNotFound(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	version := "0.11.0"
	payload := &resource.ByCatalogKindNamePayload{Catalog: "catalog-official", Kind: "task", Name: "tekton", Pipelinesversion: &version}
	_, err := resourceSvc.ByCatalogKindName(context.Background(), payload)
	assert.Error(t, err)
	assert.EqualError(t, err, "resource not found compatible with minPipelinesVersion")
}

func TestByCatalogKindName_ResourceNotFoundError(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByCatalogKindNamePayload{Catalog: "catalog-community", Kind: "task", Name: "foo"}
	_, err := resourceSvc.ByCatalogKindName(context.Background(), payload)
	assert.Error(t, err)
	assert.EqualError(t, err, "resource not found")
}

func TestByID(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByIDPayload{ID: 1}
	res, err := resourceSvc.ByID(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, "tekton", res.Data.Name)
}

func TestByID_NotFoundError(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByIDPayload{ID: 77}
	_, err := resourceSvc.ByID(context.Background(), payload)
	assert.Error(t, err)
	assert.EqualError(t, err, "resource not found")
}

func TestCreationRawURL(t *testing.T) {
	url := "https://ghe.myhost.com/org/repo/tree/main/task/name/0.1/name.yaml"
	replacer := getStringReplacer(url, "github")
	rawUrl := replacer.Replace(url)
	expected := "https://raw.ghe.myhost.com/org/repo/main/task/name/0.1/name.yaml"
	assert.Equal(t, expected, rawUrl)
}

func TestDeprecationByVersionID(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByVersionIDPayload{VersionID: 10}
	res, err := resourceSvc.ByVersionID(context.Background(), payload)
	assert.NoError(t, err)

	assert.Equal(t, true, *res.Data.Deprecated)
}

func TestLatestVersionDeprecationByID(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByIDPayload{ID: 7}
	res, err := resourceSvc.ByID(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, true, *res.Data.LatestVersion.Deprecated)
}

func TestCreationRawURLBitbucket(t *testing.T) {
	url := "https://bitbucket.org/org/catalog/src/main/task/name/0.1/name.yaml"
	replacer := getStringReplacer(url, "bitbucket")
	rawUrl := replacer.Replace(url)
	expected := "https://bitbucket.org/org/catalog/raw/main/task/name/0.1/name.yaml"
	assert.Equal(t, expected, rawUrl)
}

func TestCreationRawURLGitlab(t *testing.T) {
	url := "https://gitlab.com/org/catalog/-/blob/main/task/name/0.1/name.yaml"
	replacer := getStringReplacer(url, "gitlab")
	rawUrl := replacer.Replace(url)
	expected := "https://gitlab.com/org/catalog/-/raw/main/task/name/0.1/name.yaml"
	assert.Equal(t, expected, rawUrl)
}

func TestCreationRawURLGitlabEnterprise(t *testing.T) {
	url := "https://gitlab.myhost.com/org/catalog/-/blob/main/task/name/0.1/name.yaml"
	replacer := getStringReplacer(url, "gitlab")
	rawUrl := replacer.Replace(url)
	expected := "https://gitlab.myhost.com/org/catalog/-/raw/main/task/name/0.1/name.yaml"
	assert.Equal(t, expected, rawUrl)
}

func TestGetYamlByCatalogKindNameVersion(t *testing.T) {
	os.Setenv("CLONE_BASE_PATH", "testdata/catalog")
	defer os.Unsetenv("CLONE_BASE_PATH")

	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceService := New(tc)
	payload := &resource.GetRawYamlByCatalogKindNameVersionPayload{Catalog: "catalog-official", Kind: "task", Name: "tkn", Version: "0.1"}
	yamlFile, err := resourceService.GetRawYamlByCatalogKindNameVersion(context.Background(), payload)

	assert.NoError(t, err)
	assert.NotNil(t, yamlFile)
}
