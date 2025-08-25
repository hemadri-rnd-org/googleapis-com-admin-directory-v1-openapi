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

func Directory_chromeosdevices_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		customerIdVal, ok := args["customerId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: customerId"), nil
		}
		customerId, ok := customerIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: customerId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["includeChildOrgunits"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("includeChildOrgunits=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["orderBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orderBy=%v", val))
		}
		if val, ok := args["orgUnitPath"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orgUnitPath=%v", val))
		}
		if val, ok := args["pageToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageToken=%v", val))
		}
		if val, ok := args["projection"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("projection=%v", val))
		}
		if val, ok := args["query"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("query=%v", val))
		}
		if val, ok := args["sortOrder"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sortOrder=%v", val))
		}
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
		url := fmt.Sprintf("%s/admin/directory/v1/customer/%s/devices/chromeos%s", cfg.BaseURL, customerId, queryString)
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
		var result models.ChromeOsDevices
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

func CreateDirectory_chromeosdevices_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_admin_directory_v1_customer_customerId_devices_chromeos",
		mcp.WithDescription("Retrieves a paginated list of Chrome OS devices within an account."),
		mcp.WithString("customerId", mcp.Required(), mcp.Description("The unique ID for the customer's Google Workspace account. As an account administrator, you can also use the `my_customer` alias to represent your account's `customerId`. The `customerId` is also returned as part of the [Users resource](/admin-sdk/directory/v1/reference/users).")),
		mcp.WithBoolean("includeChildOrgunits", mcp.Description("Return devices from all child orgunits, as well as the specified org unit. If this is set to true, 'orgUnitPath' must be provided.")),
		mcp.WithNumber("maxResults", mcp.Description("Maximum number of results to return.")),
		mcp.WithString("orderBy", mcp.Description("Device property to use for sorting results.")),
		mcp.WithString("orgUnitPath", mcp.Description("The full path of the organizational unit (minus the leading `/`) or its unique ID.")),
		mcp.WithString("pageToken", mcp.Description("The `pageToken` query parameter is used to request the next page of query results. The follow-on request's `pageToken` query parameter is the `nextPageToken` from your previous response.")),
		mcp.WithString("projection", mcp.Description("Restrict information returned to a set of selected fields.")),
		mcp.WithString("query", mcp.Description("Search string in the format given at https://developers.google.com/admin-sdk/directory/v1/list-query-operators")),
		mcp.WithString("sortOrder", mcp.Description("Whether to return results in ascending or descending order. Must be used with the `orderBy` parameter.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Directory_chromeosdevices_listHandler(cfg),
	}
}
