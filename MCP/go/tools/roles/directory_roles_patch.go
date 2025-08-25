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

func Directory_roles_patchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		roleIdVal, ok := args["roleId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: roleId"), nil
		}
		roleId, ok := roleIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: roleId"), nil
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
		var requestBody models.Role
		
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
		url := fmt.Sprintf("%s/admin/directory/v1/customer/%s/roles/%s%s", cfg.BaseURL, customer, roleId, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
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
		var result models.Role
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

func CreateDirectory_roles_patchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_admin_directory_v1_customer_customer_roles_roleId",
		mcp.WithDescription("Patches a role."),
		mcp.WithString("customer", mcp.Required(), mcp.Description("Immutable ID of the Google Workspace account.")),
		mcp.WithString("roleId", mcp.Required(), mcp.Description("Immutable ID of the role.")),
		mcp.WithString("kind", mcp.Description("Input parameter: The type of the API resource. This is always `admin#directory#role`.")),
		mcp.WithString("roleDescription", mcp.Description("Input parameter: A short description of the role.")),
		mcp.WithString("roleId", mcp.Description("Input parameter: ID of the role.")),
		mcp.WithString("roleName", mcp.Description("Input parameter: Name of the role.")),
		mcp.WithArray("rolePrivileges", mcp.Description("Input parameter: The set of privileges that are granted to this role.")),
		mcp.WithString("etag", mcp.Description("Input parameter: ETag of the resource.")),
		mcp.WithBoolean("isSuperAdminRole", mcp.Description("Input parameter: Returns `true` if the role is a super admin role.")),
		mcp.WithBoolean("isSystemRole", mcp.Description("Input parameter: Returns `true` if this is a pre-defined system role.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Directory_roles_patchHandler(cfg),
	}
}
