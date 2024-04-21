/*
Keycloak Admin REST API

Testing KeyAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/miguoliang/keycloakadminclient"
)

func Test_main_KeyAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test KeyAPIService AdminRealmsRealmKeysGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		resp, httpRes, err := apiClient.KeyAPI.AdminRealmsRealmKeysGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
