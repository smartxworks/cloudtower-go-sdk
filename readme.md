# Cloudtower Go SDK

Golang 环境下的 Cloudtower SDK，适用于 golang 1.16 及以上版本

- [源码地址](https://github.com/smartxworks/cloudtower-go-sdk)
- [下载地址](https://github.com/smartxworks/cloudtower-go-sdk/releases)

## 安装

```shell
go get github.com/smartxworks/cloudtower-go-sdk/v2
```

## 使用

> 样例中使用了两个工具库 [pointy](https://go.openly.dev/pointy) 与 [go-funk](https://github.com/thoas/go-funk)，pointy 用于快速创建一个原始类型的指针，go-funk 则提供了一些工具方法，例如 `Map`、`Filter`、`Reduce` 等。

### 创建实例

#### 创建 `ApiClient` 实例

```go
import (
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)
transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
client := apiclient.New(transport, strfmt.Default)
```

> 如果需要使用 https，可以安装证书，或者忽略证书验证

```go
import (
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)
tlsClient, err := httptransport.TLSClient(httptransport.TLSClientOptions{
	InsecureSkipVerify: true,
})
if err != nil {
	fmt.Print(err)
	return
}
transport := httptransport.NewWithClient("192.168.29.157", "/v2/api", []string{"https"}, tlsClient)
client := apiclient.New(transport, strfmt.Default)
```

### 发送请求

#### 引入对应的 `client` 包

> 根据不同用途的操作引入创建相关的 `client` 包

```go
import (
  	vm "github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
)
```

#### 鉴权

可以使用 `NewWithUserConfig` 来创建一个具有鉴权信息的 `Client`

```go
import (
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

client, err := apiclient.NewWithUserConfig(apiclient.ClientConfig{
	Host:     "localhost:8090",
	BasePath: "v2/api",
	Schemes:  []string{"http"},
}, apiclient.UserConfig{
	Name:     "Name",
	Password: "Password",
	Source:   models.UserSourceLOCAL,
})
```

也可以创建 `Client` 后手动添加鉴权信息

```go
import (
  User "github.com/smartxworks/cloudtower-go-sdk/v2/client/user"
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
> | 参数名   | 类型                | 是否必须 | 说明                                |
> | -------- | ------------------- | -------- | ----------------------------------- |
> | context  | Context             | 是       | 用于控制中断                        |
> | client   | \*client.Cloudtower | 是       | 查询所使用的 client 实例            |
> | id       | string              | 是       | 需查询的 task 的 id                 |
> | interval | time.duration       | 是       | 每次查询后的等待时间，最小时间为 1s |

```go
task := *startRes.Payload[0].TaskID
err = utils.WaitTask(context.TODO(), client, task, 1*time.Second)
if err != nil {
	return err
}
```

> 如果是复数任务则可以通过 `WaitTasks`，接受复数个 task id，其余与 `WaitTask` 相同。
>
> - 方法参数说明
>
> | 参数名   | 类型                | 是否必须 | 说明                                |
> | -------- | ------------------- | -------- | ----------------------------------- |
> | context  | Context             | 是       | 用于控制中断                        |
> | client   | \*client.Cloudtower | 是       | 查询所使用的 client 实例            |
> | ids      | []string            | 是       | 需查询的 task 的 id 列表            |
> | interval | time.duration       | 是       | 每次查询后的等待时间，最小时间为 1s |

```go
tasks := funk.Map(startRes.Payload, func(tvm *models.WithTaskVM) string {
	return *tvm.TaskID
}).([]string)
err = utils.WaitTask(context.TODO(), client, tasks, 1*time.Second)
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
package main

import (
	"fmt"

	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"

	httptransport "github.com/go-openapi/runtime/client"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	vms, err := getAllVms(client)
	if err != nil {
		panic(err.Error())
	}
  // handle queried vms
}

func getAllVms(
	client *apiclient.Cloudtower) ([]*models.VM, error) {
	getAllVmsParams := vm.NewGetVmsParams()
	getAllVmsParams.RequestBody = &models.GetVmsRequestBody{}
	vmsRes, err := client.VM.GetVms(getAllVmsParams)
	if err != nil {
		return nil, err
	}
	return vmsRes.Payload, nil
}
```

#### 分页获取虚拟机

```go
package main

import (
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	vms, err := getVmsWithPagination(client)
	if err != nil {
		panic(err.Error())
	}
  // handle queried vms
}

func getVmsWithPagination(
	client *apiclient.Cloudtower,
  from int32,
  to int32) ([]*models.VM, error) {
	getVmsWithPaginationParams := vm.NewGetVmsParams()
	getVmsWithPaginationParams.RequestBody = &models.GetVmsRequestBody{
		First: pointy.Int32(from+1),
		Skip:  pointy.Int32(to-from),
	}
	vmsRes, err := client.VM.GetVms(getVmsWithPaginationParams)
	if err != nil {
		return nil, err
	}
	return vmsRes.Payload, nil
}
```

#### 获取所有已开机虚拟机

```go
package main

import (
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	vms, err := getAllRunningVms(client)
	if err != nil {
		panic(err.Error())
	}
  // handle queried vms
}

func getAllRunningVms(
	client *apiclient.Cloudtower) ([]*models.VM, error) {
	getAllRunningVmsParams := vm.NewGetVmsParams()
	getAllRunningVmsParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			Status: models.VMStatusRUNNING.Pointer(),
		},
	}
	vmsRes, err := client.VM.GetVms(getAllRunningVmsParams)
	if err != nil {
		return nil, err
	}
	return vmsRes.Payload, nil
}
```

#### 获取名称或描述中包含特定字符串的虚拟机

```go
package main

import (
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	vms, err := getVmsMatchStr(client, "matchStr")
	if err != nil {
		panic(err.Error())
	}
  // handle queried vms
}

func getVmsMatchStr(
	client *apiclient.Cloudtower,
	match string) ([]*models.VM, error) {
	getAllVmNameMatchStrParams := vm.NewGetVmsParams()
	getAllVmNameMatchStrParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			NameContains: pointy.String(match),
		},
	}
	vmsRes, err := client.VM.GetVms(getAllVmNameMatchStrParams)
	if err != nil {
		return nil, err
	}
	return vmsRes.Payload, nil
}
```

#### 获取所有 vcpu > n 的虚拟机

```go
package main

