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

func Directory_users_insertHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["resolveConflictAccount"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("resolveConflictAccount=%v", val))
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
		var requestBody models.User
		
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
		url := fmt.Sprintf("%s/admin/directory/v1/users%s", cfg.BaseURL, queryString)
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
		var result models.User
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

func CreateDirectory_users_insertTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_admin_directory_v1_users",
		mcp.WithDescription("Creates a user. Mutate calls immediately following user creation might sometimes fail as the user isn't fully created due to propagation delay in our backends. Check the error details for the "User creation is not complete" message to see if this is the case. Retrying the calls after some time can help in this case."),
		mcp.WithBoolean("resolveConflictAccount", mcp.Description("Optional. If set to `true`, the option selected for [handling unmanaged user accounts](https://support.google.com/a/answer/11112794) will apply. Default: `false`")),
		mcp.WithString("thumbnailPhotoUrl", mcp.Description("Input parameter: Output only. The URL of the user's profile photo. The URL might be temporary or private.")),
		mcp.WithString("creationTime", mcp.Description("Input parameter: User's G Suite account creation time. (Read-only)")),
		mcp.WithString("lastLoginTime", mcp.Description("Input parameter: User's last login time. (Read-only)")),
		mcp.WithString("locations", mcp.Description("Input parameter: The user's locations. The maximum allowed data size for this field is 10KB.")),
		mcp.WithObject("customSchemas", mcp.Description("Input parameter: Custom fields of the user. The key is a `schema_name` and its values are `'field_name': 'field_value'`.")),
		mcp.WithString("orgUnitPath", mcp.Description("Input parameter: The full path of the parent organization associated with the user. If the parent organization is the top-level, it is represented as a forward slash (`/`).")),
		mcp.WithString("keywords", mcp.Description("Input parameter: The list of the user's keywords. The maximum allowed data size for this field is 1KB.")),
		mcp.WithString("recoveryPhone", mcp.Description("Input parameter: Recovery phone of the user. The phone number must be in the E.164 format, starting with the plus sign (+). Example: *+16506661212*.")),
		mcp.WithString("hashFunction", mcp.Description("Input parameter: Stores the hash format of the `password` property. The following `hashFunction` values are allowed: * `MD5` - Accepts simple hex-encoded values. * `SHA-1` - Accepts simple hex-encoded values. * `crypt` - Compliant with the [C crypt library](https://en.wikipedia.org/wiki/Crypt_%28C%29). Supports the DES, MD5 (hash prefix `$1$`), SHA-256 (hash prefix `$5$`), and SHA-512 (hash prefix `$6$`) hash algorithms. If rounds are specified as part of the prefix, they must be 10,000 or fewer.")),
		mcp.WithString("deletionTime", mcp.Description("")),
		mcp.WithObject("name", mcp.Description("")),
		mcp.WithString("externalIds", mcp.Description("Input parameter: The list of external IDs for the user, such as an employee or network ID. The maximum allowed data size for this field is 2KB.")),
		mcp.WithBoolean("agreedToTerms", mcp.Description("Input parameter: Output only. This property is `true` if the user has completed an initial login and accepted the Terms of Service agreement.")),
		mcp.WithArray("aliases", mcp.Description("Input parameter: Output only. The list of the user's alias email addresses.")),
		mcp.WithString("relations", mcp.Description("Input parameter: The list of the user's relationships to other users. The maximum allowed data size for this field is 2KB.")),
		mcp.WithString("primaryEmail", mcp.Description("Input parameter: The user's primary email address. This property is required in a request to create a user account. The `primaryEmail` must be unique and cannot be an alias of another user.")),
		mcp.WithString("gender", mcp.Description("Input parameter: The user's gender. The maximum allowed data size for this field is 1KB.")),
		mcp.WithString("posixAccounts", mcp.Description("Input parameter: The list of [POSIX](https://www.opengroup.org/austin/papers/posix_faq.html) account information for the user.")),
		mcp.WithBoolean("isMailboxSetup", mcp.Description("Input parameter: Output only. Indicates if the user's Google mailbox is created. This property is only applicable if the user has been assigned a Gmail license.")),
		mcp.WithString("organizations", mcp.Description("Input parameter: The list of organizations the user belongs to. The maximum allowed data size for this field is 10KB.")),
		mcp.WithString("kind", mcp.Description("Input parameter: Output only. The type of the API resource. For Users resources, the value is `admin#directory#user`.")),
		mcp.WithString("websites", mcp.Description("Input parameter: The user's websites. The maximum allowed data size for this field is 2KB.")),
		mcp.WithString("languages", mcp.Description("Input parameter: The user's languages. The maximum allowed data size for this field is 1KB.")),
		mcp.WithString("phones", mcp.Description("Input parameter: The list of the user's phone numbers. The maximum allowed data size for this field is 1KB.")),
		mcp.WithString("etag", mcp.Description("Input parameter: Output only. ETag of the resource.")),
		mcp.WithString("thumbnailPhotoEtag", mcp.Description("Input parameter: Output only. ETag of the user's photo (Read-only)")),
		mcp.WithString("emails", mcp.Description("Input parameter: The list of the user's email addresses. The maximum allowed data size for this field is 10KB. This excludes `publicKeyEncryptionCertificates`.")),
		mcp.WithString("id", mcp.Description("Input parameter: The unique ID for the user. A user `id` can be used as a user request URI's `userKey`.")),
		mcp.WithBoolean("ipWhitelisted", mcp.Description("Input parameter: If `true`, the user's IP address is subject to a deprecated IP address [`allowlist`](https://support.google.com/a/answer/60752) configuration.")),
		mcp.WithBoolean("includeInGlobalAddressList", mcp.Description("Input parameter: Indicates if the user's profile is visible in the Google Workspace global address list when the contact sharing feature is enabled for the domain. For more information about excluding user profiles, see the [administration help center](https://support.google.com/a/answer/1285988).")),
		mcp.WithBoolean("isAdmin", mcp.Description("Input parameter: Output only. Indicates a user with super admininistrator privileges. The `isAdmin` property can only be edited in the [Make a user an administrator](/admin-sdk/directory/v1/guides/manage-users.html#make_admin) operation ( [makeAdmin](/admin-sdk/directory/v1/reference/users/makeAdmin.html) method). If edited in the user [insert](/admin-sdk/directory/v1/reference/users/insert.html) or [update](/admin-sdk/directory/v1/reference/users/update.html) methods, the edit is ignored by the API service.")),
		mcp.WithArray("nonEditableAliases", mcp.Description("Input parameter: Output only. The list of the user's non-editable alias email addresses. These are typically outside the account's primary domain or sub-domain.")),
		mcp.WithBoolean("isEnrolledIn2Sv", mcp.Description("Input parameter: Output only. Is enrolled in 2-step verification (Read-only)")),
		mcp.WithBoolean("archived", mcp.Description("Input parameter: Indicates if user is archived.")),
		mcp.WithString("notes", mcp.Description("Input parameter: Notes for the user.")),
		mcp.WithString("sshPublicKeys", mcp.Description("Input parameter: A list of SSH public keys.")),
		mcp.WithString("addresses", mcp.Description("Input parameter: The list of the user's addresses. The maximum allowed data size for this field is 10KB.")),
		mcp.WithString("recoveryEmail", mcp.Description("Input parameter: Recovery email of the user.")),
		mcp.WithString("customerId", mcp.Description("Input parameter: Output only. The customer ID to [retrieve all account users](/admin-sdk/directory/v1/guides/manage-users.html#get_all_users). You can use the alias `my_customer` to represent your account's `customerId`. As a reseller administrator, you can use the resold customer account's `customerId`. To get a `customerId`, use the account's primary domain in the `domain` parameter of a [users.list](/admin-sdk/directory/v1/reference/users/list) request.")),
		mcp.WithString("suspensionReason", mcp.Description("Input parameter: Output only. Has the reason a user account is suspended either by the administrator or by Google at the time of suspension. The property is returned only if the `suspended` property is `true`.")),
		mcp.WithBoolean("isEnforcedIn2Sv", mcp.Description("Input parameter: Output only. Is 2-step verification enforced (Read-only)")),
		mcp.WithBoolean("suspended", mcp.Description("Input parameter: Indicates if user is suspended.")),
		mcp.WithBoolean("isDelegatedAdmin", mcp.Description("Input parameter: Output only. Indicates if the user is a delegated administrator. Delegated administrators are supported by the API but cannot create or undelete users, or make users administrators. These requests are ignored by the API service. Roles and privileges for administrators are assigned using the [Admin console](https://support.google.com/a/answer/33325).")),
		mcp.WithBoolean("changePasswordAtNextLogin", mcp.Description("Input parameter: Indicates if the user is forced to change their password at next login. This setting doesn't apply when [the user signs in via a third-party identity provider](https://support.google.com/a/answer/60224).")),
		mcp.WithString("password", mcp.Description("Input parameter: User's password")),
		mcp.WithString("ims", mcp.Description("Input parameter: The list of the user's Instant Messenger (IM) accounts. A user account can have multiple ims properties. But, only one of these ims properties can be the primary IM contact. The maximum allowed data size for this field is 2KB.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Directory_users_insertHandler(cfg),
	}
}
