# RELEASE NOTE

## release 日期 2025-06-09

v2.20.0 release (tower version 4.6.0)

### feature

- 为 model 类新增 MarshalOpts struct 字段，用于控制字段是否会被显式的序列化为 null 还是被忽略，默认被忽略：
  - 通过设置 `MarshalOpts.${field}_Explicit_Null_When_Empty` 为 true 来显式设置字段在被设置为 nil 的时候被序列化为 null，否则会被忽略。
- [VM], [VMVolume]: 新增 `used_size` 与 `used_size_usage` 字段，用于获取虚拟机与虚拟卷的已使用容量
- [ResourceChangeApi]: 新增 [GetResourceChange] API 用于获取 Resource Change Events
- [PciDeviceApi]: 新增 [GetPciDevices] API 用于获取 PCI 设备列表
- [TaskApi]: [CreateTask] 新增 `status` 字段，用于指定任务状态
- [ReplicationPlanApi]: 新增 [GetReplicationPlan] API 用于获取复制计划列表
- [SmtpServerApi]: 新增 [GetSmtpServer] API 用于获取 SMTP 服务器列表
- [ReplicaVMApi]: 新增 [GetReplicaVms] API 用于获取复制虚拟机列表
- [NetworkPolicyRuleServiceApi]:
  - 新增 [GetNetworkPolicyRuleServices] API 用于获取「网络安全」的服务资源
  - 新增 [CreateNetworkPolicyRuleService] API 用于创建「网络安全」的服务资源
  - 新增 [UpdateNetworkPolicyRuleService] API 用于更新「网络安全」的服务资源
  - 新增 [DeleteNetworkPolicyRuleService] API 用于删除「网络安全」的服务资源
- [SecurityPolicyApi]: [SecurityPolicyIngressEgressInput] 支持配置 `service_ids` 用于指定「网络安全」的服务资源
- [SecurityPolicyApi]: [IPSecurityPolicy] 支持配置 `ip_block`，用于从白名单/黑名单中排除部分 IP
- [IsolationPolicyApi]:
  - 新增 [CreateIsolationPolicy] API 用于创建「隔离策略」
  - 新增 [UpdateIsolationPolicy] API 用于更新「隔离策略」
  - 新增 [DeleteIsolationPolicy] API 用于删除「隔离策略」
- [VirtualPrivateCloudExternalSubnetGroupApi]: 新增 [GetVirtualPrivateCloudExternalSubnetGroups] API 用于获取「外部子网组」
- [VirtualPrivateCloudEdgeGatewayApi]: 新增 [GetVirtualPrivateCloudEdgeGateways] API 用于获取「边缘网关」
- [VirtualPrivateCloudEdgeGatewayGroupApi]: 新增 [GetVirtualPrivateCloudEdgeGatewayGroups] API 用于获取「边缘网关组」
- [VirtualPrivateCloudNatGatewayApi]: [VirtualPrivateCloudNatGatewayCreateParams] 新增 `external_subnet_group_id` 用于配置外部子网组；新增 `external_ips` 用于配置主备转换地址。
- [VirtualPrivateCloudNatGatewayApi]: [VirtualPrivateCloudRouteGatewayUpdateParams] 新增 `external_ips` 用于配置主备转换地址。
- [VirtualPrivateCloudRouterGatewayApi]:
  - [VirtualPrivateCloudRouterGatewayCreateParams] 新增 `external_subnet_group_id` 用于配置外部子网组；新增 `external_ips` 用于配置主备转换地址。
  - [VirtualPrivateCloudRouterGatewayUpdateParams] 新增 `external_ips` 用于配置主备转换地址。
- [VirtualPrivateCloudFloatingIPApi]: 新增 [BatchCreateVirtualPrivateCloudFloatingIPs] API，用于批量分配浮动 IP。

### bugfix

- [VMApi]: [GetVMVncInfo] 修复生成的 vnc redirect url
- [VMApi]: [AddVMNic] [UpdateVMNic] [UpdateVMNicBasicInfo] 修复虚拟机网卡无法正确编辑 ip

## release 日期 2025-02-17

v2.19.0 release (tower version 4.5.0)

### breaking change