import (
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	vms, err := getVmshasNMoreCpuCore(client, 4)
	if err != nil {
		panic(err.Error())
	}
  // handle queried vms
}

func getVmshasNMoreCpuCore(
	client *apiclient.Cloudtower,
	n int32) ([]*models.VM, error) {
	getAllVmCoreGtNParams := vm.NewGetVmsParams()
	getAllVmCoreGtNParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			VcpuGt: pointy.Int32(n),
		},
	}
	vmsRes, err := client.VM.GetVms(getAllVmCoreGtNParams)
	if err != nil {
		return nil, err
	}
	return vmsRes.Payload, nil
}
```

### 从模版创建虚拟机

#### 仅指定 id

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	createdVm, err := createVmFromTemplate(client, "templateId", "clusterId", "vm_name")
	if err != nil {
		panic(err.Error())
	}
  // handle created vm
}

func createVmFromTemplate(
	client *apiclient.Cloudtower,
	templateId string,
	clusterId string,
	name string) (*models.VM, error) {
	createVmFromTemplateParams := vm.NewCreateVMFromContentLibraryTemplateParams()
	createVmFromTemplateParams.RequestBody = []*models.VMCreateVMFromContentLibraryTemplateParams{
		{
			TemplateID: &templateId,
			ClusterID:  &clusterId,
			Name:       &name,
			IsFullCopy: pointy.Bool(false),
		},
	}
	createRes, err := client.VM.CreateVMFromContentLibraryTemplate(createVmFromTemplateParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := createRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

#### 配置与模板不同的虚拟盘参数

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	createdVm, err := createVmFromTemplate(client, "templateId", "clusterId", "vm_name")
	if err != nil {
		panic(err.Error())
	}
  // handle created vm
}

func createVmFromTemplate(
	client *apiclient.Cloudtower,
	templateId string,
	clusterId string,
	name string) (*models.VM, error) {
	createVmFromTemplateParams := vm.NewCreateVMFromContentLibraryTemplateParams()
	createVmFromTemplateParams.RequestBody = []*models.VMCreateVMFromContentLibraryTemplateParams{
		{
			TemplateID: &templateId,
			ClusterID:  &clusterId,
			Name:       &name,
			DiskOperate: &models.VMDiskOperate{
				RemoveDisks: &models.VMDiskOperateRemoveDisks{
					DiskIndex: []int32{2, 3},
				},
				ModifyDisks: []*models.DiskOperateModifyDisk{
					{
						DiskIndex:  pointy.Int32(0),
						VMVolumeID: pointy.String("vmVolumeId1"),
					},
				},
				NewDisks: &models.VMDiskParams{
					MountCdRoms: []*models.VMCdRomParams{
						{
							Index:                 pointy.Int32(2),
							Boot:                  pointy.Int32(0),
							ContentLibraryImageID: pointy.String("contentLibraryImageId"),
						},
					},
					MountDisks: []*models.MountDisksParams{
						{
							Index:      pointy.Int32(3),
							Bus:        models.BusVIRTIO.Pointer(),
							Boot:       pointy.Int32(1),
							VMVolumeID: pointy.String("vmVolumeId2"),
						},
					},
					MountNewCreateDisks: []*models.MountNewCreateDisksParams{
						{
							VMVolume: &models.MountNewCreateDisksParamsVMVolume{
								ElfStoragePolicy: models.VMVolumeElfStoragePolicyTypeREPLICA2THINPROVISION.Pointer(),
								Size:             pointy.Int64(4 * 1024 * 1024 * 1024),
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
	createRes, err := client.VM.CreateVMFromContentLibraryTemplate(createVmFromTemplateParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := createRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

#### 配置与模版不同的网卡参数

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	createdVm, err := createVmFromTemplate(client, "templateId", "clusterId", "vm_name")
	if err != nil {
		panic(err.Error())
	}
  // handle created vm
}

func createVmFromTemplate(
	client *apiclient.Cloudtower,
	templateId string,
	clusterId string,
	name string) (*models.VM, error) {
	createVmFromTemplateParams := vm.NewCreateVMFromContentLibraryTemplateParams()
	createVmFromTemplateParams.RequestBody = []*models.VMCreateVMFromContentLibraryTemplateParams{
		{
			TemplateID: &templateId,
			ClusterID:  &clusterId,
			Name:       &name,
			VMNics: []*models.VMNicParams{
				{
					ConnectVlanID: pointy.String("vlanId2"),
					Enabled:       pointy.Bool(true),
					Model:         models.VMNicModelVIRTIO.Pointer(),
				},
			},
			IsFullCopy: pointy.Bool(false),
		},
	}
	createRes, err := client.VM.CreateVMFromContentLibraryTemplate(createVmFromTemplateParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := createRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

### 创建空白虚拟机

#### 简单创建

```go
package main

