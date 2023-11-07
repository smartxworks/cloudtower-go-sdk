# RELEASE NOTE

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
  - [RemoveVMNicByWhere] 新方法，支持使用 where 条件来筛选需要的网卡，对于 VmNic 而言，可以配合使用 vm + mac_address 的筛选， `{vm:<vm_where>, mac_address:<mac_address>}` 的形式来较为稳定的筛选出需要的 VmNic
- [NicApi] [UpdateNic] 支持通过 `nic_user_usage` 更新网口用途
- [LabelApi] 支持为 GPU 设备打标签

## release 日期 2023-09-01

v2.11.0 release (tower version 3.2.0)

### breaking change

- [VMApi]: [VMImportParams]: 由于 `VMNicParams` 支持配置 qos 配置，现在设计网卡的参数由 `VMNicParams` 更新为了 `VMImportNicParams`，类型和过去一致，没有 qos 支持。

### update

- feature: [VMVolumeApi]: [ImportVMVolume] [ExportVMVolume]: 现在支持虚拟卷的导入与导出
- feature: [ContentLibraryVMTemplateApi]: [ImportContentLibraryVMTemplate] [ExportContentLibraryVMTemplate]: 现在支持虚拟机模板的导入与导出
- feature: [VMApi]: [UpdateVMNicQosOption]: 现在支持通过 api 更新现有虚拟机网卡的 qos 属性
- optimize: [VMNicParams]: 现在支持配置 `qos` 参数，用于在创建虚拟机（直接创建，模板创建，快照重建，克隆创建）时配置虚拟机网卡的 qos 参数

### bugfix

- [VMApi]: [CreateVMFromTemplate] [CreateVMFromContentLibraryTemplate] [CloneVM] [RebuildVMFromSnapshot]: 修复了磁盘限速无法正确配置的问题，
  目前从模板创建，快照重建，克隆虚拟机时，仅支持配置与模板，快照或源虚拟机一致的 `io_policy` 并进行创建，原参数的 `io_policy` 被弃用
- [IscsiTargetApi] [IscsiLunApi] [NvmfSubsystemApi] [NvmfNamespaceApi]:
  [CreateIscsiTarget] [UpdateIscsiTarget] [CreateIscsiLun] [UpdateIscsiLun]
  [CreateNvmfSubsystem] [UpdateNvmfSubsystem] [CreateNvmfNamespace] [UpdateNvmfNamespace]:
  `bps_wr_maxLength` 可以被正常设置了
- [VMApi]: [CreateVMFromTemplate] [CreateVMFromContentLibraryTemplate]: 修复了单位无法被正常应用的 bug
- [GlobalAlertRuleApi]: [UpdateCustomizeAlertRule] 修复了特例报警无法正常更新的问题

## release 日期 2023-08-01

v2.10.0 release (tower version 3.1.0)

- feature: [AlertNotifierApi] 支持更新，删除以及创建新的报警通知配置
- optimize: [utils] 优化了 WaitTask 以及 WaitTasks，在任务失败的时候会返回失败任务的原因

## release 日期 2023-07-18

v2.9.1 release (tower version 3.0.0)

- optimize: optimize: 优化 NewWithUserConfig 登陆，当 UserConfig.Source 为 UserSourceLDAP 时，自动切换成当前的 LDAP 登录源进行登陆，减少迁移成本。

## release 日期 2023-07-03

v2.9.0 release (tower version 3.0.0)

- feature: [SecurityGroupApi] 支持创建，更新与删除安全组
- feature: [SecurityPolicy] 支持创建，更新与删除自定义安全策略
- feature: [OvfApi], [VMExportFileApi], [VMApi] 支持虚拟机的导入与导出
- feature: [VlanApi] 支持 trunk vlan 的创建与编辑
- feature: [UserApi] [Login] 支持使用 authn_id 登陆，旧 LDAP 登陆方式被废弃
- optimize: 为 [Host], [Nic], [UsbDevice], [VMVolume], [VMVolumeSnapshot] 添加了 `EntityAsyncStatus` 已判断资源目前的状态

