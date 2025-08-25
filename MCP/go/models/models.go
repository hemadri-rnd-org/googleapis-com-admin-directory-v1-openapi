package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Role represents the Role schema from the OpenAPI specification
type Role struct {
	Rolename string `json:"roleName,omitempty"` // Name of the role.
	Roleprivileges []map[string]interface{} `json:"rolePrivileges,omitempty"` // The set of privileges that are granted to this role.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Issuperadminrole bool `json:"isSuperAdminRole,omitempty"` // Returns `true` if the role is a super admin role.
	Issystemrole bool `json:"isSystemRole,omitempty"` // Returns `true` if this is a pre-defined system role.
	Kind string `json:"kind,omitempty"` // The type of the API resource. This is always `admin#directory#role`.
	Roledescription string `json:"roleDescription,omitempty"` // A short description of the role.
	Roleid string `json:"roleId,omitempty"` // ID of the role.
}

// GroupAlias represents the GroupAlias schema from the OpenAPI specification
type GroupAlias struct {
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Id string `json:"id,omitempty"` // The unique ID of the group.
	Kind string `json:"kind,omitempty"` // The type of the API resource. For Alias resources, the value is `admin#directory#alias`.
	Primaryemail string `json:"primaryEmail,omitempty"` // The primary email address of the group.
	Alias string `json:"alias,omitempty"` // The alias email address.
}

// FailureInfo represents the FailureInfo schema from the OpenAPI specification
type FailureInfo struct {
	Errorcode string `json:"errorCode,omitempty"` // Canonical code for why the update failed to apply.
	Errormessage string `json:"errorMessage,omitempty"` // Failure reason message.
	Printer Printer `json:"printer,omitempty"` // Printer configuration.
	Printerid string `json:"printerId,omitempty"` // Id of a failed printer.
}

// PrintServerFailureInfo represents the PrintServerFailureInfo schema from the OpenAPI specification
type PrintServerFailureInfo struct {
	Errorcode string `json:"errorCode,omitempty"` // Canonical code for why the update failed to apply.
	Errormessage string `json:"errorMessage,omitempty"` // Failure reason message.
	Printserver PrintServer `json:"printServer,omitempty"` // Configuration for a print server.
	Printserverid string `json:"printServerId,omitempty"` // ID of a failed print server.
}

// Asp represents the Asp schema from the OpenAPI specification
type Asp struct {
	Lasttimeused string `json:"lastTimeUsed,omitempty"` // The time when the ASP was last used. Expressed in [Unix time](https://en.wikipedia.org/wiki/Epoch_time) format.
	Name string `json:"name,omitempty"` // The name of the application that the user, represented by their `userId`, entered when the ASP was created.
	Userkey string `json:"userKey,omitempty"` // The unique ID of the user who issued the ASP.
	Codeid int `json:"codeId,omitempty"` // The unique ID of the ASP.
	Creationtime string `json:"creationTime,omitempty"` // The time when the ASP was created. Expressed in [Unix time](https://en.wikipedia.org/wiki/Epoch_time) format.
	Etag string `json:"etag,omitempty"` // ETag of the ASP.
	Kind string `json:"kind,omitempty"` // The type of the API resource. This is always `admin#directory#asp`.
}

// UserPosixAccount represents the UserPosixAccount schema from the OpenAPI specification
type UserPosixAccount struct {
	Shell string `json:"shell,omitempty"` // The path to the login shell for this account.
	Primary bool `json:"primary,omitempty"` // If this is user's primary account within the SystemId.
	Systemid string `json:"systemId,omitempty"` // System identifier for which account Username or Uid apply to.
	Username string `json:"username,omitempty"` // The username of the account.
	Uid string `json:"uid,omitempty"` // The POSIX compliant user ID.
	Accountid string `json:"accountId,omitempty"` // A POSIX account field identifier.
	Gid string `json:"gid,omitempty"` // The default group ID.
	Gecos string `json:"gecos,omitempty"` // The GECOS (user information) for this account.
	Homedirectory string `json:"homeDirectory,omitempty"` // The path to the home directory for this account.
	Operatingsystemtype string `json:"operatingSystemType,omitempty"` // The operating system type for this account.
}

// ChangeChromeOsDeviceStatusResult represents the ChangeChromeOsDeviceStatusResult schema from the OpenAPI specification
type ChangeChromeOsDeviceStatusResult struct {
	ErrorField Status `json:"error,omitempty"` // The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
	Response ChangeChromeOsDeviceStatusSucceeded `json:"response,omitempty"` // Response for a successful ChromeOS device status change.
	Deviceid string `json:"deviceId,omitempty"` // The unique ID of the ChromeOS device.
}

// OsUpdateStatus represents the OsUpdateStatus schema from the OpenAPI specification
type OsUpdateStatus struct {
	State string `json:"state,omitempty"` // The update state of an OS update.
	Targetkioskappversion string `json:"targetKioskAppVersion,omitempty"` // New required platform version from the pending updated kiosk app.
	Targetosversion string `json:"targetOsVersion,omitempty"` // New platform version of the OS image being downloaded and applied. It is only set when update status is UPDATE_STATUS_DOWNLOAD_IN_PROGRESS or UPDATE_STATUS_NEED_REBOOT. Note this could be a dummy "0.0.0.0" for UPDATE_STATUS_NEED_REBOOT for some edge cases, e.g. update engine is restarted without a reboot.
	Updatechecktime string `json:"updateCheckTime,omitempty"` // Date and time of the last update check.
	Updatetime string `json:"updateTime,omitempty"` // Date and time of the last successful OS update.
	Reboottime string `json:"rebootTime,omitempty"` // Date and time of the last reboot.
}

// ListPrinterModelsResponse represents the ListPrinterModelsResponse schema from the OpenAPI specification
type ListPrinterModelsResponse struct {
	Printermodels []PrinterModel `json:"printerModels,omitempty"` // Printer models that are currently allowed to be configured for ChromeOs. Some printers may be added or removed over time.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
}

// UserLanguage represents the UserLanguage schema from the OpenAPI specification
type UserLanguage struct {
	Customlanguage string `json:"customLanguage,omitempty"` // Other language. User can provide their own language name if there is no corresponding ISO 639 language code. If this is set, `languageCode` can't be set.
	Languagecode string `json:"languageCode,omitempty"` // ISO 639 string representation of a language. See [Language Codes](/admin-sdk/directory/v1/languages) for the list of supported codes. Valid language codes outside the supported set will be accepted by the API but may lead to unexpected behavior. Illegal values cause `SchemaException`. If this is set, `customLanguage` can't be set.
	Preference string `json:"preference,omitempty"` // Optional. If present, controls whether the specified `languageCode` is the user's preferred language. If `customLanguage` is set, this can't be set. Allowed values are `preferred` and `not_preferred`.
}

// UserRelation represents the UserRelation schema from the OpenAPI specification
type UserRelation struct {
	Customtype string `json:"customType,omitempty"` // Custom Type.
	TypeField string `json:"type,omitempty"` // The relation of the user. Some of the possible values are mother father sister brother manager assistant partner.
	Value string `json:"value,omitempty"` // The name of the relation.
}

// BatchChangeChromeOsDeviceStatusRequest represents the BatchChangeChromeOsDeviceStatusRequest schema from the OpenAPI specification
type BatchChangeChromeOsDeviceStatusRequest struct {
	Changechromeosdevicestatusaction string `json:"changeChromeOsDeviceStatusAction,omitempty"` // Required. The action to take on the ChromeOS device in order to change its status.
	Deprovisionreason string `json:"deprovisionReason,omitempty"` // Optional. The reason behind a device deprovision. Must be provided if 'changeChromeOsDeviceStatusAction' is set to 'CHANGE_CHROME_OS_DEVICE_STATUS_ACTION_DEPROVISION'. Otherwise, omit this field.
	Deviceids []string `json:"deviceIds,omitempty"` // Required. List of the IDs of the ChromeOS devices to change. Maximum 50.
}

// FeatureInstance represents the FeatureInstance schema from the OpenAPI specification
type FeatureInstance struct {
	Feature Feature `json:"feature,omitempty"` // JSON template for Feature object in Directory API.
}

// ChromeOsDevices represents the ChromeOsDevices schema from the OpenAPI specification
type ChromeOsDevices struct {
	Chromeosdevices []ChromeOsDevice `json:"chromeosdevices,omitempty"` // A list of Chrome OS Device objects.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // Kind of resource this is.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token used to access the next page of this result. To access the next page, use this token's value in the `pageToken` query string of this request.
}

// BatchDeletePrintServersRequest represents the BatchDeletePrintServersRequest schema from the OpenAPI specification
type BatchDeletePrintServersRequest struct {
	Printserverids []string `json:"printServerIds,omitempty"` // A list of print server IDs that should be deleted (max `100` per batch).
}

// MobileDeviceAction represents the MobileDeviceAction schema from the OpenAPI specification
type MobileDeviceAction struct {
	Action string `json:"action,omitempty"` // The action to be performed on the device.
}

