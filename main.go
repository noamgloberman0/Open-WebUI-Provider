package main

import (
        "context"
        "fmt"
        "os"
        openapi "github.com/noamgloberman0/Open-WebUI-Provider/openwebui-go-client" //replace your-project-name
)

func main() {
        cfg := openapi.NewConfiguration()
        cfg.AddDefaultHeader("Authorization", "Bearer f466e11c74554ff1adf1d2d36c5b5fd5") //example Auth header.
        client := openapi.NewAPIClient(cfg)

        // Example: Call an API endpoint (replace with your actual endpoint)
        result, _, err := client.AudioAPI.GetModelsApiV1AudioModelsGet(context.Background()).Execute()
        if err != nil {
                fmt.Println("Error:", err)
                return
        }
        fmt.Println("Result:", result)
}
