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
var _ datasource.DataSource = &BumpyMajorDataSource{}

func NewBumpyMajorDataSource() datasource.DataSource {
	return &BumpyMajorDataSource{}
}

// BumpyMajorDataSource defines the data source implementation.
type BumpyMajorDataSource struct {
	client *http.Client
}

// BumpyMajorDataSourceModel describes the data source data model.
type BumpyMajorDataSourceModel struct {
	Version types.String `tfsdk:"version"`
}

func (d *BumpyMajorDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_major"
}

func (d *BumpyMajorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {

	resp.Schema = schema.Schema{
		MarkdownDescription: "BumpyMajor data source",
		Attributes: map[string]schema.Attribute{
			"version": schema.StringAttribute{
				MarkdownDescription: "BumpyMajor bumped version",
				Required:            true,
			},
		},
	}
}

func (d *BumpyMajorDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *BumpyMajorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BumpyMajorDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	c, err := client.New("http://localhost:8080", "5s")
	if err != nil {
		return
	}

	majorVersion, err := c.BumpMajor(map[string]string{"version": data.Version.ValueString()})
	if err != nil {
		return
	}

	data.Version = types.StringValue(majorVersion)

	tflog.Trace(ctx, "read a data source")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