// MobileDevice represents the MobileDevice schema from the OpenAPI specification
type MobileDevice struct {
	Unknownsourcesstatus bool `json:"unknownSourcesStatus,omitempty"` // Unknown sources enabled or disabled on device (Read-only)
	Releaseversion string `json:"releaseVersion,omitempty"` // Mobile Device release version version (Read-only)
	Serialnumber string `json:"serialNumber,omitempty"` // The device's serial number.
	Meid string `json:"meid,omitempty"` // The device's MEID number.
	Manufacturer string `json:"manufacturer,omitempty"` // Mobile Device manufacturer (Read-only)
	Otheraccountsinfo []string `json:"otherAccountsInfo,omitempty"` // The list of accounts added on device (Read-only)
	Adbstatus bool `json:"adbStatus,omitempty"` // Adb (USB debugging) enabled or disabled on device (Read-only)
	Email []string `json:"email,omitempty"` // The list of the owner's email addresses. If your application needs the current list of user emails, use the [get](/admin-sdk/directory/v1/reference/mobiledevices/get.html) method. For additional information, see the [retrieve a user](/admin-sdk/directory/v1/guides/manage-users#get_user) method.
	Lastsync string `json:"lastSync,omitempty"` // Date and time the device was last synchronized with the policy settings in the G Suite administrator control panel (Read-only)
	Encryptionstatus string `json:"encryptionStatus,omitempty"` // Mobile Device Encryption Status (Read-only)
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Name []string `json:"name,omitempty"` // The list of the owner's user names. If your application needs the current list of device owner names, use the [get](/admin-sdk/directory/v1/reference/mobiledevices/get.html) method. For more information about retrieving mobile device user information, see the [Developer's Guide](/admin-sdk/directory/v1/guides/manage-users#get_user).
	Applications []map[string]interface{} `json:"applications,omitempty"` // The list of applications installed on an Android mobile device. It is not applicable to Google Sync and iOS devices. The list includes any Android applications that access Google Workspace data. When updating an applications list, it is important to note that updates replace the existing list. If the Android device has two existing applications and the API updates the list with five applications, the is now the updated list of five applications.
	Buildnumber string `json:"buildNumber,omitempty"` // The device's operating system build number.
	Defaultlanguage string `json:"defaultLanguage,omitempty"` // The default locale used on the device.
	Hardwareid string `json:"hardwareId,omitempty"` // The IMEI/MEID unique identifier for Android hardware. It is not applicable to Google Sync devices. When adding an Android mobile device, this is an optional property. When updating one of these devices, this is a read-only property.
	Useragent string `json:"userAgent,omitempty"` // Gives information about the device such as `os` version. This property can be [updated](/admin-sdk/directory/v1/reference/mobiledevices/update.html). For more information, see the [Developer's Guide](/admin-sdk/directory/v1/guides/manage-mobile-devices#update_mobile_device).
	Privilege string `json:"privilege,omitempty"` // DMAgentPermission (Read-only)
	Model string `json:"model,omitempty"` // The mobile device's model name, for example Nexus S. This property can be [updated](/admin-sdk/directory/v1/reference/mobiledevices/update.html). For more information, see the [Developer's Guide](/admin-sdk/directory/v1/guides/manage-mobile=devices#update_mobile_device).
	Managedaccountisonownerprofile bool `json:"managedAccountIsOnOwnerProfile,omitempty"` // Boolean indicating if this account is on owner/primary profile or not.
	Hardware string `json:"hardware,omitempty"` // Mobile Device Hardware (Read-only)
	TypeField string `json:"type,omitempty"` // The type of mobile device.
	Deviceid string `json:"deviceId,omitempty"` // The serial number for a Google Sync mobile device. For Android and iOS devices, this is a software generated unique identifier.
	Imei string `json:"imei,omitempty"` // The device's IMEI number.
	Kernelversion string `json:"kernelVersion,omitempty"` // The device's kernel version.
	Basebandversion string `json:"basebandVersion,omitempty"` // The device's baseband version.
	Networkoperator string `json:"networkOperator,omitempty"` // Mobile Device mobile or network operator (if available) (Read-only)
	Os string `json:"os,omitempty"` // The mobile device's operating system, for example IOS 4.3 or Android 2.3.5. This property can be [updated](/admin-sdk/directory/v1/reference/mobiledevices/update.html). For more information, see the [Developer's Guide](/admin-sdk/directory/v1/guides/manage-mobile-devices#update_mobile_device).
	Developeroptionsstatus bool `json:"developerOptionsStatus,omitempty"` // Developer options enabled or disabled on device (Read-only)
	Bootloaderversion string `json:"bootloaderVersion,omitempty"` // Mobile Device Bootloader version (Read-only)
	Kind string `json:"kind,omitempty"` // The type of the API resource. For Mobiledevices resources, the value is `admin#directory#mobiledevice`.
	Brand string `json:"brand,omitempty"` // Mobile Device Brand (Read-only)
	Resourceid string `json:"resourceId,omitempty"` // The unique ID the API service uses to identify the mobile device.
	Firstsync string `json:"firstSync,omitempty"` // Date and time the device was first synchronized with the policy settings in the G Suite administrator control panel (Read-only)
	Status string `json:"status,omitempty"` // The device's status.
	Wifimacaddress string `json:"wifiMacAddress,omitempty"` // The device's MAC address on Wi-Fi networks.
	Securitypatchlevel string `json:"securityPatchLevel,omitempty"` // Mobile Device Security patch level (Read-only)
	Supportsworkprofile bool `json:"supportsWorkProfile,omitempty"` // Work profile supported on device (Read-only)
	Devicecompromisedstatus string `json:"deviceCompromisedStatus,omitempty"` // The compromised device status.
	Devicepasswordstatus string `json:"devicePasswordStatus,omitempty"` // DevicePasswordStatus (Read-only)
}

// VerificationCode represents the VerificationCode schema from the OpenAPI specification
type VerificationCode struct {
	Userid string `json:"userId,omitempty"` // The obfuscated unique ID of the user.
	Verificationcode string `json:"verificationCode,omitempty"` // A current verification code for the user. Invalidated or used verification codes are not returned as part of the result.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // The type of the resource. This is always `admin#directory#verificationCode`.
}

// MobileDevices represents the MobileDevices schema from the OpenAPI specification
type MobileDevices struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token used to access next page of this result.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // Kind of resource this is.
	Mobiledevices []MobileDevice `json:"mobiledevices,omitempty"` // A list of Mobile Device objects.
}

// Tokens represents the Tokens schema from the OpenAPI specification
type Tokens struct {
	Kind string `json:"kind,omitempty"` // The type of the API resource. This is always `admin#directory#tokenList`.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Items []Token `json:"items,omitempty"` // A list of Token resources.
}

// Domains represents the Domains schema from the OpenAPI specification
type Domains struct {
	Domainname string `json:"domainName,omitempty"` // The domain name of the customer.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Isprimary bool `json:"isPrimary,omitempty"` // Indicates if the domain is a primary domain (Read-only).
	Kind string `json:"kind,omitempty"` // Kind of resource this is.
	Verified bool `json:"verified,omitempty"` // Indicates the verification state of a domain. (Read-only).
	Creationtime string `json:"creationTime,omitempty"` // Creation time of the domain. Expressed in [Unix time](https://en.wikipedia.org/wiki/Epoch_time) format. (Read-only).
	Domainaliases []DomainAlias `json:"domainAliases,omitempty"` // A list of domain alias objects. (Read-only)
}

// UserSshPublicKey represents the UserSshPublicKey schema from the OpenAPI specification
type UserSshPublicKey struct {
	Expirationtimeusec string `json:"expirationTimeUsec,omitempty"` // An expiration time in microseconds since epoch.
	Fingerprint string `json:"fingerprint,omitempty"` // A SHA-256 fingerprint of the SSH public key. (Read-only)
	Key string `json:"key,omitempty"` // An SSH public key.
}

