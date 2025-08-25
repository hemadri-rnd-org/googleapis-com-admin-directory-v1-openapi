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

func Admin_customers_chrome_printers_createHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		var requestBody models.Printer
		
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
		url := fmt.Sprintf("%s/admin/directory/v1/%s/chrome/printers%s", cfg.BaseURL, parent, queryString)
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
		var result models.Printer
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

func CreateAdmin_customers_chrome_printers_createTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_admin_directory_v1_parent_chrome_printers",
		mcp.WithDescription("Creates a printer under given Organization Unit."),
		mcp.WithString("parent", mcp.Required(), mcp.Description("Required. The name of the customer. Format: customers/{customer_id}")),
		mcp.WithString("makeAndModel", mcp.Description("Input parameter: Editable. Make and model of printer. e.g. Lexmark MS610de Value must be in format as seen in ListPrinterModels response.")),
		mcp.WithBoolean("useDriverlessConfig", mcp.Description("Input parameter: Editable. flag to use driverless configuration or not. If it's set to be true, make_and_model can be ignored")),
		mcp.WithString("orgUnitId", mcp.Description("Input parameter: Organization Unit that owns this printer (Only can be set during Printer creation)")),
		mcp.WithArray("auxiliaryMessages", mcp.Description("Input parameter: Output only. Auxiliary messages about issues with the printer configuration if any.")),
		mcp.WithString("displayName", mcp.Description("Input parameter: Editable. Name of printer.")),
		mcp.WithString("name", mcp.Description("Input parameter: The resource name of the Printer object, in the format customers/{customer-id}/printers/{printer-id} (During printer creation leave empty)")),
		mcp.WithString("uri", mcp.Description("Input parameter: Editable. Printer URI.")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. Time when printer was created.")),
		mcp.WithString("description", mcp.Description("Input parameter: Editable. Description of printer.")),
		mcp.WithString("id", mcp.Description("Input parameter: Id of the printer. (During printer creation leave empty)")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Admin_customers_chrome_printers_createHandler(cfg),
	}
}
