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

func Directory_resources_calendars_insertHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		customerVal, ok := args["customer"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: customer"), nil
		}
		customer, ok := customerVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: customer"), nil
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
		var requestBody models.CalendarResource
		
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
		url := fmt.Sprintf("%s/admin/directory/v1/customer/%s/resources/calendars%s", cfg.BaseURL, customer, queryString)
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
		var result models.CalendarResource
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

func CreateDirectory_resources_calendars_insertTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_admin_directory_v1_customer_customer_resources_calendars",
		mcp.WithDescription("Inserts a calendar resource."),
		mcp.WithString("customer", mcp.Required(), mcp.Description("The unique ID for the customer's Google Workspace account. As an account administrator, you can also use the `my_customer` alias to represent your account's customer ID.")),
		mcp.WithString("resourceName", mcp.Description("Input parameter: The name of the calendar resource. For example, \"Training Room 1A\".")),
		mcp.WithString("resourceType", mcp.Description("Input parameter: The type of the calendar resource, intended for non-room resources.")),
		mcp.WithString("resourceCategory", mcp.Description("Input parameter: The category of the calendar resource. Either CONFERENCE_ROOM or OTHER. Legacy data is set to CATEGORY_UNKNOWN.")),
		mcp.WithString("generatedResourceName", mcp.Description("Input parameter: The read-only auto-generated name of the calendar resource which includes metadata about the resource such as building name, floor, capacity, etc. For example, \"NYC-2-Training Room 1A (16)\".")),
		mcp.WithString("userVisibleDescription", mcp.Description("Input parameter: Description of the resource, visible to users and admins.")),
		mcp.WithString("floorName", mcp.Description("Input parameter: Name of the floor a resource is located on.")),
		mcp.WithNumber("capacity", mcp.Description("Input parameter: Capacity of a resource, number of seats in a room.")),
		mcp.WithString("kind", mcp.Description("Input parameter: The type of the resource. For calendar resources, the value is `admin#directory#resources#calendars#CalendarResource`.")),
		mcp.WithString("resourceDescription", mcp.Description("Input parameter: Description of the resource, visible only to admins.")),
		mcp.WithString("featureInstances", mcp.Description("Input parameter: Instances of features for the calendar resource.")),
		mcp.WithString("resourceEmail", mcp.Description("Input parameter: The read-only email for the calendar resource. Generated as part of creating a new calendar resource.")),
		mcp.WithString("buildingId", mcp.Description("Input parameter: Unique ID for the building a resource is located in.")),
		mcp.WithString("floorSection", mcp.Description("Input parameter: Name of the section within a floor a resource is located in.")),
		mcp.WithString("resourceId", mcp.Description("Input parameter: The unique ID for the calendar resource.")),
		mcp.WithString("etags", mcp.Description("Input parameter: ETag of the resource.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Directory_resources_calendars_insertHandler(cfg),
	}
}