// ChromeOsDevice represents the ChromeOsDevice schema from the OpenAPI specification
type ChromeOsDevice struct {
	Platformversion string `json:"platformVersion,omitempty"` // The Chrome device's platform version.
	Manufacturedate string `json:"manufactureDate,omitempty"` // (Read-only) The date the device was manufactured in yyyy-mm-dd format.
	Devicelicensetype string `json:"deviceLicenseType,omitempty"` // Output only. Device license type.
	Recentusers []map[string]interface{} `json:"recentUsers,omitempty"` // A list of recent device users, in descending order, by last login time.
	Tpmversioninfo map[string]interface{} `json:"tpmVersionInfo,omitempty"` // Trusted Platform Module (TPM) (Read-only)
	Firstenrollmenttime string `json:"firstEnrollmentTime,omitempty"` // Date and time for the first time the device was enrolled.
	Firmwareversion string `json:"firmwareVersion,omitempty"` // The Chrome device's firmware version.
	Devicefiles []map[string]interface{} `json:"deviceFiles,omitempty"` // A list of device files to download (Read-only)
	Status string `json:"status,omitempty"` // The status of the device.
	Meid string `json:"meid,omitempty"` // The Mobile Equipment Identifier (MEID) or the International Mobile Equipment Identity (IMEI) for the 3G mobile card in a mobile device. A MEID/IMEI is typically used when adding a device to a wireless carrier's post-pay service plan. If the device does not have this information, this property is not included in the response. For more information on how to export a MEID/IMEI list, see the [Developer's Guide](/admin-sdk/directory/v1/guides/manage-chrome-devices.html#export_meid).
	Osversion string `json:"osVersion,omitempty"` // The Chrome device's operating system version.
	Supportenddate string `json:"supportEndDate,omitempty"` // Final date the device will be supported (Read-only)
	Model string `json:"model,omitempty"` // The device's model information. If the device does not have this information, this property is not included in the response.
	Activetimeranges []map[string]interface{} `json:"activeTimeRanges,omitempty"` // A list of active time ranges (Read-only).
	Cpuinfo []map[string]interface{} `json:"cpuInfo,omitempty"` // Information regarding CPU specs in the device.
	Dockmacaddress string `json:"dockMacAddress,omitempty"` // (Read-only) Built-in MAC address for the docking station that the device connected to. Factory sets Media access control address (MAC address) assigned for use by a dock. It is reserved specifically for MAC pass through device policy. The format is twelve (12) hexadecimal digits without any delimiter (uppercase letters). This is only relevant for some devices.
	Lastdeprovisiontimestamp string `json:"lastDeprovisionTimestamp,omitempty"` // (Read-only) Date and time for the last deprovision of the device.
	Serialnumber string `json:"serialNumber,omitempty"` // The Chrome device serial number entered when the device was enabled. This value is the same as the Admin console's *Serial Number* in the *Chrome OS Devices* tab.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Macaddress string `json:"macAddress,omitempty"` // The device's wireless MAC address. If the device does not have this information, it is not included in the response.
	Systemramtotal string `json:"systemRamTotal,omitempty"` // Total RAM on the device [in bytes] (Read-only)
	Annotateduser string `json:"annotatedUser,omitempty"` // The user of the device as noted by the administrator. Maximum length is 100 characters. Empty values are allowed.
	Autoupdateexpiration string `json:"autoUpdateExpiration,omitempty"` // (Read-only) The timestamp after which the device will stop receiving Chrome updates or support
	Orgunitpath string `json:"orgUnitPath,omitempty"` // The full parent path with the organizational unit's name associated with the device. Path names are case insensitive. If the parent organizational unit is the top-level organization, it is represented as a forward slash, `/`. This property can be [updated](/admin-sdk/directory/v1/guides/manage-chrome-devices#move_chrome_devices_to_ou) using the API. For more information about how to create an organizational structure for your device, see the [administration help center](https://support.google.com/a/answer/182433).
	Diskvolumereports []map[string]interface{} `json:"diskVolumeReports,omitempty"` // Reports of disk space and other info about mounted/connected volumes.
	Deviceid string `json:"deviceId,omitempty"` // The unique ID of the Chrome device.
	Ordernumber string `json:"orderNumber,omitempty"` // The device's order number. Only devices directly purchased from Google have an order number.
	Lastknownnetwork []map[string]interface{} `json:"lastKnownNetwork,omitempty"` // Contains last known network (Read-only)
	Notes string `json:"notes,omitempty"` // Notes about this device added by the administrator. This property can be [searched](https://support.google.com/chrome/a/answer/1698333) with the [list](/admin-sdk/directory/v1/reference/chromeosdevices/list) method's `query` parameter. Maximum length is 500 characters. Empty values are allowed.
	Lastenrollmenttime string `json:"lastEnrollmentTime,omitempty"` // Date and time the device was last enrolled (Read-only)
	Systemramfreereports []map[string]interface{} `json:"systemRamFreeReports,omitempty"` // Reports of amounts of available RAM memory (Read-only)
	Ethernetmacaddress string `json:"ethernetMacAddress,omitempty"` // The device's MAC address on the ethernet network interface.
	Deprovisionreason string `json:"deprovisionReason,omitempty"` // (Read-only) Deprovision reason.
	Kind string `json:"kind,omitempty"` // The type of resource. For the Chromeosdevices resource, the value is `admin#directory#chromeosdevice`.
	Cpustatusreports []map[string]interface{} `json:"cpuStatusReports,omitempty"` // Reports of CPU utilization and temperature (Read-only)
	Ethernetmacaddress0 string `json:"ethernetMacAddress0,omitempty"` // (Read-only) MAC address used by the Chromebook’s internal ethernet port, and for onboard network (ethernet) interface. The format is twelve (12) hexadecimal digits without any delimiter (uppercase letters). This is only relevant for some devices.
	Willautorenew bool `json:"willAutoRenew,omitempty"` // Determines if the device will auto renew its support after the support end date. This is a read-only property.
	Orgunitid string `json:"orgUnitId,omitempty"` // The unique ID of the organizational unit. orgUnitPath is the human readable version of orgUnitId. While orgUnitPath may change by renaming an organizational unit within the path, orgUnitId is unchangeable for one organizational unit. This property can be [updated](/admin-sdk/directory/v1/guides/manage-chrome-devices#move_chrome_devices_to_ou) using the API. For more information about how to create an organizational structure for your device, see the [administration help center](https://support.google.com/a/answer/182433).
	Bootmode string `json:"bootMode,omitempty"` // The boot mode for the device. The possible values are: * `Verified`: The device is running a valid version of the Chrome OS. * `Dev`: The devices's developer hardware switch is enabled. When booted, the device has a command line shell. For an example of a developer switch, see the [Chromebook developer information](https://www.chromium.org/chromium-os/developer-information-for-chrome-os-devices/samsung-series-5-chromebook#TOC-Developer-switch).
	Lastsync string `json:"lastSync,omitempty"` // Date and time the device was last synchronized with the policy settings in the G Suite administrator control panel (Read-only)
	Osupdatestatus OsUpdateStatus `json:"osUpdateStatus,omitempty"` // Contains information regarding the current OS update status.
	Screenshotfiles []map[string]interface{} `json:"screenshotFiles,omitempty"` // A list of screenshot files to download. Type is always "SCREENSHOT_FILE". (Read-only)
	Annotatedlocation string `json:"annotatedLocation,omitempty"` // The address or location of the device as noted by the administrator. Maximum length is `200` characters. Empty values are allowed.
	Annotatedassetid string `json:"annotatedAssetId,omitempty"` // The asset identifier as noted by an administrator or specified during enrollment.
}

// Feature represents the Feature schema from the OpenAPI specification
type Feature struct {
	Etags string `json:"etags,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // Kind of resource this is.
	Name string `json:"name,omitempty"` // The name of the feature.
}

// Privilege represents the Privilege schema from the OpenAPI specification
type Privilege struct {
	Kind string `json:"kind,omitempty"` // The type of the API resource. This is always `admin#directory#privilege`.
	Privilegename string `json:"privilegeName,omitempty"` // The name of the privilege.
	Serviceid string `json:"serviceId,omitempty"` // The obfuscated ID of the service this privilege is for. This value is returned with [`Privileges.list()`](/admin-sdk/directory/v1/reference/privileges/list).
	Servicename string `json:"serviceName,omitempty"` // The name of the service this privilege is for.
	Childprivileges []Privilege `json:"childPrivileges,omitempty"` // A list of child privileges. Privileges for a service form a tree. Each privilege can have a list of child privileges; this list is empty for a leaf privilege.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Isouscopable bool `json:"isOuScopable,omitempty"` // If the privilege can be restricted to an organization unit.
}

// UserIm represents the UserIm schema from the OpenAPI specification
type UserIm struct {
	TypeField string `json:"type,omitempty"` // Each entry can have a type which indicates standard types of that entry. For example instant messengers could be of home work etc. In addition to the standard type an entry can have a custom type and can take any value. Such types should have the CUSTOM value as type and also have a customType value.
	Customprotocol string `json:"customProtocol,omitempty"` // Custom protocol.
	Customtype string `json:"customType,omitempty"` // Custom type.
	Im string `json:"im,omitempty"` // Instant messenger id.
	Primary bool `json:"primary,omitempty"` // If this is user's primary im. Only one entry could be marked as primary.
	Protocol string `json:"protocol,omitempty"` // Protocol used in the instant messenger. It should be one of the values from ImProtocolTypes map. Similar to type it can take a CUSTOM value and specify the custom name in customProtocol field.
}

// RoleAssignment represents the RoleAssignment schema from the OpenAPI specification
type RoleAssignment struct {
	Orgunitid string `json:"orgUnitId,omitempty"` // If the role is restricted to an organization unit, this contains the ID for the organization unit the exercise of this role is restricted to.
	Roleassignmentid string `json:"roleAssignmentId,omitempty"` // ID of this roleAssignment.
	Roleid string `json:"roleId,omitempty"` // The ID of the role that is assigned.
	Scopetype string `json:"scopeType,omitempty"` // The scope in which this role is assigned.
	Assignedto string `json:"assignedTo,omitempty"` // The unique ID of the entity this role is assigned to—either the `user_id` of a user, the `group_id` of a group, or the `uniqueId` of a service account as defined in [Identity and Access Management (IAM)](https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts).
	Assigneetype string `json:"assigneeType,omitempty"` // Output only. The type of the assignee (`USER` or `GROUP`).
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // The type of the API resource. This is always `admin#directory#roleAssignment`.
}

// UserPhone represents the UserPhone schema from the OpenAPI specification
type UserPhone struct {
	Value string `json:"value,omitempty"` // Phone number.
	Customtype string `json:"customType,omitempty"` // Custom Type.
	Primary bool `json:"primary,omitempty"` // If this is user's primary phone or not.
	TypeField string `json:"type,omitempty"` // Each entry can have a type which indicates standard types of that entry. For example phone could be of home_fax work mobile etc. In addition to the standard type an entry can have a custom type and can give it any name. Such types should have the CUSTOM value as type and also have a customType value.
}

// UserCustomProperties represents the UserCustomProperties schema from the OpenAPI specification
type UserCustomProperties struct {
}

// CalendarResources represents the CalendarResources schema from the OpenAPI specification
type CalendarResources struct {
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Items []CalendarResource `json:"items,omitempty"` // The CalendarResources in this page of results.
	Kind string `json:"kind,omitempty"` // Identifies this as a collection of CalendarResources. This is always `admin#directory#resources#calendars#calendarResourcesList`.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The continuation token, used to page through large result sets. Provide this value in a subsequent request to return the next page of results.
}

// Group represents the Group schema from the OpenAPI specification
type Group struct {
	Name string `json:"name,omitempty"` // The group's display name.
	Id string `json:"id,omitempty"` // Read-only. The unique ID of a group. A group `id` can be used as a group request URI's `groupKey`.
	Admincreated bool `json:"adminCreated,omitempty"` // Read-only. Value is `true` if this group was created by an administrator rather than a user.
	Description string `json:"description,omitempty"` // An extended description to help users determine the purpose of a group. For example, you can include information about who should join the group, the types of messages to send to the group, links to FAQs about the group, or related groups. Maximum length is `4,096` characters.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // The type of the API resource. For Groups resources, the value is `admin#directory#group`.
	Directmemberscount string `json:"directMembersCount,omitempty"` // The number of users that are direct members of the group. If a group is a member (child) of this group (the parent), members of the child group are not counted in the `directMembersCount` property of the parent group.
	Noneditablealiases []string `json:"nonEditableAliases,omitempty"` // Read-only. The list of the group's non-editable alias email addresses that are outside of the account's primary domain or subdomains. These are functioning email addresses used by the group. This is a read-only property returned in the API's response for a group. If edited in a group's POST or PUT request, the edit is ignored.
	Aliases []string `json:"aliases,omitempty"` // Read-only. The list of a group's alias email addresses. To add, update, or remove a group's aliases, use the `groups.aliases` methods. If edited in a group's POST or PUT request, the edit is ignored.
	Email string `json:"email,omitempty"` // The group's email address. If your account has multiple domains, select the appropriate domain for the email address. The `email` must be unique. This property is required when creating a group. Group email addresses are subject to the same character usage rules as usernames, see the [help center](https://support.google.com/a/answer/9193374) for details.
}

