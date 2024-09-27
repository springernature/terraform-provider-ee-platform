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
	"os"
	"terraform-provider-ee-platform/internal/client"
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
	TeamsApi types.String `tfsdk:"teams_api"`
}

func (p *eePlatformProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manage EE Platform resources with terraform",
		Attributes: map[string]schema.Attribute{
			"teams_api": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

// Configure prepares a EE Platform API client for data sources and resources.
func (p *eePlatformProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	teamsApi := os.Getenv("TEAMS_API")

	var data EEPlatformProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if teamsApi == "" {
		if !data.TeamsApi.IsNull() {
			teamsApi = data.TeamsApi.ValueString()
		} else {
			resp.Diagnostics.AddError(
				"Missing teams API endpoint",
				"While configuring the provider, the teams API endpoint was not found in the TEAMS_API environment variable or provider configuration block teams_api attribute.",
			)
			return
		}
	}

	resp.DataSourceData = client.NewTeamsClient(teamsApi)
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
