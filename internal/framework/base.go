package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	testHelper "github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

var (
	defaultKey   = "automation"
	defaultValue = "terraform"
)

type ResourceWithConfigure struct{}

// SetTagsAll calculates the new value for the `tags_all` attribute.
func (r *ResourceWithConfigure) SetTagsAll(ctx context.Context, request resource.ModifyPlanRequest, response *resource.ModifyPlanResponse) {
	// If the entire plan is null, the resource is planned for destruction.
	if request.Plan.Raw.IsNull() {
		return
	}

	defaultTagsConfig := r.getDefaultTags()

	var planTags types.Map
	response.Diagnostics.Append(request.Plan.GetAttribute(ctx, path.Root("tags"), &planTags)...)

	if response.Diagnostics.HasError() {
		return
	}

	actualTags := flex.ExpandFrameworkMapString(ctx, planTags, &response.Diagnostics)

	if !planTags.IsUnknown() {
		if !mapHasUnknownElements(planTags) {
			response.Diagnostics.Append(response.Plan.SetAttribute(ctx, path.Root("tags_all"), flex.FlattenFrameworkMapString(ctx, mergeTags(defaultTagsConfig, actualTags), &response.Diagnostics))...)
		} else {
			response.Diagnostics.Append(response.Plan.SetAttribute(ctx, path.Root("tags_all"), types.MapUnknown(types.StringType))...)
		}
	} else {
		response.Diagnostics.Append(response.Plan.SetAttribute(ctx, path.Root("tags_all"), types.MapUnknown(types.StringType))...)
	}

}

func (r *ResourceWithConfigure) getDefaultTags() map[string]interface{} {
	return map[string]interface{}{
		defaultKey: defaultValue,
	}
}

func mergeTags(t1, t2 map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for k, v := range t1 {
		result[k] = v
	}

	for k, v := range t2 {
		result[k] = v
	}

	return result
}

func mapHasUnknownElements(m types.Map) bool {
	for _, v := range m.Elements() {
		if v.IsUnknown() {
			return true
		}
	}

	return false
}

// Acceptance test helpers

/*
VerifyTagsAll adds a test helper to verify the tags generated are as expected
Default + User added tags
*/
func VerifyTagsAll(resourceName string, m map[string]string) []testHelper.TestCheckFunc {
	var resultFunc []testHelper.TestCheckFunc
	// Add check for default tag
	resultFunc = append(resultFunc, testHelper.TestCheckResourceAttr(resourceName, fmt.Sprintf("tags_all.%s", defaultKey), defaultValue))
	for k, v := range m {
		resultFunc = append(resultFunc, testHelper.TestCheckResourceAttr(resourceName, fmt.Sprintf("tags_all.%s", k), v))
	}
	return resultFunc
}
