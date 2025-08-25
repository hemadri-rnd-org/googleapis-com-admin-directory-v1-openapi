package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/admin-sdk-api/mcp-server/config"
	"github.com/admin-sdk-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Directory_members_deleteHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/admin/directory/v1/groups/%s/members/%s%s", cfg.BaseURL, groupKey, memberKey, queryString)
		req, err := http.NewRequest("DELETE", url, nil)
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
		var result map[string]interface{}
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

func CreateDirectory_members_deleteTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_admin_directory_v1_groups_groupKey_members_memberKey",
		mcp.WithDescription("Removes a member from a group."),
		mcp.WithString("groupKey", mcp.Required(), mcp.Description("Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group ID.")),
		mcp.WithString("memberKey", mcp.Required(), mcp.Description("Identifies the group member in the API request. A group member can be a user or another group. The value can be the member's (group or user) primary email address, alias, or unique ID.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Directory_members_deleteHandler(cfg),
	}
}
