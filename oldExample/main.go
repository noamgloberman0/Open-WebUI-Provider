package main

import (
	"context"
	"fmt"
	"os"

	openapi "github.com/noamgloberman0/openwebui-go-client"
)

func getConfiguration(serverURL string) *openapi.Configuration {
	cfg := openapi.Configuration(openapi.Configuration{
		Servers: []openapi.ServerConfiguration{
			{
				URL:         serverURL,
				Description: "No description provided",
			},
		},
		DefaultHeader: map[string]string{
			"Authorization": "Bearer " + os.Getenv("API_KEY"),
		},
	})
	return &cfg
}

func main() {

	// Set up the Open WebUI client
	serverURL := "http://localhost:3000"
	client := openapi.NewAPIClient(getConfiguration(serverURL))

	// Example: Call an API endpoint (replace with your actual endpoint)
	result, _, err := client.UsersAPI.GetUsersApiV1UsersGet(context.Background()).Execute()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("result:", result[0].Name)

	// Ensure CreateNewModelRequest is correctly defined in the openapi package
	meta := openapi.ModelMeta{
		Description: &openapi.Description{String: func() *string { s := "A test model created via Go client"; return &s }()},
	}

	// Example Params (replace with actual parameters your model needs)
	params := map[string]interface{}{
		"context_length": 2048,
	}

	// Example Access Control (optional)
	accessControl := map[string]interface{}{
		"public": true,
	}

	// Create the ModelForm
	modelForm := openapi.ModelForm{
		Id:            "my-unique-model-id",
		Name:          "My Go Created Model",
		Meta:          meta,
		Params:        params,
		AccessControl: accessControl,
	}

	modelModel, _, err := client.ModelsAPI.CreateNewModelApiV1ModelsCreatePost(context.Background()).ModelForm(modelForm).Execute()

	if modelModel != nil {
		fmt.Println("Model Created: ", modelModel)
	} else {
		fmt.Println("Model Creation Failed")
	}
	// fmt.Print("Model Created: ", modelModel, resp)
	fmt.Println("Error:", err.Error())
}
