# RELEASE NOTE

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
- feature: [OvfApi], [VmExportFileApi], [VmApi] 支持虚拟机的导入与导出
- feature: [VlanApi] 支持 trunk vlan 的创建与编辑
- feature: [UserApi] [Login] 支持使用 authn_id 登陆，旧 LDAP 登陆方式被废弃
- optimize: 为 [Host], [Nic], [UsbDevice], [VmVolume], [VmVolumeSnapshot] 添加了 `EntityAsyncStatus` 已判断资源目前的状态

## release 日期 2022-05-04

v2.8.0 release

- optimize: VlanApi: [VmVlanCreationParams], [VmVlanUpdationParamsData], [ManagementVlanUpdationParamsData] 限制 `VlanId` 范围为 0~4095

## release 日期 2022-03-22

v2.7.0 release

- optimize: VmApi: [delete_vm] 更新参数类型为 `VmDeleteParams`，添加 `effect` 允许删除相关的快照
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
- feature: IscsiLunSnapshotApi: [create_iscsi_lun_snapshot] 增加了同步创建 lun 快照的选项。
- feature: ClusterApi: [get_meta_leader]: 增加了获取集群 meta leader 的 api
- optimize: 增加 header 定义，可以从返回值中获取对应的 XTowerRequestID
- optimize: [NestedHost]: 嵌套的主机类型额外返回 management_ip

## release 日期 2022-11-18

v2.4.0 release

- feature:CloudTowerApplicationApi: [cloud_tower_application_api] 新增 CloudTowerApplicationApi;
  - [get_cloud_tower_applications] 获取应用;
  - [upload_cloud_tower_application_package] 上传应用包;
  - [delete_cloud_tower_application_package] 删除应用包;
  - [deploy_cloud_tower_application] 部署应用;
  - [upgrade_cloud_tower_application] 升级应用;
  - [uninstall_cloud_tower_application] 删除应用;
- feature:CloudTowerApplicationPackageApi: [cloud_tower_application_package_api] 新增 CloudTowerApplicationPackageApi;
  - [get_cloud_tower_application_packages] 获取应用包.
- optimize: 存储容量, 内存容量相关的 api 参数都允许传入 `${field}_unit` 形式的参数来为输入参数设置单位，类型为 `ByteUnit`，默认为 `ByteUnit.B`;
- optimize: 带宽相关的 api 参数都允许传入 `${field}_unit` 形式的参数来为输入参数设置单位，类型为 `BpsUnit`，默认为 `BpsUnit.Bps`.

## release 日期 2022-09-05

v2.3.0 release

- feature:VmVolumeSnapshotApi: [get_vm_volume_snapshots] 新增虚拟卷快照查询 api
- feature:VmVolumeSnapshotApi: [create_vm_volume_snapshot] 新增创建虚拟卷快照 api
- feature:VmVolumeSnapshotApi: [delete_vm_volume_snapshot] 新增删除虚拟卷快照 api
- feature:VmVolumeApi: [clone_vm_volume] 新增克隆虚拟卷 api
- feature:VmVolumeApi: [rebuild_vm_volume] 新增通过虚拟卷快照重建虚拟卷 api
- feature:VmVolumeApi: [rollback_vm_volume] 新增回滚虚拟卷至指定虚拟卷快照 api
- feature:VmVolumeApi: [update_vm_volume] 新增编辑虚拟卷 api
- feature:UserApi: [get_my_info] 新增查询当前 client 对应用户 api
- feature:VersionApi: [get_api_info] 新增查询当前 api 版本 api
- feature:VmApi: 新增内容库镜像支持，[vm_cd_rom_params] 支持传入 `content_library_image_id` 来挂载内容库镜像
- optimize: 优化 `WaitTask`, `WaitTasks` 方法，允许外部传入的 `context` 以控制超时，并且在没有搜索到对应 `taskId` 的 task 情况下，尝试等待 task 被创建或直到超时

## release 日期 2022-08-12

v2.2.0 release

- feature:VmApi: [create_vm_from_content_library_template] 新增通过内容库模板创建虚拟机 api
- bugfix: 正确生成嵌套类型的数字类型

## release 日期 2022-07-08

v2.1.0 release

