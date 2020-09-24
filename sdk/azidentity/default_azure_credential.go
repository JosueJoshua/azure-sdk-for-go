// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

const (
	developerSignOnClientID = "04b07795-8ddb-461a-bbee-02f9e1bf7b46"
)

// DefaultAzureCredentialOptions contains options for configuring how credentials are acquired.
type DefaultAzureCredentialOptions struct {
	// set this field to true in order to exclude the AzureCLICredential from the set of
	// credentials that will be used to authenticate with
	ExcludeAzureCLICredential bool
	// set this field to true in order to exclude the EnvironmentCredential from the set of
	// credentials that will be used to authenticate with
	ExcludeEnvironmentCredential bool
	// set this field to true in order to exclude the ManagedIdentityCredential from the set of
	// credentials that will be used to authenticate with
	ExcludeMSICredential bool
}

// NewDefaultAzureCredential provides a default TokenCredential chain configuration for applications that will be deployed to Azure.  Credentials that can authenticate
// for the following cases will be tried, in the following order:
// - Environment Credential
// - Managed Identity Credential
// - Azure CLI Credential
// Consult the documentation for these credential types for more information on how they attempt authentication.
func NewDefaultAzureCredential(options *DefaultAzureCredentialOptions) (azcore.Credential, error) {
	var creds []azcore.TokenCredential
	errMsg := ""

	if options == nil {
		options = &DefaultAzureCredentialOptions{}
	}

	if !options.ExcludeEnvironmentCredential {
		envCred, err := NewEnvironmentCredential(nil)
		if err == nil {
			// we know NewEnvironmentCredential will return a credential that implements the azcore.TokenCredential interface
			creds = append(creds, envCred.(azcore.TokenCredential))
		} else {
			errMsg += err.Error()
		}
	}

	if !options.ExcludeMSICredential {
		msiCred, err := NewManagedIdentityCredential("", nil)
		if err == nil {
			// we know managedIdentityCredential implements the azcore.TokenCredential interface
			creds = append(creds, msiCred.(azcore.TokenCredential))
		} else {
			errMsg += err.Error()
		}
	}

	if !options.ExcludeAzureCLICredential {
		cliCred, err := NewAzureCLICredential(nil)
		if err == nil {
			// we know AzureCLICredential implements the azcore.TokenCredential interface
			creds = append(creds, cliCred.(azcore.TokenCredential))
		} else {
			errMsg += err.Error()
		}
	}

	// if no credentials are added to the slice of TokenCredentials then return a CredentialUnavailableError
	if len(creds) == 0 {
		err := &CredentialUnavailableError{CredentialType: "Default Azure Credential", Message: errMsg}
		logCredentialError(err.CredentialType, err)
		return nil, err
	}
	azcore.Log().Write(LogCredential, "Azure Identity => NewDefaultAzureCredential() invoking NewChainedTokenCredential()")
	return NewChainedTokenCredential(creds...)
}
