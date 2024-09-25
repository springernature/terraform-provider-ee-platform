// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
type teamsDataSource struct{}

// Metadata returns the data source type name.
func (d *teamsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_teams"
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

// Read refreshes the Terraform state with the latest data.
func (d *teamsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	state := teamsDataSourceModel{
		Teams: map[string]teamsModel{
			"team1": {
				ID:         types.StringValue("team1"),
				Name:       types.StringValue("Team 1"),
				Department: types.StringValue("dept A"),
			},
			"team2": {
				ID:         types.StringValue("team2"),
				Name:       types.StringValue("Team 2"),
				Department: types.StringValue("dept A"),
				Domain:     types.StringValue("domain A"),
				SnPaasOrg:  types.StringValue("snpaas-org-a"),
			},
		},
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// teamsDataSourceModel maps the data source schema data.
type teamsDataSourceModel struct {
	Teams map[string]teamsModel `tfsdk:"teams"`
}

// teamsModel maps teams schema data.
type teamsModel struct {
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Department types.String `tfsdk:"department"`
	Domain     types.String `tfsdk:"domain"`
	SnPaasOrg  types.String `tfsdk:"snpaas_org"`
}
