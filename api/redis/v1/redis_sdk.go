// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package redis provides methods and message types of the redis v1 API.
package redis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ fmt.Stringer
	_ json.Unmarshaler
	_ url.URL
	_ net.IP
	_ http.Header
	_ bytes.Reader
	_ time.Time
	_ = strings.Join

	_ scw.ScalewayRequest
	_ marshaler.Duration
	_ scw.File
	_ = parameter.AddToQuery
	_ = namegenerator.GetRandomName
)

// API: managed Database for Redis™ API.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type AvailableClusterSettingPropertyType string

const (
	AvailableClusterSettingPropertyTypeUNKNOWN = AvailableClusterSettingPropertyType("UNKNOWN")
	AvailableClusterSettingPropertyTypeBOOLEAN = AvailableClusterSettingPropertyType("BOOLEAN")
	AvailableClusterSettingPropertyTypeINT     = AvailableClusterSettingPropertyType("INT")
	AvailableClusterSettingPropertyTypeSTRING  = AvailableClusterSettingPropertyType("STRING")
)

func (enum AvailableClusterSettingPropertyType) String() string {
	if enum == "" {
		// return default value if empty
		return "UNKNOWN"
	}
	return string(enum)
}

func (enum AvailableClusterSettingPropertyType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AvailableClusterSettingPropertyType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AvailableClusterSettingPropertyType(AvailableClusterSettingPropertyType(tmp).String())
	return nil
}

type ClusterStatus string

const (
	ClusterStatusUnknown      = ClusterStatus("unknown")
	ClusterStatusReady        = ClusterStatus("ready")
	ClusterStatusProvisioning = ClusterStatus("provisioning")
	ClusterStatusConfiguring  = ClusterStatus("configuring")
	ClusterStatusDeleting     = ClusterStatus("deleting")
	ClusterStatusError        = ClusterStatus("error")
	ClusterStatusAutohealing  = ClusterStatus("autohealing")
	ClusterStatusLocked       = ClusterStatus("locked")
	ClusterStatusSuspended    = ClusterStatus("suspended")
	ClusterStatusInitializing = ClusterStatus("initializing")
)

func (enum ClusterStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ClusterStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ClusterStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ClusterStatus(ClusterStatus(tmp).String())
	return nil
}

type ListClustersRequestOrderBy string

const (
	ListClustersRequestOrderByCreatedAtAsc  = ListClustersRequestOrderBy("created_at_asc")
	ListClustersRequestOrderByCreatedAtDesc = ListClustersRequestOrderBy("created_at_desc")
	ListClustersRequestOrderByNameAsc       = ListClustersRequestOrderBy("name_asc")
	ListClustersRequestOrderByNameDesc      = ListClustersRequestOrderBy("name_desc")
)

func (enum ListClustersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListClustersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListClustersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListClustersRequestOrderBy(ListClustersRequestOrderBy(tmp).String())
	return nil
}

type NodeTypeStock string

const (
	NodeTypeStockUnknown    = NodeTypeStock("unknown")
	NodeTypeStockLowStock   = NodeTypeStock("low_stock")
	NodeTypeStockOutOfStock = NodeTypeStock("out_of_stock")
	NodeTypeStockAvailable  = NodeTypeStock("available")
)

func (enum NodeTypeStock) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum NodeTypeStock) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NodeTypeStock) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NodeTypeStock(NodeTypeStock(tmp).String())
	return nil
}

type PrivateNetworkProvisioningMode string

const (
	PrivateNetworkProvisioningModeStatic = PrivateNetworkProvisioningMode("static")
	PrivateNetworkProvisioningModeIpam   = PrivateNetworkProvisioningMode("ipam")
)

func (enum PrivateNetworkProvisioningMode) String() string {
	if enum == "" {
		// return default value if empty
		return "static"
	}
	return string(enum)
}

func (enum PrivateNetworkProvisioningMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PrivateNetworkProvisioningMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PrivateNetworkProvisioningMode(PrivateNetworkProvisioningMode(tmp).String())
	return nil
}

// ACLRule: acl rule.
type ACLRule struct {
	// ID: ID of the rule.
	ID string `json:"id"`
	// IPCidr: iPv4 network address of the rule.
	IPCidr *scw.IPNet `json:"ip_cidr"`
	// Description: description of the rule.
	Description *string `json:"description"`
}