import (
  	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	createParams := vm.NewCreateVMParams()
	createParams.RequestBody = []*models.VMCreationParams{
		{
			ClusterID:  pointy.String("clusterId"),
			Name:       pointy.String("test_vm_name"),
			Ha:         pointy.Bool(true),
			CPUCores:   pointy.Int32(4),
			CPUSockets: pointy.Int32(2),
			Memory:     pointy.Int64(8 * 1024 * 1024 * 1024),
			Vcpu:       pointy.Int32(4 * 2),
			Status:     models.VMStatusSTOPPED.Pointer(),
			Firmware:   models.VMFirmwareBIOS.Pointer(),
			VMNics: []*models.VMNicParams{
				{ConnectVlanID: pointy.String("vlanId1")},
			},
			VMDisks: &models.VMDiskParams{
				MountCdRoms: []*models.VMCdRomParams{
					{
						Boot:  pointy.Int32(0),
						Index: pointy.Int32(0),
					},
				},
			},
		},
	}
	createdVm, err := createVm(client, createParams)
	if err != nil {
		panic(err.Error())
	}
  // handle created vm
}

func createVm(
	client *apiclient.Cloudtower,
	createParams *vm.CreateVMParams) (*models.VM, error) {

	createRes, err := client.VM.CreateVM(createParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := createRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}

```

#### 创建时配置虚拟盘

##### CD-ROM 加载 ISO

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	createParams := vm.NewCreateVMParams()
	createParams.RequestBody = []*models.VMCreationParams{
		{
			ClusterID:  pointy.String("clusterId"),
			Name:       pointy.String("test_vm_name"),
			Ha:         pointy.Bool(true),
			CPUCores:   pointy.Int32(4),
			CPUSockets: pointy.Int32(2),
			Memory:     pointy.Int64(8 * 1024 * 1024 * 1024),
			Vcpu:       pointy.Int32(4 * 2),
			Status:     models.VMStatusSTOPPED.Pointer(),
			Firmware:   models.VMFirmwareBIOS.Pointer(),
			VMNics: []*models.VMNicParams{
				{ConnectVlanID: pointy.String("vlanId1")},
			},
			VMDisks: &models.VMDiskParams{
				MountCdRoms: []*models.VMCdRomParams{
					{
						Index:      pointy.Int32(0),
						Boot:       pointy.Int32(0),
						ElfImageID: pointy.String("elfImageId"),
					},
				},
			},
		},
	}
	createdVm, err := createVm(client, createParams)
	if err != nil {
		panic(err.Error())
	}
  // handle created vm
}

func createVm(
	client *apiclient.Cloudtower,
	createParams *vm.CreateVMParams) (*models.VM, error) {

	createRes, err := client.VM.CreateVM(createParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := createRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 挂载虚拟卷为虚拟盘

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	createParams := vm.NewCreateVMParams()
	createParams.RequestBody = []*models.VMCreationParams{
		{
			ClusterID:  pointy.String("clusterId"),
			Name:       pointy.String("test_vm_name"),
			Ha:         pointy.Bool(true),
			CPUCores:   pointy.Int32(4),
			CPUSockets: pointy.Int32(2),
			Memory:     pointy.Int64(8 * 1024 * 1024 * 1024),
			Vcpu:       pointy.Int32(4 * 2),
			Status:     models.VMStatusSTOPPED.Pointer(),
			Firmware:   models.VMFirmwareBIOS.Pointer(),
			VMNics: []*models.VMNicParams{
				{ConnectVlanID: pointy.String("vlanId1")},
			},
			VMDisks: &models.VMDiskParams{
				MountDisks: []*models.MountDisksParams{
					{
						Boot:       pointy.Int32(0),
						Bus:        models.BusVIRTIO.Pointer(),
						VMVolumeID: pointy.String("vmVolumeId1"),
						Index:      pointy.Int32(0),
					},
				},
			},
		},
	}
	createdVm, err := createVm(client, createParams)
	if err != nil {
		panic(err.Error())
	}
  // handle created vm
}

func createVm(
	client *apiclient.Cloudtower,
	createParams *vm.CreateVMParams) (*models.VM, error) {

	createRes, err := client.VM.CreateVM(createParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := createRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 挂载新增虚拟盘

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	createParams := vm.NewCreateVMParams()
	createParams.RequestBody = []*models.VMCreationParams{
		{
			ClusterID:  pointy.String("clusterId"),
			Name:       pointy.String("test_vm_name"),
			Ha:         pointy.Bool(true),
			CPUCores:   pointy.Int32(4),
			CPUSockets: pointy.Int32(2),
			Memory:     pointy.Int64(8 * 1024 * 1024 * 1024),
			Vcpu:       pointy.Int32(4 * 2),
			Status:     models.VMStatusSTOPPED.Pointer(),
			Firmware:   models.VMFirmwareBIOS.Pointer(),
			VMNics: []*models.VMNicParams{
				{ConnectVlanID: pointy.String("vlanId1")},
			},
			VMDisks: &models.VMDiskParams{
				MountNewCreateDisks: []*models.MountNewCreateDisksParams{
					{
						Boot: pointy.Int32(0),
						Bus:  models.BusVIRTIO.Pointer(),
						VMVolume: &models.MountNewCreateDisksParamsVMVolume{
							Size:             pointy.Int64(10 * 1024 * 1024 * 1024),
							ElfStoragePolicy: models.VMVolumeElfStoragePolicyTypeREPLICA2THINPROVISION.Pointer(),
							Name:             pointy.String("new_vm_disk_name"),
						},
					},
				},
			},
		},
	}
	createdVm, err := createVm(client, createParams)
	if err != nil {
		panic(err.Error())
	}
  // handle created vm
}

func createVm(
	client *apiclient.Cloudtower,
	createParams *vm.CreateVMParams) (*models.VM, error) {

	createRes, err := client.VM.CreateVM(createParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := createRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

#### 创建时配置虚拟网卡

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	createParams := vm.NewCreateVMParams()
	createParams.RequestBody = []*models.VMCreationParams{
		{
			ClusterID:  pointy.String("clusterId"),
			Name:       pointy.String("test_vm_name"),
			Ha:         pointy.Bool(true),
			CPUCores:   pointy.Int32(4),
			CPUSockets: pointy.Int32(2),
			Memory:     pointy.Int64(8 * 1024 * 1024 * 1024),
			Vcpu:       pointy.Int32(4 * 2),
			Status:     models.VMStatusSTOPPED.Pointer(),
			Firmware:   models.VMFirmwareBIOS.Pointer(),
			VMNics: []*models.VMNicParams{
				{
					ConnectVlanID: pointy.String("vlanId1"),
					Mirror:        pointy.Bool(true),
					Enabled:       pointy.Bool(false),
					Model:         models.VMNicModelE1000.Pointer(),
				},
			},
			VMDisks: &models.VMDiskParams{
				MountCdRoms: []*models.VMCdRomParams{
					{
						Boot:  pointy.Int32(0),
						Index: pointy.Int32(0),
					},
				},
			},
		},
	}
	createdVm, err := createVm(client, createParams)
	if err != nil {
		panic(err.Error())
	}
  // handle created vm
}

func createVm(
	client *apiclient.Cloudtower,
	createParams *vm.CreateVMParams) (*models.VM, error) {

	createRes, err := client.VM.CreateVM(createParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := createRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

### 编辑虚拟机

#### 编辑基本信息

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	updateParams := vm.NewUpdateVMParams()
	updateParams.RequestBody = &models.VMUpdateParams{
		Where: &models.VMWhereInput{
			ID: pointy.String("vmId"),
		},
		Data: &models.VMUpdateParamsData{
			Name:        pointy.String("updated_name"),
			Description: pointy.String("updated description"),
			Ha:          pointy.Bool(true),
			CPUCores:    pointy.Int32(2),
			CPUSockets:  pointy.Int32(8),
			Vcpu:        pointy.Int32(2 * 8),
			Memory:      pointy.Int64(16 * 1024 * 1024 * 1024),
		},
	}
	updatedVm, err := updateVm(client, updateParams)
	if err != nil {
		panic(err.Error())
	}
  // handle updated vm
}

func updateVm(
	client *apiclient.Cloudtower,
	updateParams *vm.UpdateVMParams) (*models.VM, error) {
	updateRes, err := client.VM.UpdateVM(updateParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := updateRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

#### CD-ROM 编辑

##### 添加 CD-ROM

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	addCdRomParams := vm.NewAddVMCdRomParams()
	addCdRomParams.RequestBody = &models.VMAddCdRomParams{
		Where: &models.VMWhereInput{
			ID: pointy.String("vmId"),
		},
		Data: &models.VMAddCdRomParamsData{
			VMCdRoms: []*models.VMCdRomParams{
				{
					Index:      pointy.Int32(0),
					ElfImageID: pointy.String("elfImageId"),
					Boot:       pointy.Int32(0),
				},
			},
		},
	}
	updatedVm, err := addVmCdRom(client, addCdRomParams)
	if err != nil {
		panic(err.Error())
	}
  // handle updated vm
}

func addVmCdRom(
	client *apiclient.Cloudtower,
	addCdRomParams *vm.AddVMCdRomParams) (*models.VM, error) {
	updateRes, err := client.VM.AddVMCdRom(addCdRomParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := updateRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 删除 CD-ROM

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	updateParams := vm.NewRemoveVMCdRomParams()
	updateParams.RequestBody = &models.VMRemoveCdRomParams{
		Where: &models.VMWhereInput{
			ID: pointy.String("vmId"),
		},
		Data: &models.VMRemoveCdRomParamsData{
			CdRomIds: []string{"cdRomId1", "cdRomId2"},
		},
	}
	updatedVm, err := removeVmCdRom(client, updateParams)
	if err != nil {
		panic(err.Error())
	}
  // handle updated vm
}

func removeVmCdRom(
	client *apiclient.Cloudtower,
	removeCdRomParams *vm.RemoveVMCdRomParams) (*models.VM, error) {
	updateRes, err := client.VM.RemoveVMCdRom(removeCdRomParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := updateRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

#### 虚拟卷操作

##### 添加新虚拟卷

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	updateParams := vm.NewAddVMDiskParams()
	updateParams.RequestBody = &models.VMAddDiskParams{
		Where: &models.VMWhereInput{
			ID: pointy.String("vmId"),
		},
		Data: &models.VMAddDiskParamsData{
			VMDisks: &models.VMAddDiskParamsDataVMDisks{
				MountNewCreateDisks: []*models.MountNewCreateDisksParams{
					{
						VMVolume: &models.MountNewCreateDisksParamsVMVolume{
							ElfStoragePolicy: models.VMVolumeElfStoragePolicyTypeREPLICA2THINPROVISION.Pointer(),
							Size:             pointy.Int64(10 * 1024 * 1024 * 1024),
							Name:             pointy.String("new_disk_name"),
						},
						Boot: pointy.Int32(1),
						Bus:  models.BusVIRTIO.Pointer(),
					},
				},
			},
		},
	}
	updatedVm, err := addVmDisk(client, updateParams)
	if err != nil {
		panic(err.Error())
	}
  // handle created vm
}

func addVmDisk(
	client *apiclient.Cloudtower,
	addVMDiskParams *vm.AddVMDiskParams) (*models.VM, error) {
	updateRes, err := client.VM.AddVMDisk(addVMDiskParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := updateRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 挂载已存在虚拟卷为虚拟盘

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	updateParams := vm.NewAddVMDiskParams()
	updateParams.RequestBody = &models.VMAddDiskParams{
		Where: &models.VMWhereInput{
			ID: pointy.String("vmId"),
		},
		Data: &models.VMAddDiskParamsData{
			VMDisks: &models.VMAddDiskParamsDataVMDisks{
				MountDisks: []*models.MountDisksParams{
					{
						Index:      pointy.Int32(0),
						VMVolumeID: pointy.String("vmVolumeId"),
						Boot:       pointy.Int32(0),
						Bus:        models.BusVIRTIO.Pointer(),
					},
				},
			},
		},
	}
	updatedVm, err := addVmDisk(client, updateParams)
	if err != nil {
		panic(err.Error())
	}
  // handle updated vm
}

func addVmDisk(
	client *apiclient.Cloudtower,
	addVMDiskParams *vm.AddVMDiskParams) (*models.VM, error) {
	updateRes, err := client.VM.AddVMDisk(addVMDiskParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := updateRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 卸载虚拟盘

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	updateParams := vm.NewRemoveVMDiskParams()
	updateParams.RequestBody = &models.VMRemoveDiskParams{
		Where: &models.VMWhereInput{
			ID: pointy.String("vmId"),
		},
		Data: &models.VMRemoveDiskParamsData{
			DiskIds: []string{"diskId1", "diskId2"},
		},
	}
	updatedVm, err := removeVmDisk(client, updateParams)
	if err != nil {
		panic(err.Error())
	}
  // handle updated vm
}

func removeVmDisk(
	client *apiclient.Cloudtower,
	removeVmDiskParams *vm.RemoveVMDiskParams) (*models.VM, error) {
	updateRes, err := client.VM.RemoveVMDisk(removeVmDiskParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := updateRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

#### 网卡操作

##### 添加网卡

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	updateParams := vm.NewAddVMNicParams()
	updateParams.RequestBody = &models.VMAddNicParams{
		Where: &models.VMWhereInput{
			ID: pointy.String("vmId"),
		},
		Data: &models.VMAddNicParamsData{
			VMNics: []*models.VMNicParams{
				{
					ConnectVlanID: pointy.String("vlanId2"),
					Enabled:       pointy.Bool(true),
					Model:         models.VMNicModelVIRTIO.Pointer(),
				},
			},
		},
	}
	updatedVm, err := addVmNic(client, updateParams)
	if err != nil {
		panic(err.Error())
	}
  // handle updated vm
}

func addVmNic(
	client *apiclient.Cloudtower,
	addVMNicParams *vm.AddVMNicParams) (*models.VM, error) {
	updateRes, err := client.VM.AddVMNic(addVMNicParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := updateRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 编辑网卡

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	updateParams := vm.NewUpdateVMNicParams()
	updateParams.RequestBody = &models.VMUpdateNicParams{
		Where: &models.VMWhereInput{
			ID: pointy.String("vmId"),
		},
		Data: &models.VMUpdateNicParamsData{
			NicIndex: pointy.Int32(0),
			Enabled:  pointy.Bool(false),
			Mirror:   pointy.Bool(false),
		},
	}
	updatedVm, err := updateVmNic(client, updateParams)
	if err != nil {
		panic(err.Error())
	}
  // handle updated vm
}

func updateVmNic(
	client *apiclient.Cloudtower,
	updateVMNicParams *vm.UpdateVMNicParams) (*models.VM, error) {
	updateRes, err := client.VM.UpdateVMNic(updateVMNicParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := updateRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 移除网卡

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")
	updateParams := vm.NewRemoveVMNicParams()
	updateParams.RequestBody = &models.VMRemoveNicParams{
		Where: &models.VMWhereInput{
			ID: pointy.String("vmId"),
		},
		Data: &models.VMRemoveNicParamsData{
			NicIndex: []int32{0, 1},
		},
	}
	updatedVm, err := removeVmNic(client, updateParams)
	if err != nil {
		panic(err.Error())
	}
  // handle updated vm
}

func removeVmNic(
	client *apiclient.Cloudtower,
	removeVmDiskParams *vm.RemoveVMNicParams) (*models.VM, error) {
	updateRes, err := client.VM.RemoveVMNic(removeVmDiskParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := updateRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}

```

#### 虚拟机迁移

##### 迁移至指定主机

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	migratedVm, err := migrateVmToHost(client, "vmId", "hostId")
	if err != nil {
		panic(err.Error())
	}
	// handle migrated vm
}

