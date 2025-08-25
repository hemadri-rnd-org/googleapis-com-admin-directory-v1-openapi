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

func Directory_users_watchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody models.Channel
		
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
		url := fmt.Sprintf("%s/admin/directory/v1/users/watch%s", cfg.BaseURL, queryString)
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
		var result models.Channel
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

func CreateDirectory_users_watchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_admin_directory_v1_users_watch",
		mcp.WithDescription("Watches for changes in users list."),
		mcp.WithString("customFieldMask", mcp.Description("Comma-separated list of schema names. All fields from these schemas are fetched. This should only be set when projection=custom.")),
		mcp.WithString("customer", mcp.Description("Immutable ID of the Google Workspace account. In case of multi-domain, to fetch all users for a customer, fill this field instead of domain.")),
		mcp.WithString("domain", mcp.Description("Name of the domain. Fill this field to get users from only this domain. To return all users in a multi-domain fill customer field instead.\"")),
		mcp.WithString("event", mcp.Description("Events to watch for.")),
		mcp.WithNumber("maxResults", mcp.Description("Maximum number of results to return.")),
		mcp.WithString("orderBy", mcp.Description("Column to use for sorting results")),
		mcp.WithString("pageToken", mcp.Description("Token to specify next page in the list")),
		mcp.WithString("projection", mcp.Description("What subset of fields to fetch for this user.")),
		mcp.WithString("query", mcp.Description("Query string search. Should be of the form \"\". Complete documentation is at https: //developers.google.com/admin-sdk/directory/v1/guides/search-users")),
		mcp.WithString("showDeleted", mcp.Description("If set to true, retrieves the list of deleted users. (Default: false)")),
		mcp.WithString("sortOrder", mcp.Description("Whether to return results in ascending or descending order.")),
		mcp.WithString("viewType", mcp.Description("Whether to fetch the administrator-only or domain-wide public view of the user. For more information, see [Retrieve a user as a non-administrator](/admin-sdk/directory/v1/guides/manage-users#retrieve_users_non_admin).")),
		mcp.WithString("id", mcp.Description("Input parameter: A UUID or similar unique string that identifies this channel.")),
		mcp.WithString("resourceId", mcp.Description("Input parameter: An opaque ID that identifies the resource being watched on this channel. Stable across different API versions.")),
		mcp.WithString("token", mcp.Description("Input parameter: An arbitrary string delivered to the target address with each notification delivered over this channel. Optional.")),
		mcp.WithString("type", mcp.Description("Input parameter: The type of delivery mechanism used for this channel.")),
		mcp.WithString("kind", mcp.Description("Input parameter: Identifies this as a notification channel used to watch for changes to a resource, which is `api#channel`.")),
		mcp.WithString("address", mcp.Description("Input parameter: The address where notifications are delivered for this channel.")),
		mcp.WithObject("params", mcp.Description("Input parameter: Additional parameters controlling delivery channel behavior. Optional. For example, `params.ttl` specifies the time-to-live in seconds for the notification channel, where the default is 2 hours and the maximum TTL is 2 days.")),
		mcp.WithBoolean("payload", mcp.Description("Input parameter: A Boolean value to indicate whether payload is wanted. Optional.")),
		mcp.WithString("resourceUri", mcp.Description("Input parameter: A version-specific identifier for the watched resource.")),
		mcp.WithString("expiration", mcp.Description("Input parameter: Date and time of notification channel expiration, expressed as a Unix timestamp, in milliseconds. Optional.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Directory_users_watchHandler(cfg),
	}
}
