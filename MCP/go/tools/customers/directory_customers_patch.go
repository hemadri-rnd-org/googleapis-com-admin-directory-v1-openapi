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

func Directory_customers_patchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		customerKeyVal, ok := args["customerKey"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: customerKey"), nil
		}
		customerKey, ok := customerKeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: customerKey"), nil
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
		var requestBody models.Customer
		
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
		url := fmt.Sprintf("%s/admin/directory/v1/customers/%s%s", cfg.BaseURL, customerKey, queryString)
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
		var result models.Customer
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

func CreateDirectory_customers_patchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_admin_directory_v1_customers_customerKey",
		mcp.WithDescription("Patches a customer."),
		mcp.WithString("customerKey", mcp.Required(), mcp.Description("Id of the customer to be updated")),
		mcp.WithString("customerCreationTime", mcp.Description("Input parameter: The customer's creation time (Readonly)")),
		mcp.WithString("customerDomain", mcp.Description("Input parameter: The customer's primary domain name string. Do not include the `www` prefix when creating a new customer.")),
		mcp.WithString("phoneNumber", mcp.Description("Input parameter: The customer's contact phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.")),
		mcp.WithString("alternateEmail", mcp.Description("Input parameter: The customer's secondary contact email address. This email address cannot be on the same domain as the `customerDomain`")),
		mcp.WithString("etag", mcp.Description("Input parameter: ETag of the resource.")),
		mcp.WithString("kind", mcp.Description("Input parameter: Identifies the resource as a customer. Value: `admin#directory#customer`")),
		mcp.WithString("id", mcp.Description("Input parameter: The unique ID for the customer's Google Workspace account. (Readonly)")),
		mcp.WithString("language", mcp.Description("Input parameter: The customer's ISO 639-2 language code. See the [Language Codes](/admin-sdk/directory/v1/languages) page for the list of supported codes. Valid language codes outside the supported set will be accepted by the API but may lead to unexpected behavior. The default value is `en`.")),
		mcp.WithObject("postalAddress", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Directory_customers_patchHandler(cfg),
	}
}
