// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/client/alert"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/alert_notifier"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/alert_rule"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/api_info"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/application"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/brick_topo"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/cloud_tower_application"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/cloud_tower_application_package"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/cluster"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/cluster_image"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/cluster_settings"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/cluster_topo"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/cluster_upgrade_history"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/consistency_group"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/consistency_group_snapshot"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/content_library_image"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/content_library_vm_template"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/datacenter"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/deploy"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/discovered_host"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/disk"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/elf_data_store"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/elf_image"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/elf_storage_policy"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/entity_filter"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/everoute_cluster"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/everoute_license"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/everoute_package"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/global_alert_rule"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/global_settings"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/graph"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/host"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/ipmi"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/iscsi_connection"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/iscsi_lun"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/iscsi_lun_snapshot"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/iscsi_target"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/isolation_policy"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/label"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/license"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/log_collection"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/log_service_config"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/metrics"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/namespace_group"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/nfs_export"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/nfs_inode"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/nic"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/node_topo"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/nvmf_namespace"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/nvmf_namespace_snapshot"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/nvmf_subsystem"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/organization"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/ovf"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/pmem_dimm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/rack_topo"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/report_task"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/report_template"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/security_group"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/security_policy"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/snapshot_group"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/snapshot_plan"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/snapshot_plan_task"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/snmp_transport"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/snmp_trap_receiver"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/svt_image"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/system_audit_log"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/table_reporter"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/task"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/upload_task"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/usb_device"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/user"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/user_audit_log"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/user_role_next"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vcenter_account"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vds"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/view"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vlan"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_disk"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_entity_filter_result"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_export_file"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_folder"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_nic"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_placement_group"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_snapshot"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_template"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_volume"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_volume_snapshot"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vsphere_esxi_account"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/witness"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/witness_service"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/zone"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/zone_topo"
)