// User represents the User schema from the OpenAPI specification
type User struct {
	Organizations interface{} `json:"organizations,omitempty"` // The list of organizations the user belongs to. The maximum allowed data size for this field is 10KB.
	Kind string `json:"kind,omitempty"` // Output only. The type of the API resource. For Users resources, the value is `admin#directory#user`.
	Websites interface{} `json:"websites,omitempty"` // The user's websites. The maximum allowed data size for this field is 2KB.
	Languages interface{} `json:"languages,omitempty"` // The user's languages. The maximum allowed data size for this field is 1KB.
	Phones interface{} `json:"phones,omitempty"` // The list of the user's phone numbers. The maximum allowed data size for this field is 1KB.
	Etag string `json:"etag,omitempty"` // Output only. ETag of the resource.
	Thumbnailphotoetag string `json:"thumbnailPhotoEtag,omitempty"` // Output only. ETag of the user's photo (Read-only)
	Emails interface{} `json:"emails,omitempty"` // The list of the user's email addresses. The maximum allowed data size for this field is 10KB. This excludes `publicKeyEncryptionCertificates`.
	Id string `json:"id,omitempty"` // The unique ID for the user. A user `id` can be used as a user request URI's `userKey`.
	Ipwhitelisted bool `json:"ipWhitelisted,omitempty"` // If `true`, the user's IP address is subject to a deprecated IP address [`allowlist`](https://support.google.com/a/answer/60752) configuration.
	Includeinglobaladdresslist bool `json:"includeInGlobalAddressList,omitempty"` // Indicates if the user's profile is visible in the Google Workspace global address list when the contact sharing feature is enabled for the domain. For more information about excluding user profiles, see the [administration help center](https://support.google.com/a/answer/1285988).
	Isadmin bool `json:"isAdmin,omitempty"` // Output only. Indicates a user with super admininistrator privileges. The `isAdmin` property can only be edited in the [Make a user an administrator](/admin-sdk/directory/v1/guides/manage-users.html#make_admin) operation ( [makeAdmin](/admin-sdk/directory/v1/reference/users/makeAdmin.html) method). If edited in the user [insert](/admin-sdk/directory/v1/reference/users/insert.html) or [update](/admin-sdk/directory/v1/reference/users/update.html) methods, the edit is ignored by the API service.
	Noneditablealiases []string `json:"nonEditableAliases,omitempty"` // Output only. The list of the user's non-editable alias email addresses. These are typically outside the account's primary domain or sub-domain.
	Isenrolledin2sv bool `json:"isEnrolledIn2Sv,omitempty"` // Output only. Is enrolled in 2-step verification (Read-only)
	Archived bool `json:"archived,omitempty"` // Indicates if user is archived.
	Notes interface{} `json:"notes,omitempty"` // Notes for the user.
	Sshpublickeys interface{} `json:"sshPublicKeys,omitempty"` // A list of SSH public keys.
	Addresses interface{} `json:"addresses,omitempty"` // The list of the user's addresses. The maximum allowed data size for this field is 10KB.
	Recoveryemail string `json:"recoveryEmail,omitempty"` // Recovery email of the user.
	Customerid string `json:"customerId,omitempty"` // Output only. The customer ID to [retrieve all account users](/admin-sdk/directory/v1/guides/manage-users.html#get_all_users). You can use the alias `my_customer` to represent your account's `customerId`. As a reseller administrator, you can use the resold customer account's `customerId`. To get a `customerId`, use the account's primary domain in the `domain` parameter of a [users.list](/admin-sdk/directory/v1/reference/users/list) request.
	Suspensionreason string `json:"suspensionReason,omitempty"` // Output only. Has the reason a user account is suspended either by the administrator or by Google at the time of suspension. The property is returned only if the `suspended` property is `true`.
	Isenforcedin2sv bool `json:"isEnforcedIn2Sv,omitempty"` // Output only. Is 2-step verification enforced (Read-only)
	Suspended bool `json:"suspended,omitempty"` // Indicates if user is suspended.
	Isdelegatedadmin bool `json:"isDelegatedAdmin,omitempty"` // Output only. Indicates if the user is a delegated administrator. Delegated administrators are supported by the API but cannot create or undelete users, or make users administrators. These requests are ignored by the API service. Roles and privileges for administrators are assigned using the [Admin console](https://support.google.com/a/answer/33325).
	Changepasswordatnextlogin bool `json:"changePasswordAtNextLogin,omitempty"` // Indicates if the user is forced to change their password at next login. This setting doesn't apply when [the user signs in via a third-party identity provider](https://support.google.com/a/answer/60224).
	Password string `json:"password,omitempty"` // User's password
	Ims interface{} `json:"ims,omitempty"` // The list of the user's Instant Messenger (IM) accounts. A user account can have multiple ims properties. But, only one of these ims properties can be the primary IM contact. The maximum allowed data size for this field is 2KB.
	Thumbnailphotourl string `json:"thumbnailPhotoUrl,omitempty"` // Output only. The URL of the user's profile photo. The URL might be temporary or private.
	Creationtime string `json:"creationTime,omitempty"` // User's G Suite account creation time. (Read-only)
	Lastlogintime string `json:"lastLoginTime,omitempty"` // User's last login time. (Read-only)
	Locations interface{} `json:"locations,omitempty"` // The user's locations. The maximum allowed data size for this field is 10KB.
	Customschemas map[string]interface{} `json:"customSchemas,omitempty"` // Custom fields of the user. The key is a `schema_name` and its values are `'field_name': 'field_value'`.
	Orgunitpath string `json:"orgUnitPath,omitempty"` // The full path of the parent organization associated with the user. If the parent organization is the top-level, it is represented as a forward slash (`/`).
	Keywords interface{} `json:"keywords,omitempty"` // The list of the user's keywords. The maximum allowed data size for this field is 1KB.
	Recoveryphone string `json:"recoveryPhone,omitempty"` // Recovery phone of the user. The phone number must be in the E.164 format, starting with the plus sign (+). Example: *+16506661212*.
	Hashfunction string `json:"hashFunction,omitempty"` // Stores the hash format of the `password` property. The following `hashFunction` values are allowed: * `MD5` - Accepts simple hex-encoded values. * `SHA-1` - Accepts simple hex-encoded values. * `crypt` - Compliant with the [C crypt library](https://en.wikipedia.org/wiki/Crypt_%28C%29). Supports the DES, MD5 (hash prefix `$1$`), SHA-256 (hash prefix `$5$`), and SHA-512 (hash prefix `$6$`) hash algorithms. If rounds are specified as part of the prefix, they must be 10,000 or fewer.
	Deletiontime string `json:"deletionTime,omitempty"`
	Name UserName `json:"name,omitempty"`
	Externalids interface{} `json:"externalIds,omitempty"` // The list of external IDs for the user, such as an employee or network ID. The maximum allowed data size for this field is 2KB.
	Agreedtoterms bool `json:"agreedToTerms,omitempty"` // Output only. This property is `true` if the user has completed an initial login and accepted the Terms of Service agreement.
	Aliases []string `json:"aliases,omitempty"` // Output only. The list of the user's alias email addresses.
	Relations interface{} `json:"relations,omitempty"` // The list of the user's relationships to other users. The maximum allowed data size for this field is 2KB.
	Primaryemail string `json:"primaryEmail,omitempty"` // The user's primary email address. This property is required in a request to create a user account. The `primaryEmail` must be unique and cannot be an alias of another user.
	Gender interface{} `json:"gender,omitempty"` // The user's gender. The maximum allowed data size for this field is 1KB.
	Posixaccounts interface{} `json:"posixAccounts,omitempty"` // The list of [POSIX](https://www.opengroup.org/austin/papers/posix_faq.html) account information for the user.
	Ismailboxsetup bool `json:"isMailboxSetup,omitempty"` // Output only. Indicates if the user's Google mailbox is created. This property is only applicable if the user has been assigned a Gmail license.
}

// DirectoryChromeosdevicesIssueCommandRequest represents the DirectoryChromeosdevicesIssueCommandRequest schema from the OpenAPI specification
type DirectoryChromeosdevicesIssueCommandRequest struct {
	Commandtype string `json:"commandType,omitempty"` // The type of command.
	Payload string `json:"payload,omitempty"` // The payload for the command, provide it only if command supports it. The following commands support adding payload: * `SET_VOLUME`: Payload is a stringified JSON object in the form: { "volume": 50 }. The volume has to be an integer in the range [0,100]. * `DEVICE_START_CRD_SESSION`: Payload is optionally a stringified JSON object in the form: { "ackedUserPresence": true }. `ackedUserPresence` is a boolean. By default, `ackedUserPresence` is set to `false`. To start a Chrome Remote Desktop session for an active device, set `ackedUserPresence` to `true`. * `REBOOT`: Payload is a stringified JSON object in the form: { "user_session_delay_seconds": 300 }. The delay has to be in the range [0, 300]. * `FETCH_SUPPORT_PACKET`: Payload is optionally a stringified JSON object in the form: {"supportPacketDetails":{ "issueCaseId": optional_support_case_id_string, "issueDescription": optional_issue_description_string, "requestedDataCollectors": []}} The list of available `data_collector_enums` are as following: Chrome System Information (1), Crash IDs (2), Memory Details (3), UI Hierarchy (4), Additional ChromeOS Platform Logs (5), Device Event (6), Intel WiFi NICs Debug Dump (7), Touch Events (8), Lacros (9), Lacros System Information (10), ChromeOS Flex Logs (11), DBus Details (12), ChromeOS Network Routes (13), ChromeOS Shill (Connection Manager) Logs (14), Policies (15), ChromeOS System State and Logs (16), ChromeOS System Logs (17), ChromeOS Chrome User Logs (18), ChromeOS Bluetooth (19), ChromeOS Connected Input Devices (20), ChromeOS Traffic Counters (21), ChromeOS Virtual Keyboard (22), ChromeOS Network Health (23). See more details in [help article](https://support.google.com/chrome/a?p=remote-log).
}

