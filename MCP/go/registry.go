package main

import (
	"github.com/admin-sdk-api/mcp-server/config"
	"github.com/admin-sdk-api/mcp-server/models"
	tools_orgunits "github.com/admin-sdk-api/mcp-server/tools/orgunits"
	tools_customer "github.com/admin-sdk-api/mcp-server/tools/customer"
	tools_groups "github.com/admin-sdk-api/mcp-server/tools/groups"
	tools_tokens "github.com/admin-sdk-api/mcp-server/tools/tokens"
	tools_users "github.com/admin-sdk-api/mcp-server/tools/users"
	tools_twostepverification "github.com/admin-sdk-api/mcp-server/tools/twostepverification"
	tools_chromeosdevices "github.com/admin-sdk-api/mcp-server/tools/chromeosdevices"
	tools_mobiledevices "github.com/admin-sdk-api/mcp-server/tools/mobiledevices"
	tools_resources "github.com/admin-sdk-api/mcp-server/tools/resources"
	tools_customers "github.com/admin-sdk-api/mcp-server/tools/customers"
	tools_domainaliases "github.com/admin-sdk-api/mcp-server/tools/domainaliases"
	tools_members "github.com/admin-sdk-api/mcp-server/tools/members"
	tools_verificationcodes "github.com/admin-sdk-api/mcp-server/tools/verificationcodes"
	tools_roleassignments "github.com/admin-sdk-api/mcp-server/tools/roleassignments"
	tools_channels "github.com/admin-sdk-api/mcp-server/tools/channels"
	tools_domains "github.com/admin-sdk-api/mcp-server/tools/domains"
	tools_roles "github.com/admin-sdk-api/mcp-server/tools/roles"
	tools_privileges "github.com/admin-sdk-api/mcp-server/tools/privileges"
	tools_schemas "github.com/admin-sdk-api/mcp-server/tools/schemas"
	tools_asps "github.com/admin-sdk-api/mcp-server/tools/asps"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_orgunits.CreateDirectory_orgunits_listTool(cfg),
		tools_orgunits.CreateDirectory_orgunits_insertTool(cfg),
		tools_customer.CreateAdmin_customer_devices_chromeos_batchchangestatusTool(cfg),
		tools_groups.CreateDirectory_groups_aliases_listTool(cfg),
		tools_groups.CreateDirectory_groups_aliases_insertTool(cfg),
		tools_tokens.CreateDirectory_tokens_listTool(cfg),
		tools_users.CreateDirectory_users_watchTool(cfg),
		tools_twostepverification.CreateDirectory_twostepverification_turnoffTool(cfg),
		tools_chromeosdevices.CreateDirectory_chromeosdevices_listTool(cfg),
		tools_mobiledevices.CreateDirectory_mobiledevices_deleteTool(cfg),
		tools_mobiledevices.CreateDirectory_mobiledevices_getTool(cfg),
		tools_resources.CreateDirectory_resources_features_deleteTool(cfg),
		tools_resources.CreateDirectory_resources_features_getTool(cfg),
		tools_resources.CreateDirectory_resources_features_patchTool(cfg),
		tools_resources.CreateDirectory_resources_features_updateTool(cfg),
		tools_customers.CreateAdmin_customers_chrome_printers_listprintermodelsTool(cfg),
		tools_mobiledevices.CreateDirectory_mobiledevices_listTool(cfg),
		tools_domainaliases.CreateDirectory_domainaliases_deleteTool(cfg),
		tools_domainaliases.CreateDirectory_domainaliases_getTool(cfg),
		tools_members.CreateDirectory_members_deleteTool(cfg),
		tools_members.CreateDirectory_members_getTool(cfg),
		tools_members.CreateDirectory_members_patchTool(cfg),
		tools_members.CreateDirectory_members_updateTool(cfg),
		tools_groups.CreateDirectory_groups_aliases_deleteTool(cfg),
		tools_verificationcodes.CreateDirectory_verificationcodes_invalidateTool(cfg),
		tools_groups.CreateDirectory_groups_insertTool(cfg),
		tools_groups.CreateDirectory_groups_listTool(cfg),
		tools_customers.CreateDirectory_customers_getTool(cfg),
		tools_customers.CreateDirectory_customers_patchTool(cfg),
		tools_customers.CreateDirectory_customers_updateTool(cfg),
		tools_verificationcodes.CreateDirectory_verificationcodes_listTool(cfg),
		tools_users.CreateDirectory_users_insertTool(cfg),
		tools_users.CreateDirectory_users_listTool(cfg),
		tools_resources.CreateDirectory_resources_buildings_listTool(cfg),
		tools_resources.CreateDirectory_resources_buildings_insertTool(cfg),
		tools_groups.CreateDirectory_groups_deleteTool(cfg),
		tools_groups.CreateDirectory_groups_getTool(cfg),
		tools_groups.CreateDirectory_groups_patchTool(cfg),
		tools_groups.CreateDirectory_groups_updateTool(cfg),
		tools_members.CreateDirectory_members_listTool(cfg),
		tools_members.CreateDirectory_members_insertTool(cfg),
		tools_verificationcodes.CreateDirectory_verificationcodes_generateTool(cfg),
		tools_roleassignments.CreateDirectory_roleassignments_deleteTool(cfg),
		tools_roleassignments.CreateDirectory_roleassignments_getTool(cfg),
		tools_channels.CreateAdmin_channels_stopTool(cfg),
		tools_domains.CreateDirectory_domains_deleteTool(cfg),
		tools_domains.CreateDirectory_domains_getTool(cfg),
		tools_customers.CreateAdmin_customers_chrome_printservers_batchdeleteprintserversTool(cfg),
		tools_chromeosdevices.CreateDirectory_chromeosdevices_movedevicestoouTool(cfg),
		tools_customers.CreateAdmin_customers_chrome_printservers_listTool(cfg),
		tools_customers.CreateAdmin_customers_chrome_printservers_createTool(cfg),
		tools_customers.CreateAdmin_customers_chrome_printers_listTool(cfg),
		tools_customers.CreateAdmin_customers_chrome_printers_createTool(cfg),
		tools_mobiledevices.CreateDirectory_mobiledevices_actionTool(cfg),
		tools_chromeosdevices.CreateDirectory_chromeosdevices_actionTool(cfg),
		tools_resources.CreateDirectory_resources_buildings_getTool(cfg),
		tools_resources.CreateDirectory_resources_buildings_patchTool(cfg),
		tools_resources.CreateDirectory_resources_buildings_updateTool(cfg),
		tools_resources.CreateDirectory_resources_buildings_deleteTool(cfg),
		tools_resources.CreateDirectory_resources_calendars_deleteTool(cfg),
		tools_resources.CreateDirectory_resources_calendars_getTool(cfg),
		tools_resources.CreateDirectory_resources_calendars_patchTool(cfg),
		tools_resources.CreateDirectory_resources_calendars_updateTool(cfg),
		tools_users.CreateDirectory_users_makeadminTool(cfg),
		tools_customer.CreateAdmin_customer_devices_chromeos_commands_getTool(cfg),
		tools_customer.CreateAdmin_customer_devices_chromeos_issuecommandTool(cfg),
		tools_users.CreateDirectory_users_aliases_watchTool(cfg),
		tools_users.CreateDirectory_users_updateTool(cfg),
		tools_users.CreateDirectory_users_deleteTool(cfg),
		tools_users.CreateDirectory_users_getTool(cfg),
		tools_users.CreateDirectory_users_patchTool(cfg),
		tools_roles.CreateDirectory_roles_listTool(cfg),
		tools_roles.CreateDirectory_roles_insertTool(cfg),
		tools_users.CreateDirectory_users_aliases_deleteTool(cfg),
		tools_users.CreateDirectory_users_signoutTool(cfg),
		tools_domains.CreateDirectory_domains_listTool(cfg),
		tools_domains.CreateDirectory_domains_insertTool(cfg),
		tools_privileges.CreateDirectory_privileges_listTool(cfg),
		tools_schemas.CreateDirectory_schemas_listTool(cfg),
		tools_schemas.CreateDirectory_schemas_insertTool(cfg),
		tools_roles.CreateDirectory_roles_deleteTool(cfg),
		tools_roles.CreateDirectory_roles_getTool(cfg),
		tools_roles.CreateDirectory_roles_patchTool(cfg),
		tools_roles.CreateDirectory_roles_updateTool(cfg),
		tools_users.CreateDirectory_users_undeleteTool(cfg),
		tools_customers.CreateAdmin_customers_chrome_printers_batchcreateprintersTool(cfg),
		tools_members.CreateDirectory_members_hasmemberTool(cfg),
		tools_users.CreateDirectory_users_aliases_listTool(cfg),
		tools_users.CreateDirectory_users_aliases_insertTool(cfg),
		tools_roleassignments.CreateDirectory_roleassignments_listTool(cfg),
		tools_roleassignments.CreateDirectory_roleassignments_insertTool(cfg),
		tools_domainaliases.CreateDirectory_domainaliases_listTool(cfg),
		tools_domainaliases.CreateDirectory_domainaliases_insertTool(cfg),
		tools_customers.CreateAdmin_customers_chrome_printservers_deleteTool(cfg),
		tools_customers.CreateAdmin_customers_chrome_printservers_getTool(cfg),
		tools_customers.CreateAdmin_customers_chrome_printservers_patchTool(cfg),
		tools_asps.CreateDirectory_asps_listTool(cfg),
		tools_users.CreateDirectory_users_photos_deleteTool(cfg),
		tools_users.CreateDirectory_users_photos_getTool(cfg),
		tools_users.CreateDirectory_users_photos_patchTool(cfg),
		tools_users.CreateDirectory_users_photos_updateTool(cfg),
		tools_tokens.CreateDirectory_tokens_deleteTool(cfg),
		tools_tokens.CreateDirectory_tokens_getTool(cfg),
		tools_resources.CreateDirectory_resources_features_renameTool(cfg),
		tools_schemas.CreateDirectory_schemas_patchTool(cfg),
		tools_schemas.CreateDirectory_schemas_updateTool(cfg),
		tools_schemas.CreateDirectory_schemas_deleteTool(cfg),
		tools_schemas.CreateDirectory_schemas_getTool(cfg),
		tools_chromeosdevices.CreateDirectory_chromeosdevices_getTool(cfg),
		tools_chromeosdevices.CreateDirectory_chromeosdevices_patchTool(cfg),
		tools_chromeosdevices.CreateDirectory_chromeosdevices_updateTool(cfg),
		tools_asps.CreateDirectory_asps_getTool(cfg),
		tools_asps.CreateDirectory_asps_deleteTool(cfg),
		tools_orgunits.CreateDirectory_orgunits_deleteTool(cfg),
		tools_orgunits.CreateDirectory_orgunits_getTool(cfg),
		tools_orgunits.CreateDirectory_orgunits_patchTool(cfg),
		tools_orgunits.CreateDirectory_orgunits_updateTool(cfg),
		tools_customers.CreateAdmin_customers_chrome_printservers_batchcreateprintserversTool(cfg),
		tools_customers.CreateAdmin_customers_chrome_printers_batchdeleteprintersTool(cfg),
		tools_resources.CreateDirectory_resources_calendars_listTool(cfg),
		tools_resources.CreateDirectory_resources_calendars_insertTool(cfg),
		tools_resources.CreateDirectory_resources_features_listTool(cfg),
		tools_resources.CreateDirectory_resources_features_insertTool(cfg),
	}
}
