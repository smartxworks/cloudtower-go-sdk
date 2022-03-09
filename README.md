# Cloudtower Go SDK

Golang 环境下的 Cloudtower SDK，适用于 golang 1.16 及以上版本

- [源码地址](https://github.com/Sczlog/cloudtower-go-sdk)
- [下载地址](https://github.com/Sczlog/cloudtower-go-sdk/releases)
- [通用指南](https://cloudtower-api-doc.vercel.app)

## 安装

```shell
go get https://github.com/Sczlog/cloudtower-go-sdk
```

## 使用

> 样例中使用了两个工具库 [pointy](https://github.com/openlyinc/pointy) 与 [go-funk](https://github.com/thoas/go-funk)，pointy 用于快速创建一个原始类型的指针，go-funk 则提供了一些工具方法，例如 `Map`、`Filter`、`Reduce` 等。

### 创建实例

#### 创建 `ApiClient` 实例

```go
import (
	apiclient "github.com/Sczlog/cloudtower-go-sdk/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)
transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
client := apiclient.New(transport, strfmt.Default)
```

### 发送请求

#### 引入对应的 `client` 包

> 根据不同用途的操作引入创建相关的 `client` 包

```go
import (
  	vm "github.com/Sczlog/cloudtower-go-sdk/client/vm"
)
```

#### 鉴权

```go
import (
  User "github.com/Sczlog/cloudtower-go-sdk/client/user"
)
loginParams := User.NewLoginParams()
loginParams.RequestBody = &models.LoginInput{
	Username: pointy.String("username"),
	Password: pointy.String("password"),
	Source:   models.NewUserSource(models.UserSourceLOCAL),
}
logRes, err := client.User.Login(loginParams)
if err != nil {
	return err
}
transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", *logRes.Payload.Data.Token)
```

#### 获取资源

```go
getVmParams := vm.NewGetVmsParams();
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
}
vmsRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
vms := vmsRes.Payload
```

#### 更新资源

> 资源更新会产生相关的异步任务，当异步任务结束时，代表资源操作完成且数据已更新。

```go
target_vm := vmsRes.Payload[0]
vmStartParams := vm.NewStartVMParams()
vmStartParams.RequestBody = &models.VMStartParams{
	Where: &models.VMWhereInput{
		ID: target_vm.ID,
	},
}
startRes, err := client.VM.StartVM(vmStartParams)
if err != nil {
	return err
}
```

> 可以通过提供的工具方法 `WaitTask` 同步等待异步任务结束，如果任务失败或超时，都会返回一个异常，轮询间隔 5s，超时时间为 300s。
>
> - 方法参数说明
>
> | 参数名 | 类型                | 是否必须 | 说明                     |
> | ------ | ------------------- | -------- | ------------------------ |
> | client | \*client.Cloudtower | 是       | 查询所使用的 client 实例 |
> | id     | string              | 是       | 需查询的 task 的 id      |

```go
task := *startRes.Payload[0].TaskID
err = utils.WaitTask(client, task)
if err != nil {
	return err
}
```

> 如果是复数任务则可以通过 `WaitTasks`，接受复数个 task id，其余与 `WaitTask` 相同。
>
> - 方法参数说明
>
> | 参数名 | 类型                | 是否必须 | 说明                     |
> | ------ | ------------------- | -------- | ------------------------ |
> | client | \*client.Cloudtower | 是       | 查询所使用的 client 实例 |
> | ids    | []string            | 是       | 需查询的 task 的 id 列表 |

```go
tasks := funk.Map(startRes.Payload, func(tvm *models.WithTaskVM) string {
	return *tvm.TaskID
}).([]string)
err = utils.WaitTasks(client, tasks)
if err != nil {
	return err
}
```

#### 其他

##### 设置返回信息的语言

> 可以设置请求 params 中的 `ContentLanguage` 项设置返回值的语言，可选值为 `["en-US", "zh-CN"]`，默认值为 `en-US`，不在可选值范围内的语言会返回一个 HTTP 400 错误

```go
getTaskDefaultParams := task.NewGetTasksParams()
getTaskDefaultParams.RequestBody = &models.GetTasksRequestBody{
	First: pointy.Int32(10),
}
// 此时得到的 alerts 中的 message, solution, cause, impact 将被转换为英文描述。
taskDefaultRes, err := client.Task.GetTasks(getTaskDefaultParams)

getTaskZhParams := task.NewGetTasksParams()
getTaskZhParams.RequestBody = &models.GetTasksRequestBody{
	First: pointy.Int32(10),
}
// 此时得到的 alerts 中的 message, solution, cause, impact 将被转换为中文描述。
getTaskZhParams.ContentLanguage = pointy.String("zh-CN")
taskZhRes, err := client.Task.GetTasks(getTaskZhParams)
```

## 操作示例

### 获取虚拟机

#### 获取所有虚拟机

```go
getAllVmsParams := vm.NewGetVmsParams()
getAllVmsParams.RequestBody = &models.GetVmsRequestBody{}
vmsRes, err := client.VM.GetVms(getAllVmsParams)
if err != nil {
	return err
}
vms := vmsRes.Payload
```

#### 分页获取虚拟机

```go
getVmsFrom51to100Params := vm.NewGetVmsParams()
getVmsFrom51to100Params.RequestBody = &models.GetVmsRequestBody{
	First: pointy.Int32(50),
	Skip:  pointy.Int32(50),
}
vmsRes, err := client.VM.GetVms(getVmsFrom51to100Params)
if err != nil {
	return err
}
vms := vmsRes.Payload
```

#### 获取所有已开机虚拟机

```go
getAllRunningVmsParams := vm.NewGetVmsParams()
getAllRunningVmsParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		Status: models.VMStatusRUNNING.Pointer(),
	},
}
vmsRes, err := client.VM.GetVms(getAllRunningVmsParams)
if err != nil {
	return err
}
vms := vmsRes.Payload
```

#### 获取名称或描述中包含特定字符串的虚拟机

```go
getAllVmNameMatchStrParams := vm.NewGetVmsParams()
getAllVmNameMatchStrParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		NameContains: pointy.String("STR"),
	},
}
vmsRes, err := client.VM.GetVms(getAllVmNameMatchStrParams)
if err != nil {
	return err
}
```

#### 获取所有 vcpu > n 的虚拟机

```go
getAllVmCoreGtNParams := vm.NewGetVmsParams()
getAllVmCoreGtNParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		VcpuGt: pointy.Int32(n),
	},
}
vmsRes, err := client.VM.GetVms(getAllVmCoreGtNParams)
if err != nil {
	return err
}
vms := vmsRes.Payload
```

### 从模版创建虚拟机

#### 仅指定 id

```go
createVmFromTemplateParams := vm.NewCreateVMFromTemplateParams()
createVmFromTemplateParams.RequestBody = []*models.VMCreateVMFromTemplateParams{
	{
		TemplateID: pointy.String("template_id"),
		ClusterID:  pointy.String("target_cluster_id"),
		Name:       pointy.String("vm_name"),
		IsFullCopy: pointy.Bool(false),
	},
}
createRes, err := client.VM.CreateVMFromTemplate(createVmFromTemplateParams)
if err != nil {
	return err
}
withTaskVm := createRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
createdVm := queryRes.Payload[0]
```

#### 配置与模板不同的虚拟盘参数

```go
createVmFromTemplateParams := vm.NewCreateVMFromTemplateParams()
createVmFromTemplateParams.RequestBody = []*models.VMCreateVMFromTemplateParams{
	{
		TemplateID: pointy.String("template_id"),
		ClusterID:  pointy.String("target_cluster_id"),
		Name:       pointy.String("vm_name"),
		DiskOperate: &models.VMCreateVMFromTemplateParamsDiskOperate{
			RemoveDisks: &models.VMCreateVMFromTemplateParamsDiskOperateRemoveDisks{
				DiskIndex: []int32{0, 1},
			},
			ModifyDisks: []*models.VMCreateVMFromTemplateParamsDiskOperateModifyDisksItems0{
				{
					DiskIndex:  pointy.Int32(2),
					VMVolumeID: pointy.String("target_volume_id"),
				},
			},
			NewDisks: &models.VMDiskParams{
				MountCdRoms: []*models.VMCdRomParams{
					{
						Boot:       pointy.Int32(0),
						ElfImageID: pointy.String("target_image_id"),
					},
				},
				MountDisks: []*models.MountDisksParams{
					{
						Bus:        models.BusVIRTIO.Pointer(),
						Boot:       pointy.Int32(1),
						VMVolumeID: pointy.String("target_volume_id"),
					},
				},
				MountNewCreateDisks: []*models.MountNewCreateDisksParams{
					{
						VMVolume: &models.MountNewCreateDisksParamsVMVolume{
							ElfStoragePolicy: models.VMVolumeElfStoragePolicyTypeREPLICA2THINPROVISION.Pointer(),
							Size:             pointy.Float64(4 * 1024 * 1024 * 1024),
							Name:             pointy.String("disk_name"),
						},
						Boot: pointy.Int32(2),
						Bus:  models.BusVIRTIO.Pointer(),
					},
				},
			},
		},
		IsFullCopy: pointy.Bool(false),
	},
}
createRes, err := client.VM.CreateVMFromTemplate(createVmFromTemplateParams)
if err != nil {
	return err
}
withTaskVm := createRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
createdVm := queryRes.Payload[0]
```

#### 配置与模版不同的网卡参数

```go
createVmFromTemplateParams := vm.NewCreateVMFromTemplateParams()
createVmFromTemplateParams.RequestBody = []*models.VMCreateVMFromTemplateParams{
	{
		TemplateID: pointy.String("template_id"),
		ClusterID:  pointy.String("target_cluster_id"),
		Name:       pointy.String("vm_name"),
		VMNics: []*models.VMNicParams{
			{
				ConnectVlanID: pointy.String("alternate_vlan_id"),
				Enabled:       pointy.Bool(true),
				Model:         models.VMNicModelSRIOV.Pointer(),
			},
		},
		IsFullCopy: pointy.Bool(false),
	},
}
createRes, err := client.VM.CreateVMFromTemplate(createVmFromTemplateParams)
if err != nil {
	return err
}
withTaskVm := createRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
createdVm := queryRes.Payload[0]
```

### 创建空白虚拟机

#### 简单创建

```go
createParams := vm.NewCreateVMParams()
createParams.RequestBody = []*models.VMCreationParams{
	{
		ClusterID:  pointy.String("target_cluster_id"),
		Name:       pointy.String("test_vm_name"),
		Ha:         pointy.Bool(true),
		CPUCores:   pointy.Int32(4),
		CPUSockets: pointy.Int32(2),
		Memory:     pointy.Float64(8 * 1024 * 1024 * 1024),
		Vcpu:       pointy.Int32(4 * 2),
		Status:     models.VMStatusSTOPPED.Pointer(),
		Firmware:   models.VMFirmwareBIOS.Pointer(),
		VMNics: []*models.VMNicParams{
			{ConnectVlanID: pointy.String("target_vlan_id")},
		},
    VMDisks: &models.VMDiskParams{},
	},
}
createRes, err := client.VM.CreateVM(createParams)
if err != nil {
	return err
}
withTaskVm := createRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
createdVm := queryRes.Payload[0]
```

#### 创建时配置虚拟盘

##### CD-ROM 加载 ISO

```go
createParams := vm.NewCreateVMParams()
createParams.RequestBody = []*models.VMCreationParams{
	{
		ClusterID:  pointy.String("target_cluster_id"),
		Name:       pointy.String("test_vm_name"),
		Ha:         pointy.Bool(true),
		CPUCores:   pointy.Int32(4),
		CPUSockets: pointy.Int32(2),
		Memory:     pointy.Float64(8 * 1024 * 1024 * 1024),
		Vcpu:       pointy.Int32(4 * 2),
		Status:     models.VMStatusSTOPPED.Pointer(),
		Firmware:   models.VMFirmwareBIOS.Pointer(),
		VMNics: []*models.VMNicParams{
			{ConnectVlanID: pointy.String("target_vlan_id")},
		},
		VMDisks: &models.VMDiskParams{
			MountCdRoms: []*models.VMCdRomParams{
				{
					Boot:       pointy.Int32(1),
					ElfImageID: pointy.String("target_elf_image_id"),
				},
			},
		},
	},
}
createRes, err := client.VM.CreateVM(createParams)
if err != nil {
	return err
}
withTaskVm := createRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
  panic(err.Error())
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
	createdVm := queryRes.Payload[0]
```

##### 挂载虚拟卷为虚拟盘

```go
createParams := vm.NewCreateVMParams()
createParams.RequestBody = []*models.VMCreationParams{
	{
		ClusterID:  pointy.String("target_cluster_id"),
		Name:       pointy.String("test_vm_name"),
		Ha:         pointy.Bool(true),
		CPUCores:   pointy.Int32(4),
		CPUSockets: pointy.Int32(2),
		Memory:     pointy.Float64(8 * 1024 * 1024 * 1024),
		Vcpu:       pointy.Int32(4 * 2),
		Status:     models.VMStatusSTOPPED.Pointer(),
		Firmware:   models.VMFirmwareBIOS.Pointer(),
		VMNics: []*models.VMNicParams{
			{ConnectVlanID: pointy.String("target_vlan_id")},
		},
		VMDisks: &models.VMDiskParams{
			MountDisks: []*models.MountDisksParams{
				{
					Boot:       pointy.Int32(0),
					Bus:        models.BusVIRTIO.Pointer(),
					VMVolumeID: pointy.String("target_volume_id"),
				},
			},
		},
	},
}
createRes, err := client.VM.CreateVM(createParams)
if err != nil {
	return err
}
withTaskVm := createRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
createdVm := queryRes.Payload[0]
```

##### 挂载新增虚拟盘

```go
createParams := vm.NewCreateVMParams()
createParams.RequestBody = []*models.VMCreationParams{
	{
		ClusterID:  pointy.String("target_cluster_id"),
		Name:       pointy.String("test_vm_name"),
		Ha:         pointy.Bool(true),
		CPUCores:   pointy.Int32(4),
		CPUSockets: pointy.Int32(2),
		Memory:     pointy.Float64(8 * 1024 * 1024 * 1024),
		Vcpu:       pointy.Int32(4 * 2),
		Status:     models.VMStatusSTOPPED.Pointer(),
		Firmware:   models.VMFirmwareBIOS.Pointer(),
		VMNics: []*models.VMNicParams{
			{ConnectVlanID: pointy.String("target_vlan_id")},
		},
		VMDisks: &models.VMDiskParams{
			MountNewCreateDisks: []*models.MountNewCreateDisksParams{
				{
					Boot: pointy.Int32(0),
					Bus:  models.BusIDE.Pointer(),
					VMVolume: &models.MountNewCreateDisksParamsVMVolume{
						Size:             pointy.Float64(20 * 1024 * 1024 * 1024),
						ElfStoragePolicy: models.VMVolumeElfStoragePolicyTypeREPLICA2THINPROVISION.Pointer(),
						Name:             pointy.String("new_vm_disk_name"),
					},
				},
			},
		},
	},
}
createRes, err := client.VM.CreateVM(createParams)
if err != nil {
	return err
}
withTaskVm := createRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
createdVm := queryRes.Payload[0]
```

#### 创建时配置虚拟网卡

```go
createParams := vm.NewCreateVMParams()
createParams.RequestBody = []*models.VMCreationParams{
	{
		ClusterID:  pointy.String("target_cluster_id"),
		Name:       pointy.String("test_vm_name"),
		Ha:         pointy.Bool(true),
		CPUCores:   pointy.Int32(4),
		CPUSockets: pointy.Int32(2),
		Memory:     pointy.Float64(8 * 1024 * 1024 * 1024),
		Vcpu:       pointy.Int32(4 * 2),
		Status:     models.VMStatusSTOPPED.Pointer(),
		Firmware:   models.VMFirmwareBIOS.Pointer(),
		VMNics: []*models.VMNicParams{
			{
				ConnectVlanID: pointy.String("target_vlan_id"),
				Mirror:        pointy.Bool(true),
				Model:         models.VMNicModelSRIOV.Pointer(),
			},
		},
    VMDisks: &models.VMDiskParams{},
	},
}
createRes, err := client.VM.CreateVM(createParams)
if err != nil {
	return err
}
withTaskVm := createRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
createdVm := queryRes.Payload[0]
```

### 编辑虚拟机

#### 编辑基本信息

```go
updateParams := vm.NewUpdateVMParams()
updateParams.RequestBody = &models.VMUpdateParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
	Data: &models.VMUpdateParamsData{
		Name:        pointy.String("updated_name"),
		Description: pointy.String("updated description"),
		Ha:          pointy.Bool(true),
		CPUCores:    pointy.Int32(2),
		CPUSockets:  pointy.Int32(8),
		Vcpu:        pointy.Int32(2 * 8),
		Memory:      pointy.Float64(16 * 1024 * 1024 * 1024),
	},
}
updateRes, err := client.VM.UpdateVM(updateParams)
if err != nil {
	return err
}
withTaskVm := updateRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
updatedVm := queryRes.Payload[0]
```

#### CD-ROM 编辑

##### 添加 CD-ROM

```go
updateParams := vm.NewAddVMCdRomParams()
updateParams.RequestBody = &models.VMAddCdRomParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
	Data: &models.VMAddCdRomParamsData{
		VMCdRoms: []*models.VMCdRomParams{
			{
				ElfImageID: pointy.String("target_elf_image_id"),
				Boot:       pointy.Int32(0),
			},
		},
	},
}
updateRes, err := client.VM.AddVMCdRom(updateParams)
if err != nil {
	return err
}
withTaskVm := updateRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
updatedVm := queryRes.Payload[0]
```

##### 删除 CD-ROM

```go
updateParams := vm.NewRemoveVMCdRomParams()
updateParams.RequestBody = &models.VMRemoveCdRomParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
	Data: &models.VMRemoveCdRomParamsData{
		CdRomIds: []string{"id1", "id2"},
	},
}
updateRes, err := client.VM.RemoveVMCdRom(updateParams)
if err != nil {
	return err
}
withTaskVm := updateRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
updatedVm := queryRes.Payload[0]
```

#### 虚拟卷操作

##### 添加新虚拟卷

```go
updateParams := vm.NewAddVMDiskParams()
updateParams.RequestBody = &models.VMAddDiskParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
	Data: &models.VMAddDiskParamsData{
		VMDisks: &models.VMAddDiskParamsDataVMDisks{
			MountNewCreateDisks: []*models.MountNewCreateDisksParams{
				{
					VMVolume: &models.MountNewCreateDisksParamsVMVolume{
						ElfStoragePolicy: models.VMVolumeElfStoragePolicyTypeREPLICA2THINPROVISION.Pointer(),
						Size:             pointy.Float64(20 * 1024 * 1024 * 1024),
						Name:             pointy.String("new_disk_name"),
					},
					Boot: pointy.Int32(1),
					Bus:  models.BusVIRTIO.Pointer(),
				},
			},
		},
	},
}
updateRes, err := client.VM.AddVMDisk(updateParams)
if err != nil {
	return err
}
withTaskVm := updateRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
updatedVm := queryRes.Payload[0]
```

##### 挂载已存在虚拟卷为虚拟盘

```go
updateParams := vm.NewAddVMDiskParams()
updateParams.RequestBody = &models.VMAddDiskParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
	Data: &models.VMAddDiskParamsData{
		VMDisks: &models.VMAddDiskParamsDataVMDisks{
			MountDisks: []*models.MountDisksParams{
				{
					VMVolumeID: pointy.String("target_vm_volume_id"),
					Boot:       pointy.Int32(1),
					Bus:        models.BusVIRTIO.Pointer(),
				},
			},
		},
	},
}
updateRes, err := client.VM.AddVMDisk(updateParams)
if err != nil {
	return err
}
withTaskVm := updateRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
updatedVm := queryRes.Payload[0]
```

##### 卸载虚拟盘

```go
updateParams := vm.NewRemoveVMDiskParams()
updateParams.RequestBody = &models.VMRemoveDiskParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
	Data: &models.VMRemoveDiskParamsData{
		DiskIds: []string{"id1", "id2"},
	},
}
updateRes, err := client.VM.RemoveVMDisk(updateParams)
if err != nil {
	return err
}
withTaskVm := updateRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
updatedVm := queryRes.Payload[0]
```

#### 网卡操作

##### 添加网卡

```go
updateParams := vm.NewAddVMNicParams()
updateParams.RequestBody = &models.VMAddNicParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
	Data: &models.VMAddNicParamsData{
		VMNics: []*models.VMNicParams{
			{
				ConnectVlanID: pointy.String("target_vlan_id"),
				Enabled:       pointy.Bool(true),
				Model:         models.VMNicModelVIRTIO.Pointer(),
				IPAddress:     pointy.String("ip_address"),
				Gateway:       pointy.String("gateway"),
				SubnetMask:    pointy.String("subnetmask"),
			},
			{
				ConnectVlanID: pointy.String("target_vlan_id_2"),
				Enabled:       pointy.Bool(true),
				Mirror:        pointy.Bool(true),
				Model:         models.VMNicModelVIRTIO.Pointer(),
				IPAddress:     pointy.String("ip_address_2"),
				Gateway:       pointy.String("gateway"),
				SubnetMask:    pointy.String("subnetmask"),
			},
		},
	},
}
updateRes, err := client.VM.AddVMNic(updateParams)
if err != nil {
	return err
}
withTaskVm := updateRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
updatedVm := queryRes.Payload[0]
```

##### 编辑网卡

```go
updateParams := vm.NewUpdateVMNicParams()
updateParams.RequestBody = &models.VMUpdateNicParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
	Data: &models.VMUpdateNicParamsData{
		NicIndex:   pointy.Int32(0),
		Enabled:    pointy.Bool(false),
		Mirror:     pointy.Bool(false),
		MacAddress: pointy.String("new_mac_address"),
		IPAddress:  pointy.String("new_ip"),
		Gateway:    pointy.String("new_gateway"),
		SubnetMask: pointy.String("new_subnet_mask"),
	},
}
updateRes, err := client.VM.UpdateVMNic(updateParams)
if err != nil {
	return err
}
withTaskVm := updateRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
updatedVm := queryRes.Payload[0]
```

##### 移除网卡

```go
updateParams := vm.NewRemoveVMNicParams()
updateParams.RequestBody = &models.VMRemoveNicParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
	Data: &models.VMRemoveNicParamsData{
		NicIndex: []int32{0, 1},
	},
}
updateRes, err := client.VM.RemoveVMNic(updateParams)
if err != nil {
	return err
}
withTaskVm := updateRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
updatedVm := queryRes.Payload[0]
```

#### 虚拟机迁移

##### 迁移至指定主机

```go
migrateParams := vm.NewMigRateVMParams()
migrateParams.RequestBody = &models.VMMigrateParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
	Data: &models.VMMigrateParamsData{
		HostID: pointy.String("host_id"),
	},
}
migrateRes, err := client.VM.MigRateVM(migrateParams)
if err != nil {
	return err
}
withTaskVm := migrateRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
migratedVm := queryRes.Payload[0]
```

##### 自动调度到合适的主机

```go
migrateParams := vm.NewMigRateVMParams()
migrateParams.RequestBody = &models.VMMigrateParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
}
migrateRes, err := client.VM.MigRateVM(migrateParams)
if err != nil {
	return err
}
withTaskVm := migrateRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
migratedVm := queryRes.Payload[0]
```

### 虚拟机电源操作

#### 虚拟机开机:

##### 指定虚拟机开机，自动调度到合适的虚拟机

```go
startParams := vm.NewStartVMParams()
startParams.RequestBody = &models.VMStartParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
}
startRes, err := client.VM.StartVM(startParams)
if err != nil {
	return err
}
withTaskVm := startRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
startedVm := queryRes.Payload[0]
```

##### 批量虚拟机开机，自动调度到合适的虚拟机

```go
startParams := vm.NewStartVMParams()
startParams.RequestBody = &models.VMStartParams{
	Where: &models.VMWhereInput{
		IDIn: []string{"vm_id_1", "vm_id_2"},
	},
}
startRes, err := client.VM.StartVM(startParams)
if err != nil {
	return err
}
withTaskVms := startRes.Payload
taskIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
	return *vm.TaskID
}).([]string)
if !valid {
	return fmt.Errorf("failed to parse task ids")
}
vmIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
	return *vm.Data.ID
}).([]string)
if !valid {
	return fmt.Errorf("failed to parse vm ids")
}
err = utils.WaitTasks(client, taskIds)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		IDIn: vmIds,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
startedVms := queryRes.Payload
```

##### 开机至指定主机

```go
startParams := vm.NewStartVMParams()
startParams.RequestBody = &models.VMStartParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
	Data: &models.VMStartParamsData{
		HostID: pointy.String("host_id"),
	},
}
startRes, err := client.VM.StartVM(startParams)
if err != nil {
	return err
}
withTaskVm := startRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
startedVm := queryRes.Payload[0]
```

#### 虚拟机关机

##### 指定虚拟机关机

```go
shutdownParams := vm.NewShutDownVMParams()
shutdownParams.RequestBody = &models.VMOperateParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
}
shutdownRes, err := client.VM.ShutDownVM(shutdownParams)
if err != nil {
	return err
}
withTaskVm := shutdownRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
shutdownVm := queryRes.Payload[0]
```

##### 批量虚拟机关机

```go
shutdownParams := vm.NewShutDownVMParams()
shutdownParams.RequestBody = &models.VMOperateParams{
	Where: &models.VMWhereInput{
		IDIn: []string{"vm_id_1", "vm_id_2"},
	},
}
startRes, err := client.VM.ShutDownVM(shutdownParams)
if err != nil {
	return err
}
withTaskVms := startRes.Payload
taskIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
	return *vm.TaskID
}).([]string)
if !valid {
	return fmt.Errorf("failed to parse task ids")
}
vmIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
	return *vm.Data.ID
}).([]string)
if !valid {
	return fmt.Errorf("failed to parse vm ids")
}
err = utils.WaitTasks(client, taskIds)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		IDIn: vmIds,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
shutdownVms := queryRes.Payload
```

##### 强制关机指定虚拟机

```go
shutdownParams := vm.NewForceShutDownVMParams()
shutdownParams.RequestBody = &models.VMOperateParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
}
shutdownRes, err := client.VM.ForceShutDownVM(shutdownParams)
if err != nil {
	return err
}
withTaskVm := shutdownRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
shutdownVm := queryRes.Payload[0]
```

##### 强制关机批量虚拟机

```go
shutdownParams := vm.NewForceShutDownVMParams()
shutdownParams.RequestBody = &models.VMOperateParams{
	Where: &models.VMWhereInput{
		IDIn: []string{"vm_id_1", "vm_id_2"},
	},
}
startRes, err := client.VM.ForceShutDownVM(shutdownParams)
if err != nil {
	return err
}
withTaskVms := startRes.Payload
taskIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
	return *vm.TaskID
}).([]string)
if !valid {
	return fmt.Errorf("failed to parse task ids")
}
vmIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
	return *vm.Data.ID
}).([]string)
if !valid {
	return fmt.Errorf("failed to parse vm ids")
}
err = utils.WaitTasks(client, taskIds)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		IDIn: vmIds,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
shutdownVms := queryRes.Payload
```

#### 虚拟机重启

##### 重启指定虚拟机

```go
restartParams := vm.NewRestartVMParams()
restartParams.RequestBody = &models.VMOperateParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
}
restartRes, err := client.VM.RestartVM(restartParams)
if err != nil {
	return err
}
withTaskVm := restartRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
restartedVm := queryRes.Payload[0]
```

##### 重启批量虚拟机

```go
restartParams := vm.NewRestartVMParams()
restartParams.RequestBody = &models.VMOperateParams{
	Where: &models.VMWhereInput{
		IDIn: []string{"vm_id_1", "vm_id_2"},
	},
}
startRes, err := client.VM.RestartVM(restartParams)
if err != nil {
	return err
}
withTaskVms := startRes.Payload
taskIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
	return *vm.TaskID
}).([]string)
if !valid {
	return fmt.Errorf("failed to parse task ids")
}
vmIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
	return *vm.Data.ID
}).([]string)
if !valid {
	return fmt.Errorf("failed to parse vm ids")
}
err = utils.WaitTasks(client, taskIds)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		IDIn: vmIds,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
restartedVms := queryRes.Payload
```

##### 强制重启指定虚拟机

```go
restartParams := vm.NewForceRestartVMParams()
restartParams.RequestBody = &models.VMOperateParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
}
restartRes, err := client.VM.ForceRestartVM(restartParams)
if err != nil {
	return err
}
withTaskVm := restartRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
restartedVm := queryRes.Payload[0]
```

##### 强制重启批量虚拟机

```go
restartParams := vm.NewForceRestartVMParams()
restartParams.RequestBody = &models.VMOperateParams{
	Where: &models.VMWhereInput{
		IDIn: []string{"vm_id_1", "vm_id_2"},
	},
}
startRes, err := client.VM.ForceRestartVM(restartParams)
if err != nil {
	return err
}
withTaskVms := startRes.Payload
taskIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
	return *vm.TaskID
}).([]string)
if !valid {
	return fmt.Errorf("failed to parse task ids")
}
vmIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
	return *vm.Data.ID
}).([]string)
if !valid {
	return fmt.Errorf("failed to parse vm ids")
}
err = utils.WaitTasks(client, taskIds)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		IDIn: vmIds,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
restartedVms := queryRes.Payload
```

#### 虚拟机暂停

##### 暂停指定虚拟机

```go
suspendParams := vm.NewSuspendVMParams()
suspendParams.RequestBody = &models.VMOperateParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
}
restartRes, err := client.VM.SuspendVM(suspendParams)
if err != nil {
	return err
}
withTaskVm := restartRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
suspendedVm := queryRes.Payload[0]
```

##### 暂停批量虚拟机

```go
suspendParams := vm.NewSuspendVMParams()
suspendParams.RequestBody = &models.VMOperateParams{
	Where: &models.VMWhereInput{
		IDIn: []string{"vm_id_1", "vm_id_2"},
	},
}
startRes, err := client.VM.SuspendVM(suspendParams)
if err != nil {
	return err
}
withTaskVms := startRes.Payload
taskIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
	return *vm.TaskID
}).([]string)
if !valid {
	return fmt.Errorf("failed to parse task ids")
}
vmIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
	return *vm.Data.ID
}).([]string)
if !valid {
	return fmt.Errorf("failed to parse vm ids")
}
err = utils.WaitTasks(client, taskIds)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		IDIn: vmIds,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
suspendedVms := queryRes.Payload
```

#### 虚拟机恢复

##### 恢复指定虚拟机

```go
resumeParams := vm.NewResumeVMParams()
resumeParams.RequestBody = &models.VMOperateParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
}
restartRes, err := client.VM.ResumeVM(resumeParams)
if err != nil {
	return err
}
withTaskVm := restartRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
resumedVm := queryRes.Payload[0]
```

##### 恢复批量虚拟机

```go
resumeParams := vm.NewResumeVMParams()
resumeParams.RequestBody = &models.VMOperateParams{
	Where: &models.VMWhereInput{
		IDIn: []string{"vm_id_1", "vm_id_2"},
	},
}
startRes, err := client.VM.ResumeVM(resumeParams)
if err != nil {
	return err
}
withTaskVms := startRes.Payload
taskIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
	return *vm.TaskID
}).([]string)
if !valid {
	return fmt.Errorf("failed to parse task ids")
}
vmIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
	return *vm.Data.ID
}).([]string)
if !valid {
	return fmt.Errorf("failed to parse vm ids")
}
err = utils.WaitTasks(client, taskIds)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		IDIn: vmIds,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
resumedVms := queryRes.Payload
```

### 删除虚拟机

#### 回收站

##### 移入回收站

```go
moveToBinParams := vm.NewMoveVMToRecycleBinParams()
moveToBinParams.RequestBody = &models.VMOperateParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
}
restartRes, err := client.VM.MoveVMToRecycleBin(moveToBinParams)
if err != nil {
	return err
}
withTaskVm := restartRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
vmInBin := queryRes.Payload[0]
```

##### 从回收站恢复

```go
recoveryParams := vm.NewRecoverVMFromRecycleBinParams()
recoveryParams.RequestBody = &models.VMOperateParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
}
restartRes, err := client.VM.RecoverVMFromRecycleBin(recoveryParams)
if err != nil {
	return err
}
withTaskVm := restartRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
getVmParams := vm.NewGetVmsParams()
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: withTaskVm.Data.ID,
	},
}
queryRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
recoveredVm := queryRes.Payload[0]
```

#### 永久删除

```go
deleteParams := vm.NewDeleteVMParams()
deleteParams.RequestBody = &models.VMOperateParams{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
}
restartRes, err := client.VM.DeleteVM(deleteParams)
if err != nil {
	return err
}
withTaskVm := restartRes.Payload[0]
err = utils.WaitTask(client, withTaskVm.TaskID)
if err != nil {
	return err
}
```

## 场景示例

### 虚拟机备份

```go
package main

import (
	"fmt"

	apiclient "github.com/Sczlog/cloudtower-go-sdk/client"
	"github.com/Sczlog/cloudtower-go-sdk/client/iscsi_lun_snapshot"
	"github.com/Sczlog/cloudtower-go-sdk/client/user"
	"github.com/Sczlog/cloudtower-go-sdk/client/vm"
	"github.com/Sczlog/cloudtower-go-sdk/client/vm_snapshot"
	"github.com/Sczlog/cloudtower-go-sdk/models"
	"github.com/Sczlog/cloudtower-go-sdk/utils"
	"github.com/openlyinc/pointy"
	"github.com/thoas/go-funk"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func create_vm_snapshot(
  client *apiclient.Cloudtower,
	targetVmName string,
	targetSnapshotName string,
	consistentType models.ConsistentType) (*models.VMSnapshot, []*models.IscsiLunSnapshot, error) {
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			Name: &targetVmName,
		},
		First: pointy.Int32(1),
	}
	// 1. 获取所需备份的虚拟机的信息，这里我们需要vm的id来构建创建snapshot的参数
	getVmRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, nil, err
	}
	targetVm := getVmRes.Payload[0]
	vmToolStatus := *targetVm.VMToolsStatus
	// vm 已安装并启动 VMTools 时，consistent_type 可以使用 FILE_SYSTEM_CONSISTENT 代表文件系统一致性快照
	if vmToolStatus != models.VMToolsStatusRUNNING && consistentType == models.ConsistentTypeFILESYSTEMCONSISTENT {
		consistentType = models.ConsistentTypeCRASHCONSISTENT
	}
	createSnapshotParams := vm_snapshot.NewCreateVMSnapshotParams()
	createSnapshotParams.RequestBody = &models.VMSnapshotCreationParams{
		Data: []*models.VMSnapshotCreationParamsDataItems0{
			{
				VMID:           targetVm.ID,
				Name:           &targetSnapshotName,
				ConsistentType: consistentType.Pointer(),
			},
		},
	}
	// 2. 创建虚拟机快照
	createRes, err := client.VMSnapshot.CreateVMSnapshot(createSnapshotParams)
	if err != nil {
		return nil, nil, err
	}
	withTaskSnapshot := createRes.Payload[0]
	// 3. 等待Task完成
	err = utils.WaitTask(client, withTaskSnapshot.TaskID)
	if err != nil {
		return nil, nil, err
	}
	getSnapshotParams := vm_snapshot.NewGetVMSnapshotsParams()
	getSnapshotParams.RequestBody = &models.GetVMSnapshotsRequestBody{
		Where: &models.VMSnapshotWhereInput{
			ID: withTaskSnapshot.Data.ID,
		},
	}
	// 4. 根据返回的id查询生成的虚拟机快照
	getSnapshotRes, err := client.VMSnapshot.GetVMSnapshots(getSnapshotParams)
	if err != nil {
		return nil, nil, err
	}
	createdSnapshot := getSnapshotRes.Payload[0]
	// 5. 根据返回的snapshot中的vm_disks包含了快照的虚拟盘信息
	// type 为 DISK 表示对应一个卷，其中会包含一个 snapshot_local_id 则表示该虚拟卷对应的lun快照的 local_id
	// type 为 CD-ROM则代表为被挂载的CD-ROM，不会产生lun快照
	lunSnapshotIds := funk.Map(funk.Filter(createdSnapshot.VMDisks, func(disk *models.NestedFrozenDisks) bool {
		return *disk.Type == models.VMDiskTypeDISK
	}), func(disk *models.NestedFrozenDisks) string {
		return *disk.SnapshotLocalID
	}).([]string)
	getLunSnapshotParams := iscsi_lun_snapshot.NewGetIscsiLunSnapshotsParams()
	getLunSnapshotParams.RequestBody = &models.GetIscsiLunSnapshotsRequestBody{
		Where: &models.IscsiLunSnapshotWhereInput{
			IDIn: lunSnapshotIds,
		},
	}
	getLunSnapshotRes, err := client.IscsiLunSnapshot.GetIscsiLunSnapshots(getLunSnapshotParams)
	if err != nil {
		return nil, nil, err
	}
	return createdSnapshot, getLunSnapshotRes.Payload, nil
}
```

### Dashboard 构建

#### 定义工具方法

```go
import (
	"fmt"

	apiclient "github.com/Sczlog/cloudtower-go-sdk/client"
	"github.com/Sczlog/cloudtower-go-sdk/client/alert"
	"github.com/Sczlog/cloudtower-go-sdk/client/cluster"
	"github.com/Sczlog/cloudtower-go-sdk/client/disk"
	"github.com/Sczlog/cloudtower-go-sdk/client/host"
	"github.com/Sczlog/cloudtower-go-sdk/client/user"
	"github.com/Sczlog/cloudtower-go-sdk/models"
	"github.com/openlyinc/pointy"
	"github.com/thoas/go-funk"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

var ByteUnits []string = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB"}
var HzUnits []string = []string{"Hz", "KHz", "MHz", "GHz", "THz"}

func formatUnit(base float64, units []string, step int32) string {
	length := len(units)
	if length == 0 {
		panic("No unit provided")
	}
	if base <= 0 {
		return fmt.Sprintf("0%s", units[0])
	}
	for i, unit := range units {
		if base < float64(step) || i == length-1 {
			return fmt.Sprintf("%0.2f%s", base, unit)
		}
		base = base / float64(step)
	}
	return fmt.Sprintf("%0.2f%s", base, units[length-1])
}
```

#### 构建报警信息

```go
type AlertInfo struct {
	Critical []*models.Alert
	Notice   []*models.Alert
	Info     []*models.Alert
}

func NewAlertInfo(critial []*models.Alert, notice []*models.Alert, info []*models.Alert) *AlertInfo {
	return &AlertInfo{
		Critical: critial,
		Notice:   notice,
		Info:     info,
	}
}

func buildAlertInfo(client *apiclient.Cloudtower, clusterIds []string) (*AlertInfo, error) {
	getAlertParams := alert.NewGetAlertsParams()
	if len(clusterIds) == 0 {
		getAlertParams.RequestBody = &models.GetAlertsRequestBody{
			Where: &models.AlertWhereInput{
				Ended: pointy.Bool(false),
			},
		}
	} else {
		getAlertParams.RequestBody = &models.GetAlertsRequestBody{
			Where: &models.AlertWhereInput{
				Ended: pointy.Bool(false),
				Cluster: &models.ClusterWhereInput{
					IDIn: clusterIds,
				},
			},
		}
	}
	res, err := client.Alert.GetAlerts(getAlertParams)
	if err != nil {
		return nil, err
	}
	var critial []*models.Alert = []*models.Alert{}
	var notice []*models.Alert = []*models.Alert{}
	var info []*models.Alert = []*models.Alert{}
	for _, alert := range res.Payload {
		switch *alert.Severity {
		case "CRITICAL":
			critial = append(critial, alert)
		case "NOTICE":
			notice = append(notice, alert)
		case "INFO":
			info = append(info, alert)
		}
	}
	return NewAlertInfo(critial, notice, info), nil
}
```

#### 构建硬盘信息

> 这里以机械硬盘为例

```go
type DiskInfo struct {
	Healthy int32
	Warning int32
	Error   int32
	Total   int32
}

func NewDiskInfo() *DiskInfo {
	return &DiskInfo{
		Healthy: 0,
		Warning: 0,
		Error:   0,
		Total:   0,
	}
}

func buildHddInfo(client *apiclient.Cloudtower, clusterIds []string) (*DiskInfo, error) {
	getDiskParams := disk.NewGetDisksParams()
	if len(clusterIds) == 0 {
		getDiskParams.RequestBody = &models.GetDisksRequestBody{}
	} else {
		getDiskParams.RequestBody = &models.GetDisksRequestBody{
			Where: &models.DiskWhereInput{
				Host: &models.HostWhereInput{
					Cluster: &models.ClusterWhereInput{
						IDIn: clusterIds,
					},
				},
			},
		}
	}
	res, err := client.Disk.GetDisks(getDiskParams)
	if err != nil {
		return nil, err
	}
	hddInfo := NewDiskInfo()
	for _, disk := range res.Payload {
		if *disk.Type == models.DiskTypeHDD {
			if funk.Contains(
				[]models.DiskHealthStatus{
					models.DiskHealthStatusHEALTHY,
					models.DiskHealthStatusSUBHEALTHY,
					models.DiskHealthStatusSMARTFAILED,
				},
				*disk.HealthStatus,
			) {
				hddInfo.Error++
			} else if funk.Contains(
				[]models.DiskUsageStatus{
					models.DiskUsageStatusUNMOUNTED,
					models.DiskUsageStatusPARTIALMOUNTED,
				},
				*disk.UsageStatus,
			) {
				hddInfo.Warning++
			} else {
				hddInfo.Healthy++
			}
			hddInfo.Total++
		}
	}
	return hddInfo, nil
}
```

#### 构建性能指标

> 获取指定集群的 CPU 核数，CPU 频率总数，CPU 使用率，内存总量，内存使用量，存储资源总量，存储资源已使用量，存储资源失效量与存储资源可用量。

```go
type CpuInfo struct {
	TotalCore uint32
	TotalInHz uint64
	Total     string
	UsedInHz  uint64
	Used      string
	Usage     string
}

func NewCpuInfo() *CpuInfo {
	return &CpuInfo{
		TotalCore: 0,
		TotalInHz: 0,
		UsedInHz:  0,
	}
}

func (info *CpuInfo) compute() *CpuInfo {
	info.Total = formatUnit(float64(info.TotalInHz), HzUnits, 1000)
	info.Used = formatUnit(float64(info.UsedInHz), HzUnits, 1000)
	info.Usage = fmt.Sprintf("%0.2f%%", float64(info.UsedInHz)/float64(info.TotalInHz))
	return info
}

type MemoryInfo struct {
	TotalInByte uint64
	Total       string
	UsedInByte  uint64
	Used        string
	Usage       string
}

func NewMemoryInfo() *MemoryInfo {
	return &MemoryInfo{
		TotalInByte: 0,
		UsedInByte:  0,
	}
}

func (info *MemoryInfo) compute() *MemoryInfo {
	info.Total = formatUnit(float64(info.TotalInByte), ByteUnits, 1024)
	info.Used = formatUnit(float64(info.UsedInByte), ByteUnits, 1024)
	info.Usage = fmt.Sprintf("%0.2f%%", float64(info.UsedInByte)/float64(info.TotalInByte))
	return info
}

type StorageInfo struct {
	TotalInByte     uint64
	Total           string
	UsedInByte      uint64
	Used            string
	InvalidInByte   uint64
	Invalid         string
	AvailableInByte uint64
	Available       string
}

func NewStorageInfo() *StorageInfo {
	return &StorageInfo{
		TotalInByte:     0,
		UsedInByte:      0,
		InvalidInByte:   0,
		AvailableInByte: 0,
	}
}

func (info *StorageInfo) compute() *StorageInfo {
	info.AvailableInByte = info.TotalInByte - info.UsedInByte - info.InvalidInByte
	info.Total = formatUnit(float64(info.TotalInByte), ByteUnits, 1024)
	info.Used = formatUnit(float64(info.UsedInByte), ByteUnits, 1024)
	info.Invalid = formatUnit(float64(info.InvalidInByte), ByteUnits, 1024)
	info.Available = formatUnit(float64(info.AvailableInByte), ByteUnits, 1024)
	return info
}

type MetricInfo struct {
	Storage *StorageInfo
	Memory  *MemoryInfo
	Cpu     *CpuInfo
}

func buildMetricsInfo(client *apiclient.Cloudtower, clusters []*models.Cluster, clusterIds []string) (*MetricInfo, error) {
	memory := NewMemoryInfo()
	storage := NewStorageInfo()
	cpu := NewCpuInfo()
	getHostParams := host.NewGetHostsParams()
	if len(clusterIds) == 0 {
		getHostParams.RequestBody = &models.GetHostsRequestBody{}
	} else {
		getHostParams.RequestBody = &models.GetHostsRequestBody{
			Where: &models.HostWhereInput{
				Cluster: &models.ClusterWhereInput{
					IDIn: clusterIds,
				},
			},
		}
	}
	hosts, err := client.Host.GetHosts(getHostParams)
	if err != nil {
		return nil, err
	}
	clusterIdMap := make(map[string]*models.Cluster)
	for _, cluster := range clusters {
		if _, ok := clusterIdMap[*cluster.ID]; !ok {
			clusterIdMap[*cluster.ID] = cluster
		}
		if *cluster.Type == models.ClusterTypeSMTXOS {
			cpu.TotalCore += uint32(*cluster.TotalCPUCores)
			cpu.TotalInHz += uint64(*cluster.TotalCPUHz)
			cpu.UsedInHz += uint64(*cluster.UsedCPUHz)
			if cluster.Hypervisor != nil && *cluster.Hypervisor == models.HypervisorVMWARE {
				memory.TotalInByte += uint64(*cluster.TotalMemoryBytes)
				memory.UsedInByte += uint64(*cluster.UsedMemoryBytes)
			}
		}
		storage.TotalInByte += uint64(*cluster.TotalDataCapacity)
		storage.UsedInByte += uint64(*cluster.UsedDataSpace)
		storage.InvalidInByte += uint64(*cluster.FailureDataSpace)
	}
	for _, host := range hosts.Payload {
		cluster, ok := clusterIdMap[*host.Cluster.ID]
		if ok {
			if *cluster.Hypervisor == models.HypervisorELF {
				memory.TotalInByte += uint64(*host.TotalMemoryBytes)
				memory.UsedInByte += uint64(*host.UsedMemoryBytes)
			}
		}
	}
	storage.compute()
	cpu.compute()
	memory.compute()
	return &MetricInfo{
		Memory:  memory,
		Cpu:     cpu,
		Storage: storage,
	}, nil
}
```

#### 构建 Dashboard

```go
type DashboardInfo struct {
	Metric *MetricInfo
	Hdd    *DiskInfo
	Alert  *AlertInfo
}

func BuildDashboard(client *apiclient.Cloudtower, clusterIds []string) (*DashboardInfo, error) {
	getClusterParams := cluster.NewGetClustersParams()
	if len(clusterIds) == 0 {
		getClusterParams.RequestBody = &models.GetClustersRequestBody{}
	} else {
		getClusterParams.RequestBody = &models.GetClustersRequestBody{
			Where: &models.ClusterWhereInput{
				IDIn: clusterIds,
			},
		}
	}
	res, err := client.Cluster.GetClusters(getClusterParams)
	if err != nil {
		return nil, err
	}
	metrics, err := buildMetricsInfo(client, res.Payload, clusterIds)
	if err != nil {
		return nil, err
	}
	hdd, err := buildHddInfo(client, clusterIds)
	if err != nil {
		return nil, err
	}
	alert, err := buildAlertInfo(client, clusterIds)
	if err != nil {
		return nil, err
	}
	return &DashboardInfo{
		Metric: metrics,
		Hdd:    hdd,
		Alert:  alert,
	}, nil
}
```