// ListPrintersResponse represents the ListPrintersResponse schema from the OpenAPI specification
type ListPrintersResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
	Printers []Printer `json:"printers,omitempty"` // List of printers. If `org_unit_id` was given in the request, then only printers visible for this OU will be returned. If `org_unit_id` was not given in the request, then all printers will be returned.
}

// CalendarResource represents the CalendarResource schema from the OpenAPI specification
type CalendarResource struct {
	Uservisibledescription string `json:"userVisibleDescription,omitempty"` // Description of the resource, visible to users and admins.
	Floorname string `json:"floorName,omitempty"` // Name of the floor a resource is located on.
	Capacity int `json:"capacity,omitempty"` // Capacity of a resource, number of seats in a room.
	Kind string `json:"kind,omitempty"` // The type of the resource. For calendar resources, the value is `admin#directory#resources#calendars#CalendarResource`.
	Resourcedescription string `json:"resourceDescription,omitempty"` // Description of the resource, visible only to admins.
	Featureinstances interface{} `json:"featureInstances,omitempty"` // Instances of features for the calendar resource.
	Resourceemail string `json:"resourceEmail,omitempty"` // The read-only email for the calendar resource. Generated as part of creating a new calendar resource.
	Buildingid string `json:"buildingId,omitempty"` // Unique ID for the building a resource is located in.
	Floorsection string `json:"floorSection,omitempty"` // Name of the section within a floor a resource is located in.
	Resourceid string `json:"resourceId,omitempty"` // The unique ID for the calendar resource.
	Etags string `json:"etags,omitempty"` // ETag of the resource.
	Resourcename string `json:"resourceName,omitempty"` // The name of the calendar resource. For example, "Training Room 1A".
	Resourcetype string `json:"resourceType,omitempty"` // The type of the calendar resource, intended for non-room resources.
	Resourcecategory string `json:"resourceCategory,omitempty"` // The category of the calendar resource. Either CONFERENCE_ROOM or OTHER. Legacy data is set to CATEGORY_UNKNOWN.
	Generatedresourcename string `json:"generatedResourceName,omitempty"` // The read-only auto-generated name of the calendar resource which includes metadata about the resource such as building name, floor, capacity, etc. For example, "NYC-2-Training Room 1A (16)".
}

// Members represents the Members schema from the OpenAPI specification
type Members struct {
	Members []Member `json:"members,omitempty"` // A list of member objects.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token used to access next page of this result.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // Kind of resource this is.
}

// Building represents the Building schema from the OpenAPI specification
type Building struct {
	Buildingname string `json:"buildingName,omitempty"` // The building name as seen by users in Calendar. Must be unique for the customer. For example, "NYC-CHEL". The maximum length is 100 characters.
	Coordinates BuildingCoordinates `json:"coordinates,omitempty"` // Public API: Resources.buildings
	Description string `json:"description,omitempty"` // A brief description of the building. For example, "Chelsea Market".
	Etags string `json:"etags,omitempty"` // ETag of the resource.
	Floornames []string `json:"floorNames,omitempty"` // The display names for all floors in this building. The floors are expected to be sorted in ascending order, from lowest floor to highest floor. For example, ["B2", "B1", "L", "1", "2", "2M", "3", "PH"] Must contain at least one entry.
	Kind string `json:"kind,omitempty"` // Kind of resource this is.
	Address BuildingAddress `json:"address,omitempty"` // Public API: Resources.buildings
	Buildingid string `json:"buildingId,omitempty"` // Unique identifier for the building. The maximum length is 100 characters.
}

// BatchDeletePrintServersResponse represents the BatchDeletePrintServersResponse schema from the OpenAPI specification
type BatchDeletePrintServersResponse struct {
	Failedprintservers []PrintServerFailureInfo `json:"failedPrintServers,omitempty"` // A list of update failures.
	Printserverids []string `json:"printServerIds,omitempty"` // A list of print server IDs that were successfully deleted.
}

// Customer represents the Customer schema from the OpenAPI specification
type Customer struct {
	Customercreationtime string `json:"customerCreationTime,omitempty"` // The customer's creation time (Readonly)
	Customerdomain string `json:"customerDomain,omitempty"` // The customer's primary domain name string. Do not include the `www` prefix when creating a new customer.
	Phonenumber string `json:"phoneNumber,omitempty"` // The customer's contact phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
	Alternateemail string `json:"alternateEmail,omitempty"` // The customer's secondary contact email address. This email address cannot be on the same domain as the `customerDomain`
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // Identifies the resource as a customer. Value: `admin#directory#customer`
	Id string `json:"id,omitempty"` // The unique ID for the customer's Google Workspace account. (Readonly)
	Language string `json:"language,omitempty"` // The customer's ISO 639-2 language code. See the [Language Codes](/admin-sdk/directory/v1/languages) page for the list of supported codes. Valid language codes outside the supported set will be accepted by the API but may lead to unexpected behavior. The default value is `en`.
	Postaladdress CustomerPostalAddress `json:"postalAddress,omitempty"`
}

// CreatePrintServerRequest represents the CreatePrintServerRequest schema from the OpenAPI specification
type CreatePrintServerRequest struct {
	Parent string `json:"parent,omitempty"` // Required. The [unique ID](https://developers.google.com/admin-sdk/directory/reference/rest/v1/customers) of the customer's Google Workspace account. Format: `customers/{id}`
	Printserver PrintServer `json:"printServer,omitempty"` // Configuration for a print server.
}

// Printer represents the Printer schema from the OpenAPI specification
type Printer struct {
	Auxiliarymessages []AuxiliaryMessage `json:"auxiliaryMessages,omitempty"` // Output only. Auxiliary messages about issues with the printer configuration if any.
	Displayname string `json:"displayName,omitempty"` // Editable. Name of printer.
	Name string `json:"name,omitempty"` // The resource name of the Printer object, in the format customers/{customer-id}/printers/{printer-id} (During printer creation leave empty)
	Uri string `json:"uri,omitempty"` // Editable. Printer URI.
	Createtime string `json:"createTime,omitempty"` // Output only. Time when printer was created.
	Description string `json:"description,omitempty"` // Editable. Description of printer.
	Id string `json:"id,omitempty"` // Id of the printer. (During printer creation leave empty)
	Makeandmodel string `json:"makeAndModel,omitempty"` // Editable. Make and model of printer. e.g. Lexmark MS610de Value must be in format as seen in ListPrinterModels response.
	Usedriverlessconfig bool `json:"useDriverlessConfig,omitempty"` // Editable. flag to use driverless configuration or not. If it's set to be true, make_and_model can be ignored
	Orgunitid string `json:"orgUnitId,omitempty"` // Organization Unit that owns this printer (Only can be set during Printer creation)
}

// UserPhoto represents the UserPhoto schema from the OpenAPI specification
type UserPhoto struct {
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Height int `json:"height,omitempty"` // Height of the photo in pixels.
	Id string `json:"id,omitempty"` // The ID the API uses to uniquely identify the user.
	Kind string `json:"kind,omitempty"` // The type of the API resource. For Photo resources, this is `admin#directory#user#photo`.
	Mimetype string `json:"mimeType,omitempty"` // The MIME type of the photo. Allowed values are `JPEG`, `PNG`, `GIF`, `BMP`, `TIFF`, and web-safe base64 encoding.
	Photodata string `json:"photoData,omitempty"` // The user photo's upload data in [web-safe Base64](https://en.wikipedia.org/wiki/Base64#URL_applications) format in bytes. This means: * The slash (/) character is replaced with the underscore (_) character. * The plus sign (+) character is replaced with the hyphen (-) character. * The equals sign (=) character is replaced with the asterisk (*). * For padding, the period (.) character is used instead of the RFC-4648 baseURL definition which uses the equals sign (=) for padding. This is done to simplify URL-parsing. * Whatever the size of the photo being uploaded, the API downsizes it to 96x96 pixels.
	Primaryemail string `json:"primaryEmail,omitempty"` // The user's primary email address.
	Width int `json:"width,omitempty"` // Width of the photo in pixels.
}

// Features represents the Features schema from the OpenAPI specification
type Features struct {
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Features []Feature `json:"features,omitempty"` // The Features in this page of results.
	Kind string `json:"kind,omitempty"` // Kind of resource this is.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The continuation token, used to page through large result sets. Provide this value in a subsequent request to return the next page of results.
}

// Schema represents the Schema schema from the OpenAPI specification
type Schema struct {
	Fields []SchemaFieldSpec `json:"fields,omitempty"` // A list of fields in the schema.
	Kind string `json:"kind,omitempty"` // Kind of resource this is.
	Schemaid string `json:"schemaId,omitempty"` // The unique identifier of the schema (Read-only)
	Schemaname string `json:"schemaName,omitempty"` // The schema's name. Each `schema_name` must be unique within a customer. Reusing a name results in a `409: Entity already exists` error.
	Displayname string `json:"displayName,omitempty"` // Display name for the schema.
	Etag string `json:"etag,omitempty"` // The ETag of the resource.
}

// Privileges represents the Privileges schema from the OpenAPI specification
type Privileges struct {
	Kind string `json:"kind,omitempty"` // The type of the API resource. This is always `admin#directory#privileges`.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Items []Privilege `json:"items,omitempty"` // A list of Privilege resources.
}

