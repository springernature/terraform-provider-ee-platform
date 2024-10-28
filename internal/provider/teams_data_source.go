// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/springernature/ee-platform/pkg/api_client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &teamsDataSource{}
)

// NewTeamsDataSource is a helper function to simplify the provider implementation.
func NewTeamsDataSource() datasource.DataSource {
	return &teamsDataSource{}
}

// teamsDataSource is the data source implementation.
type teamsDataSource struct {
	platformClient *api_client.APIClient
}

// Metadata returns the data source type name.
func (d *teamsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_teams"
}

func (d *teamsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	platformClient, ok := req.ProviderData.(*api_client.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *api_client.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.platformClient = platformClient
}

// Schema defines the schema for the data source.
func (d *teamsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves all teams",
		Attributes: map[string]schema.Attribute{
			"teams": schema.MapNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required:            true,
							Description:         "Unique identifier of the team",
							MarkdownDescription: "Unique identifier of the team",
						},
						"snpaas_org": schema.StringAttribute{
							Optional:            true,
							Description:         "SNPaaS Cloud Foundry organization that the team deploys to",
							MarkdownDescription: "SNPaaS Cloud Foundry organization that the team deploys to",
						},
					},
				},
			},
		},
	}
}

func (d *teamsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	teams, _, err := d.platformClient.DefaultAPI.ListTeams(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Unable to fetch teams", err.Error())
	}

	state := teamsDataSourceModel{
		Teams: make(map[string]teamsModel),
	}

	for _, team := range teams.GetTeams() {
		state.Teams[team.GetId()] = teamsModel{
			ID:        types.StringValue(team.GetId()),
			SnPaasOrg: types.StringValue(team.GetSnpaasOrg()),
		}
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

type teamsDataSourceModel struct {
	Teams map[string]teamsModel `tfsdk:"teams"`
}

type teamsModel struct {
	ID        types.String `tfsdk:"id"`
	SnPaasOrg types.String `tfsdk:"snpaas_org"`
}