- [TaskApi]: [UpdateTask]: Descripton 字段由 *string 更新为 *TaskDescription

### feature

- [BackupPlanApi]: 新增 GetBackupRestorePointMetadata API 用于获取备份恢复点元数据;
- [NtpApi]: 新增 GetNtpServiceURL API 用于获取 Ntp 服务 URL;
- [ClusterApi]: 新增 GetClusterStorageInfo API 用于获取集群存储信息;
- [ObservabilityApi]: 新增 ClearSystemServiceAlertNotificationConfig API 用于清除系统服务报警信息

### optimize

- [BackupPlanApi]: [CreateBackupPlan]: 在 IncrementalPeriod 为 weekly 时，校验 IncrementalWeekdays 是否已输入;
- 为以下资源新增返回字段
  - [BackupPlan]: 新增 Vms, Phase, LastExecuteStatusMessage, LastManualExecuteStatusMessage, BackupRestorePointCount, ValidSizeOfRestorePoint, BackupTotalSize, LogicalSize, BackupDelayOption, DeleteStrategy, BackupPlanExecutions 和 BackupRestorePoints 字段;
  - [NestedVirtualPrivateCloudService]: 新增 InternalCidr 和 TepIpPools 字段;
  - [VirtualPrivateCloud]: 新增 VpcService 字段;
  - [SecurityPolicy]: 新增 IsBlocklist 字段;
  - [ContentLibraryVmTemplate]: 新增 VmDisks, VmNics, ClockOffset, Cpu, CpuModel, Firmware, Ha, IoPolicy, LocalCreatedAt, MaxBandwidth, MaxBandwidthPolicy, MaxIops, MaxIopsPolicy, TemplateConfig, VideoType, WinOpt 和 ZbsStorageInfo 字段;
  - [VirtualPrivateCloudExternalSubnets]: 新增 EdgeGateway 和 Exclusive 字段;
- [TaskApi]: [CreateTask], [UpdateTask]: 支持传入 StartedAt 和 FinishedAt 字段;
- [ErrorBody]: 为 ErrorBody 实现了 stringer 接口，方便打印错误信息;

### bugfix

- [IscsiLunApi]: [CopyIscsiLun]: 修复跨集群克隆 iscsi lun 失败;
- [ObservabilityApi]: [DisassociateSystemServiceFromObsService]: 修复解除关系错误时无法正确返回错误信息。

## release 日期 2024-10-14

v2.18.0 release (tower version 4.4.0)

### feature

- [VMApi]: [CreateVMFromContentLibraryTemplateBatch]: 支持批量通过内容库模板创建虚拟机
- 新增备份相关 API
  - [BackupPlanExecutionApi]: 备份计划执行记录查询
  - [BackupPlanApi]: 备份计划管理 API
  - [BackupRestoreExecutionApi]: 备份计划恢复点执行记录查询
  - [BackupRestorePointApi]: 备份计划恢复点管理 API
  - [BackupServiceApi]: 备份服务 API
  - [BackupStoreRepositoryApi]: 备份存储库 API
  - [BackupTargetExecutionApi]: 备份虚拟机执行记录查询

### optimize

- [VMApi]: [InstallVMTools]: 优化虚拟机镜像挂载，现在无需输入虚拟机工具镜像 ID 即可挂载
- [LogCollectionApi]: [ForceStopLogCollection]: 优化执行，不再会将日志收集任务置为失败
- [ContentLanguage]: 支持配置为 `*` 以同时返回所有支持的语言，返回值依旧是 string 类型，可以被 JSON 序列化为语言和实际值的键值对
- [SecurityPolicyApi]: 支持配置 alg_protocol
- [HostApi]: [CreateHost]: 支持添加主机时配置主机账户密码
- [SecurityPolicyIngressEgressInput]: 新增 type 字段 `SecurityPolicyFlowControlType` 以支持配置全部流量

### bugfix

- [SecurityGroupApi]: [UpdateSecurityGroup]: 修复无法正确更新安全组名称与描述

## release 日期 2024-08-21

v2.17.0 release (tower version 4.3.0)

### feature

