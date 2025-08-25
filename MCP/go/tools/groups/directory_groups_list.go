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

func Directory_groups_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["customer"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("customer=%v", val))
		}
		if val, ok := args["domain"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("domain=%v", val))
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
		if val, ok := args["query"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("query=%v", val))
		}
		if val, ok := args["sortOrder"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sortOrder=%v", val))
		}
		if val, ok := args["userKey"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("userKey=%v", val))
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
		url := fmt.Sprintf("%s/admin/directory/v1/groups%s", cfg.BaseURL, queryString)
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
		var result models.Groups
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

func CreateDirectory_groups_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_admin_directory_v1_groups",
		mcp.WithDescription("Retrieves all groups of a domain or of a user given a userKey (paginated)."),
		mcp.WithString("customer", mcp.Description("The unique ID for the customer's Google Workspace account. In case of a multi-domain account, to fetch all groups for a customer, use this field instead of `domain`. You can also use the `my_customer` alias to represent your account's `customerId`. The `customerId` is also returned as part of the [Users](/admin-sdk/directory/v1/reference/users) resource. You must provide either the `customer` or the `domain` parameter.")),
		mcp.WithString("domain", mcp.Description("The domain name. Use this field to get groups from only one domain. To return all domains for a customer account, use the `customer` query parameter instead.")),
		mcp.WithNumber("maxResults", mcp.Description("Maximum number of results to return. Max allowed value is 200.")),
		mcp.WithString("orderBy", mcp.Description("Column to use for sorting results")),
		mcp.WithString("pageToken", mcp.Description("Token to specify next page in the list")),
		mcp.WithString("query", mcp.Description("Query string search. Should be of the form \"\". Complete documentation is at https: //developers.google.com/admin-sdk/directory/v1/guides/search-groups")),
		mcp.WithString("sortOrder", mcp.Description("Whether to return results in ascending or descending order. Only of use when orderBy is also used")),
		mcp.WithString("userKey", mcp.Description("Email or immutable ID of the user if only those groups are to be listed, the given user is a member of. If it's an ID, it should match with the ID of the user object.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Directory_groups_listHandler(cfg),
	}
}