// ACLRuleSpec: acl rule spec.
type ACLRuleSpec struct {
	// IPCidr: iPv4 network address of the rule.
	IPCidr scw.IPNet `json:"ip_cidr"`
	// Description: description of the rule.
	Description string `json:"description"`
}

// AddACLRulesResponse: add acl rules response.
type AddACLRulesResponse struct {
	// ACLRules: ACL Rules enabled for the Database Instance.
	ACLRules []*ACLRule `json:"acl_rules"`
	// TotalCount: total count of ACL rules of the Database Instance.
	TotalCount uint32 `json:"total_count"`
}

// AddEndpointsResponse: add endpoints response.
type AddEndpointsResponse struct {
	// Endpoints: endpoints defined on the Database Instance.
	Endpoints []*Endpoint `json:"endpoints"`
	// TotalCount: total count of endpoints of the Database Instance.
	TotalCount uint32 `json:"total_count"`
}

// AvailableClusterSetting: available cluster setting.
type AvailableClusterSetting struct {
	// Name: name of the setting.
	Name string `json:"name"`
	// DefaultValue: default value of the setting.
	DefaultValue *string `json:"default_value"`
	// Type: type of setting.
	// Default value: UNKNOWN
	Type AvailableClusterSettingPropertyType `json:"type"`
	// Description: description of the setting.
	Description string `json:"description"`
	// MaxValue: optional maximum value of the setting.
	MaxValue *int64 `json:"max_value"`
	// MinValue: optional minimum value of the setting.
	MinValue *int64 `json:"min_value"`
	// Regex: optional validation rule of the setting.
	Regex *string `json:"regex"`
	// Deprecated: defines whether or not the setting is deprecated.
	Deprecated bool `json:"deprecated"`
}

// Cluster: cluster.
type Cluster struct {
	// ID: UUID of the Database Instance.
	ID string `json:"id"`
	// Name: name of the Database Instance.
	Name string `json:"name"`
	// ProjectID: project ID the Database Instance belongs to.
	ProjectID string `json:"project_id"`
	// Status: status of the Database Instance.
	// Default value: unknown
	Status ClusterStatus `json:"status"`
	// Version: redis™ engine version of the Database Instance.
	Version string `json:"version"`
	// Endpoints: list of Database Instance endpoints.
	Endpoints []*Endpoint `json:"endpoints"`
	// Tags: list of tags applied to the Database Instance.
	Tags []string `json:"tags"`
	// NodeType: node type of the Database Instance.
	NodeType string `json:"node_type"`
	// CreatedAt: creation date (Format ISO 8601).
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: update date (Format ISO 8601).
	UpdatedAt *time.Time `json:"updated_at"`
	// TLSEnabled: defines whether or not TLS is enabled.
	TLSEnabled bool `json:"tls_enabled"`
	// ClusterSettings: list of Database Instance settings.
	ClusterSettings []*ClusterSetting `json:"cluster_settings"`
	// ACLRules: list of ACL rules.
	ACLRules []*ACLRule `json:"acl_rules"`
	// ClusterSize: number of nodes of the Database Instance cluster.
	ClusterSize uint32 `json:"cluster_size"`
	// Zone: zone of the Database Instance.
	Zone scw.Zone `json:"zone"`
	// UserName: name of the user associated to the cluster.
	UserName string `json:"user_name"`
	// UpgradableVersions: list of engine versions the Database Instance can upgrade to.
	UpgradableVersions []string `json:"upgradable_versions"`
}

// ClusterMetricsResponse: cluster metrics response.
type ClusterMetricsResponse struct {
	// Timeseries: time series of metrics of a given cluster.
	Timeseries []*scw.TimeSeries `json:"timeseries"`
}

// ClusterSetting: cluster setting.
type ClusterSetting struct {
	// Value: value of the setting.
	Value string `json:"value"`
	// Name: name of the setting.
	Name string `json:"name"`
}

// ClusterSettingsResponse: cluster settings response.
type ClusterSettingsResponse struct {
	// Settings: settings configured for a given Database Instance.
	Settings []*ClusterSetting `json:"settings"`
}