- 新增 VPC 相关 API
  - [VirtualPrivateCloudApi]: VPC 服务管理
  - [VirtualPrivateCloudClusterBindingApi]: VPC 服务与集群关联关系
  - [VirtualPrivateCloudExternalSubnetApi]: VPC 网络外部子网
  - [VirtualPrivateCloudFloatingIpApi]: VPC 网络浮动 IP 管理
  - [VirtualPrivateCloudNatGatewayApi]: VPC 网络 NAT 网关管理
  - [VirtualPrivateCloudRouterGatewayApi]: VPC 网络路由网关管理
  - [VirtualPrivateCloudRouteTableApi]: VPC 网络路由表管理
  - [VirtualPrivateCloudSecurityGroupApi]: VPC 网络安全组管理
  - [VirtualPrivateCloudSecurityPolicyApi]: VPC 网络安全策略管理
  - [VirtualPrivateCloudSubnetApi]: VPC 子网管理
- [VMApi]: [UpdateVMNicVpc]: 更新虚拟机的 VPC 网卡

### update

- [ContentLibraryImage]: 新增 `IscsiLuns` 字段，用于记录分发至存算分离节点中的镜像的 LUN。
- [ContentLibraryImageApi]: 支持使用 `IscsiLuns` 进行搜索
- [VMApi]:
  - [CreateVM], [CreateVMFromContentLibraryTemplate], [CreateVMFromTemplate], [CloneVM], [RebuildVMFromSnapshot]: 支持配置虚拟机的 VPC 网卡
  - [AddVMNic], [UpdateVMNic]: 支持配置 VPC 网卡

### bugfix

- [VMApi]: [UpdateVMDisk], [AddVMDisk], [ExpandVMDisk], [RemoveVMDisk], [AddVMCdRom], [RemoveVMCdRom], [EjectIsoFromVMCdRom], [ToggleVMCdRomDisable]: 修复被修改虚拟机的虚拟盘会丢失限速设置以及 CD-ROM 的禁用的问题
- [VMApi]: [CreateVMFromContentLibraryTemplate]: 修复当内容库模板仅分发在存算分离的集群时，无法创建虚拟机的问题。

## release 日期 2024-07-16

v2.16.0 release (tower version 4.2.0)

### update

- [HostApi]:
  - [CreateHost]: 新增 `Vdses` 字段适配网络融合，新增 `ZbsSpec` 字段适配 zbs 560 以上版本集群的添加
- [VMApi]:
  - [GetVms]: 新增 `BiosUUID` 字段

## release 日期 2024-05-30

v2.15.1 release (tower version 4.1.0)

### bugfix

- [VM], [VMVolume], [IscsiLun], [NvmfNameSpace], [NfsInode]: 更新 `UniqueLogicalSize` 类型为 `float64`

## release 日期 2024-05-11

v2.15.0 release (tower version 4.1.0)

### update

- [CommonHeader]: 新增默认返回 Header 类
- [NestedLabel], [LabelApi]: `LabelApi` 将会返回一致的 `Label` 相关的属性，NestedLabel 额外返回 `key` 和 `value`。
- [VM], [VMVolume], [IscsiLun], [NvmfNameSpace], [NfsInode]: 新增返回 `UniqueLogicalSize` 用于表示资源的独占逻辑容量

### bugfix

- [VMApi]
  - [AddVMNic], [RemoveVMNic], [RemoveVMNicByWhere], [UpdateVMNic], [UpdateVMNicBasicInfo], [UpdateVMNicAdvanceInfo], [UpdateVMNicQosOption]: 修复更新虚拟机网卡类 API 编辑后丢失部分网卡信息的问题
- [VMPlacementGroupApi]
  - [CreateVMPlacementGroup], [UpdateVMPlacementGroup]: 修复了创建与更新放置组时，`VMVMPolicyEnabled` 为 false 时无法更新组内虚拟机成员

## release 日期 2024-01-04

v2.14.0 release (tower version 4.0.0)

### update

- [ContentLibraryImageApi]:
  - [ImportContentLibraryImage]: 新增通过 url 导入内容库镜像 API
- [GpuDeviceApi]:
  - [GetDetailVMInfoByGpuDevices]: 新增获取 GPU 设备关联的虚拟机 API
