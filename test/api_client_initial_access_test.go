/*
Keycloak Admin REST API

Testing ClientInitialAccessAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package keycloakadminclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/miguoliang/keycloakadminclient"
)

func Test_keycloakadminclient_ClientInitialAccessAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ClientInitialAccessAPIService AdminRealmsRealmClientsInitialAccessGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		resp, httpRes, err := apiClient.ClientInitialAccessAPI.AdminRealmsRealmClientsInitialAccessGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientInitialAccessAPIService AdminRealmsRealmClientsInitialAccessIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var id string

		httpRes, err := apiClient.ClientInitialAccessAPI.AdminRealmsRealmClientsInitialAccessIdDelete(context.Background(), realm, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientInitialAccessAPIService AdminRealmsRealmClientsInitialAccessPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		resp, httpRes, err := apiClient.ClientInitialAccessAPI.AdminRealmsRealmClientsInitialAccessPost(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}