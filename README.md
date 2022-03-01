# Cloudtower Go SDK

Golang 环境下的 Cloudtower SDK，适用于 Java 1.8 及以上版本

- [源码地址](https://github.com/Sczlog/cloudtower-go-sdk)
- [下载地址](https://github.com/Sczlog/cloudtower-go-sdk/releases)
- [通用指南](https://cloudtower-api-doc.vercel.app)

## 安装

```shell
go get https://github.com/smartxworks/cloudtower-python-sdk
```

## 使用

### 创建实例

#### 创建 `ApiClient` 实例

```go
import (
	apiclient "github.com/Sczlog/cloudtower-go-sdk/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)
transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
c := apiclient.New(transport, strfmt.Default)
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
logResponse, err := c.User.Login(loginParams)
if err != nil {
	panic(err.Error())
}
transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", *logResponse.Payload.Data.Token)
```

#### 获取资源

```go
getVmParams := vm.NewGetVmsParams();
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
    // 使用pointy创建指针类型
		ID: pointy.String("vm_id"),
	},
}
vmsResponse, err := client.VM.GetVms(getVmParams, nil)

if err != nil {
	panic(err.Error())
}
```

#### 更新资源

> 资源更新会产生相关的异步任务，当异步任务结束时，代表资源操作完成且数据已更新。

```go
target_vm := vmsResponse.GetPayload()[0]
vmStartParams := vm.NewStartVMParams()
vmStartParams.RequestBody = &models.VMStartParams{
	Where: &models.VMWhereInput{
		ID: target_vm.ID,
	},
}
startResponse, err := client.VM.StartVM(vmStartParams, nil)
if err != nil {
	panic(err.Error())
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
task := *startResponse.Payload[0].TaskID
err = utils.WaitTask(client, task)
if err != nil {
	panic(err.Error())
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
tasks := funk.Map(startResponse.Payload, func(tvm *models.WithTaskVM) string {
	return *tvm.TaskID
}).([]string)
err = utils.WaitTasks(client, tasks)
if err != nil {
	panic(err.Error())
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
taskDefaultResponse, err := client.Task.GetTasks(getTaskDefaultParams, nil)

getTaskZhParams := task.NewGetTasksParams()
getTaskZhParams.RequestBody = &models.GetTasksRequestBody{
	First: pointy.Int32(10),
}
// 此时得到的 alerts 中的 message, solution, cause, impact 将被转换为中文描述。
getTaskZhParams.ContentLanguage = pointy.String("zh-CN")
taskZhResponse, err := client.Task.GetTasks(getTaskZhParams, nil)
```