// ClusterVersion: cluster version.
type ClusterVersion struct {
	// Version: redis™ engine version.
	Version string `json:"version"`
	// EndOfLifeAt: date of End of Life.
	EndOfLifeAt *time.Time `json:"end_of_life_at"`
	// AvailableSettings: cluster settings available to be updated.
	AvailableSettings []*AvailableClusterSetting `json:"available_settings"`
	// LogoURL: redis™ logo url.
	LogoURL string `json:"logo_url"`
	// Zone: zone of the Redis™ Database Instance.
	Zone scw.Zone `json:"zone"`
}

// Endpoint: endpoint.
type Endpoint struct {
	// Port: TCP port of the endpoint.
	Port uint32 `json:"port"`
	// PrivateNetwork: private Network details.
	// Precisely one of PrivateNetwork, PublicNetwork must be set.
	PrivateNetwork *PrivateNetwork `json:"private_network,omitempty"`
	// PublicNetwork: public network details.
	// Precisely one of PrivateNetwork, PublicNetwork must be set.
	PublicNetwork *PublicNetwork `json:"public_network,omitempty"`
	// IPs: list of IPv4 addresses of the endpoint.
	IPs []net.IP `json:"ips"`
	// ID: UUID of the endpoint.
	ID string `json:"id"`
}

// EndpointSpec: endpoint spec.
type EndpointSpec struct {
	// PrivateNetwork: private Network specification details.
	// Precisely one of PrivateNetwork, PublicNetwork must be set.
	PrivateNetwork *EndpointSpecPrivateNetworkSpec `json:"private_network,omitempty"`
	// PublicNetwork: public network specification details.
	// Precisely one of PrivateNetwork, PublicNetwork must be set.
	PublicNetwork *EndpointSpecPublicNetworkSpec `json:"public_network,omitempty"`
}

// EndpointSpecPrivateNetworkSpec: endpoint spec. private network spec.
type EndpointSpecPrivateNetworkSpec struct {
	// ID: UUID of the Private Network to connect to the Database Instance.
	ID string `json:"id"`
	// ServiceIPs: endpoint IPv4 address with a CIDR notation. You must provide at least one IPv4 per node.
	ServiceIPs []scw.IPNet `json:"service_ips"`
	// IpamConfig: automated configuration of your Private Network endpoint with Scaleway IPAM service.
	IpamConfig *EndpointSpecPrivateNetworkSpecIpamConfig `json:"ipam_config"`
}

type EndpointSpecPrivateNetworkSpecIpamConfig struct {
}

// EndpointSpecPublicNetworkSpec: endpoint spec. public network spec.
type EndpointSpecPublicNetworkSpec struct {
}

// ListClusterVersionsResponse: list cluster versions response.
type ListClusterVersionsResponse struct {
	// Versions: list of available Redis™ engine versions.
	Versions []*ClusterVersion `json:"versions"`
	// TotalCount: total count of available Redis™ engine versions.
	TotalCount uint32 `json:"total_count"`
}

// ListClustersResponse: list clusters response.
type ListClustersResponse struct {
	// Clusters: list all Database Instances.
	Clusters []*Cluster `json:"clusters"`
	// TotalCount: total count of Database Instances.
	TotalCount uint32 `json:"total_count"`
}

// ListNodeTypesResponse: list node types response.
type ListNodeTypesResponse struct {
	// NodeTypes: types of node.
	NodeTypes []*NodeType `json:"node_types"`
	// TotalCount: total count of node types available.
	TotalCount uint32 `json:"total_count"`
}

// NodeType: node type.
type NodeType struct {
	// Name: node type name.
	Name string `json:"name"`
	// StockStatus: current stock status of the node type.
	// Default value: unknown
	StockStatus NodeTypeStock `json:"stock_status"`
	// Description: current specifications of the offer.
	Description string `json:"description"`
	// Vcpus: number of virtual CPUs.
	Vcpus uint32 `json:"vcpus"`
	// Memory: quantity of RAM.
	Memory scw.Size `json:"memory"`
	// Disabled: defines whether node type is currently disabled or not.
	Disabled bool `json:"disabled"`
	// Beta: defines whether node type is currently in beta.
	Beta bool `json:"beta"`
	// Zone: zone of the node type.
	Zone scw.Zone `json:"zone"`
}

