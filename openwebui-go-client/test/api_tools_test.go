/*
FastAPI

Testing ToolsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_ToolsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ToolsAPIService CreateNewToolsApiV1ToolsCreatePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolsAPI.CreateNewToolsApiV1ToolsCreatePost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolsAPIService DeleteToolsByIdApiV1ToolsIdIdDeleteDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ToolsAPI.DeleteToolsByIdApiV1ToolsIdIdDeleteDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolsAPIService ExportToolsApiV1ToolsExportGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolsAPI.ExportToolsApiV1ToolsExportGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolsAPIService GetToolListApiV1ToolsListGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolsAPI.GetToolListApiV1ToolsListGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolsAPIService GetToolsApiV1ToolsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolsAPI.GetToolsApiV1ToolsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolsAPIService GetToolsByIdApiV1ToolsIdIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ToolsAPI.GetToolsByIdApiV1ToolsIdIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolsAPIService GetToolsUserValvesByIdApiV1ToolsIdIdValvesUserGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ToolsAPI.GetToolsUserValvesByIdApiV1ToolsIdIdValvesUserGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolsAPIService GetToolsUserValvesSpecByIdApiV1ToolsIdIdValvesUserSpecGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ToolsAPI.GetToolsUserValvesSpecByIdApiV1ToolsIdIdValvesUserSpecGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolsAPIService GetToolsValvesByIdApiV1ToolsIdIdValvesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ToolsAPI.GetToolsValvesByIdApiV1ToolsIdIdValvesGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolsAPIService GetToolsValvesSpecByIdApiV1ToolsIdIdValvesSpecGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ToolsAPI.GetToolsValvesSpecByIdApiV1ToolsIdIdValvesSpecGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolsAPIService UpdateToolsByIdApiV1ToolsIdIdUpdatePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ToolsAPI.UpdateToolsByIdApiV1ToolsIdIdUpdatePost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolsAPIService UpdateToolsUserValvesByIdApiV1ToolsIdIdValvesUserUpdatePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ToolsAPI.UpdateToolsUserValvesByIdApiV1ToolsIdIdValvesUserUpdatePost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolsAPIService UpdateToolsValvesByIdApiV1ToolsIdIdValvesUpdatePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ToolsAPI.UpdateToolsValvesByIdApiV1ToolsIdIdValvesUpdatePost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
