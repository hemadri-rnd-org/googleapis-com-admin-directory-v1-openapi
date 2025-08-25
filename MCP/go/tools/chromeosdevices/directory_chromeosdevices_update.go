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

func Directory_chromeosdevices_updateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["projection"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("projection=%v", val))
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
		var requestBody models.ChromeOsDevice
		
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
		url := fmt.Sprintf("%s/admin/directory/v1/customer/%s/devices/chromeos/%s%s", cfg.BaseURL, customerId, deviceId, queryString)
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
		var result models.ChromeOsDevice
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

func CreateDirectory_chromeosdevices_updateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_admin_directory_v1_customer_customerId_devices_chromeos_deviceId",
		mcp.WithDescription("Updates a device's updatable properties, such as `annotatedUser`, `annotatedLocation`, `notes`, `orgUnitPath`, or `annotatedAssetId`."),
		mcp.WithString("customerId", mcp.Required(), mcp.Description("The unique ID for the customer's Google Workspace account. As an account administrator, you can also use the `my_customer` alias to represent your account's `customerId`. The `customerId` is also returned as part of the [Users resource](/admin-sdk/directory/v1/reference/users).")),
		mcp.WithString("deviceId", mcp.Required(), mcp.Description("The unique ID of the device. The `deviceId`s are returned in the response from the [chromeosdevices.list](/admin-sdk/v1/reference/chromeosdevices/list) method.")),
		mcp.WithString("projection", mcp.Description("Restrict information returned to a set of selected fields.")),
		mcp.WithString("manufactureDate", mcp.Description("Input parameter: (Read-only) The date the device was manufactured in yyyy-mm-dd format.")),
		mcp.WithString("deviceLicenseType", mcp.Description("Input parameter: Output only. Device license type.")),
		mcp.WithArray("recentUsers", mcp.Description("Input parameter: A list of recent device users, in descending order, by last login time.")),
		mcp.WithObject("tpmVersionInfo", mcp.Description("Input parameter: Trusted Platform Module (TPM) (Read-only)")),
		mcp.WithString("firstEnrollmentTime", mcp.Description("Input parameter: Date and time for the first time the device was enrolled.")),
		mcp.WithString("firmwareVersion", mcp.Description("Input parameter: The Chrome device's firmware version.")),
		mcp.WithArray("deviceFiles", mcp.Description("Input parameter: A list of device files to download (Read-only)")),
		mcp.WithString("status", mcp.Description("Input parameter: The status of the device.")),
		mcp.WithString("meid", mcp.Description("Input parameter: The Mobile Equipment Identifier (MEID) or the International Mobile Equipment Identity (IMEI) for the 3G mobile card in a mobile device. A MEID/IMEI is typically used when adding a device to a wireless carrier's post-pay service plan. If the device does not have this information, this property is not included in the response. For more information on how to export a MEID/IMEI list, see the [Developer's Guide](/admin-sdk/directory/v1/guides/manage-chrome-devices.html#export_meid).")),
		mcp.WithString("osVersion", mcp.Description("Input parameter: The Chrome device's operating system version.")),
		mcp.WithString("supportEndDate", mcp.Description("Input parameter: Final date the device will be supported (Read-only)")),
		mcp.WithString("model", mcp.Description("Input parameter: The device's model information. If the device does not have this information, this property is not included in the response.")),
		mcp.WithArray("activeTimeRanges", mcp.Description("Input parameter: A list of active time ranges (Read-only).")),
		mcp.WithArray("cpuInfo", mcp.Description("Input parameter: Information regarding CPU specs in the device.")),
		mcp.WithString("dockMacAddress", mcp.Description("Input parameter: (Read-only) Built-in MAC address for the docking station that the device connected to. Factory sets Media access control address (MAC address) assigned for use by a dock. It is reserved specifically for MAC pass through device policy. The format is twelve (12) hexadecimal digits without any delimiter (uppercase letters). This is only relevant for some devices.")),
		mcp.WithString("lastDeprovisionTimestamp", mcp.Description("Input parameter: (Read-only) Date and time for the last deprovision of the device.")),
		mcp.WithString("serialNumber", mcp.Description("Input parameter: The Chrome device serial number entered when the device was enabled. This value is the same as the Admin console's *Serial Number* in the *Chrome OS Devices* tab.")),
		mcp.WithString("etag", mcp.Description("Input parameter: ETag of the resource.")),
		mcp.WithString("macAddress", mcp.Description("Input parameter: The device's wireless MAC address. If the device does not have this information, it is not included in the response.")),
		mcp.WithString("systemRamTotal", mcp.Description("Input parameter: Total RAM on the device [in bytes] (Read-only)")),
		mcp.WithString("annotatedUser", mcp.Description("Input parameter: The user of the device as noted by the administrator. Maximum length is 100 characters. Empty values are allowed.")),
		mcp.WithString("autoUpdateExpiration", mcp.Description("Input parameter: (Read-only) The timestamp after which the device will stop receiving Chrome updates or support")),
		mcp.WithString("orgUnitPath", mcp.Description("Input parameter: The full parent path with the organizational unit's name associated with the device. Path names are case insensitive. If the parent organizational unit is the top-level organization, it is represented as a forward slash, `/`. This property can be [updated](/admin-sdk/directory/v1/guides/manage-chrome-devices#move_chrome_devices_to_ou) using the API. For more information about how to create an organizational structure for your device, see the [administration help center](https://support.google.com/a/answer/182433).")),
		mcp.WithArray("diskVolumeReports", mcp.Description("Input parameter: Reports of disk space and other info about mounted/connected volumes.")),
		mcp.WithString("deviceId", mcp.Description("Input parameter: The unique ID of the Chrome device.")),
		mcp.WithString("orderNumber", mcp.Description("Input parameter: The device's order number. Only devices directly purchased from Google have an order number.")),
		mcp.WithArray("lastKnownNetwork", mcp.Description("Input parameter: Contains last known network (Read-only)")),
		mcp.WithString("notes", mcp.Description("Input parameter: Notes about this device added by the administrator. This property can be [searched](https://support.google.com/chrome/a/answer/1698333) with the [list](/admin-sdk/directory/v1/reference/chromeosdevices/list) method's `query` parameter. Maximum length is 500 characters. Empty values are allowed.")),
		mcp.WithString("lastEnrollmentTime", mcp.Description("Input parameter: Date and time the device was last enrolled (Read-only)")),
		mcp.WithArray("systemRamFreeReports", mcp.Description("Input parameter: Reports of amounts of available RAM memory (Read-only)")),
		mcp.WithString("ethernetMacAddress", mcp.Description("Input parameter: The device's MAC address on the ethernet network interface.")),
		mcp.WithString("deprovisionReason", mcp.Description("Input parameter: (Read-only) Deprovision reason.")),
		mcp.WithString("kind", mcp.Description("Input parameter: The type of resource. For the Chromeosdevices resource, the value is `admin#directory#chromeosdevice`.")),
		mcp.WithArray("cpuStatusReports", mcp.Description("Input parameter: Reports of CPU utilization and temperature (Read-only)")),
		mcp.WithString("ethernetMacAddress0", mcp.Description("Input parameter: (Read-only) MAC address used by the Chromebook’s internal ethernet port, and for onboard network (ethernet) interface. The format is twelve (12) hexadecimal digits without any delimiter (uppercase letters). This is only relevant for some devices.")),
		mcp.WithBoolean("willAutoRenew", mcp.Description("Input parameter: Determines if the device will auto renew its support after the support end date. This is a read-only property.")),
		mcp.WithString("orgUnitId", mcp.Description("Input parameter: The unique ID of the organizational unit. orgUnitPath is the human readable version of orgUnitId. While orgUnitPath may change by renaming an organizational unit within the path, orgUnitId is unchangeable for one organizational unit. This property can be [updated](/admin-sdk/directory/v1/guides/manage-chrome-devices#move_chrome_devices_to_ou) using the API. For more information about how to create an organizational structure for your device, see the [administration help center](https://support.google.com/a/answer/182433).")),
		mcp.WithString("bootMode", mcp.Description("Input parameter: The boot mode for the device. The possible values are: * `Verified`: The device is running a valid version of the Chrome OS. * `Dev`: The devices's developer hardware switch is enabled. When booted, the device has a command line shell. For an example of a developer switch, see the [Chromebook developer information](https://www.chromium.org/chromium-os/developer-information-for-chrome-os-devices/samsung-series-5-chromebook#TOC-Developer-switch).")),
		mcp.WithString("lastSync", mcp.Description("Input parameter: Date and time the device was last synchronized with the policy settings in the G Suite administrator control panel (Read-only)")),
		mcp.WithObject("osUpdateStatus", mcp.Description("Input parameter: Contains information regarding the current OS update status.")),
		mcp.WithArray("screenshotFiles", mcp.Description("Input parameter: A list of screenshot files to download. Type is always \"SCREENSHOT_FILE\". (Read-only)")),
		mcp.WithString("annotatedLocation", mcp.Description("Input parameter: The address or location of the device as noted by the administrator. Maximum length is `200` characters. Empty values are allowed.")),
		mcp.WithString("annotatedAssetId", mcp.Description("Input parameter: The asset identifier as noted by an administrator or specified during enrollment.")),
		mcp.WithString("platformVersion", mcp.Description("Input parameter: The Chrome device's platform version.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Directory_chromeosdevices_updateHandler(cfg),
	}
}
