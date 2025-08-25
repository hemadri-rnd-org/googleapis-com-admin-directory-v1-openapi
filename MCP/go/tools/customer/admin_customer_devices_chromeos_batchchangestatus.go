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

func Admin_customer_devices_chromeos_batchchangestatusHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		var requestBody models.BatchChangeChromeOsDeviceStatusRequest
		
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
		url := fmt.Sprintf("%s/admin/directory/v1/customer/%s/devices/chromeos:batchChangeStatus%s", cfg.BaseURL, customerId, queryString)
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
		var result models.BatchChangeChromeOsDeviceStatusResponse
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

func CreateAdmin_customer_devices_chromeos_batchchangestatusTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_admin_directory_v1_customer_customerId_devices_chromeos_batchChangeStatus",
		mcp.WithDescription("Changes the status of a batch of ChromeOS devices. For more information about changing a ChromeOS device state [Repair, repurpose, or retire ChromeOS devices](https://support.google.com/chrome/a/answer/3523633)."),
		mcp.WithString("customerId", mcp.Required(), mcp.Description("Required. Immutable ID of the Google Workspace account.")),
		mcp.WithString("changeChromeOsDeviceStatusAction", mcp.Description("Input parameter: Required. The action to take on the ChromeOS device in order to change its status.")),
		mcp.WithString("deprovisionReason", mcp.Description("Input parameter: Optional. The reason behind a device deprovision. Must be provided if 'changeChromeOsDeviceStatusAction' is set to 'CHANGE_CHROME_OS_DEVICE_STATUS_ACTION_DEPROVISION'. Otherwise, omit this field.")),
		mcp.WithArray("deviceIds", mcp.Description("Input parameter: Required. List of the IDs of the ChromeOS devices to change. Maximum 50.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Admin_customer_devices_chromeos_batchchangestatusHandler(cfg),
	}
}
