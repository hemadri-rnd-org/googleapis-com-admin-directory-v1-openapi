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

func Directory_roleassignments_getHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		customerVal, ok := args["customer"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: customer"), nil
		}
		customer, ok := customerVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: customer"), nil
		}
		roleAssignmentIdVal, ok := args["roleAssignmentId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: roleAssignmentId"), nil
		}
		roleAssignmentId, ok := roleAssignmentIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: roleAssignmentId"), nil
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
		url := fmt.Sprintf("%s/admin/directory/v1/customer/%s/roleassignments/%s%s", cfg.BaseURL, customer, roleAssignmentId, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.RoleAssignment
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

func CreateDirectory_roleassignments_getTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_admin_directory_v1_customer_customer_roleassignments_roleAssignmentId",
		mcp.WithDescription("Retrieves a role assignment."),
		mcp.WithString("customer", mcp.Required(), mcp.Description("The unique ID for the customer's Google Workspace account. In case of a multi-domain account, to fetch all groups for a customer, use this field instead of `domain`. You can also use the `my_customer` alias to represent your account's `customerId`. The `customerId` is also returned as part of the [Users](/admin-sdk/directory/v1/reference/users) resource. You must provide either the `customer` or the `domain` parameter.")),
		mcp.WithString("roleAssignmentId", mcp.Required(), mcp.Description("Immutable ID of the role assignment.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Directory_roleassignments_getHandler(cfg),
	}
}