// BuildingAddress represents the BuildingAddress schema from the OpenAPI specification
type BuildingAddress struct {
	Postalcode string `json:"postalCode,omitempty"` // Optional. Postal code of the address.
	Regioncode string `json:"regionCode,omitempty"` // Required. CLDR region code of the country/region of the address.
	Sublocality string `json:"sublocality,omitempty"` // Optional. Sublocality of the address.
	Addresslines []string `json:"addressLines,omitempty"` // Unstructured address lines describing the lower levels of an address.
	Administrativearea string `json:"administrativeArea,omitempty"` // Optional. Highest administrative subdivision which is used for postal addresses of a country or region.
	Languagecode string `json:"languageCode,omitempty"` // Optional. BCP-47 language code of the contents of this address (if known).
	Locality string `json:"locality,omitempty"` // Optional. Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In regions of the world where localities are not well defined or do not fit into this structure well, leave locality empty and use addressLines.
}

// BatchChangeChromeOsDeviceStatusResponse represents the BatchChangeChromeOsDeviceStatusResponse schema from the OpenAPI specification
type BatchChangeChromeOsDeviceStatusResponse struct {
	Changechromeosdevicestatusresults []ChangeChromeOsDeviceStatusResult `json:"changeChromeOsDeviceStatusResults,omitempty"` // The results for each of the ChromeOS devices provided in the request.
}

// UserAbout represents the UserAbout schema from the OpenAPI specification
type UserAbout struct {
	Contenttype string `json:"contentType,omitempty"` // About entry can have a type which indicates the content type. It can either be plain or html. By default, notes contents are assumed to contain plain text.
	Value string `json:"value,omitempty"` // Actual value of notes.
}

// Member represents the Member schema from the OpenAPI specification
type Member struct {
	Id string `json:"id,omitempty"` // The unique ID of the group member. A member `id` can be used as a member request URI's `memberKey`.
	Kind string `json:"kind,omitempty"` // The type of the API resource. For Members resources, the value is `admin#directory#member`.
	Role string `json:"role,omitempty"` // The member's role in a group. The API returns an error for cycles in group memberships. For example, if `group1` is a member of `group2`, `group2` cannot be a member of `group1`. For more information about a member's role, see the [administration help center](https://support.google.com/a/answer/167094).
	Status string `json:"status,omitempty"` // Status of member (Immutable)
	TypeField string `json:"type,omitempty"` // The type of group member.
	Delivery_settings string `json:"delivery_settings,omitempty"` // Defines mail delivery preferences of member. This field is only supported by `insert`, `update`, and `get` methods.
	Email string `json:"email,omitempty"` // The member's email address. A member can be a user or another group. This property is required when adding a member to a group. The `email` must be unique and cannot be an alias of another group. If the email address is changed, the API automatically reflects the email address changes.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
}

// PrintServer represents the PrintServer schema from the OpenAPI specification
type PrintServer struct {
	Orgunitid string `json:"orgUnitId,omitempty"` // ID of the organization unit (OU) that owns this print server. This value can only be set when the print server is initially created. If it's not populated, the print server is placed under the root OU. The `org_unit_id` can be retrieved using the [Directory API](/admin-sdk/directory/reference/rest/v1/orgunits).
	Uri string `json:"uri,omitempty"` // Editable. Print server URI.
	Createtime string `json:"createTime,omitempty"` // Output only. Time when the print server was created.
	Description string `json:"description,omitempty"` // Editable. Description of the print server (as shown in the Admin console).
	Displayname string `json:"displayName,omitempty"` // Editable. Display name of the print server (as shown in the Admin console).
	Id string `json:"id,omitempty"` // Immutable. ID of the print server. Leave empty when creating.
	Name string `json:"name,omitempty"` // Immutable. Resource name of the print server. Leave empty when creating. Format: `customers/{customer.id}/printServers/{print_server.id}`
}

// UserMakeAdmin represents the UserMakeAdmin schema from the OpenAPI specification
type UserMakeAdmin struct {
	Status bool `json:"status,omitempty"` // Indicates the administrator status of the user.
}

// ChangeChromeOsDeviceStatusSucceeded represents the ChangeChromeOsDeviceStatusSucceeded schema from the OpenAPI specification
type ChangeChromeOsDeviceStatusSucceeded struct {
}

// UserName represents the UserName schema from the OpenAPI specification
type UserName struct {
	Displayname string `json:"displayName,omitempty"` // The user's display name. Limit: 256 characters.
	Familyname string `json:"familyName,omitempty"` // The user's last name. Required when creating a user account.
	Fullname string `json:"fullName,omitempty"` // The user's full name formed by concatenating the first and last name values.
	Givenname string `json:"givenName,omitempty"` // The user's first name. Required when creating a user account.
}

// BatchCreatePrintersResponse represents the BatchCreatePrintersResponse schema from the OpenAPI specification
type BatchCreatePrintersResponse struct {
	Failures []FailureInfo `json:"failures,omitempty"` // A list of create failures. Printer IDs are not populated, as printer were not created.
	Printers []Printer `json:"printers,omitempty"` // A list of successfully created printers with their IDs populated.
}

// VerificationCodes represents the VerificationCodes schema from the OpenAPI specification
type VerificationCodes struct {
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Items []VerificationCode `json:"items,omitempty"` // A list of verification code resources.
	Kind string `json:"kind,omitempty"` // The type of the resource. This is always `admin#directory#verificationCodesList`.
}

// Aliases represents the Aliases schema from the OpenAPI specification
type Aliases struct {
	Aliases []interface{} `json:"aliases,omitempty"`
	Etag string `json:"etag,omitempty"`
	Kind string `json:"kind,omitempty"`
}

// Asps represents the Asps schema from the OpenAPI specification
type Asps struct {
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Items []Asp `json:"items,omitempty"` // A list of ASP resources.
	Kind string `json:"kind,omitempty"` // The type of the API resource. This is always `admin#directory#aspList`.
}

// DomainAlias represents the DomainAlias schema from the OpenAPI specification
type DomainAlias struct {
	Parentdomainname string `json:"parentDomainName,omitempty"` // The parent domain name that the domain alias is associated with. This can either be a primary or secondary domain name within a customer.
	Verified bool `json:"verified,omitempty"` // Indicates the verification state of a domain alias. (Read-only)
	Creationtime string `json:"creationTime,omitempty"` // The creation time of the domain alias. (Read-only).
	Domainaliasname string `json:"domainAliasName,omitempty"` // The domain alias name.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // Kind of resource this is.
}

// DomainAliases represents the DomainAliases schema from the OpenAPI specification
type DomainAliases struct {
	Domainaliases []DomainAlias `json:"domainAliases,omitempty"` // A list of domain alias objects.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // Kind of resource this is.
}

// BatchCreatePrintServersRequest represents the BatchCreatePrintServersRequest schema from the OpenAPI specification
type BatchCreatePrintServersRequest struct {
	Requests []CreatePrintServerRequest `json:"requests,omitempty"` // Required. A list of `PrintServer` resources to be created (max `50` per batch).
}

// UserGender represents the UserGender schema from the OpenAPI specification
type UserGender struct {
	Addressmeas string `json:"addressMeAs,omitempty"` // AddressMeAs. A human-readable string containing the proper way to refer to the profile owner by humans for example he/him/his or they/them/their.
	Customgender string `json:"customGender,omitempty"` // Custom gender.
	TypeField string `json:"type,omitempty"` // Gender.
}

// Buildings represents the Buildings schema from the OpenAPI specification
type Buildings struct {
	Kind string `json:"kind,omitempty"` // Kind of resource this is.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The continuation token, used to page through large result sets. Provide this value in a subsequent request to return the next page of results.
	Buildings []Building `json:"buildings,omitempty"` // The Buildings in this page of results.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
}

// CreatePrinterRequest represents the CreatePrinterRequest schema from the OpenAPI specification
type CreatePrinterRequest struct {
	Parent string `json:"parent,omitempty"` // Required. The name of the customer. Format: customers/{customer_id}
	Printer Printer `json:"printer,omitempty"` // Printer configuration.
}

// BatchDeletePrintersRequest represents the BatchDeletePrintersRequest schema from the OpenAPI specification
type BatchDeletePrintersRequest struct {
	Printerids []string `json:"printerIds,omitempty"` // A list of Printer.id that should be deleted. Max 100 at a time.
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Code int `json:"code,omitempty"` // The status code, which should be an enum value of google.rpc.Code.
	Details []map[string]interface{} `json:"details,omitempty"` // A list of messages that carry the error details. There is a common set of message types for APIs to use.
	Message string `json:"message,omitempty"` // A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.
}

// Token represents the Token schema from the OpenAPI specification
type Token struct {
	Scopes []string `json:"scopes,omitempty"` // A list of authorization scopes the application is granted.
	Userkey string `json:"userKey,omitempty"` // The unique ID of the user that issued the token.
	Anonymous bool `json:"anonymous,omitempty"` // Whether the application is registered with Google. The value is `true` if the application has an anonymous Client ID.
	Clientid string `json:"clientId,omitempty"` // The Client ID of the application the token is issued to.
	Displaytext string `json:"displayText,omitempty"` // The displayable name of the application the token is issued to.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // The type of the API resource. This is always `admin#directory#token`.
	Nativeapp bool `json:"nativeApp,omitempty"` // Whether the token is issued to an installed application. The value is `true` if the application is installed to a desktop or mobile device.
}

// BatchDeletePrintersResponse represents the BatchDeletePrintersResponse schema from the OpenAPI specification
type BatchDeletePrintersResponse struct {
	Failedprinters []FailureInfo `json:"failedPrinters,omitempty"` // A list of update failures.
	Printerids []string `json:"printerIds,omitempty"` // A list of Printer.id that were successfully deleted.
}

