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

func Admin_customers_chrome_printservers_createHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		parentVal, ok := args["parent"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: parent"), nil
		}
		parent, ok := parentVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: parent"), nil
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
		var requestBody models.PrintServer
		
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
		url := fmt.Sprintf("%s/admin/directory/v1/%s/chrome/printServers%s", cfg.BaseURL, parent, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
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
		var result models.PrintServer
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

func CreateAdmin_customers_chrome_printservers_createTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_admin_directory_v1_parent_chrome_printServers",
		mcp.WithDescription("Creates a print server."),
		mcp.WithString("parent", mcp.Required(), mcp.Description("Required. The [unique ID](https://developers.google.com/admin-sdk/directory/reference/rest/v1/customers) of the customer's Google Workspace account. Format: `customers/{id}`")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. Time when the print server was created.")),
		mcp.WithString("description", mcp.Description("Input parameter: Editable. Description of the print server (as shown in the Admin console).")),
		mcp.WithString("displayName", mcp.Description("Input parameter: Editable. Display name of the print server (as shown in the Admin console).")),
		mcp.WithString("id", mcp.Description("Input parameter: Immutable. ID of the print server. Leave empty when creating.")),
		mcp.WithString("name", mcp.Description("Input parameter: Immutable. Resource name of the print server. Leave empty when creating. Format: `customers/{customer.id}/printServers/{print_server.id}`")),
		mcp.WithString("orgUnitId", mcp.Description("Input parameter: ID of the organization unit (OU) that owns this print server. This value can only be set when the print server is initially created. If it's not populated, the print server is placed under the root OU. The `org_unit_id` can be retrieved using the [Directory API](/admin-sdk/directory/reference/rest/v1/orgunits).")),
		mcp.WithString("uri", mcp.Description("Input parameter: Editable. Print server URI.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Admin_customers_chrome_printservers_createHandler(cfg),
	}
}