- [HostApi]:
  - [EnterMaintenanceMode]: 新增进入维护模式 API
  - [EnterMaintenanceModePreCheck]: 新增进入维护模式预检 API
  - [EnterMaintenanceModePrecheckResult]: 新增获取进入维护模式预检结果 API
  - [ExitMaintenanceMode]: 新增离开维护模式 API
  - [ExitMaintenanceModePrecheckResult]: 新增离开维护模式预检 API
  - [PowerOffHost]: 新增主机电源操作 API，用于关闭、重启主机
- [IscsiLunAPi]:
  - [CopyIscsiLun]: 新增复制 iscsi lun API
- [NicApi]:
  - [GetNics]: 额外返回 iommu_status 以及关联的虚拟机
- [TaskApi]:
  - [CreateTask]: 新增创建 Task API
  - [UpdateTask]: 新增更新 Task API
- [UserAuditLogApi]:
  - [CreateUserAuditLog]: 新增创建事件审计 API
- [VMApi]:
  - [GetVMGpuDeviceInfo]: 新增根据虚拟机获取其挂载的 GPU 设备信息 API
  - [GetVMVncInfo]: 新增获取 vm vnc 信息 API
  - [MigrateVMAcrossCluster]\: 优化了报错逻辑，当目标主机存在且不在目标集群上时，直接报错
  - [RebuildVMFromSnapshot]: 重建虚拟机 API 支持配置 PCI 网卡
  - [CreateVM], [CreateVMFromTemplate], [CreateVMFromContentLibraryTemplate], [CloneVM], [RebuildVMFromSnapshot]: 创建虚拟机时支持配置虚拟机所属用户

### bugfix

- [ContentLibraryImageApi]:
  - [updateContentLibraryImage]: 修复了更新内容库镜像时，没有传递 name 会失败的问题
- [ElfImageApi]:
  - [updateElfImage]: 修复了更新镜像时，没有传递 name 会失败的问题
- [VMApi]:
  - [createVMFromContentLibraryTemplate]\: 修复模板卷存储策略包含三副本时，非完全克隆虚拟机失败，提示非完全克隆无法修改存储策略的问题

## release 日期 2023-11-07

v2.13.0 release (tower version 3.4.0)

### update

- [ROLEACTION] 新增 `SMTX_INSPECTOR`

## release 日期 2023-10-19

v2.12.0 release (tower version 3.3.0)

### update

- [GpuDeviceApi] 新增 GPU 设备的支持，新增以下 API：
  - [GetGpuDevices] 获取 GPU 设备列表
  - [GetGpuDeviceConnections] 获取 GPU 设备数量
  - [SwitchGpuDeviceSriov] 切换 GPU 设备 sr-iov 开启
  - [UpdateGpuDeviceDescription] 更新 GPU 设备描述
  - [UpdateGpuDeviceUsage] 更新 GPU 设备用途
- [VMApi] 新增 GPU 设备支持：
  - 支持创建虚拟机，克隆虚拟机，从内容库模板创建虚拟机时额外配置 GPU 设备，需要指定主机
  - 新增 [AddVMGpuDevice] 为已有虚拟机挂载 GPU 设备
  - 新增 [RemoveVMGpuDevice] 为已有虚拟机卸载 GPU 设备
- [VMApi] 新增 PCI 网卡支持：
  - 支持创建虚拟机，克隆虚拟机，从内容库模板创建虚拟机时额外配置 PCI 网卡，需要指定主机
  - 新增 [AddVMPciNic] 支持为已有虚拟机挂载 PCI 网卡
  - 新增 [RemoveVMPciNic] 支持为已有的虚拟机卸载 PCI 网卡
- [VMApi] 优化卸载网卡:
  - [RemoveVMNic] `nic_index` 作为删除标记不够稳定，弃用
  - [RemoveVMNicByWhere] 新方法，支持使用 where 条件来筛选需要的网卡，对于 VMNic 而言，可以配合使用 vm + mac_address 的筛选， `{vm:<vm_where>, mac_address:<mac_address>}` 的形式来较为稳定的筛选出需要的 VMNic
- [NicApi] [UpdateNic] 支持通过 `nic_user_usage`