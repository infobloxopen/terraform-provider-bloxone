package utils

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ImportStatePassthroughInt32ID parses the import ID as an int32 and saves it to the "id" attribute.
// resource.ImportStatePassthroughID cannot be used for resources with a numeric "id" because it
// writes the import ID as a string.
func ImportStatePassthroughInt32ID(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 32)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", fmt.Sprintf("Expected a numeric ID, got %q: %s", req.ID, err))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.Int32Value(int32(id)))...)
}

// ImportStatePassthroughInt64ID parses the import ID as an int64 and saves it to the "id" attribute.
// resource.ImportStatePassthroughID cannot be used for resources with a numeric "id" because it
// writes the import ID as a string.
func ImportStatePassthroughInt64ID(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", fmt.Sprintf("Expected a numeric ID, got %q: %s", req.ID, err))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.Int64Value(id))...)
}