// PrivateNetwork: private network.
type PrivateNetwork struct {
	// ID: UUID of the Private Network.
	ID string `json:"id"`
	// ServiceIPs: list of IPv4 CIDR notation addresses of the endpoint.
	ServiceIPs []scw.IPNet `json:"service_ips"`
	// Zone: zone of the Private Network.
	Zone scw.Zone `json:"zone"`
	// ProvisioningMode: how your endpoint ips are provisioned.
	// Default value: static
	ProvisioningMode PrivateNetworkProvisioningMode `json:"provisioning_mode"`
}

type PublicNetwork struct {
}

// SetACLRulesResponse: set acl rules response.
type SetACLRulesResponse struct {
	// ACLRules: ACL Rules enabled for the Database Instance.
	ACLRules []*ACLRule `json:"acl_rules"`
}

// SetEndpointsResponse: set endpoints response.
type SetEndpointsResponse struct {
	// Endpoints: endpoints defined on the Database Instance.
	Endpoints []*Endpoint `json:"endpoints"`
}

// Service API

// Zones list localities the api is available in
func (s *API) Zones() []scw.Zone {
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw1, scw.ZonePlWaw2}
}

type CreateClusterRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ProjectID: project ID in which to create the Database Instance.
	ProjectID string `json:"project_id"`
	// Name: name of the Database Instance.
	Name string `json:"name"`
	// Version: redis™ engine version of the Database Instance.
	Version string `json:"version"`
	// Tags: tags to apply to the Database Instance.
	Tags []string `json:"tags"`
	// NodeType: type of node to use for the Database Instance.
	NodeType string `json:"node_type"`
	// UserName: name of the user created upon Database Instance creation.
	UserName string `json:"user_name"`
	// Password: password of the user.
	Password string `json:"password"`
	// ClusterSize: number of nodes in the Redis™ cluster.
	ClusterSize *int32 `json:"cluster_size"`
	// ACLRules: list of ACLRuleSpec used to secure your publicly exposed cluster.
	ACLRules []*ACLRuleSpec `json:"acl_rules"`
	// Endpoints: zero or multiple EndpointSpec used to expose your cluster publicly and inside Private Networks.
	// Zero or multiple EndpointSpec used to expose your cluster publicly and inside private networks. If no EndpoindSpec is given the cluster will be publicly exposed by default.
	Endpoints []*EndpointSpec `json:"endpoints"`
	// TLSEnabled: defines whether or not TLS is enabled.
	TLSEnabled bool `json:"tls_enabled"`
	// ClusterSettings: list of advanced settings to be set upon Database Instance initialization.
	ClusterSettings []*ClusterSetting `json:"cluster_settings"`
}

// CreateCluster: create a Redis™ Database Instance.
// Create a new Redis™ Database Instance (Redis™ cluster). You must set the `zone`, `project_id`, `version`, `node_type`, `user_name` and `password` parameters. Optionally you can define `acl_rules`, `endpoints`, `tls_enabled` and `cluster_settings`.
func (s *API) CreateCluster(req *CreateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("ins")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateClusterRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance to update.
	ClusterID string `json:"-"`
	// Name: name of the Database Instance.
	Name *string `json:"name"`
	// Tags: database Instance tags.
	Tags *[]string `json:"tags"`
	// UserName: name of the Database Instance user.
	UserName *string `json:"user_name"`
	// Password: password of the Database Instance user.
	Password *string `json:"password"`
}

// UpdateCluster: update a Redis™ Database Instance.
// Update the parameters of a Redis™ Database Instance (Redis™ cluster), including `name`, `tags`, `user_name` and `password`.
func (s *API) UpdateCluster(req *UpdateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetClusterRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster.
	ClusterID string `json:"-"`
}

