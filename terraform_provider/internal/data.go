package provider

import (
	"bumpy/package/client"
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &BumpyDataSource{}

func NewBumpyDataSource() datasource.DataSource {
	return &BumpyDataSource{}
}

// BumpyDataSource defines the data source implementation.
type BumpyDataSource struct {
	client *http.Client
}

// BumpyDataSourceModel describes the data source data model.
type BumpyDataSourceModel struct {
	MajorVersion types.String `tfsdk:"major_version"`
	MinorVersion types.String `tfsdk:"minor_version"`
	PatchVersion types.String `tfsdk:"patch_version"`

	Version types.String `tfsdk:"version"`
}

func (d *BumpyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName
}

func (d *BumpyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Bumpy data source",
		Attributes: map[string]schema.Attribute{
			"major_version": schema.StringAttribute{
				MarkdownDescription: "Bumpy input major version to bump",
				Optional:            true,
			},
			"minor_version": schema.StringAttribute{
				MarkdownDescription: "Bumpy input minor version to bump",
				Optional:            true,
			},
			"patch_version": schema.StringAttribute{
				MarkdownDescription: "Bumpy input patch version to bump",
				Optional:            true,
			},
			"version": schema.StringAttribute{
				MarkdownDescription: "Bumpy bumped version",
				Computed:            true,
			},
		},
	}
}

func (d *BumpyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *BumpyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BumpyDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	c, err := client.New("http://localhost:8080", "5s")
	if err != nil {
		return
	}

	var bumpedVersion string

	if data.MajorVersion.ValueString() != "" {
		bumpedVersion, err = c.BumpMajor(map[string]string{"version": data.MajorVersion.ValueString()})
		if err != nil {
			return
		}
	}

	if data.MinorVersion.ValueString() != "" {
		bumpedVersion, err = c.BumpMinor(map[string]string{"version": data.MinorVersion.ValueString()})
		if err != nil {
			return
		}
	}

	if data.PatchVersion.ValueString() != "" {
		bumpedVersion, err = c.BumpPatch(map[string]string{"version": data.PatchVersion.ValueString()})
		if err != nil {
			return
		}
	}

	data.Version = types.StringValue(bumpedVersion)

	tflog.Trace(ctx, "read a data source")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