// Default cloudtower HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new cloudtower HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Cloudtower {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new cloudtower HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Cloudtower {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new cloudtower client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Cloudtower {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Cloudtower)
	cli.Transport = transport
	cli.Alert = alert.New(transport, formats)
	cli.AlertNotifier = alert_notifier.New(transport, formats)
	cli.AlertRule = alert_rule.New(transport, formats)
	cli.APIInfo = api_info.New(transport, formats)
	cli.Application = application.New(transport, formats)
	cli.BrickTopo = brick_topo.New(transport, formats)
	cli.CloudTowerApplication = cloud_tower_application.New(transport, formats)
	cli.CloudTowerApplicationPackage = cloud_tower_application_package.New(transport, formats)
	cli.Cluster = cluster.New(transport, formats)
	cli.ClusterImage = cluster_image.New(transport, formats)
	cli.ClusterSettings = cluster_settings.New(transport, formats)
	cli.ClusterTopo = cluster_topo.New(transport, formats)
	cli.ClusterUpgradeHistory = cluster_upgrade_history.New(transport, formats)
	cli.ConsistencyGroup = consistency_group.New(transport, formats)
	cli.ConsistencyGroupSnapshot = consistency_group_snapshot.New(transport, formats)
	cli.ContentLibraryImage = content_library_image.New(transport, formats)
	cli.ContentLibraryVMTemplate = content_library_vm_template.New(transport, formats)
	cli.Datacenter = datacenter.New(transport, formats)
	cli.Deploy = deploy.New(transport, formats)
	cli.DiscoveredHost = discovered_host.New(transport, formats)
	cli.Disk = disk.New(transport, formats)
	cli.ElfDataStore = elf_data_store.New(transport, formats)
	cli.ElfImage = elf_image.New(transport, formats)
	cli.ElfStoragePolicy = elf_storage_policy.New(transport, formats)
	cli.EntityFilter = entity_filter.New(transport, formats)
	cli.EverouteCluster = everoute_cluster.New(transport, formats)
	cli.EverouteLicense = everoute_license.New(transport, formats)
	cli.EveroutePackage = everoute_package.New(transport, formats)
	cli.GlobalAlertRule = global_alert_rule.New(transport, formats)
	cli.GlobalSettings = global_settings.New(transport, formats)
	cli.Graph = graph.New(transport, formats)
	cli.Host = host.New(transport, formats)
	cli.Ipmi = ipmi.New(transport, formats)
	cli.IscsiConnection = iscsi_connection.New(transport, formats)
	cli.IscsiLun = iscsi_lun.New(transport, formats)
	cli.IscsiLunSnapshot = iscsi_lun_snapshot.New(transport, formats)
	cli.IscsiTarget = iscsi_target.New(transport, formats)
	cli.IsolationPolicy = isolation_policy.New(transport, formats)
	cli.Label = label.New(transport, formats)
	cli.License = license.New(transport, formats)
	cli.LogCollection = log_collection.New(transport, formats)
	cli.LogServiceConfig = log_service_config.New(transport, formats)
	cli.Metrics = metrics.New(transport, formats)
	cli.NamespaceGroup = namespace_group.New(transport, formats)
	cli.NfsExport = nfs_export.New(transport, formats)
	cli.NfsInode = nfs_inode.New(transport, formats)
	cli.Nic = nic.New(transport, formats)
	cli.NodeTopo = node_topo.New(transport, formats)
	cli.NvmfNamespace = nvmf_namespace.New(transport, formats)
	cli.NvmfNamespaceSnapshot = nvmf_namespace_snapshot.New(transport, formats)
	cli.NvmfSubsystem = nvmf_subsystem.New(transport, formats)
	cli.Organization = organization.New(transport, formats)
	cli.Ovf = ovf.New(transport, formats)
	cli.PmemDimm = pmem_dimm.New(transport, formats)
	cli.RackTopo = rack_topo.New(transport, formats)
	cli.ReportTask = report_task.New(transport, formats)
	cli.ReportTemplate = report_template.New(transport, formats)
	cli.SecurityGroup = security_group.New(transport, formats)
	cli.SecurityPolicy = security_policy.New(transport, formats)
	cli.SnapshotGroup = snapshot_group.New(transport, formats)
	cli.SnapshotPlan = snapshot_plan.New(transport, formats)
	cli.SnapshotPlanTask = snapshot_plan_task.New(transport, formats)
	cli.SnmpTransport = snmp_transport.New(transport, formats)
	cli.SnmpTrapReceiver = snmp_trap_receiver.New(transport, formats)
	cli.SvtImage = svt_image.New(transport, formats)
	cli.SystemAuditLog = system_audit_log.New(transport, formats)
	cli.TableReporter = table_reporter.New(transport, formats)
	cli.Task = task.New(transport, formats)
	cli.UploadTask = upload_task.New(transport, formats)
	cli.UsbDevice = usb_device.New(transport, formats)
	cli.User = user.New(transport, formats)
	cli.UserAuditLog = user_audit_log.New(transport, formats)
	cli.UserRoleNext = user_role_next.New(transport, formats)
	cli.VcenterAccount = vcenter_account.New(transport, formats)
	cli.Vds = vds.New(transport, formats)
	cli.View = view.New(transport, formats)
	cli.Vlan = vlan.New(transport, formats)
	cli.VM = vm.New(transport, formats)
	cli.VMDisk = vm_disk.New(transport, formats)
	cli.VMEntityFilterResult = vm_entity_filter_result.New(transport, formats)
	cli.VMExportFile = vm_export_file.New(transport, formats)
	cli.VMFolder = vm_folder.New(transport, formats)
	cli.VMNic = vm_nic.New(transport, formats)
	cli.VMPlacementGroup = vm_placement_group.New(transport, formats)
	cli.VMSnapshot = vm_snapshot.New(transport, formats)
	cli.VMTemplate = vm_template.New(transport, formats)
	cli.VMVolume = vm_volume.New(transport, formats)
	cli.VMVolumeSnapshot = vm_volume_snapshot.New(transport, formats)
	cli.VsphereEsxiAccount = vsphere_esxi_account.New(transport, formats)
	cli.Witness = witness.New(transport, formats)
	cli.WitnessService = witness_service.New(transport, formats)
	cli.Zone = zone.New(transport, formats)
	cli.ZoneTopo = zone_topo.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Cloudtower is a client for cloudtower
