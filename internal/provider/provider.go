// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	zilliz "github.com/zilliztech/terraform-provider-zillizcloud/client"
	byoc "github.com/zilliztech/terraform-provider-zillizcloud/internal/provider/byoc"
	byoc_op "github.com/zilliztech/terraform-provider-zillizcloud/internal/provider/byoc_i"
)

// Ensure ZillizProvider satisfies various provider interfaces.
var _ provider.Provider = &ZillizProvider{}

// ZillizProvider defines the provider implementation.
type ZillizProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// zillizProviderModel describes the provider data model.
type zillizProviderModel struct {
	ApiKey      types.String `tfsdk:"api_key"`
	RegionId    types.String `tfsdk:"region_id"`
	HostAddress types.String `tfsdk:"host_address"`
}

func (p *ZillizProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "zillizcloud"
	resp.Version = p.version
}

func (p *ZillizProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				MarkdownDescription: "Zilliz Cloud API Key",
				Optional:            true,
				Sensitive:           true,
			},
			"region_id": schema.StringAttribute{
				MarkdownDescription: "Zilliz Cloud Region Id",
				Optional:            true,
			},
			"host_address": schema.StringAttribute{
				MarkdownDescription: "Zilliz Cloud Host Address",
				Optional:            true,
			},
		},
	}
}

func (p *ZillizProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data zillizProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Default to environment variables, but override
	// with Terraform configuration value if set.
	apiKey := os.Getenv("ZILLIZCLOUD_API_KEY")
	if !data.ApiKey.IsNull() {
		apiKey = data.ApiKey.ValueString()
	}

	hostAddress := os.Getenv("ZILLIZCLOUD_HOST_ADDRESS")
	if !data.HostAddress.IsNull() {
		hostAddress = data.HostAddress.ValueString()
	}

	client, err := zilliz.NewClient(
		zilliz.WithApiKey(apiKey),
		zilliz.WithCloudRegionId(data.RegionId.ValueString()),
		zilliz.WithHostAddress(hostAddress),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create client: %v", err.Error())
		return
	}

	// Zilliz client for data sources and resources
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *ZillizProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewClusterResource,
		byoc.NewBYOCProjectResource,
		byoc_op.NewBYOCOpProjectSettingsResource,
		byoc_op.NewBYOCOpProjectResource,
		byoc_op.NewBYOCOpProjectAgentResource,
		NewUserResource,
		NewUserRoleResource,
		NewDatabaseResource,
		NewCollectionResource,
		NewIndexResource,
		NewAliasResource,
		NewPartitionsResource,
	}
}

func (p *ZillizProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewCloudProvidersDataSource,
		NewCloudRegionsDataSource,
		NewProjectDataSource,
		NewClustersDataSource,
		NewClusterDataSource,
		byoc.NewExternalIdDataSource,
		byoc_op.NewBYOCOpProjectSettingsData,
		NewUsersDataSource,
		NewRolesDataSource,
		NewDatabasesDataSource,
		NewCollectionsDataSource,
		NewIndexesDataSource,
		NewAliasesDataSource,
		NewPartitionsDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ZillizProvider{
			version: version,
		}
	}
}
