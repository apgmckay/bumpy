package provider

import (
	"context"
	"fmt"
	"net/http"

	client "github.com/apgmckay/bumpy-client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type bumpyVersionDataSource struct {
	Version    types.String `tfsdk:"version"`
	PreRelease types.String `tfsdk:"pre_release"`
	Build      types.String `tfsdk:"build"`
	Result     types.String `tfsdk:"result"`
}

var bumpVersionDataSourceAttributes = map[string]schema.Attribute{
	"version": schema.StringAttribute{
		MarkdownDescription: "Bumpy version to bump",
		Required:            true,
	},
	"pre_release": schema.StringAttribute{
		MarkdownDescription: "A pre release string to attach to the version",
		Optional:            true,
	},
	"build": schema.StringAttribute{
		MarkdownDescription: "A build string to attach to the version",
		Optional:            true,
	},
	"result": schema.StringAttribute{
		MarkdownDescription: "Bumpy version bump result",
		Computed:            true,
	},
}

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &BumpyMajorVersionDataSource{}

func NewBumpyMajorVersionDataSource() datasource.DataSource {
	return &BumpyMajorVersionDataSource{}
}

// BumpyMajorVersionDataSource defines the data source implementation.
type BumpyMajorVersionDataSource struct {
	client *http.Client
}

// BumpyMajorVersionDataSourceModel describes the data source data model.
type BumpyMajorVersionDataSourceModel struct {
	bumpyVersionDataSource
}

func (d *BumpyMajorVersionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = fmt.Sprintf("%s_%s_%s", req.ProviderTypeName, "major", "version")
}

func (d *BumpyMajorVersionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Bumpy data source",
		Attributes:          bumpVersionDataSourceAttributes,
	}
}

func (d *BumpyMajorVersionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*http.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *BumpyMajorVersionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BumpyMajorVersionDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	c, err := client.New("http://localhost:8080", "5s")
	if err != nil {
		return
	}

	var bumpedVersion string

	bumpedVersion, err = c.BumpMajor(map[string]string{
		"version":     data.Version.ValueString(),
		"pre-release": data.PreRelease.ValueString(),
		"build":       data.Build.ValueString(),
	})

	if err != nil {
		return
	}

	data.Result = types.StringValue(bumpedVersion)

	tflog.Trace(ctx, "read a data source")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &BumpyMinorVersionDataSource{}

func NewBumpyMinorVersionDataSource() datasource.DataSource {
	return &BumpyMinorVersionDataSource{}
}

// BumpyMinorVersionDataSource defines the data source implementation.
type BumpyMinorVersionDataSource struct {
	client *http.Client
}

// BumpyMinorVersionDataSourceModel describes the data source data model.
type BumpyMinorVersionDataSourceModel struct {
	bumpyVersionDataSource
}

func (d *BumpyMinorVersionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = fmt.Sprintf("%s_%s_%s", req.ProviderTypeName, "minor", "version")
}

func (d *BumpyMinorVersionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Bumpy data source",
		Attributes:          bumpVersionDataSourceAttributes,
	}
}

func (d *BumpyMinorVersionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*http.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *BumpyMinorVersionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BumpyMinorVersionDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	c, err := client.New("http://localhost:8080", "5s")
	if err != nil {
		return
	}

	var bumpedVersion string

	bumpedVersion, err = c.BumpMinor(map[string]string{
		"version":     data.Version.ValueString(),
		"pre-release": data.PreRelease.ValueString(),
		"build":       data.Build.ValueString(),
	})

	if err != nil {
		return
	}

	data.Result = types.StringValue(bumpedVersion)

	tflog.Trace(ctx, "read a data source")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &BumpyPatchVersionDataSource{}

func NewBumpyPatchVersionDataSource() datasource.DataSource {
	return &BumpyPatchVersionDataSource{}
}

// BumpyPatchVersionDataSource defines the data source implementation.
type BumpyPatchVersionDataSource struct {
	client *http.Client
}

// BumpyPatchVersionDataSourceModel describes the data source data model.
type BumpyPatchVersionDataSourceModel struct {
	bumpyVersionDataSource
}

func (d *BumpyPatchVersionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = fmt.Sprintf("%s_%s_%s", req.ProviderTypeName, "patch", "version")
}

func (d *BumpyPatchVersionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Bumpy data source",
		Attributes:          bumpVersionDataSourceAttributes,
	}
}

func (d *BumpyPatchVersionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*http.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *BumpyPatchVersionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BumpyPatchVersionDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	c, err := client.New("http://localhost:8080", "5s")
	if err != nil {
		return
	}

	var bumpedVersion string

	bumpedVersion, err = c.BumpPatch(map[string]string{
		"version":     data.Version.ValueString(),
		"pre-release": data.PreRelease.ValueString(),
		"build":       data.Build.ValueString(),
	})
	if err != nil {
		return
	}

	data.Result = types.StringValue(bumpedVersion)

	tflog.Trace(ctx, "read a data source")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