func migrateVmToHost(
	client *apiclient.Cloudtower,
	vmId string,
	hostId string) (*models.VM, error) {
	migrateParams := vm.NewMigRateVMParams()
	migrateParams.RequestBody = &models.VMMigrateParams{
		Where: &models.VMWhereInput{
			ID: pointy.String(vmId),
		},
		Data: &models.VMMigrateParamsData{
			HostID: pointy.String(hostId),
		},
	}
	migrateRes, err := client.VM.MigRateVM(migrateParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := migrateRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 自动调度到合适的主机

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	migratedVm, err := migrateVmAutoSchedule(client, "vmId")
	if err != nil {
		panic(err.Error())
	}
	// handle migrated vm
}

func migrateVmAutoSchedule(
	client *apiclient.Cloudtower,
	vmId string) (*models.VM, error) {
	migrateParams := vm.NewMigRateVMParams()
	migrateParams.RequestBody = &models.VMMigrateParams{
		Where: &models.VMWhereInput{
			ID: pointy.String(vmId),
		},
	}
	migrateRes, err := client.VM.MigRateVM(migrateParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := migrateRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

### 虚拟机电源操作

#### 虚拟机开机:

##### 指定虚拟机开机，自动调度到合适的虚拟机

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	startedVm, err := startVm(client, "vmId")
	if err != nil {
		panic(err.Error())
	}
  // handle started vm
}

func startVm(
	client *apiclient.Cloudtower,
	vmId string) (*models.VM, error) {
	startParams := vm.NewStartVMParams()
	startParams.RequestBody = &models.VMStartParams{
		Where: &models.VMWhereInput{
			ID: pointy.String(vmId),
		},
	}
	startRes, err := client.VM.StartVM(startParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := startRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 批量虚拟机开机，自动调度到合适的虚拟机

```go
package main

import (
	"fmt"
	"time"
	"context"

	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"
	"github.com/thoas/go-funk"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	startedVms, err := startVmsByQuery(client, &models.VMWhereInput{
		IDIn: []string{"vmId"},
	})
	if err != nil {
		panic(err.Error())
	}
  // handle started vms
}

func startVmsByQuery(client *apiclient.Cloudtower,
	where *models.VMWhereInput) ([]*models.VM, error) {
	startParams := vm.NewStartVMParams()
	startParams.RequestBody = &models.VMStartParams{
		Where: where,
	}
	startRes, err := client.VM.StartVM(startParams)
	if err != nil {
		return nil, err
	}
	withTaskVms := startRes.Payload
	taskIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
		return *vm.TaskID
	}).([]string)
	if !valid {
		return nil, fmt.Errorf("failed to parse task ids")
	}
	vmIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
		return *vm.Data.ID
	}).([]string)
	if !valid {
		return nil, fmt.Errorf("failed to parse vm ids")
	}
	err = utils.WaitTasks(context.TODO(), client, taskIds, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			IDIn: vmIds,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload, nil
}
```

##### 开机至指定主机

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	startedVm, err := startVmOnHost(client, "vmId", "hostId2")
	if err != nil {
		panic(err.Error())
	}
  // handle started vm
}

func startVmOnHost(
	client *apiclient.Cloudtower,
	vmId string,
	hostId string) (*models.VM, error) {
	startParams := vm.NewStartVMParams()
	startParams.RequestBody = &models.VMStartParams{
		Where: &models.VMWhereInput{
			ID: pointy.String(vmId),
		},
		Data: &models.VMStartParamsData{
			HostID: pointy.String(hostId),
		},
	}
	startRes, err := client.VM.StartVM(startParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := startRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

#### 虚拟机关机

##### 指定虚拟机关机

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	shutdownVm, err := shutdownVm(client, "vmId")
	if err != nil {
		panic(err.Error())
	}
  // handle shutdown vm
}

func shutdownVm(
	client *apiclient.Cloudtower,
	vmId string) (*models.VM, error) {
	shutdownParams := vm.NewShutDownVMParams()
	shutdownParams.RequestBody = &models.VMOperateParams{
		Where: &models.VMWhereInput{
			ID: pointy.String(vmId),
		},
	}
	shutdownRes, err := client.VM.ShutDownVM(shutdownParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := shutdownRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 批量虚拟机关机

```go
package main

import (
	"fmt"
	"time"
	"context"

	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"
	"github.com/thoas/go-funk"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	shutdownVms, err := shutdownVmsByQuery(client, &models.VMWhereInput{
		IDIn: []string{"vmId"},
	})
	if err != nil {
		panic(err.Error())
	}
  // handle shutdown vms
}

func shutdownVmsByQuery(client *apiclient.Cloudtower,
	where *models.VMWhereInput) ([]*models.VM, error) {
	shutdownParams := vm.NewShutDownVMParams()
	shutdownParams.RequestBody = &models.VMOperateParams{
		Where: where,
	}
	shutdownRes, err := client.VM.ShutDownVM(shutdownParams)
	if err != nil {
		return nil, err
	}
	withTaskVms := shutdownRes.Payload
	taskIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
		return *vm.TaskID
	}).([]string)
	if !valid {
		return nil, fmt.Errorf("failed to parse task ids")
	}
	vmIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
		return *vm.Data.ID
	}).([]string)
	if !valid {
		return nil, fmt.Errorf("failed to parse vm ids")
	}
	err = utils.WaitTasks(context.TODO(), client, taskIds, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			IDIn: vmIds,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload, nil
}
```

##### 强制关机指定虚拟机

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	shutdownVm, err := forceShutdownVm(client, "vmId")
	if err != nil {
		panic(err.Error())
	}
  // handle shutdown vm
}

func forceShutdownVm(
	client *apiclient.Cloudtower,
	vmId string) (*models.VM, error) {
	shutdownParams := vm.NewForceShutDownVMParams()
	shutdownParams.RequestBody = &models.VMOperateParams{
		Where: &models.VMWhereInput{
			ID: pointy.String(vmId),
		},
	}
	shutdownRes, err := client.VM.ForceShutDownVM(shutdownParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := shutdownRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 强制关机批量虚拟机

```go
package main

import (
	"fmt"
	"time"
	"context"

	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"
	"github.com/thoas/go-funk"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	shutdownVms, err := forceshutdownVmsByQuery(client, &models.VMWhereInput{
		IDIn: []string{"vmId"},
	})
	if err != nil {
		panic(err.Error())
	}
  // handle shutdown vms
}


func forceshutdownVmsByQuery(client *apiclient.Cloudtower,
	where *models.VMWhereInput) ([]*models.VM, error) {
	shutdownParams := vm.NewForceShutDownVMParams()
	shutdownParams.RequestBody = &models.VMOperateParams{
		Where: where,
	}
	shutdownRes, err := client.VM.ForceShutDownVM(shutdownParams)
	if err != nil {
		return nil, err
	}
	withTaskVms := shutdownRes.Payload
	taskIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
		return *vm.TaskID
	}).([]string)
	if !valid {
		return nil, fmt.Errorf("failed to parse task ids")
	}
	vmIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
		return *vm.Data.ID
	}).([]string)
	if !valid {
		return nil, fmt.Errorf("failed to parse vm ids")
	}
	err = utils.WaitTasks(context.TODO(), client, taskIds, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			IDIn: vmIds,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload, nil
}

```

#### 虚拟机重启

##### 重启指定虚拟机

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	restartedVm, err := restartVm(client, "vmId")
	if err != nil {
		panic(err.Error())
	}
  // handle restarted vm
}

func restartVm(
	client *apiclient.Cloudtower,
	vmId string) (*models.VM, error) {
	restartParams := vm.NewRestartVMParams()
	restartParams.RequestBody = &models.VMOperateParams{
		Where: &models.VMWhereInput{
			ID: pointy.String(vmId),
		},
	}
	restartRes, err := client.VM.RestartVM(restartParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := restartRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 重启批量虚拟机

```go
package main

import (
	"fmt"
	"time"
	"context"

	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"
	"github.com/thoas/go-funk"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

  	restartedVms, err := restartVmsByQuery(client, &models.VMWhereInput{
		IDIn: []string{"vmId"},
	})
	if err != nil {
		panic(err.Error())
	}
  // handle restarted vms
}

func restartVmsByQuery(client *apiclient.Cloudtower,
	where *models.VMWhereInput) ([]*models.VM, error) {
	restartParams := vm.NewRestartVMParams()
	restartParams.RequestBody = &models.VMOperateParams{
		Where: where,
	}
	restartRes, err := client.VM.RestartVM(restartParams)
	if err != nil {
		return nil, err
	}
	withTaskVms := restartRes.Payload
	taskIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
		return *vm.TaskID
	}).([]string)
	if !valid {
		return nil, fmt.Errorf("failed to parse task ids")
	}
	vmIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
		return *vm.Data.ID
	}).([]string)
	if !valid {
		return nil, fmt.Errorf("failed to parse vm ids")
	}
	err = utils.WaitTasks(context.TODO(), client, taskIds, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			IDIn: vmIds,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload, nil
}
```

##### 强制重启指定虚拟机

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	restartedVm, err := forceRestartVm(client, "vmId")
	if err != nil {
		panic(err.Error())
	}
  // handle restarted vm
}

func forceRestartVm(
	client *apiclient.Cloudtower,
	vmId string) (*models.VM, error) {
	restartParams := vm.NewForceRestartVMParams()
	restartParams.RequestBody = &models.VMOperateParams{
		Where: &models.VMWhereInput{
			ID: pointy.String(vmId),
		},
	}
	shutdownRes, err := client.VM.ForceRestartVM(restartParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := shutdownRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 强制重启批量虚拟机

```go
package main

import (
	"fmt"
	"time"
	"context"

	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"
	"github.com/thoas/go-funk"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	restartedVms, err := forceRestartVmsByQuery(client, &models.VMWhereInput{
		IDIn: []string{"vmId"},
	})
	if err != nil {
		panic(err.Error())
	}
  // handle restarted vms
}

func forceRestartVmsByQuery(client *apiclient.Cloudtower,
	where *models.VMWhereInput) ([]*models.VM, error) {
	restartParams := vm.NewForceRestartVMParams()
	restartParams.RequestBody = &models.VMOperateParams{
		Where: where,
	}
	startRes, err := client.VM.ForceRestartVM(restartParams)
	if err != nil {
		return nil, err
	}
	withTaskVms := startRes.Payload
	taskIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
		return *vm.TaskID
	}).([]string)
	if !valid {
		return nil, fmt.Errorf("failed to parse task ids")
	}
	vmIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
		return *vm.Data.ID
	}).([]string)
	if !valid {
		return nil, fmt.Errorf("failed to parse vm ids")
	}
	err = utils.WaitTasks(context.TODO(), client, taskIds, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			IDIn: vmIds,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload, nil
}
```

#### 虚拟机暂停

##### 暂停指定虚拟机

```go
package main

import (
	"time"
	"context"

	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	suspendedVm, err := suspendVm(client, "vmId")
	if err != nil {
		panic(err.Error())
	}
  // handle suspended vm
}

func suspendVm(
	client *apiclient.Cloudtower,
	vmId string) (*models.VM, error) {
	suspendParams := vm.NewSuspendVMParams()
	suspendParams.RequestBody = &models.VMOperateParams{
		Where: &models.VMWhereInput{
			ID: pointy.String(vmId),
		},
	}
	shutdownRes, err := client.VM.SuspendVM(suspendParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := shutdownRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 暂停批量虚拟机

```go
package main

import (
	"fmt"
	"time"
	"context"

	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"
	"github.com/thoas/go-funk"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	suspendedVms, err := suspendVmsByQuery(client, &models.VMWhereInput{
		IDIn: []string{"vmId"},
	})
	if err != nil {
		panic(err.Error())
	}
	// handle restarted vms
	print(suspendedVms)
}

func suspendVmsByQuery(client *apiclient.Cloudtower,
	where *models.VMWhereInput) ([]*models.VM, error) {
	suspendParams := vm.NewSuspendVMParams()
	suspendParams.RequestBody = &models.VMOperateParams{
		Where: where,
	}
	suspendRes, err := client.VM.SuspendVM(suspendParams)
	if err != nil {
		return nil, err
	}
	withTaskVms := suspendRes.Payload
	taskIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
		return *vm.TaskID
	}).([]string)
	if !valid {
		return nil, fmt.Errorf("failed to parse task ids")
	}
	vmIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
		return *vm.Data.ID
	}).([]string)
	if !valid {
		return nil, fmt.Errorf("failed to parse vm ids")
	}
	err = utils.WaitTasks(context.TODO(), client, taskIds, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			IDIn: vmIds,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload, nil
}

```

#### 虚拟机恢复

##### 恢复指定虚拟机

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	resumedVm, err := resumeVm(client, "vmId")
	if err != nil {
		panic(err.Error())
	}
  // handle resumed vm
}

func resumeVm(
	client *apiclient.Cloudtower,
	vmId string) (*models.VM, error) {
	resumeVmParams := vm.NewResumeVMParams()
	resumeVmParams.RequestBody = &models.VMOperateParams{
		Where: &models.VMWhereInput{
			ID: pointy.String(vmId),
		},
	}
	resumeRes, err := client.VM.ResumeVM(resumeVmParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := resumeRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 恢复批量虚拟机

```go
package main

import (
	"fmt"
	"time"
	"context"

	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"
	"github.com/thoas/go-funk"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	resumedVms, err := resumeVmsByQuery(client, &models.VMWhereInput{
		IDIn: []string{"vmId"},
	})
	if err != nil {
		panic(err.Error())
	}
  // handle resumed vms
}

func resumeVmsByQuery(client *apiclient.Cloudtower,
	where *models.VMWhereInput) ([]*models.VM, error) {
	resumeParams := vm.NewResumeVMParams()
	resumeParams.RequestBody = &models.VMOperateParams{
		Where: where,
	}
	resumeRes, err := client.VM.ResumeVM(resumeParams)
	if err != nil {
		return nil, err
	}
	withTaskVms := resumeRes.Payload
	taskIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
		return *vm.TaskID
	}).([]string)
	if !valid {
		return nil, fmt.Errorf("failed to parse task ids")
	}
	vmIds, valid := funk.Map(withTaskVms, func(vm *models.WithTaskVM) string {
		return *vm.Data.ID
	}).([]string)
	if !valid {
		return nil, fmt.Errorf("failed to parse vm ids")
	}
	err = utils.WaitTasks(context.TODO(), client, taskIds, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			IDIn: vmIds,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload, nil
}

```

### 删除虚拟机

#### 回收站

##### 移入回收站

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	vmInRecycleBin, err := moveVmToRecycleBin(client, "vmId")

	if err != nil {
		panic(err.Error())
	}
	// fmt.Print(vms)
	print(vmInRecycleBin)
}

func moveVmToRecycleBin(
	client *apiclient.Cloudtower,
	vmId string) (*models.VM, error) {
	moveParams := vm.NewMoveVMToRecycleBinParams()
	moveParams.RequestBody = &models.VMOperateParams{
		Where: &models.VMWhereInput{
			ID: pointy.String(vmId),
		},
	}
	moveRes, err := client.VM.MoveVMToRecycleBin(moveParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := moveRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

##### 从回收站恢复

```go
package main

import (
	"time"
	"context"

	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	vmInRecycleBin, err := recoverVmFromRecycleBin(client, "vmId")

	if err != nil {
		panic(err.Error())
	}
	// fmt.Print(vms)
	print(vmInRecycleBin)
}

func recoverVmFromRecycleBin(
	client *apiclient.Cloudtower,
	vmId string) (*models.VM, error) {
	recoverParams := vm.NewRecoverVMFromRecycleBinParams()
	recoverParams.RequestBody = &models.VMOperateParams{
		Where: &models.VMWhereInput{
			ID: pointy.String(vmId),
		},
	}
	recoverRes, err := client.VM.RecoverVMFromRecycleBin(recoverParams)
	if err != nil {
		return nil, err
	}
	withTaskVm := recoverRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return nil, err
	}
	getVmParams := vm.NewGetVmsParams()
	getVmParams.RequestBody = &models.GetVmsRequestBody{
		Where: &models.VMWhereInput{
			ID: withTaskVm.Data.ID,
		},
	}
	queryRes, err := client.VM.GetVms(getVmParams)
	if err != nil {
		return nil, err
	}
	return queryRes.Payload[0], nil
}
```

#### 永久删除

```go
package main

import (
	"time"
	"context"
	"go.openly.dev/pointy"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "token")

	err := deleteVm(client, "vmId")

	if err != nil {
		panic(err.Error())
	}
}

func deleteVm(
	client *apiclient.Cloudtower,
	vmId string) error {
	deleteParams := vm.NewDeleteVMParams()
	deleteParams.RequestBody = &models.VMOperateParams{
		Where: &models.VMWhereInput{
			ID: pointy.String(vmId),
		},
	}
	deleteRes, err := client.VM.DeleteVM(deleteParams)
	if err != nil {
		return err
	}
	withTaskVm := deleteRes.Payload[0]
	err = utils.WaitTask(context.TODO(), client, withTaskVm.TaskID, 1*time.Second)
	if err != nil {
		return err
	}
	return nil
}
```

## 场景示例

### 虚拟机备份

```go
package main

import (
	"fmt"
	"time"
	"context"

	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/iscsi_lun_snapshot"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/user"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/vm_snapshot"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/smartxworks/cloudtower-go-sdk/v2/utils"
	"go.openly.dev/pointy"
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
	err = utils.WaitTask(context.TODO(), client, withTaskSnapshot.TaskID, 1*time.Second)
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
			NameIn: lunSnapshotIds,
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

	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/alert"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/cluster"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/disk"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/host"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/user"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"go.openly.dev/pointy"
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