// DirectoryChromeosdevicesCommand represents the DirectoryChromeosdevicesCommand schema from the OpenAPI specification
type DirectoryChromeosdevicesCommand struct {
	Commandid string `json:"commandId,omitempty"` // Unique ID of a device command.
	Commandresult DirectoryChromeosdevicesCommandResult `json:"commandResult,omitempty"` // The result of executing a command.
	Issuetime string `json:"issueTime,omitempty"` // The timestamp when the command was issued by the admin.
	Payload string `json:"payload,omitempty"` // The payload that the command specified, if any.
	State string `json:"state,omitempty"` // Indicates the command state.
	TypeField string `json:"type,omitempty"` // The type of the command.
	Commandexpiretime string `json:"commandExpireTime,omitempty"` // The time at which the command will expire. If the device doesn't execute the command within this time the command will become expired.
}

// BuildingCoordinates represents the BuildingCoordinates schema from the OpenAPI specification
type BuildingCoordinates struct {
	Latitude float64 `json:"latitude,omitempty"` // Latitude in decimal degrees.
	Longitude float64 `json:"longitude,omitempty"` // Longitude in decimal degrees.
}

// ChromeOsDeviceAction represents the ChromeOsDeviceAction schema from the OpenAPI specification
type ChromeOsDeviceAction struct {
	Action string `json:"action,omitempty"` // Action to be taken on the Chrome OS device.
	Deprovisionreason string `json:"deprovisionReason,omitempty"` // Only used when the action is `deprovision`. With the `deprovision` action, this field is required. *Note*: The deprovision reason is audited because it might have implications on licenses for perpetual subscription customers.
}

// DirectoryChromeosdevicesCommandResult represents the DirectoryChromeosdevicesCommandResult schema from the OpenAPI specification
type DirectoryChromeosdevicesCommandResult struct {
	Commandresultpayload string `json:"commandResultPayload,omitempty"` // The payload for the command result. The following commands respond with a payload: * `DEVICE_START_CRD_SESSION`: Payload is a stringified JSON object in the form: { "url": url }. The URL provides a link to the Chrome Remote Desktop session.
	Errormessage string `json:"errorMessage,omitempty"` // The error message with a short explanation as to why the command failed. Only present if the command failed.
	Executetime string `json:"executeTime,omitempty"` // The time at which the command was executed or failed to execute.
	Result string `json:"result,omitempty"` // The result of the command.
}

// AuxiliaryMessage represents the AuxiliaryMessage schema from the OpenAPI specification
type AuxiliaryMessage struct {
	Fieldmask string `json:"fieldMask,omitempty"` // Field that this message concerns.
	Severity string `json:"severity,omitempty"` // Message severity
	Auxiliarymessage string `json:"auxiliaryMessage,omitempty"` // Human readable message in English. Example: "Given printer is invalid or no longer supported."
}

// UserExternalId represents the UserExternalId schema from the OpenAPI specification
type UserExternalId struct {
	Customtype string `json:"customType,omitempty"` // Custom type.
	TypeField string `json:"type,omitempty"` // The type of the Id.
	Value string `json:"value,omitempty"` // The value of the id.
}

// UserKeyword represents the UserKeyword schema from the OpenAPI specification
type UserKeyword struct {
	TypeField string `json:"type,omitempty"` // Each entry can have a type which indicates standard type of that entry. For example keyword could be of type occupation or outlook. In addition to the standard type an entry can have a custom type and can give it any name. Such types should have the CUSTOM value as type and also have a customType value.
	Value string `json:"value,omitempty"` // Keyword.
	Customtype string `json:"customType,omitempty"` // Custom Type.
}

// Domains2 represents the Domains2 schema from the OpenAPI specification
type Domains2 struct {
	Domains []Domains `json:"domains,omitempty"` // A list of domain objects.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // Kind of resource this is.
}

// Users represents the Users schema from the OpenAPI specification
type Users struct {
	Users []User `json:"users,omitempty"` // A list of user objects.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // Kind of resource this is.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token used to access next page of this result.
	Trigger_event string `json:"trigger_event,omitempty"` // Event that triggered this response (only used in case of Push Response)
}

// RoleAssignments represents the RoleAssignments schema from the OpenAPI specification
type RoleAssignments struct {
	Kind string `json:"kind,omitempty"` // The type of the API resource. This is always `admin#directory#roleAssignments`.
	Nextpagetoken string `json:"nextPageToken,omitempty"`
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Items []RoleAssignment `json:"items,omitempty"` // A list of RoleAssignment resources.
}

// ListPrintServersResponse represents the ListPrintServersResponse schema from the OpenAPI specification
type ListPrintServersResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token that can be sent as `page_token` in a request to retrieve the next page. If this field is omitted, there are no subsequent pages.
	Printservers []PrintServer `json:"printServers,omitempty"` // List of print servers.
}

// Empty represents the Empty schema from the OpenAPI specification
type Empty struct {
}

// CustomerPostalAddress represents the CustomerPostalAddress schema from the OpenAPI specification
type CustomerPostalAddress struct {
	Addressline2 string `json:"addressLine2,omitempty"` // Address line 2 of the address.
	Addressline3 string `json:"addressLine3,omitempty"` // Address line 3 of the address.
	Contactname string `json:"contactName,omitempty"` // The customer contact's name.
	Countrycode string `json:"countryCode,omitempty"` // This is a required property. For `countryCode` information see the [ISO 3166 country code elements](https://www.iso.org/iso/country_codes.htm).
	Organizationname string `json:"organizationName,omitempty"` // The company or company division name.
	Locality string `json:"locality,omitempty"` // Name of the locality. An example of a locality value is the city of `San Francisco`.
	Postalcode string `json:"postalCode,omitempty"` // The postal code. A postalCode example is a postal zip code such as `10009`. This is in accordance with - http: //portablecontacts.net/draft-spec.html#address_element.
	Addressline1 string `json:"addressLine1,omitempty"` // A customer's physical address. The address can be composed of one to three lines.
	Region string `json:"region,omitempty"` // Name of the region. An example of a region value is `NY` for the state of New York.
}

// FeatureRename represents the FeatureRename schema from the OpenAPI specification
type FeatureRename struct {
	Newname string `json:"newName,omitempty"` // New name of the feature.
}

// PrinterModel represents the PrinterModel schema from the OpenAPI specification
type PrinterModel struct {
	Manufacturer string `json:"manufacturer,omitempty"` // Manufacturer. eq. "Brother"
	Displayname string `json:"displayName,omitempty"` // Display name. eq. "Brother MFC-8840D"
	Makeandmodel string `json:"makeAndModel,omitempty"` // Make and model as represented in "make_and_model" field in Printer object. eq. "brother mfc-8840d"
}

// UserLocation represents the UserLocation schema from the OpenAPI specification
type UserLocation struct {
	Area string `json:"area,omitempty"` // Textual location. This is most useful for display purposes to concisely describe the location. For example 'Mountain View, CA', 'Near Seattle', 'US-NYC-9TH 9A209A.''
	Buildingid string `json:"buildingId,omitempty"` // Building Identifier.
	Customtype string `json:"customType,omitempty"` // Custom Type.
	Deskcode string `json:"deskCode,omitempty"` // Most specific textual code of individual desk location.
	Floorname string `json:"floorName,omitempty"` // Floor name/number.
	Floorsection string `json:"floorSection,omitempty"` // Floor section. More specific location within the floor. For example if a floor is divided into sections 'A', 'B' and 'C' this field would identify one of those values.
	TypeField string `json:"type,omitempty"` // Each entry can have a type which indicates standard types of that entry. For example location could be of types default and desk. In addition to standard type an entry can have a custom type and can give it any name. Such types should have 'custom' as type and also have a customType value.
}

// OrgUnits represents the OrgUnits schema from the OpenAPI specification
type OrgUnits struct {
	Organizationunits []OrgUnit `json:"organizationUnits,omitempty"` // A list of organizational unit objects.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // The type of the API resource. For Org Unit resources, the type is `admin#directory#orgUnits`.
}

// Channel represents the Channel schema from the OpenAPI specification
type Channel struct {
	Resourceid string `json:"resourceId,omitempty"` // An opaque ID that identifies the resource being watched on this channel. Stable across different API versions.
	Token string `json:"token,omitempty"` // An arbitrary string delivered to the target address with each notification delivered over this channel. Optional.
	TypeField string `json:"type,omitempty"` // The type of delivery mechanism used for this channel.
	Kind string `json:"kind,omitempty"` // Identifies this as a notification channel used to watch for changes to a resource, which is `api#channel`.
	Address string `json:"address,omitempty"` // The address where notifications are delivered for this channel.
	Params map[string]interface{} `json:"params,omitempty"` // Additional parameters controlling delivery channel behavior. Optional. For example, `params.ttl` specifies the time-to-live in seconds for the notification channel, where the default is 2 hours and the maximum TTL is 2 days.
	Payload bool `json:"payload,omitempty"` // A Boolean value to indicate whether payload is wanted. Optional.
	Resourceuri string `json:"resourceUri,omitempty"` // A version-specific identifier for the watched resource.
	Expiration string `json:"expiration,omitempty"` // Date and time of notification channel expiration, expressed as a Unix timestamp, in milliseconds. Optional.
	Id string `json:"id,omitempty"` // A UUID or similar unique string that identifies this channel.
}

// ChromeOsMoveDevicesToOu represents the ChromeOsMoveDevicesToOu schema from the OpenAPI specification
type ChromeOsMoveDevicesToOu struct {
	Deviceids []string `json:"deviceIds,omitempty"` // Chrome OS devices to be moved to OU
}

// UserAlias represents the UserAlias schema from the OpenAPI specification
type UserAlias struct {
	Kind string `json:"kind,omitempty"` // The type of the API resource. For Alias resources, the value is `admin#directory#alias`.
	Primaryemail string `json:"primaryEmail,omitempty"` // The user's primary email address.
	Alias string `json:"alias,omitempty"` // The alias email address.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Id string `json:"id,omitempty"` // The unique ID for the user.
}

