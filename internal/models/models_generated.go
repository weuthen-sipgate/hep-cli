// Code generated from swagger.json — DO NOT EDIT.
// Regenerate with: go run ./tools/generate

package models

import "encoding/json"

// Ensure json import is used.
var _ json.RawMessage

// AdminTables represents the AdminTables schema.
type AdminTables struct {
	Count int64 `json:"count,omitempty"`
	Data []string `json:"data,omitempty"`
}

// AgentLocationCreateSuccessResponse represents the AgentLocationCreateSuccessResponse schema.
type AgentLocationCreateSuccessResponse struct {
	Data string `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// AgentLocationDeleteSuccessResponse represents the AgentLocationDeleteSuccessResponse schema.
type AgentLocationDeleteSuccessResponse struct {
	Data string `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// AgentLocationUpdateSuccessResponse represents the AgentLocationUpdateSuccessResponse schema.
type AgentLocationUpdateSuccessResponse struct {
	Data string `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// AgentsLocation represents the AgentsLocation schema.
type AgentsLocation struct {
	Active bool `json:"active,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	ExpireDate string `json:"expire_date,omitempty"`
	GID uint32 `json:"gid,omitempty"`
	Host string `json:"host,omitempty"`
	Node string `json:"node,omitempty"`
	Path string `json:"path,omitempty"`
	Port uint16 `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	TTL uint32 `json:"ttl,omitempty"`
	Type string `json:"type,omitempty"`
	UUID string `json:"uuid,omitempty"`
	Version uint64 `json:"version,omitempty"`
}

// AgentsLocationList represents the AgentsLocationList schema.
type AgentsLocationList struct {
	Data []AgentsLocation `json:"Data,omitempty"`
}

// AliasMapStruct represents the AliasMapStruct schema.
type AliasMapStruct struct {
	Alias string `json:"alias"`
	Firstipv4 string `json:"firstipv4"`
	Firstipv6 string `json:"firstipv6"`
	Group string `json:"group,omitempty"`
	IP string `json:"ip"`
	Ipobject json.RawMessage `json:"ipobject"`
	Ipv6 bool `json:"ipv6,omitempty"`
	Lastbinaryipv4 uint32 `json:"lastbinaryipv4"`
	Lastbinaryipv6 int64 `json:"lastbinaryipv6"`
	Lastipv4 string `json:"lastipv4"`
	Lastipv6 string `json:"lastipv6"`
	Mask uint16 `json:"mask"`
	Port uint16 `json:"port"`
	Servertype string `json:"servertype,omitempty"`
	Type uint32 `json:"type,omitempty"`
}

// AliasResponse represents the AliasResponse schema.
type AliasResponse struct {
	MetadataModificationTime uint32 `json:"metadata_modification_time,omitempty"`
	TotalBytes uint32 `json:"total_bytes,omitempty"`
	TotalRows uint32 `json:"total_rows,omitempty"`
}

// AliasStruct represents the AliasStruct schema.
type AliasStruct struct {
	Alias string `json:"alias"`
	Group string `json:"group"`
	IP string `json:"ip"`
	Ipobject string `json:"ipobject,omitempty"`
	Ipv6 bool `json:"ipv6,omitempty"`
	Mask uint16 `json:"mask"`
	Port uint16 `json:"port"`
	Servertype string `json:"servertype"`
	Shardid string `json:"shardid,omitempty"`
	Status bool `json:"status"`
	Type uint32 `json:"type"`
	UUID string `json:"uuid"`
	Version uint64 `json:"version,omitempty"`
}

// AliasSwaggerStruct represents the AliasSwaggerStruct schema.
type AliasSwaggerStruct struct {
	Alias string `json:"alias"`
	Group string `json:"group"`
	IP string `json:"ip"`
	Ipobject map[string]interface{} `json:"ipobject,omitempty"`
	Ipv6 bool `json:"ipv6,omitempty"`
	Mask uint16 `json:"mask"`
	Port uint16 `json:"port"`
	Servertype string `json:"servertype"`
	Shardid string `json:"shardid,omitempty"`
	Status bool `json:"status"`
	Type uint32 `json:"type"`
	UUID string `json:"uuid"`
	Version uint64 `json:"version,omitempty"`
}

// ArchiveResponse represents the ArchiveResponse schema.
type ArchiveResponse struct {
	File string `json:"File,omitempty"`
}

// CallElement represents the CallElement schema.
type CallElement struct {
	AliasDst string `json:"aliasDst,omitempty"`
	AliasSrc string `json:"aliasSrc,omitempty"`
	CreateDate int64 `json:"create_date,omitempty"`
	Destination int64 `json:"destination,omitempty"`
	DstHost string `json:"dstHost,omitempty"`
	DstID string `json:"dstId,omitempty"`
	DstIP string `json:"dstIp,omitempty"`
	DstPort float64 `json:"dstPort,omitempty"`
	ID float64 `json:"id,omitempty"`
	Method string `json:"method,omitempty"`
	MethodText string `json:"method_text,omitempty"`
	MicroTs int64 `json:"micro_ts,omitempty"`
	MsgColor string `json:"msg_color,omitempty"`
	Prot float64 `json:"prot,omitempty"`
	Protocol float64 `json:"protocol,omitempty"`
	RuriUser string `json:"ruri_user,omitempty"`
	Sid string `json:"sid,omitempty"`
	SrcHost string `json:"srcHost,omitempty"`
	SrcID string `json:"srcId,omitempty"`
	SrcIP string `json:"srcIp,omitempty"`
	SrcPort float64 `json:"srcPort,omitempty"`
	Table string `json:"table,omitempty"`
}

// ClickhouseObject represents the ClickhouseObject schema.
type ClickhouseObject struct {
	Query string `json:"query"`
}

// ClickhouseRawQuery represents the ClickhouseRawQuery schema.
type ClickhouseRawQuery struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Total int64 `json:"total,omitempty"`
}

// CreateUserStruct represents the CreateUserStruct schema.
type CreateUserStruct struct {
	Department string `json:"department"`
	Email string `json:"email"`
	Firstname string `json:"firstname"`
	Guid string `json:"guid,omitempty"`
	Lastname string `json:"lastname"`
	Params json.RawMessage `json:"params,omitempty"`
	Partid uint16 `json:"partid"`
	Password string `json:"password"`
	Usergroup string `json:"usergroup"`
	Username string `json:"username,omitempty"`
	Version uint64 `json:"version"`
}

// DbstatLog represents the DBStatLog schema.
type DbstatLog struct {
	Critical bool `json:"critical,omitempty"`
	Error string `json:"error,omitempty"`
	Time string `json:"time,omitempty"`
}

// Dbstats represents the DBStats schema.
type Dbstats struct {
	Idle int64 `json:"Idle,omitempty"`
	InUse int64 `json:"InUse,omitempty"`
	MaxIdleClosed int64 `json:"MaxIdleClosed,omitempty"`
	MaxIdleTimeClosed int64 `json:"MaxIdleTimeClosed,omitempty"`
	MaxLifetimeClosed int64 `json:"MaxLifetimeClosed,omitempty"`
	MaxOpenConnections int64 `json:"MaxOpenConnections,omitempty"`
	OpenConnections int64 `json:"OpenConnections,omitempty"`
	WaitCount int64 `json:"WaitCount,omitempty"`
	WaitDuration int64 `json:"WaitDuration,omitempty"`
}

// DashboardElementList represents the DashboardElementList schema.
type DashboardElementList struct {
	Auth string `json:"auth,omitempty"`
	Data []DashboardElements `json:"data,omitempty"`
	Status string `json:"status,omitempty"`
	Total int64 `json:"total,omitempty"`
}

// DashboardElements represents the DashboardElements schema.
type DashboardElements struct {
	Cssclass string `json:"cssclass,omitempty"`
	Href string `json:"href,omitempty"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Owner string `json:"owner,omitempty"`
	Param string `json:"param,omitempty"`
	Shared bool `json:"shared,omitempty"`
	Type int64 `json:"type,omitempty"`
	Weight float64 `json:"weight,omitempty"`
}

// DatabaseStatistic represents the DatabaseStatistic schema.
type DatabaseStatistic struct {
	DatabaseName string `json:"database_name,omitempty"`
	DatabaseVersion string `json:"database_version,omitempty"`
	DBErrorCount int64 `json:"db_error_count,omitempty"`
	DBErrorLog []DbstatLog `json:"db_error_log,omitempty"`
	DBStats Dbstats `json:"db_stats,omitempty"`
	LastCheck string `json:"last_check,omitempty"`
	LastError string `json:"last_error,omitempty"`
	LatencyAvg int64 `json:"latency_avg,omitempty"`
	LatencyMax int64 `json:"latency_max,omitempty"`
	LatencyMin int64 `json:"latency_min,omitempty"`
	Online bool `json:"online,omitempty"`
	Primary bool `json:"primary,omitempty"`
}

// ExportActionActive represents the ExportActionActive schema.
type ExportActionActive struct {
	Data map[string]interface{} `json:"data,omitempty"`
}

// ExportActionRtpagent represents the ExportActionRTPagent schema.
type ExportActionRtpagent struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Total int64 `json:"total,omitempty"`
}

