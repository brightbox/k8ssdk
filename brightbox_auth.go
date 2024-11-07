// Copyright 2020 Brightbox Systems Ltd
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

package k8ssdk

import (
	"context"
	"fmt"
	"os"

	brightbox "github.com/brightbox/gobrightbox/v2"
	"github.com/brightbox/gobrightbox/v2/clientcredentials"
	"github.com/brightbox/gobrightbox/v2/endpoint"
	"github.com/brightbox/gobrightbox/v2/passwordcredentials"
	"github.com/hashicorp/go-cleanhttp"
	"golang.org/x/oauth2"
	klog "k8s.io/klog/v2"
)

const (
	defaultClientID     = "app-dkmch"
	defaultClientSecret = "uogoelzgt0nwawb"
	clientEnvVar        = "BRIGHTBOX_CLIENT"
	clientSecretEnvVar  = "BRIGHTBOX_CLIENT_SECRET"
	usernameEnvVar      = "BRIGHTBOX_USER_NAME"
	passwordEnvVar      = "BRIGHTBOX_PASSWORD"
	accountEnvVar       = "BRIGHTBOX_ACCOUNT"
	apiURLEnvVar        = "BRIGHTBOX_API_URL"

	defaultTimeoutSeconds = 10

	// ValidAcmeDomainStatus indicates a valid acme state
	ValidAcmeDomainStatus = "valid"
)

type authdetails struct {
	APIClient string
	APISecret string
	UserName  string
	password  string
	Account   string
	APIURL    string
}

// obtainCloudClient creates a new Brightbox client using details from
// the environment
func obtainCloudClient() (CloudAccess, error) {
	klog.V(4).Infof("obtainCloudClient")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return configureClient(
		ctx,
		authdetails{
			APIClient: getenvWithDefault(clientEnvVar,
				defaultClientID),
			APISecret: getenvWithDefault(clientSecretEnvVar,
				defaultClientSecret),
			UserName: os.Getenv(usernameEnvVar),
			password: os.Getenv(passwordEnvVar),
			Account:  os.Getenv(accountEnvVar),
			APIURL:   os.Getenv(apiURLEnvVar),
		},
	)
}

// Validate account config entries
func validateConfig(authd authdetails) error {
	klog.V(4).Infof("validateConfig")
	if authd.APIClient == defaultClientID &&
		authd.APISecret == defaultClientSecret {
		if authd.Account == "" {
			return fmt.Errorf("must specify Account with User Credentials")
		}
	} else {
		if authd.UserName != "" || authd.password != "" {
			return fmt.Errorf("User Credentials not used with API Client")
		}
	}
	return nil
}

func configureClient(ctx context.Context, authd authdetails) (CloudAccess, error) {
	klog.V(4).Infof("configureClient")
	if err := validateConfig(authd); err != nil {
		return nil, err
	}

	return authenticatedClient(ctx, authd)
}

// Authenticate the details and return a client
func authenticatedClient(authCtx context.Context, authd authdetails) (CloudAccess, error) {
	klog.V(4).Infof("configureClient")
	refreshContext := contextWithLoggedHTTPClient(context.Background())

	client, err := brightbox.Connect(refreshContext, confFromAuthd(authd))
	if err != nil {
		return nil, err
	}

	if authd.Account == "" {
		klog.V(4).Infof("Obtaining default account")
		accounts, err := client.Accounts(authCtx)
		if err != nil {
			return nil, err
		}
		authd.Account = accounts[0].ID
		klog.V(4).Infof("default account is %v", authd.Account)
	}
	return client, nil
}

func confFromAuthd(authd authdetails) brightbox.Oauth2 {
	if authd.UserName != "" || authd.password != "" {
		return &passwordcredentials.Config{
			UserName: authd.UserName,
			Password: authd.password,
			ID:       authd.APIClient,
			Secret:   authd.APISecret,
			Config: endpoint.Config{
				BaseURL: authd.APIURL,
				Account: authd.Account,
				Scopes:  endpoint.FullScope,
			},
		}
	}
	return &clientcredentials.Config{
		ID:     authd.APIClient,
		Secret: authd.APISecret,
		Config: endpoint.Config{
			BaseURL: authd.APIURL,
			Scopes:  endpoint.FullScope,
		},
	}
}

func contextWithLoggedHTTPClient(ctx context.Context) context.Context {
	client := cleanhttp.DefaultClient()
	return context.WithValue(ctx, oauth2.HTTPClient, client)
}
