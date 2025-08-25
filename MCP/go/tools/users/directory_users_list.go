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

func Directory_users_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["customFieldMask"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("customFieldMask=%v", val))
		}
		if val, ok := args["customer"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("customer=%v", val))
		}
		if val, ok := args["domain"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("domain=%v", val))
		}
		if val, ok := args["event"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("event=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["orderBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orderBy=%v", val))
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
		if val, ok := args["showDeleted"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("showDeleted=%v", val))
		}
		if val, ok := args["sortOrder"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sortOrder=%v", val))
		}
		if val, ok := args["viewType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("viewType=%v", val))
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
		url := fmt.Sprintf("%s/admin/directory/v1/users%s", cfg.BaseURL, queryString)
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
		var result models.Users
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

func CreateDirectory_users_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_admin_directory_v1_users",
		mcp.WithDescription("Retrieves a paginated list of either deleted users or all users in a domain."),
		mcp.WithString("customFieldMask", mcp.Description("A comma-separated list of schema names. All fields from these schemas are fetched. This should only be set when `projection=custom`.")),
		mcp.WithString("customer", mcp.Description("The unique ID for the customer's Google Workspace account. In case of a multi-domain account, to fetch all groups for a customer, use this field instead of `domain`. You can also use the `my_customer` alias to represent your account's `customerId`. The `customerId` is also returned as part of the [Users](/admin-sdk/directory/v1/reference/users) resource. You must provide either the `customer` or the `domain` parameter.")),
		mcp.WithString("domain", mcp.Description("The domain name. Use this field to get groups from only one domain. To return all domains for a customer account, use the `customer` query parameter instead. Either the `customer` or the `domain` parameter must be provided.")),
		mcp.WithString("event", mcp.Description("Event on which subscription is intended (if subscribing)")),
		mcp.WithNumber("maxResults", mcp.Description("Maximum number of results to return.")),
		mcp.WithString("orderBy", mcp.Description("Property to use for sorting results.")),
		mcp.WithString("pageToken", mcp.Description("Token to specify next page in the list")),
		mcp.WithString("projection", mcp.Description("What subset of fields to fetch for this user.")),
		mcp.WithString("query", mcp.Description("Query string for searching user fields. For more information on constructing user queries, see [Search for Users](/admin-sdk/directory/v1/guides/search-users).")),
		mcp.WithString("showDeleted", mcp.Description("If set to `true`, retrieves the list of deleted users. (Default: `false`)")),
		mcp.WithString("sortOrder", mcp.Description("Whether to return results in ascending or descending order, ignoring case.")),
		mcp.WithString("viewType", mcp.Description("Whether to fetch the administrator-only or domain-wide public view of the user. For more information, see [Retrieve a user as a non-administrator](/admin-sdk/directory/v1/guides/manage-users#retrieve_users_non_admin).")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Directory_users_listHandler(cfg),
	}
}
