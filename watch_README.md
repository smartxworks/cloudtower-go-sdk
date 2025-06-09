# CloudTower 资源变更监控客户端使用说明

## 概述

`ResourceChangeWatchClient` 是一个用于监控 CloudTower 平台资源变更的客户端工具。它通过轮询机制持续获取资源变更事件，并通过不同的通道（channel）将这些事件传递给使用者。

## 通道（Channel）的作用

`ResourceChangeWatchClient` 提供了三种不同类型的通道，每种通道都有特定的用途：

### 1. 资源变更事件通道 (Channel)

```go
func (c *ResourceChangeWatchClient) Channel() <-chan *models.ResourceChangeEvent
```

- **作用**：传递正常的资源变更事件
- **数据类型**：`*models.ResourceChangeEvent`
- **缓冲大小**：500
- **使用场景**：接收所有正常的资源变更通知，包括资源的创建、更新、删除等操作
- **获取方式**：通过调用 `client.Channel()` 获取

### 2. 警告事件通道 (WarningChannel)

```go
func (c *ResourceChangeWatchClient) WarningChannel() <-chan *WarningEvent
```

- **作用**：传递非致命性错误事件
- **数据类型**：`*WarningEvent`，包含 `Err` 字段
- **缓冲大小**：无缓冲
- **使用场景**：接收轮询过程中发生的非致命性错误，如临时网络问题、请求失败但未超过最大重试次数等
- **获取方式**：通过调用 `client.WarningChannel()` 获取
- **特点**：
  - 接收到警告事件后，客户端会继续运行，并根据指数退避策略重试
  - 警告通道采用非阻塞的清空机制，每次写入新警告前会尝试清空通道中的旧警告，确保只保留最新的警告信息，避免通道阻塞

### 3. 错误事件通道 (ErrorChannel)

```go
func (c *ResourceChangeWatchClient) ErrorChannel() <-chan *ErrorEvent
```

- **作用**：传递致命性错误事件
- **数据类型**：`*ErrorEvent`，包含 `Type`、`Err` 和 `CompactRevision` 字段
- **缓冲大小**：无缓冲
- **使用场景**：接收导致客户端停止工作的严重错误，如：
  - `ErrorEventTypeCompacted`：请求的修订版本已被压缩
  - `ErrorEventTypeRequest`：请求错误超过最大重试次数
- **获取方式**：通过调用 `client.ErrorChannel()` 获取
- **特点**：当错误事件发生时，客户端会自动停止（调用 `Stop()` 方法）

## CatchUp 机制

`catchedUp` 是一个原子布尔值，用于标记客户端是否已经"赶上"了最新的资源变更事件。这个机制的主要作用是确保客户端能够从指定的起始点开始，逐步获取所有变更事件，直到追赶上最新状态。

### CatchUp 机制的工作流程：

1. **初始化**：
   - 客户端启动时，`catchedUp` 被设置为 `false`
   - 可以通过 `StartRevision` 参数指定起始的修订版本

2. **追赶过程**：
   - 客户端不断轮询获取资源变更事件
   - 每次轮询后，会比较最新接收到的事件修订版本（`c.currentRevision`）与服务器当前最新修订版本（`events.Payload.CurrentRevision`）

3. **追赶完成判断**：
   ```go
   if !c.catchedUp.Load() {
       if compare, err := utils.CompareBigIntStrings(c.currentRevision, events.Payload.CurrentRevision); err != nil {
           return err
       } else if compare >= 0 {
           // 如果最新事件的修订版本大于或等于当前修订版本
           // 表示我们已经追赶上了最新数据
           c.catchedUp.Store(true)
       }
   }
   ```

4. **压缩错误处理**：
   - 当 `catchedUp` 为 `true` 时，客户端会检查是否发生了压缩错误
   - 如果当前修订版本小于服务器的压缩修订版本，表示请求的历史数据已被压缩，客户端将发送 `ErrorEventTypeCompacted` 错误事件并停止

### CatchUp 机制的意义：

1. **数据完整性**：确保客户端不会遗漏任何资源变更事件
2. **状态同步**：帮助客户端了解自己是否已经与服务器同步到最新状态
3. **错误处理**：只有在追赶上最新状态后，才会检查压缩错误，避免在初始加载阶段误报错误

## 使用示例

```go
// 创建客户端
params := &watchor.NewResourceChangeWatchClientParams{
    Client:          cloudtowerClient,
    ClientOptions:   clientOptions,
    MaxRetries:      5,
    ResourceTypes:   []string{"VM", "HOST"},
    PollingInterval: 2 * time.Second,
}
watchClient, err := watchor.NewResourceChangeWatchClient(params)
if err != nil {
    log.Fatalf("创建监控客户端失败: %v", err)
}

// 启动客户端
startParams := &watchor.ResourceChangeWatchStartParams{
    StartRevision: nil, // 从最新的修订版本开始
}
if err := watchClient.Start(startParams); err != nil {
    log.Fatalf("启动监控客户端失败: %v", err)
}

// 处理各种事件
go func() {
    for {
        select {
        case event, ok := <-watchClient.Channel():
            if !ok {
                return // 通道已关闭
            }
            log.Printf("收到资源变更事件: %v", event)
            
        case warning, ok := <-watchClient.WarningChannel():
            if !ok {
                return // 通道已关闭
            }
            log.Printf("警告: %v", warning.Err)
            
        case err, ok := <-watchClient.ErrorChannel():
            if !ok {
                return // 通道已关闭
            }
            if err.Type == watchor.ErrorEventTypeCompacted {
                log.Printf("请求的修订版本已被压缩，压缩版本: %s", *err.CompactRevision)
                // 可以使用压缩版本重新启动客户端
            } else {
                log.Printf("错误: %v", err.Err)
            }
            return // 发生错误，客户端已停止
        }
    }
}()

// 使用完毕后停止客户端
defer watchClient.Stop()
```

## 注意事项

1. 客户端使用轮询机制，而非长连接，因此会有一定的延迟
2. 错误事件会导致客户端自动停止，需要重新启动
3. 当发生压缩错误时，可以使用返回的 `CompactRevision` 重新启动客户端
4. 可以通过调整 `PollingInterval` 参数控制轮询频率，但过高的频率可能会增加服务器负担
5. 客户端使用指数退避策略处理临时错误，最大重试次数由 `MaxRetries` 参数控制

通过合理使用这三种通道和理解 CatchUp 机制，您可以构建一个健壮的资源变更监控系统，及时响应 CloudTower 平台上的资源变化。 