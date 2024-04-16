/*
Keycloak Admin REST API

Testing RolesByIDAPIService

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

func Test_keycloakadminclient_RolesByIDAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RolesByIDAPIService AdminRealmsRealmRolesByIdRoleIdCompositesClientsClientUuidGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var clientUuid string
		var roleId string

		resp, httpRes, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesClientsClientUuidGet(context.Background(), realm, clientUuid, roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesByIDAPIService AdminRealmsRealmRolesByIdRoleIdCompositesDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var roleId string

		httpRes, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesDelete(context.Background(), realm, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesByIDAPIService AdminRealmsRealmRolesByIdRoleIdCompositesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var roleId string

		resp, httpRes, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesGet(context.Background(), realm, roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesByIDAPIService AdminRealmsRealmRolesByIdRoleIdCompositesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var roleId string

		httpRes, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesPost(context.Background(), realm, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesByIDAPIService AdminRealmsRealmRolesByIdRoleIdCompositesRealmGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var roleId string

		resp, httpRes, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesRealmGet(context.Background(), realm, roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesByIDAPIService AdminRealmsRealmRolesByIdRoleIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var roleId string

		httpRes, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdDelete(context.Background(), realm, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesByIDAPIService AdminRealmsRealmRolesByIdRoleIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var roleId string

		resp, httpRes, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdGet(context.Background(), realm, roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesByIDAPIService AdminRealmsRealmRolesByIdRoleIdManagementPermissionsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var roleId string

		resp, httpRes, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdManagementPermissionsGet(context.Background(), realm, roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesByIDAPIService AdminRealmsRealmRolesByIdRoleIdManagementPermissionsPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var roleId string

		resp, httpRes, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdManagementPermissionsPut(context.Background(), realm, roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesByIDAPIService AdminRealmsRealmRolesByIdRoleIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var roleId string

		httpRes, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdPut(context.Background(), realm, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}