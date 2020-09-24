// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"context"
	"errors"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// chainedTokenCredential provides a TokenCredential implementation that chains multiple TokenCredential sources to be tried in order
// and returns the token from the first successful call to GetToken().
type chainedTokenCredential struct {
	sources []azcore.TokenCredential
}

// NewChainedTokenCredential creates a credential with the specified TokenCredential sources that will be tried in order
// until one of the GetToken calls returns an access token.
func NewChainedTokenCredential(sources ...azcore.TokenCredential) (azcore.Credential, error) {
	if len(sources) == 0 {
		credErr := &CredentialUnavailableError{CredentialType: "Chained Token Credential", Message: "Length of sources cannot be 0"}
		logCredentialError(credErr.CredentialType, credErr)
		return nil, credErr
	}
	for _, source := range sources {
		if source == nil { // cannot have a nil credential in the chain or else the application will panic when GetToken() is called on nil
			credErr := &CredentialUnavailableError{CredentialType: "Chained Token Credential", Message: "Sources cannot contain a nil TokenCredential"}
			logCredentialError(credErr.CredentialType, credErr)
			return nil, credErr
		}
	}
	return &chainedTokenCredential{sources: sources}, nil
}

// GetToken sequentially calls TokenCredential.GetToken on all the specified sources, returning the token from the first successful call to GetToken().
func (c *chainedTokenCredential) GetToken(ctx context.Context, opts azcore.TokenRequestOptions) (token *azcore.AccessToken, err error) {
	var errList []*CredentialUnavailableError
	// loop through all of the credentials provided in sources
	for _, cred := range c.sources {
		// make a GetToken request for the current credential in the loop
		token, err = cred.GetToken(ctx, opts)
		// check if we received a CredentialUnavailableError
		var credErr *CredentialUnavailableError
		if errors.As(err, &credErr) {
			// if we did receive a CredentialUnavailableError then we append it to our error slice and continue looping for a good credential
			errList = append(errList, credErr)
		} else if err != nil {
			// if we receive some other type of error then we must stop looping and process the error accordingly
			var authenticationFailed *AuthenticationFailedError
			if errors.As(err, &authenticationFailed) {
				// if the error is an AuthenticationFailedError we return the error related to the invalid credential and append all of the other error messages received prior to this point
				authErr := &AuthenticationFailedError{msg: "Received an AuthenticationFailedError, there is an invalid credential in the chain. " + createChainedErrorMessage(errList), inner: err}
				return nil, authErr
			}
			// if we receive some other error type this is unexpected and we simple return the unexpected error
			return nil, err
		} else {
			logGetTokenSuccess(c, opts)
			// if we did not receive an error then we return the token
			return token, nil
		}
	}
	// if we reach this point it means that all of the credentials in the chain returned CredentialUnavailableErrors
	credErr := &CredentialUnavailableError{CredentialType: "Chained Token Credential", Message: createChainedErrorMessage(errList)}
	// skip adding the stack trace here as it was already logged by other calls to GetToken()
	addGetTokenFailureLogs("Chained Token Credential", credErr, false)
	return nil, credErr
}

// AuthenticationPolicy implements the azcore.Credential interface on ChainedTokenCredential and sets the bearer token
func (c *chainedTokenCredential) AuthenticationPolicy(options azcore.AuthenticationPolicyOptions) azcore.Policy {
	return newBearerTokenPolicy(c, options)
}

// helper function used to chain the error messages of the CredentialUnavailableError slice
func createChainedErrorMessage(errList []*CredentialUnavailableError) string {
	msg := ""
	for _, err := range errList {
		msg += err.Error()
	}

	return msg
}
