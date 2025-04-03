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
}
