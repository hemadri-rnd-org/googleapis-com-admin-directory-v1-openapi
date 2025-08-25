package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/admin-sdk-api/mcp-server/config"
	"github.com/admin-sdk-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Directory_members_updateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		groupKeyVal, ok := args["groupKey"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: groupKey"), nil
		}
		groupKey, ok := groupKeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: groupKey"), nil
		}
		memberKeyVal, ok := args["memberKey"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: memberKey"), nil
		}
		memberKey, ok := memberKeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: memberKey"), nil
		}
		queryParams := make([]string, 0)
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("access_token=%s", cfg.BearerToken))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("oauth_token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Member
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/admin/directory/v1/groups/%s/members/%s%s", cfg.BaseURL, groupKey, memberKey, queryString)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Member
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateDirectory_members_updateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_admin_directory_v1_groups_groupKey_members_memberKey",
		mcp.WithDescription("Updates the membership of a user in the specified group."),
		mcp.WithString("groupKey", mcp.Required(), mcp.Description("Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group ID.")),
		mcp.WithString("memberKey", mcp.Required(), mcp.Description("Identifies the group member in the API request. A group member can be a user or another group. The value can be the member's (group or user) primary email address, alias, or unique ID.")),
		mcp.WithString("status", mcp.Description("Input parameter: Status of member (Immutable)")),
		mcp.WithString("type", mcp.Description("Input parameter: The type of group member.")),
		mcp.WithString("delivery_settings", mcp.Description("Input parameter: Defines mail delivery preferences of member. This field is only supported by `insert`, `update`, and `get` methods.")),
		mcp.WithString("email", mcp.Description("Input parameter: The member's email address. A member can be a user or another group. This property is required when adding a member to a group. The `email` must be unique and cannot be an alias of another group. If the email address is changed, the API automatically reflects the email address changes.")),
		mcp.WithString("etag", mcp.Description("Input parameter: ETag of the resource.")),
		mcp.WithString("id", mcp.Description("Input parameter: The unique ID of the group member. A member `id` can be used as a member request URI's `memberKey`.")),
		mcp.WithString("kind", mcp.Description("Input parameter: The type of the API resource. For Members resources, the value is `admin#directory#member`.")),
		mcp.WithString("role", mcp.Description("Input parameter: The member's role in a group. The API returns an error for cycles in group memberships. For example, if `group1` is a member of `group2`, `group2` cannot be a member of `group1`. For more information about a member's role, see the [administration help center](https://support.google.com/a/answer/167094).")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Directory_members_updateHandler(cfg),
	}
}