## release 日期 2022-05-04

v2.8.0 release

- optimize: VlanApi: [VMVlanCreationParams], [VMVlanUpdationParamsData], [ManagementVlanUpdationParamsData] 限制 `VlanId` 范围为 0~4095

## release 日期 2022-03-22

v2.7.0 release

- optimize: VMApi: [delete_vm] 更新参数类型为 `VMDeleteParams`，添加 `effect` 允许删除相关的快照
- feature: [vm_usage] 枚举添加:
  - `BUNDLE_APPLICATION`
- feature: [ROLE_ACTION] 枚举添加:
  - `MANAGE_OBSERVABILITY_PACKAGE`
  - `MANAGE_OBSERVABILITY_SERVICE`
- feature: [software_edition] 枚举添加：
  - `ENTERPRISE_PLUS`
- feature: [upload_resource_type] 枚举添加:
  - `HOST_PLUGIN_PACKAGE`
- feature: [task_type] 枚举添加:
  - `HOST_PLUGIN`

## release 日期 2023-02-20

v2.6.0 release

- feature: [vm_usage] 枚举添加 `SKS_MANAGEMENT` 与 `REGISTRY`
- feature: [ROLEACTION] 枚举添加:
  - `MANAGE_SKS_SERVICE`
  - `MANAGE_SKS_LICENSE`
  - `CONFIGURE_SKS_SERVICE`
  - `CREATE_SKS_WORKLOAD_CLUSTER`
  - `DELETE_SKS_WORKLOAD_CLUSTER`
  - `UPDATE_SKS_WORKLOAD_CLUSTER`
  - `DOWNLOAD_SKS_WORKLOAD_CLUSTER_KUBECONFIG`

## release 日期 2023-01-03

v2.5.0 release

- bugfix: [IscsiTargetCommonParams]: 修复错误的 `BpsWrMaxSize` 为 `BpsWrMaxUnit`
- feature: [IscsiLunSnapshotApi]: [CreateIscsiLunSnapshot] 增加了同步创建 lun 快照的选项。
- feature: ClusterApi: [GetMetaLeader]: 增加了获取集群 meta leader 的 api
- optimize: 增加 header 定义，可以从返回值中获取对应的 XTowerRequestID
- optimize: [NestedHost]: 嵌套的主机类型额外返回 management_ip

## release 日期 2022-11-18

v2.4.0 release

- feature: [CloudTowerApplicationApi] 新增 CloudTowerApplicationApi;
  - [GetCloudTowerApplications] 获取应用;
  - [UploadCloudTowerApplicationPackage] 上传应用包;
  - [DeleteCloudTowerApplicationPackage] 删除应用包;
  - [DeployCloudTowerApplication] 部署应用;
  - [UpgradeCloudTowerApplication] 升级应用;
  - [UninstallCloudTowerApplication] 删除应用;
- feature: [CloudTowerApplicationPackageApi] 新增 CloudTowerApplicationPackageApi;
  - [GetCloudTowerApplicationPackages] 获取应用包.
- optimize: 存储容量, 内存容量相关的 api 参数都允许传入 `${field}_unit` 形式的参数来为输入参数设置单位，类型为 `ByteUnit`，默认为 `ByteUnit.B`;
- optimize: 带宽相关的 api 参数都允许传入 `${field}_unit` 形式的参数来为输入参数设置单位，类型为 `BpsUnit`，默认为 `BpsUnit.Bps`.

## release 日期 2022-09-05

v2.3.0 release

