/*
Keycloak Admin REST API

Testing ClientAttributeCertificateAPIService

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

func Test_keycloakadminclient_ClientAttributeCertificateAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ClientAttributeCertificateAPIService AdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var clientUuid string
		var attr string

		resp, httpRes, err := apiClient.ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPost(context.Background(), realm, clientUuid, attr).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientAttributeCertificateAPIService AdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var clientUuid string
		var attr string

		resp, httpRes, err := apiClient.ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPost(context.Background(), realm, clientUuid, attr).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientAttributeCertificateAPIService AdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var clientUuid string
		var attr string

		resp, httpRes, err := apiClient.ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePost(context.Background(), realm, clientUuid, attr).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientAttributeCertificateAPIService AdminRealmsRealmClientsClientUuidCertificatesAttrGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var clientUuid string
		var attr string

		resp, httpRes, err := apiClient.ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrGet(context.Background(), realm, clientUuid, attr).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientAttributeCertificateAPIService AdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var clientUuid string
		var attr string

		resp, httpRes, err := apiClient.ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePost(context.Background(), realm, clientUuid, attr).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientAttributeCertificateAPIService AdminRealmsRealmClientsClientUuidCertificatesAttrUploadPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var clientUuid string
		var attr string

		resp, httpRes, err := apiClient.ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrUploadPost(context.Background(), realm, clientUuid, attr).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
