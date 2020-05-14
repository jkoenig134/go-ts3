package go_ts3

// instanceedit `manage_scope`
type InstanceEditRequest struct {
	ServerinstanceGuestServerqueryGroup       int    `schema:"serverinstance_guest_serverquery_group,omitempty"`
	ServerinstanceTemplateServeradminGroup    int    `schema:"serverinstance_template_serveradmin_group,omitempty"`
	ServerinstanceFiletransferPort            int    `schema:"serverinstance_filetransfer_port,omitempty"`
	ServerinstanceMaxDownloadTotalBandwidth   uint64 `schema:"serverinstance_max_download_total_bandwidth,omitempty"`
	ServerinstanceMaxUploadTotalBandwidth     uint64 `schema:"serverinstance_max_upload_total_bandwidth,omitempty"`
	ServerinstanceTemplateServerdefaultGroup  int    `schema:"serverinstance_template_serverdefault_group,omitempty"`
	ServerinstanceTemplateChanneldefaultGroup int    `schema:"serverinstance_template_channeldefault_group,omitempty"`
	ServerinstanceTemplateChanneladminGroup   int    `schema:"serverinstance_template_channeladmin_group,omitempty"`
	ServerinstanceServerqueryFloodCommands    int    `schema:"serverinstance_serverquery_flood_commands,omitempty"`
	ServerinstanceServerqueryFloodTime        int    `schema:"serverinstance_serverquery_flood_time,omitempty"`
	ServerinstanceServerqueryFloodBanTime     int    `schema:"serverinstance_serverquery_flood_ban_time,omitempty"`
}

func (c *TeamspeakHttpClient) InstanceEdit(request InstanceEditRequest) error {
	return c.requestWithParams("instanceedit", request, nil)
}

// instanceinfo `manage_scope, write_scope, read_scope`
type InstanceInfo struct {
	ServerinstanceDatabaseVersion                int    `json:"serverinstance_database_version,string"`
	ServerinstanceFiletransferPort               int    `json:"serverinstance_filetransfer_port,string"`
	ServerinstanceGuestServerqueryGroup          int    `json:"serverinstance_guest_serverquery_group,string"`
	ServerinstanceMaxDownloadTotalBandwidth      uint64 `json:"serverinstance_max_download_total_bandwidth,string"`
	ServerinstanceMaxUploadTotalBandwidth        uint64 `json:"serverinstance_max_upload_total_bandwidth,string"`
	ServerinstancePendingConnectionsPerIP        int    `json:"serverinstance_pending_connections_per_ip,string"`
	ServerinstancePermissionsVersion             int    `json:"serverinstance_permissions_version,string"`
	ServerinstanceServerqueryBanTime             int    `json:"serverinstance_serverquery_ban_time,string"`
	ServerinstanceServerqueryFloodCommands       int    `json:"serverinstance_serverquery_flood_commands,string"`
	ServerinstanceServerqueryFloodTime           int    `json:"serverinstance_serverquery_flood_time,string"`
	ServerinstanceServerqueryMaxConnectionsPerIP int    `json:"serverinstance_serverquery_max_connections_per_ip,string"`
	ServerinstanceTemplateChannelAdminGroup      int    `json:"serverinstance_template_channeladmin_group,string"`
	ServerinstanceTemplateChannelDefaultGroup    int    `json:"serverinstance_template_channeldefault_group,string"`
	ServerinstanceTemplateServerAdminGroup       int    `json:"serverinstance_template_serveradmin_group,string"`
	ServerinstanceTemplateServerDefaultGroup     int    `json:"serverinstance_template_serverdefault_group,string"`
}

func (c *TeamspeakHttpClient) InstanceInfo() (*InstanceInfo, error) {
	var infos []InstanceInfo

	err := c.request("instanceinfo", &infos)
	if err != nil {
		return nil, err
	}

	return &infos[0], err
}