- feature:VMVolumeSnapshotApi: [GetVMVolumeSnapshots] 新增虚拟卷快照查询 api
- feature:VMVolumeSnapshotApi: [CreateVMVolumeSnapshot] 新增创建虚拟卷快照 api
- feature:VMVolumeSnapshotApi: [DeleteVMVolumeSnapshot] 新增删除虚拟卷快照 api
- feature:VMVolumeApi: [CloneVMVolume] 新增克隆虚拟卷 api
- feature:VMVolumeApi: [RebuildVMVolume] 新增通过虚拟卷快照重建虚拟卷 api
- feature:VMVolumeApi: [RollbackVMVolume] 新增回滚虚拟卷至指定虚拟卷快照 api
- feature:VMVolumeApi: [UpdateVMVolume] 新增编辑虚拟卷 api
- feature:UserApi: [GetMyInfo] 新增查询当前 client 对应用户 api
- feature:VersionApi: [GetApiInfo] 新增查询当前 api 版本 api
- feature:VMApi: 新增内容库镜像支持，[vm_cd_rom_params] 支持传入 `content_library_image_id` 来挂载内容库镜像
- optimize: 优化 `WaitTask`, `WaitTasks` 方法，允许外部传入的 `context` 以控制超时，并且在没有搜索到对应 `taskId` 的 task 情况下，尝试等待 task 被创建或直到超时

## release 日期 2022-08-12

v2.2.0 release

- feature:VMApi: [CreateVMFromContentLIbraryTemplate] 新增通过内容库模板创建虚拟机 api
- bugfix: 正确生成嵌套类型的数字类型

## release 日期 2022-07-08

v2.1.0 release

- feature:ClusterApi: [UpdateClusterNetworkSetting] 新增更新集群网络配置 api
- feature:ClusterApi: [UpdateClusterVirtualizationSetting] 新增更新集群虚拟化设置 api
- feature:ClusterApi: [UpdateClusterHaSetting] 新增更新集群高可用设置 api
- feature:ClusterApi: [UpdateClusterEnableIscsiSetting] 新增更新集群块功能启用设置 api
- feature:VMApi: [MigrateVMAcrossCluster] 新增跨集群迁移虚拟机 api
- feature:VMApi: [AbortMigrateVMAcrossCluster] 新增取消跨集群迁移 api
- feature:VMApi: [StopVMInCutoverMigration] 新增关闭源虚拟机 api
- feature:VMApi: [UpdateVMHostOptions] 新增更新虚拟机 guest os 设置 api，更新 dns, hostname 与 ntp server，需要虚拟机工具的支持。
- feature:VMApi: [ResetVMGuestOsPassword] 新增更新虚拟机 guest os 用户密码 api，需要虚拟机工具的支持。
- feature:VMApi: [UpdateVMOwner] 新增更新虚拟机拥有者 api
- feature:SecurityApi: [UpdatePasswordSecurity] 新增更新密码安全设置 api
- feature:SecurityApi: [UpdateAccessRestriction] 新增更新访问限制 api
- feature:SecurityApi: [UpdateSessionTimeout] 新增更新会话超时 api
- feature:VcenterAccountApi: [UpdateVcenterAccount] 新增更新 vcenter 账号 api
- feature:VcenterAccountApi: [CreateVcenterAccount] 新增添加 vcenter 账号 api
- feature:VsphereEsxiAccountApi: [UpdateVsphereEsxiAccount] 新增更新 vsphere esxi 账号 api
- feature:SvtImageApi: [UploadSvtImage] 新增上传虚拟机镜像 api 工具
- feature:TableReporterApi: [ExportCsv] 新增导出 CSV 报表 api
- feature:UploadTaskApi: [CancelUploadTask] 新增取消上传 api
- feature:LabelApi: [AddLabelsToResources],[removeLabelsFromResources] 新增内容库模板，内容库镜像，隔离策略，安全策略添加，删除标签

- bugfix:ContentLibraryImageApi,ElfImageApi: 修复了上传类 Api 无法正确执行的问题，并优化了上传类 Api 的执行逻辑，第一次上传时会上传第一个分片而非只是创建一个上传任务，详见[示例](/examples/upload/readme.md)

- optimize:VMTemplateApi: 优化了模板创建时根据传入的 cpu 参数和模板参数计算缺省值的逻辑
- optimize:ContentLibraryImageApi: 优化了分发的逻辑，不再同时上传一个镜像至多个集群，等待上传置单个集群后再分发。
- optimize: 添加 `utils.CheckSvtImageVersion` 工具方法检查 svt 镜像版本。

