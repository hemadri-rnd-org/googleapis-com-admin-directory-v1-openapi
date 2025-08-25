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

func Directory_resources_buildings_updateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		buildingIdVal, ok := args["buildingId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: buildingId"), nil
		}
		buildingId, ok := buildingIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: buildingId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["coordinatesSource"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("coordinatesSource=%v", val))
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
		var requestBody models.Building
		
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
		url := fmt.Sprintf("%s/admin/directory/v1/customer/%s/resources/buildings/%s%s", cfg.BaseURL, customer, buildingId, queryString)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
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
		var result models.Building
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

func CreateDirectory_resources_buildings_updateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_admin_directory_v1_customer_customer_resources_buildings_buildingId",
		mcp.WithDescription("Updates a building."),
		mcp.WithString("customer", mcp.Required(), mcp.Description("The unique ID for the customer's Google Workspace account. As an account administrator, you can also use the `my_customer` alias to represent your account's customer ID.")),
		mcp.WithString("buildingId", mcp.Required(), mcp.Description("The id of the building to update.")),
		mcp.WithString("coordinatesSource", mcp.Description("Source from which Building.coordinates are derived.")),
		mcp.WithString("description", mcp.Description("Input parameter: A brief description of the building. For example, \"Chelsea Market\".")),
		mcp.WithString("etags", mcp.Description("Input parameter: ETag of the resource.")),
		mcp.WithArray("floorNames", mcp.Description("Input parameter: The display names for all floors in this building. The floors are expected to be sorted in ascending order, from lowest floor to highest floor. For example, [\"B2\", \"B1\", \"L\", \"1\", \"2\", \"2M\", \"3\", \"PH\"] Must contain at least one entry.")),
		mcp.WithString("kind", mcp.Description("Input parameter: Kind of resource this is.")),
		mcp.WithObject("address", mcp.Description("Input parameter: Public API: Resources.buildings")),
		mcp.WithString("buildingId", mcp.Description("Input parameter: Unique identifier for the building. The maximum length is 100 characters.")),
		mcp.WithString("buildingName", mcp.Description("Input parameter: The building name as seen by users in Calendar. Must be unique for the customer. For example, \"NYC-CHEL\". The maximum length is 100 characters.")),
		mcp.WithObject("coordinates", mcp.Description("Input parameter: Public API: Resources.buildings")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Directory_resources_buildings_updateHandler(cfg),
	}
}
