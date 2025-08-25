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

func Directory_tokens_deleteHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		userKeyVal, ok := args["userKey"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: userKey"), nil
		}
		userKey, ok := userKeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: userKey"), nil
		}
		clientIdVal, ok := args["clientId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: clientId"), nil
		}
		clientId, ok := clientIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: clientId"), nil
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
		url := fmt.Sprintf("%s/admin/directory/v1/users/%s/tokens/%s%s", cfg.BaseURL, userKey, clientId, queryString)
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

func CreateDirectory_tokens_deleteTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_admin_directory_v1_users_userKey_tokens_clientId",
		mcp.WithDescription("Deletes all access tokens issued by a user for an application."),
		mcp.WithString("userKey", mcp.Required(), mcp.Description("Identifies the user in the API request. The value can be the user's primary email address, alias email address, or unique user ID.")),
		mcp.WithString("clientId", mcp.Required(), mcp.Description("The Client ID of the application the token is issued to.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Directory_tokens_deleteHandler(cfg),
	}
}
