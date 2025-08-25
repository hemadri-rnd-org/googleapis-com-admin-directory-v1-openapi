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

func Admin_customer_devices_chromeos_commands_getHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		deviceIdVal, ok := args["deviceId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: deviceId"), nil
		}
		deviceId, ok := deviceIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: deviceId"), nil
		}
		commandIdVal, ok := args["commandId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: commandId"), nil
		}
		commandId, ok := commandIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: commandId"), nil
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
		url := fmt.Sprintf("%s/admin/directory/v1/customer/%s/devices/chromeos/%s/commands/%s%s", cfg.BaseURL, customerId, deviceId, commandId, queryString)
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
		var result models.DirectoryChromeosdevicesCommand
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

func CreateAdmin_customer_devices_chromeos_commands_getTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_admin_directory_v1_customer_customerId_devices_chromeos_deviceId_commands_commandId",
		mcp.WithDescription("Gets command data a specific command issued to the device."),
		mcp.WithString("customerId", mcp.Required(), mcp.Description("Immutable. ID of the Google Workspace account.")),
		mcp.WithString("deviceId", mcp.Required(), mcp.Description("Immutable. ID of Chrome OS Device.")),
		mcp.WithString("commandId", mcp.Required(), mcp.Description("Immutable. ID of Chrome OS Device Command.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Admin_customer_devices_chromeos_commands_getHandler(cfg),
	}
}
