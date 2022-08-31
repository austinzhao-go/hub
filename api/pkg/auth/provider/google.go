/*
Copyright 2022 The Tekton Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.

You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"fmt"
	"os"

	"github.com/markbates/goth/providers/google"
)

func GoogleProvider(AUTH_URL string) provider {
	googleAuth := provider{
		ClientId:     os.Getenv("GG_CLIENT_ID"),
		ClientSecret: os.Getenv("GG_CLIENT_SECRET"),
		CallbackUrl:  fmt.Sprintf(AUTH_URL, "google"),
		AuthUrl:      google.Endpoint.AuthURL,
		TokenUrl:     google.Endpoint.TokenURL,
	}
	return googleAuth
}