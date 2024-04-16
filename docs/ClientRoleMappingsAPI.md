# \ClientRoleMappingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientAvailableGet**](ClientRoleMappingsAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientAvailableGet) | **Get** /admin/realms/{realm}/groups/{group-id}/role-mappings/clients/{client}/available | Get available client-level roles that can be mapped to the user
[**AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientCompositeGet**](ClientRoleMappingsAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientCompositeGet) | **Get** /admin/realms/{realm}/groups/{group-id}/role-mappings/clients/{client}/composite | Get effective client-level role mappings This recurses any composite roles
[**AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientDelete**](ClientRoleMappingsAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientDelete) | **Delete** /admin/realms/{realm}/groups/{group-id}/role-mappings/clients/{client} | Delete client-level roles from user role mapping
[**AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientGet**](ClientRoleMappingsAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientGet) | **Get** /admin/realms/{realm}/groups/{group-id}/role-mappings/clients/{client} | Get client-level role mappings for the user, and the app
[**AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientPost**](ClientRoleMappingsAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientPost) | **Post** /admin/realms/{realm}/groups/{group-id}/role-mappings/clients/{client} | Add client-level roles to the user role mapping
[**AdminRealmsRealmUsersUserIdRoleMappingsClientsClientAvailableGet**](ClientRoleMappingsAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsClientsClientAvailableGet) | **Get** /admin/realms/{realm}/users/{user-id}/role-mappings/clients/{client}/available | Get available client-level roles that can be mapped to the user
[**AdminRealmsRealmUsersUserIdRoleMappingsClientsClientCompositeGet**](ClientRoleMappingsAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsClientsClientCompositeGet) | **Get** /admin/realms/{realm}/users/{user-id}/role-mappings/clients/{client}/composite | Get effective client-level role mappings This recurses any composite roles
[**AdminRealmsRealmUsersUserIdRoleMappingsClientsClientDelete**](ClientRoleMappingsAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsClientsClientDelete) | **Delete** /admin/realms/{realm}/users/{user-id}/role-mappings/clients/{client} | Delete client-level roles from user role mapping
[**AdminRealmsRealmUsersUserIdRoleMappingsClientsClientGet**](ClientRoleMappingsAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsClientsClientGet) | **Get** /admin/realms/{realm}/users/{user-id}/role-mappings/clients/{client} | Get client-level role mappings for the user, and the app
[**AdminRealmsRealmUsersUserIdRoleMappingsClientsClientPost**](ClientRoleMappingsAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsClientsClientPost) | **Post** /admin/realms/{realm}/users/{user-id}/role-mappings/clients/{client} | Add client-level roles to the user role mapping



## AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientAvailableGet

> []RoleRepresentation AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientAvailableGet(ctx, realm, groupId, client).Execute()

Get available client-level roles that can be mapped to the user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/miguoliang/keycloakadminclient"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	groupId := "groupId_example" // string | 
	client := "client_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientAvailableGet(context.Background(), realm, groupId, client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientAvailableGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientAvailableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]RoleRepresentation**](RoleRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientCompositeGet

> []RoleRepresentation AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientCompositeGet(ctx, realm, groupId, client).BriefRepresentation(briefRepresentation).Execute()

Get effective client-level role mappings This recurses any composite roles

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/miguoliang/keycloakadminclient"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	groupId := "groupId_example" // string | 
	client := "client_example" // string | 
	briefRepresentation := true // bool | if false, return roles with their attributes (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientCompositeGet(context.Background(), realm, groupId, client).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientCompositeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientCompositeGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientCompositeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientCompositeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **briefRepresentation** | **bool** | if false, return roles with their attributes | [default to true]

### Return type

[**[]RoleRepresentation**](RoleRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientDelete

> AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientDelete(ctx, realm, groupId, client).RoleRepresentation(roleRepresentation).Execute()

Delete client-level roles from user role mapping

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/miguoliang/keycloakadminclient"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	groupId := "groupId_example" // string | 
	client := "client_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientDelete(context.Background(), realm, groupId, client).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientGet

> []RoleRepresentation AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientGet(ctx, realm, groupId, client).Execute()

Get client-level role mappings for the user, and the app

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/miguoliang/keycloakadminclient"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	groupId := "groupId_example" // string | 
	client := "client_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientGet(context.Background(), realm, groupId, client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]RoleRepresentation**](RoleRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientPost

> AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientPost(ctx, realm, groupId, client).RoleRepresentation(roleRepresentation).Execute()

Add client-level roles to the user role mapping

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/miguoliang/keycloakadminclient"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	groupId := "groupId_example" // string | 
	client := "client_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientPost(context.Background(), realm, groupId, client).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdRoleMappingsClientsClientAvailableGet

> []RoleRepresentation AdminRealmsRealmUsersUserIdRoleMappingsClientsClientAvailableGet(ctx, realm, userId, client).Execute()

Get available client-level roles that can be mapped to the user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/miguoliang/keycloakadminclient"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	client := "client_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientAvailableGet(context.Background(), realm, userId, client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdRoleMappingsClientsClientAvailableGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsClientsClientAvailableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]RoleRepresentation**](RoleRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdRoleMappingsClientsClientCompositeGet

> []RoleRepresentation AdminRealmsRealmUsersUserIdRoleMappingsClientsClientCompositeGet(ctx, realm, userId, client).BriefRepresentation(briefRepresentation).Execute()

Get effective client-level role mappings This recurses any composite roles

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/miguoliang/keycloakadminclient"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	client := "client_example" // string | 
	briefRepresentation := true // bool | if false, return roles with their attributes (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientCompositeGet(context.Background(), realm, userId, client).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientCompositeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdRoleMappingsClientsClientCompositeGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientCompositeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsClientsClientCompositeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **briefRepresentation** | **bool** | if false, return roles with their attributes | [default to true]

### Return type

[**[]RoleRepresentation**](RoleRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdRoleMappingsClientsClientDelete

> AdminRealmsRealmUsersUserIdRoleMappingsClientsClientDelete(ctx, realm, userId, client).RoleRepresentation(roleRepresentation).Execute()

Delete client-level roles from user role mapping

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/miguoliang/keycloakadminclient"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	client := "client_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientDelete(context.Background(), realm, userId, client).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsClientsClientDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdRoleMappingsClientsClientGet

> []RoleRepresentation AdminRealmsRealmUsersUserIdRoleMappingsClientsClientGet(ctx, realm, userId, client).Execute()

Get client-level role mappings for the user, and the app

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/miguoliang/keycloakadminclient"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	client := "client_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientGet(context.Background(), realm, userId, client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdRoleMappingsClientsClientGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsClientsClientGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]RoleRepresentation**](RoleRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdRoleMappingsClientsClientPost

> AdminRealmsRealmUsersUserIdRoleMappingsClientsClientPost(ctx, realm, userId, client).RoleRepresentation(roleRepresentation).Execute()

Add client-level roles to the user role mapping

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/miguoliang/keycloakadminclient"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	client := "client_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientPost(context.Background(), realm, userId, client).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsClientsClientPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

