package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	openapi "github.com/noamgloberman0/openwebui-go-client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &openwebuiProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &openwebuiProvider{
			version: version,
		}
	}
}

// openwebuiProvider is the provider implementation.
type openwebuiProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

type openwebuiProviderModel struct {
	Host string `tfsdk:"host"`
	Key  string `tfsdk:"api_key"`
}

func getConfiguration(serverURL string, api_key string) *openapi.Configuration {
	cfg := openapi.Configuration(openapi.Configuration{
		Servers: []openapi.ServerConfiguration{
			{
				URL:         serverURL,
				Description: "No description provided",
			},
		},
		DefaultHeader: map[string]string{
			"Authorization": "Bearer " + api_key,
		},
	})
	return &cfg
}

// Metadata returns the provider type name.
func (p *openwebuiProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "openwebui"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *openwebuiProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Required: true,
			},
			"api_key": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

// Configure prepares a openwebui API client for data sources and resources.
func (p *openwebuiProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config openwebuiProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Open WebUI API Host",
			"The provider cannot create the Open WebUI API client as there is an empty configuration value for the Open WebUI API host.",
		)
	}

	if config.Key == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Unknown Open WebUI API Key",
			"The provider cannot create the Open WebUI API client as there is an empty configuration value for the Open WebUI API key.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// create the open webui client
	openWebUIConfig := getConfiguration(config.Host, config.Key)
	client := openapi.NewAPIClient(openWebUIConfig)
	resp.DataSourceData = client
	resp.ResourceData = client
}

// DataSources defines the data sources implemented in the provider.
func (p *openwebuiProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewUsersDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *openwebuiProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
