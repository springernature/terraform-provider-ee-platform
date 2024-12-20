// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/springernature/ee-platform/pkg/api_client"
	"net/url"
	"os"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &eePlatformProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &eePlatformProvider{
			version: version,
		}
	}
}

// eePlatformProvider is the provider implementation.
type eePlatformProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *eePlatformProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "ee-platform"
	resp.Version = p.version
}

type EEPlatformProviderModel struct {
	PlatformAPI types.String `tfsdk:"platform_api"`
}

func (p *eePlatformProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manage EE Platform resources with terraform",
		Attributes: map[string]schema.Attribute{
			"platform_api": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

// Configure prepares a EE Platform API client for data sources and resources.
func (p *eePlatformProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	platformAPI := os.Getenv("PLATFORM_API")

	var data EEPlatformProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if platformAPI == "" {
		if !data.PlatformAPI.IsNull() {
			platformAPI = data.PlatformAPI.ValueString()
		} else {
			resp.Diagnostics.AddError(
				"Missing teams API endpoint",
				"While configuring the provider, the teams API endpoint was not found in the PLATFORM_API environment variable or provider configuration block platform_api attribute.",
			)
			return
		}
	}

	pAURL, err := url.Parse(platformAPI)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to parse API endpoint",
			"",
		)
		return
	}

	clientConfig := api_client.NewConfiguration()

	clientConfig.Host = pAURL.Host
	clientConfig.Scheme = pAURL.Scheme
	apiClient := api_client.NewAPIClient(clientConfig)
	resp.DataSourceData = apiClient
}

// DataSources defines the data sources implemented in the provider.
func (p *eePlatformProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewTeamsDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *eePlatformProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