// GetCluster: get a Redis™ Database Instance.
// Retrieve information about a Redis™ Database Instance (Redis™ cluster). Specify the `cluster_id` and `region` in your request to get information such as `id`, `status`, `version`, `tls_enabled`, `cluster_settings`, `upgradable_versions` and `endpoints` about your cluster in the response.
func (s *API) GetCluster(req *GetClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
		Headers: http.Header{},
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListClustersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// Tags: filter by Database Instance tags.
	Tags []string `json:"-"`
	// Name: filter by Database Instance names.
	Name *string `json:"-"`
	// OrderBy: criteria to use when ordering the list.
	// Default value: created_at_asc
	OrderBy ListClustersRequestOrderBy `json:"-"`
	// ProjectID: filter by Project ID.
	ProjectID *string `json:"-"`
	// OrganizationID: filter by Organization ID.
	OrganizationID *string `json:"-"`
	// Version: filter by Redis™ engine version.
	Version *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListClusters: list Redis™ Database Instances.
// List all Redis™ Database Instances (Redis™ cluster) in the specified zone. By default, the Database Instances returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as `tags`, `name`, `organization_id` and `version`.
func (s *API) ListClusters(req *ListClustersRequest, opts ...scw.RequestOption) (*ListClustersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "version", req.Version)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListClustersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type MigrateClusterRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance to update.
	ClusterID string `json:"-"`
	// Version: redis™ engine version of the Database Instance.
	// Precisely one of ClusterSize, NodeType, Version must be set.
	Version *string `json:"version,omitempty"`
	// NodeType: type of node to use for the Database Instance.
	// Precisely one of ClusterSize, NodeType, Version must be set.
	NodeType *string `json:"node_type,omitempty"`
	// ClusterSize: number of nodes for the Database Instance.
	// Precisely one of ClusterSize, NodeType, Version must be set.
	ClusterSize *uint32 `json:"cluster_size,omitempty"`
}

// MigrateCluster: scale up a Redis™ Database Instance.
// Upgrade your standalone Redis™ Database Instance node, either by upgrading to a bigger node type (vertical scaling) or by adding more nodes to your Database Instance to increase your number of endpoints and distribute cache (horizontal scaling). Note that scaling horizontally your Redis™ Database Instance will not renew its TLS certificate. In order to refresh the TLS certificate, you must use the Renew TLS certificate endpoint.
func (s *API) MigrateCluster(req *MigrateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/migrate",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteClusterRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance to delete.
	ClusterID string `json:"-"`
}

// DeleteCluster: delete a Redis™ Database Instance.
// Delete a Redis™ Database Instance (Redis™ cluster), specified by the `region` and `cluster_id` parameters. Deleting a Database Instance is permanent, and cannot be undone. Note that upon deletion all your data will be lost.
func (s *API) DeleteCluster(req *DeleteClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
		Headers: http.Header{},
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetClusterMetricsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster.
	ClusterID string `json:"-"`
	// StartAt: start date.
	StartAt *time.Time `json:"-"`
	// EndAt: end date.
	EndAt *time.Time `json:"-"`
	// MetricName: name of the metric to gather.
	MetricName *string `json:"-"`
}

