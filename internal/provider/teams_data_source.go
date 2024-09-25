// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
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
		MarkdownDescription: "Teams data source",
		Attributes: map[string]schema.Attribute{
			"teams": schema.MapNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "id and Github team"},
						"name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Readable team name",
						},
						"cf_org": schema.StringAttribute{
							Computed:            true,
							Optional:            true,
							MarkdownDescription: "CF organization team deploys to",
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
				ID:    types.StringValue("team1"),
				Name:  types.StringValue("Team 1"),
				CfOrg: types.StringValue("cforg1"),
			},
			"team2": {
				ID:   types.StringValue("team2"),
				Name: types.StringValue("Team 2"),
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

// coffeesDataSourceModel maps the data source schema data.
type teamsDataSourceModel struct {
	Teams map[string]teamsModel `tfsdk:"teams"`
}

// teamsModel maps coffees schema data.
type teamsModel struct {
	ID    types.String `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	CfOrg types.String `tfsdk:"cf_org"`
}