## release 日期：2022-06-14

v2.0.0+rev3 release

- bugfix: `NewWithUserConfig` 不再直接 panic 而会将 error 作为返回值返回。

## release 日期：2022-06-13

v2.0.0+rev2 release

- bugfix: [Cluster],[ClusterWhereInput]: 修复 usedMemoryBytes, usedCpuUsage 的数据类型 Long -> Double
- bugfix: [Host],[HostWhereInput]: 修复 usedMemoryBytes, usedCpuUsage 的数据类型 Long -> Double
- bugfix: [Datacenter],[DatacenterWhereInput]: 修复 usedMemoryBytes, usedCpuUsage 的数据类型 Long -> Double

## release 日期：2022-06-09

v2.0.0+rev1 release

- feature：添加 `NewWithUserConfig` 来创建一个拥有登陆信息的 `Client` 实例。

## release 日期：2022-05-20

v2.0.0 release

- feature: 开放内容库相关 Api [ContentLibraryImage], [ContentLibraryVMTemplate]
- feature: 新增 [Metric] 功能，用于查询资源性能指标
- bugfix: 移除 [Witness] 中的 [management_ip] 字段

## release 日期：2022-05-13

v1.10.0 release

- feature: UserApi: [CreateRootUser]: 创建初始 root 用户，创建用户需要传递 role id，而获取 role id 需要鉴权，提供一个接口在没有账号被建立的情况下创建一个默认的 root 用户，不需要鉴权；
- feature: VMApi: [ExpandVMDisk]: 提供扩容指定虚拟盘的能力；
- feature: VMApi: [EjectIsoFromVMCdRom]: 将 iso 从 vm cdrom 中卸载；
- feature: VMApi: [toggleVMCdRomDisable]: 禁用、启用 vm cdrom；
- feature: VMApi: [UpdateVMNicBasicInfo]: 更新一个虚拟机网卡的常用信息，包括 IP 地址，网关地址，子网掩码，需要虚拟机工具的支持；
- feature: VMApi: [UpdateVMNicAdvanceInfo]: 更新一个虚拟机网卡的非常用信息，包括接入的虚拟网络，sriov 网卡的直连网卡，mac 地址，是否启用，以及是否开启镜像模式；
- feature: VMApi: [UpdateVMAdvancedOptions]: 更新虚拟机的高级信息，包括 CPU 兼容性，虚拟机时钟(UTC、LocalTime)，是否开启为 windows 优化，以及虚拟机显卡型号；
- optimize: 提供了更加完整的错误返回信息，包含了错误发生的位置，graphql operationName，并透传 tower server 状态码。
- optimize: 缩短了部分变量类型名称。
- bugfix: VMApi: [createVMFromTemplate], []: 修复了无法在使用模板创建虚拟机时挂载已存在的虚拟机和 cd-rom 的问题，修正 MountDiskParams 中的 index 为可选项。
- bugfix: 使用了更加准确的数字类型，将以 Byte，HZ 为单位的数字类型从 Double 转为了 Long；
- bugfix: 将分类错误的 endpoint 移至正确的 api 下
  - [GetIscsiConnections]: IscsiApi -> IscsiConnectionApi
  - [CreateNvmfSubsystem]: DefaultApi -> NvmfSubsystemApi
  - [DeleteNvmfSubsystem]: DefaultApi -> NvmfSubsystemApi
  - [UpdateNvmfSubsystem]: DefaultApi -> NvmfSubsystemApi
- remove: 屏蔽了不在 1.\*.0 中开放的功能：
  - 备份：
    - [BackupLicense]
    - [BackupPackage]
    - [BackupPlan]
    - [BackupPlanExecution]
    - [BackupRestoreExecution]
    - [BackupRestorePoint]
    - [BackupService]
    - [BackupStoreRepository]
    - [BackupTargetExecution]
  - 内容库：
    - [ContentLibraryImage]
    - [ContentLibraryVMTemplate]