// ExportCallData represents the ExportCallData schema.
type ExportCallData struct {
	Param map[string]interface{} `json:"param"`
	Timestamp map[string]interface{} `json:"timestamp,omitempty"`
}

// FailureResponse represents the FailureResponse schema.
type FailureResponse struct {
	Error string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Statuscode int64 `json:"statuscode,omitempty"`
}

// GlobalSettingsCreateSuccessfulResponse represents the GlobalSettingsCreateSuccessfulResponse schema.
type GlobalSettingsCreateSuccessfulResponse struct {
	Data string `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// GlobalSettingsDeleteSuccessfulResponse represents the GlobalSettingsDeleteSuccessfulResponse schema.
type GlobalSettingsDeleteSuccessfulResponse struct {
	Data string `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// GlobalSettingsStruct represents the GlobalSettingsStruct schema.
type GlobalSettingsStruct struct {
	Category string `json:"category,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
	Guid string `json:"guid,omitempty"`
	Param string `json:"param,omitempty"`
	Partid uint16 `json:"partid,omitempty"`
	Setting json.RawMessage `json:"setting,omitempty"`
	Version uint64 `json:"version,omitempty"`
}

// GlobalSettingsStructList represents the GlobalSettingsStructList schema.
type GlobalSettingsStructList struct {
	Count int64 `json:"count,omitempty"`
	Data []GlobalSettingsStruct `json:"data,omitempty"`
}

// GlobalSettingsUpdateSuccessfulResponse represents the GlobalSettingsUpdateSuccessfulResponse schema.
type GlobalSettingsUpdateSuccessfulResponse struct {
	Data string `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// HepsubCreateSuccessResponse represents the HepsubCreateSuccessResponse schema.
type HepsubCreateSuccessResponse struct {
	Data string `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// HepsubDeleteSuccessResponse represents the HepsubDeleteSuccessResponse schema.
type HepsubDeleteSuccessResponse struct {
	Data string `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// HepsubSchema represents the HepsubSchema schema.
type HepsubSchema struct {
	CreateDate string `json:"create_date,omitempty"`
	Guid string `json:"guid,omitempty"`
	HEPAlias string `json:"hep_alias,omitempty"`
	Hepid uint16 `json:"hepid,omitempty"`
	Mapping json.RawMessage `json:"mapping,omitempty"`
	Profile string `json:"profile,omitempty"`
	Version uint64 `json:"version,omitempty"`
}

// HepsubSchemaList represents the HepsubSchemaList schema.
type HepsubSchemaList struct {
	Count int64 `json:"count,omitempty"`
	Data []HepsubSchema `json:"data,omitempty"`
}

// HepsubUpdateSuccessResponse represents the HepsubUpdateSuccessResponse schema.
type HepsubUpdateSuccessResponse struct {
	Data string `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// IpaliasFileDownload represents the IPAliasFileDownload schema.
type IpaliasFileDownload struct {
	File string `json:"File,omitempty"`
}

// IpaliasFileUpload represents the IPAliasFileUpload schema.
type IpaliasFileUpload struct {
	File string `json:"File,omitempty"`
}

// InterceptionAgentRequest represents the InterceptionAgentRequest schema.
type InterceptionAgentRequest struct {
	Liid uint32 `json:"LIID,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Description string `json:"description,omitempty"`
	GID uint32 `json:"gid,omitempty"`
	Number string `json:"number,omitempty"`
	ResellerID uint32 `json:"reseller_id,omitempty"`
	SearchCallee string `json:"search_callee,omitempty"`
	SearchCaller string `json:"search_caller,omitempty"`
	SearchIP string `json:"search_ip,omitempty"`
	SIPDomain string `json:"sip_domain,omitempty"`
	SIPUsername string `json:"sip_username,omitempty"`
	TsCreate int64 `json:"ts_create"`
	TsModify int64 `json:"ts_modify"`
	TsStart int64 `json:"ts_start"`
	TsStop int64 `json:"ts_stop"`
	UUID string `json:"uuid"`
}

// InterceptionsStruct represents the InterceptionsStruct schema.
type InterceptionsStruct struct {
	CreateDate string `json:"create_date"`
	Delivery json.RawMessage `json:"delivery"`
	Description string `json:"description,omitempty"`
	GID uint32 `json:"gid,omitempty"`
	Liid uint32 `json:"liid,omitempty"`
	ModifyDate string `json:"modify_date,omitempty"`
	SearchCallee string `json:"search_callee,omitempty"`
	SearchCaller string `json:"search_caller,omitempty"`
	SearchIP string `json:"search_ip,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	Status bool `json:"status,omitempty"`
	StopDate string `json:"stop_date,omitempty"`
	UUID string `json:"uuid"`
	Version uint64 `json:"version,omitempty"`
}

// LabelData represents the LabelData schema.
type LabelData struct {
	Entries []LabelData `json:"entries,omitempty"`
	Labels map[string]interface{} `json:"labels,omitempty"`
}

// LegacyAlias represents the LegacyAlias schema.
type LegacyAlias struct {
	Group string `json:"Group,omitempty"`
	IP string `json:"IP"`
	Ipbits uint16 `json:"IPBits"`
	Port uint16 `json:"Port"`
	ServerType string `json:"ServerType,omitempty"`
	ShardID string `json:"ShardID,omitempty"`
	Type uint32 `json:"Type"`
	Ipobject string `json:"ipobject,omitempty"`
}

// LinkResponse represents the LinkResponse schema.
type LinkResponse struct {
	Data map[string]interface{} `json:"data,omitempty"`
}

// ListUsers represents the ListUsers schema.
type ListUsers struct {
	Count int64 `json:"count,omitempty"`
	Data []CreateUserStruct `json:"data,omitempty"`
}

// MappingSchema represents the MappingSchema schema.
type MappingSchema struct {
	ApplyTTLAll bool `json:"apply_ttl_all,omitempty"`
	CorrelationMapping json.RawMessage `json:"correlation_mapping,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	CreateIndex json.RawMessage `json:"create_index,omitempty"`
	CreateTable string `json:"create_table,omitempty"`
	FieldsMapping json.RawMessage `json:"fields_mapping,omitempty"`
	FieldsSettings json.RawMessage `json:"fields_settings,omitempty"`
	Guid string `json:"guid,omitempty"`
	HEPAlias string `json:"hep_alias,omitempty"`
	Hepid uint16 `json:"hepid,omitempty"`
	Partid uint16 `json:"partid,omitempty"`
	PartitionStep uint16 `json:"partition_step,omitempty"`
	Profile string `json:"profile,omitempty"`
	Retention uint16 `json:"retention,omitempty"`
	SchemaMapping json.RawMessage `json:"schema_mapping,omitempty"`
	SchemaSettings json.RawMessage `json:"schema_settings,omitempty"`
	TableName string `json:"table_name,omitempty"`
	UserMapping json.RawMessage `json:"user_mapping,omitempty"`
	Version uint64 `json:"version,omitempty"`
}

// MessageDecoded represents the MessageDecoded schema.
type MessageDecoded struct {
	Data []map[string]interface{} `json:"data,omitempty"`
}

// Node represents the Node schema.
type Node struct {
	Arhive bool `json:"arhive,omitempty"`
	DBArchive string `json:"db_archive,omitempty"`
	DBName string `json:"db_name,omitempty"`
	Host string `json:"host,omitempty"`
	Name string `json:"name,omitempty"`
	Node string `json:"node,omitempty"`
	Online bool `json:"online,omitempty"`
	Primary bool `json:"primary,omitempty"`
	TablePrefix string `json:"table_prefix,omitempty"`
	Value string `json:"value,omitempty"`
}

// NodeList represents the NodeList schema.
type NodeList struct {
	Count int64 `json:"count,omitempty"`
	Data []Node `json:"data,omitempty"`
}

// OldaliasStruct represents the OLDAliasStruct schema.
type OldaliasStruct struct {
	Arguments []map[string]interface{} `json:"Arguments,omitempty"`
	CreateDate string `json:"CreateDate,omitempty"`
	DataMaps [][]map[string]interface{} `json:"DataMaps,omitempty"`
	DB string `json:"Db,omitempty"`
	ManualResync bool `json:"ManualResync,omitempty"`
	Node string `json:"Node,omitempty"`
	Query string `json:"Query,omitempty"`
	Status bool `json:"Status,omitempty"`
	Table string `json:"Table,omitempty"`
	Type string `json:"Type,omitempty"`
}

// Pcapresponse represents the PCAPResponse schema.
type Pcapresponse struct {
	File string `json:"File,omitempty"`
}

// PrometheusObject represents the PrometheusObject schema.
type PrometheusObject struct {
	Param map[string]interface{} `json:"param,omitempty"`
	Timestamp map[string]interface{} `json:"timestamp,omitempty"`
}

// RecordingTransactionRTP represents the RecordingTransactionRTP schema.
type RecordingTransactionRTP struct {
	Active bool `json:"active,omitempty"`
	CorrelationID string `json:"correlation_id,omitempty"`
	CreateDate uint32 `json:"create_date"`
	Date string `json:"date"`
	Direction uint8 `json:"direction,omitempty"`
	DstIP string `json:"dst_ip,omitempty"`
	DstPort uint16 `json:"dst_port,omitempty"`
	Filename string `json:"filename,omitempty"`
	Liid uint32 `json:"liid,omitempty"`
	Node string `json:"node,omitempty"`
	Proto uint8 `json:"proto,omitempty"`
	RecordDatetime string `json:"record_datetime"`
	Sid uint32 `json:"sid,omitempty"`
	SrcIP string `json:"src_ip,omitempty"`
	SrcPort uint16 `json:"src_port,omitempty"`
	Ssrc uint32 `json:"ssrc,omitempty"`
	Storedir string `json:"storedir,omitempty"`
	TimeSec uint32 `json:"time_sec,omitempty"`
	TimeUsec uint32 `json:"time_usec,omitempty"`
	Type string `json:"type,omitempty"`
	UUID string `json:"uuid"`
}

// RecordingTransactionSIP represents the RecordingTransactionSIP schema.
type RecordingTransactionSIP struct {
	Active bool `json:"active,omitempty"`
	CorrelationID string `json:"correlation_id,omitempty"`
	CreateDate uint32 `json:"create_date"`
	Date string `json:"date"`
	DstIP string `json:"dst_ip,omitempty"`
	DstPort uint16 `json:"dst_port,omitempty"`
	Filename string `json:"filename,omitempty"`
	Liid uint32 `json:"liid,omitempty"`
	Node string `json:"node,omitempty"`
	Proto uint8 `json:"proto,omitempty"`
	RecordDatetime string `json:"record_datetime"`
	Sid uint32 `json:"sid,omitempty"`
	SrcIP string `json:"src_ip,omitempty"`
	SrcPort uint16 `json:"src_port,omitempty"`
	Storedir string `json:"storedir,omitempty"`
	TimeSec uint32 `json:"time_sec,omitempty"`
	TimeUsec uint32 `json:"time_usec,omitempty"`
	Type string `json:"type,omitempty"`
	UUID string `json:"uuid"`
}

// RemoteObject represents the RemoteObject schema.
type RemoteObject struct {
	Param map[string]interface{} `json:"param,omitempty"`
	Timestamp map[string]interface{} `json:"timestamp,omitempty"`
}

// RemoteResponseData represents the RemoteResponseData schema.
type RemoteResponseData struct {
	Data []map[string]interface{} `json:"Data,omitempty"`
}

// Sippresponse represents the SIPPResponse schema.
type Sippresponse struct {
	File string `json:"File,omitempty"`
}

// ScriptDataStruct represents the ScriptDataStruct schema.
type ScriptDataStruct struct {
	Data string `json:"data"`
	HEPAlias string `json:"hep_alias,omitempty"`
	Hepid uint16 `json:"hepid,omitempty"`
	Partid uint16 `json:"partid,omitempty"`
	Profile string `json:"profile,omitempty"`
	Status bool `json:"status"`
	Type string `json:"type"`
	UUID string `json:"uuid,omitempty"`
	Version uint64 `json:"version,omitempty"`
}

// SearchCallData represents the SearchCallData schema.
type SearchCallData struct {
	Data []CallElement `json:"data,omitempty"`
	Keys []string `json:"keys,omitempty"`
	Total int64 `json:"total,omitempty"`
}

// SearchCallExportPCAP represents the SearchCallExportPCAP schema.
type SearchCallExportPCAP struct {
	Param map[string]interface{} `json:"param"`
	Timestamp map[string]interface{} `json:"timestamp,omitempty"`
}

// SearchObject represents the SearchObject schema.
type SearchObject struct {
	Param map[string]interface{} `json:"param"`
	Timestamp map[string]interface{} `json:"timestamp,omitempty"`
}

// SearchObjectExportPCAP represents the SearchObjectExportPCAP schema.
type SearchObjectExportPCAP struct {
	Param map[string]interface{} `json:"param"`
	Timestamp map[string]interface{} `json:"timestamp,omitempty"`
}

// SearchTransactionData represents the SearchTransactionData schema.
type SearchTransactionData struct {
	Param map[string]interface{} `json:"param"`
	Timestamp map[string]interface{} `json:"timestamp,omitempty"`
}

// SearchTransactionQOS represents the SearchTransactionQos schema.
type SearchTransactionQOS struct {
	RTCP SearchTransactionRTPList `json:"rtcp,omitempty"`
	RTP SearchTransactionRTPList `json:"rtp,omitempty"`
}

// SearchTransactionResponse represents the SearchTransactionResponse schema.
type SearchTransactionResponse struct {
	Data []map[string]interface{} `json:"Data,omitempty"`
	Keys []string `json:"keys,omitempty"`
	Total int64 `json:"total,omitempty"`
}

// SearchTransactionRTCP represents the SearchTransactionRtcp schema.
type SearchTransactionRTCP struct {
	CaptureID string `json:"captureId,omitempty"`
	CapturePass string `json:"capturePass,omitempty"`
	CorrelationID string `json:"correlation_id,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	Dbnode string `json:"dbnode,omitempty"`
	DstIP string `json:"dstIp,omitempty"`
	DstPort float64 `json:"dstPort,omitempty"`
	ID float64 `json:"id,omitempty"`
	Node []string `json:"node,omitempty"`
	PayloadType int64 `json:"payloadType,omitempty"`
	Profile string `json:"profile,omitempty"`
	Proto string `json:"proto,omitempty"`
	Protocol int64 `json:"protocol,omitempty"`
	ProtocolFamily int64 `json:"protocolFamily,omitempty"`
	Raw string `json:"raw,omitempty"`
	Sid string `json:"sid,omitempty"`
	SrcIP string `json:"srcIp,omitempty"`
	SrcPort float64 `json:"srcPort,omitempty"`
	TimeSeconds int64 `json:"timeSeconds,omitempty"`
	TimeUseconds int64 `json:"timeUseconds,omitempty"`
}

// SearchTransactionRTP represents the SearchTransactionRtp schema.
type SearchTransactionRTP struct {
	CaptureID string `json:"captureId,omitempty"`
	CapturePass string `json:"capturePass,omitempty"`
	CorrelationID string `json:"correlation_id,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	Dbnode string `json:"dbnode,omitempty"`
	DstIP string `json:"dstIp,omitempty"`
	DstPort float64 `json:"dstPort,omitempty"`
	ID float64 `json:"id,omitempty"`
	Node []string `json:"node,omitempty"`
	PayloadType int64 `json:"payloadType,omitempty"`
	Profile string `json:"profile,omitempty"`
	Proto string `json:"proto,omitempty"`
	Protocol int64 `json:"protocol,omitempty"`
	ProtocolFamily int64 `json:"protocolFamily,omitempty"`
	Raw string `json:"raw,omitempty"`
	Sid string `json:"sid,omitempty"`
	SrcIP string `json:"srcIp,omitempty"`
	SrcPort float64 `json:"srcPort,omitempty"`
	TimeSeconds int64 `json:"timeSeconds,omitempty"`
	TimeUseconds int64 `json:"timeUseconds,omitempty"`
}

// SearchTransactionRTPList represents the SearchTransactionRtpList schema.
type SearchTransactionRTPList struct {
	Data []SearchTransactionRTP `json:"data,omitempty"`
}

// StatisticObject represents the StatisticObject schema.
type StatisticObject struct {
	Param map[string]interface{} `json:"param,omitempty"`
	Timestamp map[string]interface{} `json:"timestamp,omitempty"`
}

// StatisticSearchObject represents the StatisticSearchObject schema.
type StatisticSearchObject struct {
	Param map[string]interface{} `json:"param,omitempty"`
	Timestamp map[string]interface{} `json:"timestamp,omitempty"`
}

// StenographerResponse represents the StenographerResponse schema.
type StenographerResponse struct {
	File string `json:"File,omitempty"`
}

// SuccessResponse represents the SuccessResponse schema.
type SuccessResponse struct {
	Count int64 `json:"count"`
	Data json.RawMessage `json:"data"`
	Message string `json:"message"`
}

// TableAuthToken represents the TableAuthToken schema.
type TableAuthToken struct {
	Active bool `json:"active,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	CreatorGuid string `json:"creator_guid,omitempty"`
	ExpireDate string `json:"expire_date,omitempty"`
	Guid string `json:"guid,omitempty"`
	IPAddress string `json:"ip_address,omitempty"`
	LastusageDate string `json:"lastusage_date,omitempty"`
	LimitCalls uint32 `json:"limit_calls,omitempty"`
	Name string `json:"name,omitempty"`
	UsageCalls uint32 `json:"usage_calls,omitempty"`
	UserObject json.RawMessage `json:"user_object,omitempty"`
	Usergroup string `json:"usergroup,omitempty"`
	Version uint64 `json:"version,omitempty"`
}

// TableLogsDataV2 represents the TableLogsDataV2 schema.
type TableLogsDataV2 struct {
	Callid string `json:"callid,omitempty"`
	Captid uint32 `json:"captid,omitempty"`
	CaptureIP string `json:"capture_ip,omitempty"`
	CreateTs uint64 `json:"create_ts,omitempty"`
	Data string `json:"data,omitempty"`
	DestinationIP string `json:"destination_ip,omitempty"`
	DestinationPort uint16 `json:"destination_port,omitempty"`
	Event string `json:"event,omitempty"`
	Guid string `json:"guid,omitempty"`
	Message string `json:"message,omitempty"`
	Node string `json:"node,omitempty"`
	Proto uint16 `json:"proto,omitempty"`
	SourceIP string `json:"source_ip,omitempty"`
	SourcePort uint16 `json:"source_port,omitempty"`
	Type string `json:"type,omitempty"`
	Vlan uint8 `json:"vlan,omitempty"`
}

// TableRTPStatsV2 represents the TableRtpStatsV2 schema.
type TableRTPStatsV2 struct {
	Callid string `json:"callid,omitempty"`
	Captid uint32 `json:"captid,omitempty"`
	CaptureIP string `json:"capture_ip,omitempty"`
	CreateTs uint64 `json:"create_ts,omitempty"`
	Data string `json:"data,omitempty"`
	DestinationIP string `json:"destination_ip,omitempty"`
	DestinationPort uint16 `json:"destination_port,omitempty"`
	Event string `json:"event,omitempty"`
	FrameProtocols string `json:"frame_protocols,omitempty"`
	Guid string `json:"guid,omitempty"`
	Message string `json:"message,omitempty"`
	Node string `json:"node,omitempty"`
	Proto uint16 `json:"proto,omitempty"`
	Raw string `json:"raw,omitempty"`
	SourceIP string `json:"source_ip,omitempty"`
	SourcePort uint16 `json:"source_port,omitempty"`
	Type string `json:"type,omitempty"`
}

// TableSIPMessagesCallV2 represents the TableSipMessagesCallV2 schema.
type TableSIPMessagesCallV2 struct {
	Callid string `json:"callid,omitempty"`
	Captid uint32 `json:"captid,omitempty"`
	CaptureIP string `json:"capture_ip,omitempty"`
	CreateTs uint64 `json:"create_ts,omitempty"`
	Data string `json:"data,omitempty"`
	DestinationIP string `json:"destination_ip,omitempty"`
	DestinationPort uint16 `json:"destination_port,omitempty"`
	Event string `json:"event,omitempty"`
	Guid string `json:"guid,omitempty"`
	Message string `json:"message,omitempty"`
	Node string `json:"node,omitempty"`
	Proto uint16 `json:"proto,omitempty"`
	RegionID string `json:"region_id,omitempty"`
	SourceIP string `json:"source_ip,omitempty"`
	SourcePort uint16 `json:"source_port,omitempty"`
	Vlan uint8 `json:"vlan,omitempty"`
}

// TableSIPRegistrationAllV2 represents the TableSipRegistrationAllV2 schema.
type TableSIPRegistrationAllV2 struct {
	Callid string `json:"callid,omitempty"`
	Captid uint32 `json:"captid,omitempty"`
	CaptureIP string `json:"capture_ip,omitempty"`
	CreateTs uint64 `json:"create_ts,omitempty"`
	Data string `json:"data,omitempty"`
	DestinationIP string `json:"destination_ip,omitempty"`
	DestinationPort uint16 `json:"destination_port,omitempty"`
	Event string `json:"event,omitempty"`
	Guid string `json:"guid,omitempty"`
	Message string `json:"message,omitempty"`
	Node string `json:"node,omitempty"`
	Proto uint16 `json:"proto,omitempty"`
	RegionID string `json:"region_id,omitempty"`
	SourceIP string `json:"source_ip,omitempty"`
	SourcePort uint16 `json:"source_port,omitempty"`
	Vlan uint8 `json:"vlan,omitempty"`
}

// TableSIPTransactionCallV2 represents the TableSipTransactionCallV2 schema.
type TableSIPTransactionCallV2 struct {
	AuthUser string `json:"auth_user,omitempty"`
	Callid string `json:"callid,omitempty"`
	Captid uint32 `json:"captid,omitempty"`
	CdrFailed uint64 `json:"cdr_failed,omitempty"`
	CdrStart uint64 `json:"cdr_start,omitempty"`
	CdrStop uint64 `json:"cdr_stop,omitempty"`
	ContactUser string `json:"contact_user,omitempty"`
	Correlations []string `json:"correlations,omitempty"`
	Cseq uint16 `json:"cseq,omitempty"`
	Custom1 string `json:"custom_1,omitempty"`
	Custom2 string `json:"custom_2,omitempty"`
	Custom3 string `json:"custom_3,omitempty"`
	Custom4 string `json:"custom_4,omitempty"`
	Custom5 string `json:"custom_5,omitempty"`
	Custom6 string `json:"custom_6,omitempty"`
	CustomstringKey []string `json:"customstring_key,omitempty"`
	CustomstringValue []string `json:"customstring_value,omitempty"`
	CustomuintKey []string `json:"customuint_key,omitempty"`
	CustomuintValue []uint32 `json:"customuint_value,omitempty"`
	Data string `json:"data,omitempty"`
	DestinationIP string `json:"destination_ip,omitempty"`
	DestinationPort uint16 `json:"destination_port,omitempty"`
	Event uint16 `json:"event,omitempty"`
	ExpireRep uint16 `json:"expire_rep,omitempty"`
	ExpireReq uint16 `json:"expire_req,omitempty"`
	Family uint8 `json:"family,omitempty"`
	FromUser string `json:"from_user,omitempty"`
	GeoCc string `json:"geo_cc,omitempty"`
	Guid string `json:"guid,omitempty"`
	IpgroupIn string `json:"ipgroup_in,omitempty"`
	IpgroupOut string `json:"ipgroup_out,omitempty"`
	Ips []string `json:"ips,omitempty"`
	MethodsKey []string `json:"methods_key,omitempty"`
	MethodsValue []uint32 `json:"methods_value,omitempty"`
	MetricsKey []string `json:"metrics_key,omitempty"`
	MetricsValue []float64 `json:"metrics_value,omitempty"`
	Realm string `json:"realm,omitempty"`
	RegionID string `json:"region_id,omitempty"`
	Rrd uint16 `json:"rrd,omitempty"`
	RuriUser string `json:"ruri_user,omitempty"`
	ServerTypeIn string `json:"server_type_in,omitempty"`
	ServerTypeOut string `json:"server_type_out,omitempty"`
	SourceIP string `json:"source_ip,omitempty"`
	SourcePort uint16 `json:"source_port,omitempty"`
	Srd uint16 `json:"srd,omitempty"`
	Status uint8 `json:"status,omitempty"`
	Termcode uint16 `json:"termcode,omitempty"`
	ToUser string `json:"to_user,omitempty"`
	Uas string `json:"uas,omitempty"`
	UpdateTs uint64 `json:"update_ts,omitempty"`
	Usergroup string `json:"usergroup,omitempty"`
	Vlan uint8 `json:"vlan,omitempty"`
	Xgroup string `json:"xgroup,omitempty"`
}

// TableUserList represents the TableUserList schema.
type TableUserList struct {
	Count int64 `json:"count,omitempty"`
	Data []CreateUserStruct `json:"data,omitempty"`
}

// TableVqrtcpxrStatsV2 represents the TableVqrtcpxrStatsV2 schema.
type TableVqrtcpxrStatsV2 struct {
	Callid string `json:"callid,omitempty"`
	Captid uint32 `json:"captid,omitempty"`
	CaptureIP string `json:"capture_ip,omitempty"`
	CreateTs uint64 `json:"create_ts,omitempty"`
	Data string `json:"data,omitempty"`
	DestinationIP string `json:"destination_ip,omitempty"`
	DestinationPort uint16 `json:"destination_port,omitempty"`
	Event string `json:"event,omitempty"`
	Guid string `json:"guid,omitempty"`
	Message string `json:"message,omitempty"`
	Mos uint16 `json:"mos,omitempty"`
	Node string `json:"node,omitempty"`
	Proto uint16 `json:"proto,omitempty"`
	SourceIP string `json:"source_ip,omitempty"`
	SourcePort uint16 `json:"source_port,omitempty"`
	Type string `json:"type,omitempty"`
	Vlan uint8 `json:"vlan,omitempty"`
}

// TextResponse represents the TextResponse schema.
type TextResponse struct {
	File string `json:"File,omitempty"`
}

// UserDeleteSuccessResponse represents the UserDeleteSuccessResponse schema.
type UserDeleteSuccessResponse struct {
	Data string `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// UserDetailsResponse represents the UserDetailsResponse schema.
type UserDetailsResponse struct {
	User map[string]interface{} `json:"user,omitempty"`
}

// UserFileDownload represents the UserFileDownload schema.
type UserFileDownload struct {
	File string `json:"File,omitempty"`
}

// UserFileUpload represents the UserFileUpload schema.
type UserFileUpload struct {
	File string `json:"File,omitempty"`
}

// UserGroupList represents the UserGroupList schema.
type UserGroupList struct {
	Count int64 `json:"count,omitempty"`
	Data []string `json:"data,omitempty"`
}

// UserLegacyStruct represents the UserLegacyStruct schema.
type UserLegacyStruct struct {
	Department string `json:"Department"`
	Email string `json:"Email"`
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	Params string `json:"Params,omitempty"`
	PartID uint16 `json:"PartID"`
	Password string `json:"Password"`
	PasswordHash string `json:"PasswordHash,omitempty"`
	UserGroup string `json:"UserGroup,omitempty"`
	UserName string `json:"UserName,omitempty"`
}

// UserLogin represents the UserLogin schema.
type UserLogin struct {
	Password string `json:"password"`
	Type string `json:"type,omitempty"`
	Username string `json:"username"`
}

// UserLoginSuccessResponse represents the UserLoginSuccessResponse schema.
type UserLoginSuccessResponse struct {
	Scope string `json:"scope,omitempty"`
	Token string `json:"token,omitempty"`
	User map[string]interface{} `json:"user,omitempty"`
}

// UserSettings represents the UserSettings schema.
type UserSettings struct {
	Data json.RawMessage `json:"data,omitempty"`
	Guid string `json:"guid,omitempty"`
	Params string `json:"params,omitempty"`
	ProtocolID string `json:"protocol_id,omitempty"`
}

// UserSuccessResponse represents the UserSuccessResponse schema.
type UserSuccessResponse struct {
	Data string `json:"data"`
	Message string `json:"message"`
}

// UserUpdateSuccessResponse represents the UserUpdateSuccessResponse schema.
type UserUpdateSuccessResponse struct {
	Data string `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