// UserOrganization represents the UserOrganization schema from the OpenAPI specification
type UserOrganization struct {
	Primary bool `json:"primary,omitempty"` // If it user's primary organization.
	Symbol string `json:"symbol,omitempty"` // Symbol of the organization.
	TypeField string `json:"type,omitempty"` // Each entry can have a type which indicates standard types of that entry. For example organization could be of school work etc. In addition to the standard type an entry can have a custom type and can give it any name. Such types should have the CUSTOM value as type and also have a CustomType value.
	Customtype string `json:"customType,omitempty"` // Custom type.
	Title string `json:"title,omitempty"` // Title (designation) of the user in the organization.
	Fulltimeequivalent int `json:"fullTimeEquivalent,omitempty"` // The full-time equivalent millipercent within the organization (100000 = 100%).
	Location string `json:"location,omitempty"` // Location of the organization. This need not be fully qualified address.
	Name string `json:"name,omitempty"` // Name of the organization
	Costcenter string `json:"costCenter,omitempty"` // The cost center of the users department.
	Department string `json:"department,omitempty"` // Department within the organization.
	Description string `json:"description,omitempty"` // Description of the organization.
	Domain string `json:"domain,omitempty"` // The domain to which the organization belongs to.
}

// Alias represents the Alias schema from the OpenAPI specification
type Alias struct {
	Alias string `json:"alias,omitempty"`
	Etag string `json:"etag,omitempty"`
	Id string `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
	Primaryemail string `json:"primaryEmail,omitempty"`
}

// SchemaFieldSpec represents the SchemaFieldSpec schema from the OpenAPI specification
type SchemaFieldSpec struct {
	Displayname string `json:"displayName,omitempty"` // Display Name of the field.
	Readaccesstype string `json:"readAccessType,omitempty"` // Specifies who can view values of this field. See [Retrieve users as a non-administrator](/admin-sdk/directory/v1/guides/manage-users#retrieve_users_non_admin) for more information. Note: It may take up to 24 hours for changes to this field to be reflected.
	Indexed bool `json:"indexed,omitempty"` // Boolean specifying whether the field is indexed or not. Default: `true`.
	Kind string `json:"kind,omitempty"` // The kind of resource this is. For schema fields this is always `admin#directory#schema#fieldspec`.
	Numericindexingspec map[string]interface{} `json:"numericIndexingSpec,omitempty"` // Indexing spec for a numeric field. By default, only exact match queries will be supported for numeric fields. Setting the `numericIndexingSpec` allows range queries to be supported.
	Etag string `json:"etag,omitempty"` // The ETag of the field.
	Multivalued bool `json:"multiValued,omitempty"` // A boolean specifying whether this is a multi-valued field or not. Default: `false`.
	Fieldid string `json:"fieldId,omitempty"` // The unique identifier of the field (Read-only)
	Fieldname string `json:"fieldName,omitempty"` // The name of the field.
	Fieldtype string `json:"fieldType,omitempty"` // The type of the field.
}

// Schemas represents the Schemas schema from the OpenAPI specification
type Schemas struct {
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // Kind of resource this is.
	Schemas []Schema `json:"schemas,omitempty"` // A list of UserSchema objects.
}

// MembersHasMember represents the MembersHasMember schema from the OpenAPI specification
type MembersHasMember struct {
	Ismember bool `json:"isMember,omitempty"` // Output only. Identifies whether the given user is a member of the group. Membership can be direct or nested.
}

// Roles represents the Roles schema from the OpenAPI specification
type Roles struct {
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Items []Role `json:"items,omitempty"` // A list of Role resources.
	Kind string `json:"kind,omitempty"` // The type of the API resource. This is always `admin#directory#roles`.
	Nextpagetoken string `json:"nextPageToken,omitempty"`
}

// OrgUnit represents the OrgUnit schema from the OpenAPI specification
type OrgUnit struct {
	Orgunitid string `json:"orgUnitId,omitempty"` // The unique ID of the organizational unit.
	Parentorgunitid string `json:"parentOrgUnitId,omitempty"` // The unique ID of the parent organizational unit. Required, unless `parentOrgUnitPath` is set.
	Description string `json:"description,omitempty"` // Description of the organizational unit.
	Blockinheritance bool `json:"blockInheritance,omitempty"` // Determines if a sub-organizational unit can inherit the settings of the parent organization. The default value is `false`, meaning a sub-organizational unit inherits the settings of the nearest parent organizational unit. This field is deprecated. Setting it to `true` is no longer supported and can have _unintended consequences_. For more information about inheritance and users in an organization structure, see the [administration help center](https://support.google.com/a/answer/4352075).
	Etag string `json:"etag,omitempty"` // ETag of the resource.
	Kind string `json:"kind,omitempty"` // The type of the API resource. For Orgunits resources, the value is `admin#directory#orgUnit`.
	Orgunitpath string `json:"orgUnitPath,omitempty"` // The full path to the organizational unit. The `orgUnitPath` is a derived property. When listed, it is derived from `parentOrgunitPath` and organizational unit's `name`. For example, for an organizational unit named 'apps' under parent organization '/engineering', the orgUnitPath is '/engineering/apps'. In order to edit an `orgUnitPath`, either update the name of the organization or the `parentOrgunitPath`. A user's organizational unit determines which Google Workspace services the user has access to. If the user is moved to a new organization, the user's access changes. For more information about organization structures, see the [administration help center](https://support.google.com/a/answer/4352075). For more information about moving a user to a different organization, see [Update a user](/admin-sdk/directory/v1/guides/manage-users.html#update_user).
	Name string `json:"name,omitempty"` // The organizational unit's path name. For example, an organizational unit's name within the /corp/support/sales_support parent path is sales_support. Required.
	Parentorgunitpath string `json:"parentOrgUnitPath,omitempty"` // The organizational unit's parent path. For example, /corp/sales is the parent path for /corp/sales/sales_support organizational unit. Required, unless `parentOrgUnitId` is set.
}

// UserUndelete represents the UserUndelete schema from the OpenAPI specification
type UserUndelete struct {
	Orgunitpath string `json:"orgUnitPath,omitempty"` // OrgUnit of User
}

// UserWebsite represents the UserWebsite schema from the OpenAPI specification
type UserWebsite struct {
	Value string `json:"value,omitempty"` // Website.
	Customtype string `json:"customType,omitempty"` // Custom Type.
	Primary bool `json:"primary,omitempty"` // If this is user's primary website or not.
	TypeField string `json:"type,omitempty"` // Each entry can have a type which indicates standard types of that entry. For example website could be of home work blog etc. In addition to the standard type an entry can have a custom type and can give it any name. Such types should have the CUSTOM value as type and also have a customType value.
}

// BatchCreatePrintServersResponse represents the BatchCreatePrintServersResponse schema from the OpenAPI specification
type BatchCreatePrintServersResponse struct {
	Failures []PrintServerFailureInfo `json:"failures,omitempty"` // A list of create failures. `PrintServer` IDs are not populated, as print servers were not created.
	Printservers []PrintServer `json:"printServers,omitempty"` // A list of successfully created print servers with their IDs populated.
}

// UserAddress represents the UserAddress schema from the OpenAPI specification
type UserAddress struct {
	Postalcode string `json:"postalCode,omitempty"` // Postal code.
	TypeField string `json:"type,omitempty"` // Each entry can have a type which indicates standard values of that entry. For example address could be of home work etc. In addition to the standard type an entry can have a custom type and can take any value. Such type should have the CUSTOM value as type and also have a customType value.
	Sourceisstructured bool `json:"sourceIsStructured,omitempty"` // User supplied address was structured. Structured addresses are NOT supported at this time. You might be able to write structured addresses but any values will eventually be clobbered.
	Locality string `json:"locality,omitempty"` // Locality.
	Primary bool `json:"primary,omitempty"` // If this is user's primary address. Only one entry could be marked as primary.
	Region string `json:"region,omitempty"` // Region.
	Countrycode string `json:"countryCode,omitempty"` // Country code.
	Customtype string `json:"customType,omitempty"` // Custom type.
	Extendedaddress string `json:"extendedAddress,omitempty"` // Extended Address.
	Pobox string `json:"poBox,omitempty"` // Other parts of address.
	Streetaddress string `json:"streetAddress,omitempty"` // Street.
	Country string `json:"country,omitempty"` // Country.
	Formatted string `json:"formatted,omitempty"` // Formatted address.
}

// BatchCreatePrintersRequest represents the BatchCreatePrintersRequest schema from the OpenAPI specification
type BatchCreatePrintersRequest struct {
	Requests []CreatePrinterRequest `json:"requests,omitempty"` // A list of Printers to be created. Max 50 at a time.
}

// Groups represents the Groups schema from the OpenAPI specification
type Groups struct {
	Groups []Group `json:"groups,omitempty"` // A list of group objects.
	Kind string `json:"kind,omitempty"` // Kind of resource this is.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token used to access next page of this result.
	Etag string `json:"etag,omitempty"` // ETag of the resource.
}

// DirectoryChromeosdevicesIssueCommandResponse represents the DirectoryChromeosdevicesIssueCommandResponse schema from the OpenAPI specification
type DirectoryChromeosdevicesIssueCommandResponse struct {
	Commandid string `json:"commandId,omitempty"` // The unique ID of the issued command, used to retrieve the command status.
}

// UserEmail represents the UserEmail schema from the OpenAPI specification
type UserEmail struct {
	TypeField string `json:"type,omitempty"` // Each entry can have a type which indicates standard types of that entry. For example email could be of home, work etc. In addition to the standard type, an entry can have a custom type and can take any value Such types should have the CUSTOM value as type and also have a customType value.
	Address string `json:"address,omitempty"` // Email id of the user.
	Customtype string `json:"customType,omitempty"` // Custom Type.
	Primary bool `json:"primary,omitempty"` // If this is user's primary email. Only one entry could be marked as primary.
	Public_key_encryption_certificates map[string]interface{} `json:"public_key_encryption_certificates,omitempty"` // Public Key Encryption Certificates. Current limit: 1 per email address, and 5 per user.
}