type Cloudtower struct {
	Alert alert.ClientService

	AlertNotifier alert_notifier.ClientService

	AlertRule alert_rule.ClientService

	APIInfo api_info.ClientService

	Application application.ClientService

	BrickTopo brick_topo.ClientService

	CloudTowerApplication cloud_tower_application.ClientService

	CloudTowerApplicationPackage cloud_tower_application_package.ClientService

	Cluster cluster.ClientService

	ClusterImage cluster_image.ClientService

	ClusterSettings cluster_settings.ClientService

	ClusterTopo cluster_topo.ClientService

	ClusterUpgradeHistory cluster_upgrade_history.ClientService

	ConsistencyGroup consistency_group.ClientService

	ConsistencyGroupSnapshot consistency_group_snapshot.ClientService

	ContentLibraryImage content_library_image.ClientService

	ContentLibraryVMTemplate content_library_vm_template.ClientService

	Datacenter datacenter.ClientService

	Deploy deploy.ClientService

	DiscoveredHost discovered_host.ClientService

	Disk disk.ClientService

	ElfDataStore elf_data_store.ClientService

	ElfImage elf_image.ClientService

	ElfStoragePolicy elf_storage_policy.ClientService

	EntityFilter entity_filter.ClientService

	EverouteCluster everoute_cluster.ClientService

	EverouteLicense everoute_license.ClientService

	EveroutePackage everoute_package.ClientService

	GlobalAlertRule global_alert_rule.ClientService

	GlobalSettings global_settings.ClientService

	Graph graph.ClientService

	Host host.ClientService

	Ipmi ipmi.ClientService

	IscsiConnection iscsi_connection.ClientService

	IscsiLun iscsi_lun.ClientService

	IscsiLunSnapshot iscsi_lun_snapshot.ClientService

	IscsiTarget iscsi_target.ClientService

	IsolationPolicy isolation_policy.ClientService

	Label label.ClientService

	License license.ClientService

	LogCollection log_collection.ClientService

	LogServiceConfig log_service_config.ClientService

	Metrics metrics.ClientService

	NamespaceGroup namespace_group.ClientService

	NfsExport nfs_export.ClientService

	NfsInode nfs_inode.ClientService

	Nic nic.ClientService

	NodeTopo node_topo.ClientService

	NvmfNamespace nvmf_namespace.ClientService

	NvmfNamespaceSnapshot nvmf_namespace_snapshot.ClientService

	NvmfSubsystem nvmf_subsystem.ClientService

	Organization organization.ClientService

	Ovf ovf.ClientService

	PmemDimm pmem_dimm.ClientService

	RackTopo rack_topo.ClientService

	ReportTask report_task.ClientService

	ReportTemplate report_template.ClientService

	SecurityGroup security_group.ClientService

	SecurityPolicy security_policy.ClientService

	SnapshotGroup snapshot_group.ClientService

	SnapshotPlan snapshot_plan.ClientService

	SnapshotPlanTask snapshot_plan_task.ClientService

	SnmpTransport snmp_transport.ClientService

	SnmpTrapReceiver snmp_trap_receiver.ClientService

	SvtImage svt_image.ClientService

	SystemAuditLog system_audit_log.ClientService

	TableReporter table_reporter.ClientService

	Task task.ClientService

	UploadTask upload_task.ClientService

	UsbDevice usb_device.ClientService

	User user.ClientService

	UserAuditLog user_audit_log.ClientService

	UserRoleNext user_role_next.ClientService

	VcenterAccount vcenter_account.ClientService

	Vds vds.ClientService

	View view.ClientService

	Vlan vlan.ClientService

	VM vm.ClientService

	VMDisk vm_disk.ClientService

	VMEntityFilterResult vm_entity_filter_result.ClientService

	VMExportFile vm_export_file.ClientService

	VMFolder vm_folder.ClientService

	VMNic vm_nic.ClientService

	VMPlacementGroup vm_placement_group.ClientService

	VMSnapshot vm_snapshot.ClientService

	VMTemplate vm_template.ClientService

	VMVolume vm_volume.ClientService

	VMVolumeSnapshot vm_volume_snapshot.ClientService

	VsphereEsxiAccount vsphere_esxi_account.ClientService

	Witness witness.ClientService

	WitnessService witness_service.ClientService

	Zone zone.ClientService

	ZoneTopo zone_topo.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Cloudtower) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Alert.SetTransport(transport)
	c.AlertNotifier.SetTransport(transport)
	c.AlertRule.SetTransport(transport)
	c.APIInfo.SetTransport(transport)
	c.Application.SetTransport(transport)
	c.BrickTopo.SetTransport(transport)
	c.CloudTowerApplication.SetTransport(transport)
	c.CloudTowerApplicationPackage.SetTransport(transport)
	c.Cluster.SetTransport(transport)
	c.ClusterImage.SetTransport(transport)
	c.ClusterSettings.SetTransport(transport)
	c.ClusterTopo.SetTransport(transport)
	c.ClusterUpgradeHistory.SetTransport(transport)
	c.ConsistencyGroup.SetTransport(transport)
	c.ConsistencyGroupSnapshot.SetTransport(transport)
	c.ContentLibraryImage.SetTransport(transport)
	c.ContentLibraryVMTemplate.SetTransport(transport)
	c.Datacenter.SetTransport(transport)
	c.Deploy.SetTransport(transport)
	c.DiscoveredHost.SetTransport(transport)
	c.Disk.SetTransport(transport)
	c.ElfDataStore.SetTransport(transport)
	c.ElfImage.SetTransport(transport)
	c.ElfStoragePolicy.SetTransport(transport)
	c.EntityFilter.SetTransport(transport)
	c.EverouteCluster.SetTransport(transport)
	c.EverouteLicense.SetTransport(transport)
	c.EveroutePackage.SetTransport(transport)
	c.GlobalAlertRule.SetTransport(transport)
	c.GlobalSettings.SetTransport(transport)
	c.Graph.SetTransport(transport)
	c.Host.SetTransport(transport)
	c.Ipmi.SetTransport(transport)
	c.IscsiConnection.SetTransport(transport)
	c.IscsiLun.SetTransport(transport)
	c.IscsiLunSnapshot.SetTransport(transport)
	c.IscsiTarget.SetTransport(transport)
	c.IsolationPolicy.SetTransport(transport)
	c.Label.SetTransport(transport)
	c.License.SetTransport(transport)
	c.LogCollection.SetTransport(transport)
	c.LogServiceConfig.SetTransport(transport)
	c.Metrics.SetTransport(transport)
	c.NamespaceGroup.SetTransport(transport)
	c.NfsExport.SetTransport(transport)
	c.NfsInode.SetTransport(transport)
	c.Nic.SetTransport(transport)
	c.NodeTopo.SetTransport(transport)
	c.NvmfNamespace.SetTransport(transport)
	c.NvmfNamespaceSnapshot.SetTransport(transport)
	c.NvmfSubsystem.SetTransport(transport)
	c.Organization.SetTransport(transport)
	c.Ovf.SetTransport(transport)
	c.PmemDimm.SetTransport(transport)
	c.RackTopo.SetTransport(transport)
	c.ReportTask.SetTransport(transport)
	c.ReportTemplate.SetTransport(transport)
	c.SecurityGroup.SetTransport(transport)
	c.SecurityPolicy.SetTransport(transport)
	c.SnapshotGroup.SetTransport(transport)
	c.SnapshotPlan.SetTransport(transport)
	c.SnapshotPlanTask.SetTransport(transport)
	c.SnmpTransport.SetTransport(transport)
	c.SnmpTrapReceiver.SetTransport(transport)
	c.SvtImage.SetTransport(transport)
	c.SystemAuditLog.SetTransport(transport)
	c.TableReporter.SetTransport(transport)
	c.Task.SetTransport(transport)
	c.UploadTask.SetTransport(transport)
	c.UsbDevice.SetTransport(transport)
	c.User.SetTransport(transport)
	c.UserAuditLog.SetTransport(transport)
	c.UserRoleNext.SetTransport(transport)
	c.VcenterAccount.SetTransport(transport)
	c.Vds.SetTransport(transport)
	c.View.SetTransport(transport)
	c.Vlan.SetTransport(transport)
	c.VM.SetTransport(transport)
	c.VMDisk.SetTransport(transport)
	c.VMEntityFilterResult.SetTransport(transport)
	c.VMExportFile.SetTransport(transport)
	c.VMFolder.SetTransport(transport)
	c.VMNic.SetTransport(transport)
	c.VMPlacementGroup.SetTransport(transport)
	c.VMSnapshot.SetTransport(transport)
	c.VMTemplate.SetTransport(transport)
	c.VMVolume.SetTransport(transport)
	c.VMVolumeSnapshot.SetTransport(transport)
	c.VsphereEsxiAccount.SetTransport(transport)
	c.Witness.SetTransport(transport)
	c.WitnessService.SetTransport(transport)
	c.Zone.SetTransport(transport)
	c.ZoneTopo.SetTransport(transport)
}
