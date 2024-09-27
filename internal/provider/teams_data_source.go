// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-ee-platform/internal/client"
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
	platformClient client.PlatformClient
}

// Metadata returns the data source type name.
func (d *teamsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_teams"
}

func (d *teamsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	platformClient, ok := req.ProviderData.(client.PlatformClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected client.Teams, got: %T. Please report this issue to the provider developers.", req.ProviderData),
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
						"name": schema.StringAttribute{
							Required:            true,
							Description:         "Team name",
							MarkdownDescription: "Team name",
						},
						"department": schema.StringAttribute{
							Required:            true,
							Description:         "SN department the team is part of",
							MarkdownDescription: "SN department the team is part of",
						},
						"domain": schema.StringAttribute{
							Optional:            true,
							Description:         "SN Digital domain the team is part of",
							MarkdownDescription: "SN Digital domain the team is part of",
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
	teams, err := d.platformClient.GetTeams()
	if err != nil {
		resp.Diagnostics.AddError("Unable to fetch teams", err.Error())
	}

	state := teamsDataSourceModel{
		Teams: make(map[string]teamsModel),
	}

	for _, team := range teams {
		state.Teams[team.ID] = teamsModel{
			ID:         types.StringValue(team.ID),
			Name:       types.StringValue(team.Name),
			Department: types.StringValue(team.Department),
			Domain:     types.StringValue(team.Domain),
			SnPaasOrg:  types.StringValue(team.SnPaasOrg),
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
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Department types.String `tfsdk:"department"`
	Domain     types.String `tfsdk:"domain"`
	SnPaasOrg  types.String `tfsdk:"snpaas_org"`
}
