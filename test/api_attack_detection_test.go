/*
Keycloak Admin REST API

Testing AttackDetectionAPIService

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

func Test_keycloakadminclient_AttackDetectionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AttackDetectionAPIService AdminRealmsRealmAttackDetectionBruteForceUsersDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		httpRes, err := apiClient.AttackDetectionAPI.AdminRealmsRealmAttackDetectionBruteForceUsersDelete(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttackDetectionAPIService AdminRealmsRealmAttackDetectionBruteForceUsersUserIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var userId string

		httpRes, err := apiClient.AttackDetectionAPI.AdminRealmsRealmAttackDetectionBruteForceUsersUserIdDelete(context.Background(), realm, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttackDetectionAPIService AdminRealmsRealmAttackDetectionBruteForceUsersUserIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var userId string

		resp, httpRes, err := apiClient.AttackDetectionAPI.AdminRealmsRealmAttackDetectionBruteForceUsersUserIdGet(context.Background(), realm, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}