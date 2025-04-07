package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	openapi "github.com/noamgloberman0/openwebui-go-client"
)

var (
	_ resource.Resource                = &Group{}
	_ resource.ResourceWithConfigure   = &Group{}
	_ resource.ResourceWithImportState = &Group{}
)

func GroupResource() resource.Resource {
	return &Group{}
}

type Group struct {
	client *openapi.APIClient
}

func (r *Group) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

func (r *Group) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the resource.",
			},
			"description": schema.StringAttribute{
				Required:    true,
				Description: "A description for the resource.",
			},
			"permissions": schema.MapAttribute{
				ElementType: types.MapType{ElemType: types.BoolType},
				Required:    true,
				Description: "Permissions as a map of dynamic key-value pairs.",
			},
		},
	}
}

func (r *Group) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*openapi.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *openapi.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}
func (r *Group) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GroupModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Convert permissions to map[string]interface{}
	permissions := flattenPermissions(data.Permissions)

	response, _, err := r.client.GroupsAPI.CreateNewGroupApiV1GroupsCreatePost(ctx).GroupForm(openapi.GroupForm{
		Name:        data.Name.ValueString(),
		Description: data.Description.ValueString(),
		Permissions: permissions,
	}).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating group",
			fmt.Sprintf("Could not create group: %s", err.Error()),
		)
		return
	}

	data.ID = types.StringValue(response.Id)
	data.Name = types.StringValue(response.Name)
	data.Description = types.StringValue(response.Description)
	data.Permissions = expandPermissions(response.Permissions)

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *Group) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GroupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	group, _, err := r.client.GroupsAPI.GetGroupByIdApiV1GroupsIdIdGet(ctx, data.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading group",
			fmt.Sprintf("Could not read group %q: %v", data.ID.ValueString(), err),
		)
		return
	}

	// Populate the data model with the group information
	data.ID = types.StringValue(group.Id)
	data.Name = types.StringValue(group.Name)
	data.Description = types.StringValue(group.Description)
	data.Permissions = expandPermissions(group.Permissions)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *Group) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data GroupModel
	var state GroupModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.ID = state.ID

	// Convert permissions to map[string]interface{}
	permissions := flattenPermissions(data.Permissions)

	// Update the group
	group, _, err := r.client.GroupsAPI.UpdateGroupByIdApiV1GroupsIdIdUpdatePost(ctx, data.ID.ValueString()).GroupUpdateForm(openapi.GroupUpdateForm{
		Name:        data.Name.ValueString(),
		Description: data.Description.ValueString(),
		Permissions: permissions,
	}).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating group",
			fmt.Sprintf("Could not update group %q: %v", data.ID.ValueString(), err),
		)
		return
	}

	// Populate the data model with the updated group information
	data.Name = types.StringValue(group.Name)
	data.Description = types.StringValue(group.Description)
	data.Permissions = expandPermissions(group.Permissions)

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *Group) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GroupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.client.GroupsAPI.DeleteGroupByIdApiV1GroupsIdIdDeleteDelete(ctx, data.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting group",
			fmt.Sprintf("Could not delete group %q: %v", data.ID.ValueString(), err),
		)
		return
	}
}

func (r *Group) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

type GroupModel struct {
	ID          types.String                     `tfsdk:"id"`
	Name        types.String                     `tfsdk:"name"`
	Description types.String                     `tfsdk:"description"`
	Permissions map[string]map[string]types.Bool `tfsdk:"permissions"`
}

func flattenPermissions(permissions map[string]map[string]types.Bool) map[string]interface{} {
	flat := make(map[string]interface{})
	for key, nestedMap := range permissions {
		innerMap := make(map[string]bool)
		for innerKey, value := range nestedMap {
			innerMap[innerKey] = value.ValueBool()
		}
		flat[key] = innerMap
	}
	return flat
}

func expandPermissions(permissions map[string]interface{}) map[string]map[string]types.Bool {
	expanded := make(map[string]map[string]types.Bool)
	for key, value := range permissions {
		if nestedMap, ok := value.(map[string]interface{}); ok {
			innerMap := make(map[string]types.Bool)
			for innerKey, innerValue := range nestedMap {
				if boolValue, ok := innerValue.(bool); ok {
					innerMap[innerKey] = types.BoolValue(boolValue)
				}
			}
			expanded[key] = innerMap
		}
	}
	return expanded
}