- feature:ClusterApi: [update_cluster_network_setting] 新增更新集群网络配置 api
- feature:ClusterApi: [update_cluster_virtualization_setting] 新增更新集群虚拟化设置 api
- feature:ClusterApi: [update_cluster_ha_setting] 新增更新集群高可用设置 api
- feature:ClusterApi: [update_cluster_enable_iscsi_setting] 新增更新集群块功能启用设置 api
- feature:VmApi: [migratevm_across_cluster] 新增跨集群迁移虚拟机 api
- feature:VmApi: [abort_migrate_vm_across_cluster] 新增取消跨集群迁移 api
- feature:VmApi: [stop_vm_in_cutover_migration] 新增关闭源虚拟机 api
- feature:VmApi: [update_vmHostOptions] 新增更新虚拟机 guest os 设置 api，更新 dns, hostname 与 ntp server，需要虚拟机工具的支持。
- feature:VmApi: [reset_vmG_guest_ss_password] 新增更新虚拟机 guest os 用户密码 api，需要虚拟机工具的支持。
- feature:VmApi: [update_vm_owner] 新增更新虚拟机拥有者 api
- feature:SecurityApi: [update_password_security] 新增更新密码安全设置 api
- feature:SecurityApi: [update_access_restriction] 新增更新访问限制 api
- feature:SecurityApi: [update_session_timeout] 新增更新会话超时 api
- feature:VcenterAccountApi: [update_vcenter_account] 新增更新 vcenter 账号 api
- feature:VcenterAccountApi: [create_vcenter_account] 新增添加 vcenter 账号 api
- feature:VsphereEsxiAccountApi: [update_vsphere_esxi_account] 新增更新 vsphere esxi 账号 api
- feature:SvtImageApi: [upload_svt_image] 新增上传虚拟机镜像 api 工具
- feature:TableReporterApi: [export_csv] 新增导出 CSV 报表 api
- feature:UploadTaskApi: [cancel_upload_task] 新增取消上传 api
- feature:LabelApi: [add_labels_to_resources],[remove_labels_from_resources] 新增内容库模板，内容库镜像，隔离策略，安全策略添加，删除标签

- bugfix:ContentLibraryImageApi,ElfImageApi: 修复了上传类 Api 无法正确执行的问题，并优化了上传类 Api 的执行逻辑，第一次上传时会上传第一个分片而非只是创建一个上传任务，详见[示例](/examples/upload/readme.md)

- optimize:VmTemplateApi: 优化了模板创建时根据传入的 cpu 参数和模板参数计算缺省值的逻辑
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

- feature: UserApi: [create_root_user]: 创建初始 root 用户，创建用户需要传递 role id，而获取 role id 需要鉴权，提供一个接口在没有账号被建立的情况下创建一个默认的 root 用户，不需要鉴权；
- feature: VmApi: [expand_vm_disk]: 提供扩容指定虚拟盘的能力；
- feature: VmApi: [eject_iso_from_vm_cd_rom]: 将 iso 从 vm cdrom 中卸载；
- feature: VmApi: [toggle_vm_cd_rom_disable]: 禁用、启用 vm cdrom；
- feature: VmApi: [update_vm_nic_basic_info]: 更新一个虚拟机网卡的常用信息，包括 IP 地址，网关地址，子网掩码，需要虚拟机工具的支持；
- feature: VmApi: [update_vm_nic_advance_info]: 更新一个虚拟机网卡的非常用信息，包括接入的虚拟网络，sriov 网卡的直连网卡，mac 地址，是否启用，以及是否开启镜像模式；
- feature: VmApi: [update_vm_advanced_options]: 更新虚拟机的高级信息，包括 CPU 兼容性，虚拟机时钟(UTC、LocalTime)，是否开启为 windows 优化，以及虚拟机显卡型号；
- optimize: 提供了更加完整的错误返回信息，包含了错误发生的位置，graphql operationName，并透传 tower server 状态码。
- optimize: 缩短了部分变量类型名称。
- bugfix: VmApi: [create_vm_from_template], []: 修复了无法在使用模板创建虚拟机时挂载已存在的虚拟机和 cd-rom 的问题，修正 MountDiskParams 中的 index 为可选项。
- bugfix: 使用了更加准确的数字类型，将以 Byte，HZ 为单位的数字类型从 Double 转为了 Long；
- bugfix: 将分类错误的 endpoint 移至正确的 api 下
  - [get_iscsi_connections]: IscsiApi -> IscsiConnectionApi
  - [create_nvmf_subsystem]: DefaultApi -> NvmfSubsystemApi
  - [delete_nvmf_subsystem]: DefaultApi -> NvmfSubsystemApi
  - [update_nvmf_subsystem]: DefaultApi -> NvmfSubsystemApi
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