// GetClusterMetrics: get metrics of a Redis™ Database Instance.
// Retrieve the metrics of a Redis™ Database Instance (Redis™ cluster). You can define the period from which to retrieve metrics by specifying the `start_date` and `end_date`.
func (s *API) GetClusterMetrics(req *GetClusterMetricsRequest, opts ...scw.RequestOption) (*ClusterMetricsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "start_at", req.StartAt)
	parameter.AddToQuery(query, "end_at", req.EndAt)
	parameter.AddToQuery(query, "metric_name", req.MetricName)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/metrics",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ClusterMetricsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListNodeTypesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// IncludeDisabledTypes: defines whether or not to include disabled types.
	IncludeDisabledTypes bool `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListNodeTypes: list available node types.
// List all available node types. By default, the node types returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListNodeTypes(req *ListNodeTypesRequest, opts ...scw.RequestOption) (*ListNodeTypesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "include_disabled_types", req.IncludeDisabledTypes)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/node-types",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListNodeTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListClusterVersionsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// IncludeDisabled: defines whether or not to include disabled Redis™ engine versions.
	IncludeDisabled bool `json:"-"`
	// IncludeBeta: defines whether or not to include beta Redis™ engine versions.
	IncludeBeta bool `json:"-"`
	// IncludeDeprecated: defines whether or not to include deprecated Redis™ engine versions.
	IncludeDeprecated bool `json:"-"`
	// Version: list Redis™ engine versions that match a given name pattern.
	Version *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListClusterVersions: list available Redis™ versions.
// List the Redis™ database engine versions available. You can define additional parameters for your query, such as `include_disabled`, `include_beta`, `include_deprecated` and `version`.
func (s *API) ListClusterVersions(req *ListClusterVersionsRequest, opts ...scw.RequestOption) (*ListClusterVersionsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "include_disabled", req.IncludeDisabled)
	parameter.AddToQuery(query, "include_beta", req.IncludeBeta)
	parameter.AddToQuery(query, "include_deprecated", req.IncludeDeprecated)
	parameter.AddToQuery(query, "version", req.Version)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/cluster-versions",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListClusterVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetClusterCertificateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster.
	ClusterID string `json:"-"`
}

// GetClusterCertificate: get the TLS certificate of a cluster.
// Retrieve information about the TLS certificate of a Redis™ Database Instance (Redis™ cluster). Details like name and content are returned in the response.
func (s *API) GetClusterCertificate(req *GetClusterCertificateRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/certificate",
		Headers: http.Header{},
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RenewClusterCertificateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster.
	ClusterID string `json:"-"`
}

// RenewClusterCertificate: renew the TLS certificate of a cluster.
// Renew a TLS certificate for a Redis™ Database Instance (Redis™ cluster). Renewing a certificate means that you will not be able to connect to your Database Instance using the previous certificate. You will also need to download and update the new certificate for all database clients.
func (s *API) RenewClusterCertificate(req *RenewClusterCertificateRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/renew-certificate",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type AddClusterSettingsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance you want to add settings to.
	ClusterID string `json:"-"`
	// Settings: settings to add to the cluster.
	Settings []*ClusterSetting `json:"settings"`
}

// AddClusterSettings: add advanced settings.
// Add an advanced setting to a Redis™ Database Instance (Redis™ cluster). You must set the `name` and the `value` of each setting.
func (s *API) AddClusterSettings(req *AddClusterSettingsRequest, opts ...scw.RequestOption) (*ClusterSettingsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/settings",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ClusterSettingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteClusterSettingRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance where the settings must be set.
	ClusterID string `json:"-"`
	// SettingName: setting name to delete.
	SettingName string `json:"-"`
}

// DeleteClusterSetting: delete advanced setting.
// Delete an advanced setting in a Redis™ Database Instance (Redis™ cluster). You must specify the names of the settings you want to delete in the request body.
func (s *API) DeleteClusterSetting(req *DeleteClusterSettingRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	if fmt.Sprint(req.SettingName) == "" {
		return nil, errors.New("field SettingName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/settings/" + fmt.Sprint(req.SettingName) + "",
		Headers: http.Header{},
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetClusterSettingsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance where the settings must be set.
	ClusterID string `json:"-"`
	// Settings: settings to define for the Database Instance.
	Settings []*ClusterSetting `json:"settings"`
}

// SetClusterSettings: set advanced settings.
// Update an advanced setting for a Redis™ Database Instance (Redis™ cluster). Settings added upon database engine initalization can only be defined once, and cannot, therefore, be updated.
func (s *API) SetClusterSettings(req *SetClusterSettingsRequest, opts ...scw.RequestOption) (*ClusterSettingsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/settings",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ClusterSettingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetACLRulesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance where the ACL rules have to be set.
	ClusterID string `json:"-"`
	// ACLRules: aCLs rules to define for the cluster.
	ACLRules []*ACLRuleSpec `json:"acl_rules"`
}

// SetACLRules: set ACL rules for a cluster.
// Replace all the ACL rules of a Redis™ Database Instance (Redis™ cluster).
func (s *API) SetACLRules(req *SetACLRulesRequest, opts ...scw.RequestOption) (*SetACLRulesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/acls",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type AddACLRulesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance you want to add ACL rules to.
	ClusterID string `json:"-"`
	// ACLRules: aCLs rules to add to the cluster.
	ACLRules []*ACLRuleSpec `json:"acl_rules"`
}

// AddACLRules: add ACL rules for a cluster.
// Add an additional ACL rule to a Redis™ Database Instance (Redis™ cluster).
func (s *API) AddACLRules(req *AddACLRulesRequest, opts ...scw.RequestOption) (*AddACLRulesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/acls",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteACLRuleRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ACLID: UUID of the ACL rule you want to delete.
	ACLID string `json:"-"`
}

// DeleteACLRule: delete an ACL rule for a cluster.
// Delete an ACL rule of a Redis™ Database Instance (Redis™ cluster). You must specify the `acl_id` of the rule you want to delete in your request.
func (s *API) DeleteACLRule(req *DeleteACLRuleRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ACLID) == "" {
		return nil, errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/acls/" + fmt.Sprint(req.ACLID) + "",
		Headers: http.Header{},
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetACLRuleRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ACLID: UUID of the ACL rule you want to get.
	ACLID string `json:"-"`
}

// GetACLRule: get an ACL rule.
// Retrieve information about an ACL rule of a Redis™ Database Instance (Redis™ cluster). You must specify the `acl_id` of the rule in your request.
func (s *API) GetACLRule(req *GetACLRuleRequest, opts ...scw.RequestOption) (*ACLRule, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ACLID) == "" {
		return nil, errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/acls/" + fmt.Sprint(req.ACLID) + "",
		Headers: http.Header{},
	}

	var resp ACLRule

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetEndpointsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance where the endpoints have to be set.
	ClusterID string `json:"-"`
	// Endpoints: endpoints to define for the Database Instance.
	Endpoints []*EndpointSpec `json:"endpoints"`
}

// SetEndpoints: set endpoints for a cluster.
// Update an endpoint for a Redis™ Database Instance (Redis™ cluster). You must specify the `cluster_id` and the `endpoints` parameters in your request.
func (s *API) SetEndpoints(req *SetEndpointsRequest, opts ...scw.RequestOption) (*SetEndpointsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/endpoints",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetEndpointsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type AddEndpointsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance you want to add endpoints to.
	ClusterID string `json:"-"`
	// Endpoints: endpoints to add to the Database Instance.
	Endpoints []*EndpointSpec `json:"endpoints"`
}

// AddEndpoints: add endpoints for a cluster.
// Add a new endpoint for a Redis™ Database Instance (Redis™ cluster). You can add `private_network` or `public_network` specifications to the body of the request.
func (s *API) AddEndpoints(req *AddEndpointsRequest, opts ...scw.RequestOption) (*AddEndpointsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/endpoints",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddEndpointsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteEndpointRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// EndpointID: UUID of the endpoint you want to delete.
	EndpointID string `json:"-"`
}

// DeleteEndpoint: delete an endpoint for a cluster.
// Delete the endpoint of a Redis™ Database Instance (Redis™ cluster). You must specify the `region` and `endpoint_id` parameters of the endpoint you want to delete. Note that might need to update any environment configurations that point to the deleted endpoint.
func (s *API) DeleteEndpoint(req *DeleteEndpointRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.EndpointID) == "" {
		return nil, errors.New("field EndpointID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
		Headers: http.Header{},
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetEndpointRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// EndpointID: UUID of the endpoint you want to get.
	EndpointID string `json:"-"`
}

// GetEndpoint: get an endpoint.
// Retrieve information about a Redis™ Database Instance (Redis™ cluster) endpoint. Full details about the endpoint, like `ips`, `port`, `private_network` and `public_network` specifications are returned in the response.
func (s *API) GetEndpoint(req *GetEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.EndpointID) == "" {
		return nil, errors.New("field EndpointID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
		Headers: http.Header{},
	}

	var resp Endpoint

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateEndpointRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	EndpointID string `json:"-"`

	// Precisely one of PrivateNetwork, PublicNetwork must be set.
	PrivateNetwork *EndpointSpecPrivateNetworkSpec `json:"private_network,omitempty"`

	// Precisely one of PrivateNetwork, PublicNetwork must be set.
	PublicNetwork *EndpointSpecPublicNetworkSpec `json:"public_network,omitempty"`
}

// UpdateEndpoint: update an endpoint.
// Update information about a Redis™ Database Instance (Redis™ cluster) endpoint. Full details about the endpoint, like `ips`, `port`, `private_network` and `public_network` specifications are returned in the response.
func (s *API) UpdateEndpoint(req *UpdateEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.EndpointID) == "" {
		return nil, errors.New("field EndpointID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Endpoint

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListClustersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListClustersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListClustersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Clusters = append(r.Clusters, results.Clusters...)
	r.TotalCount += uint32(len(results.Clusters))
	return uint32(len(results.Clusters)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNodeTypesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNodeTypesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListNodeTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.NodeTypes = append(r.NodeTypes, results.NodeTypes...)
	r.TotalCount += uint32(len(results.NodeTypes))
	return uint32(len(results.NodeTypes)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListClusterVersionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListClusterVersionsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListClusterVersionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Versions = append(r.Versions, results.Versions...)
	r.TotalCount += uint32(len(results.Versions))
	return uint32(len(results.Versions)), nil
}
