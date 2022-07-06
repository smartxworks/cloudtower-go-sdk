## 文件上传

CloudTower 的文件上传需要首先创建一个上传任务，然后使用 Api 分片上传文件，分片大小为 8MB，SDK 简化了上传步骤，将上传类资源的分片上传和创建上传任务封装为同一个 api，这里以上传虚拟机工具镜像为例。

```go
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/pkg/errors"
	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	image "github.com/smartxworks/cloudtower-go-sdk/v2/client/content_library_image"
	models "github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

type CloudTowerWriter struct {
	pos          int
	server       string
	clusterId    string
	uploadTaskId string
	name         string
	size         int64
	client       apiclient.Cloudtower
	description  string
}

func (c *CloudTowerWriter) Write(p []byte) (n int, err error) {
	params := image.NewCreateContentLibraryImageParamsWithTimeout(24 * time.Hour)
	clusterWhere, _ := json.Marshal(map[string]string{
		"id": c.clusterId,
	})
	clusterWhereStr := string(clusterWhere)
	params.Clusters = clusterWhereStr
	if c.pos == 0 {
		params.Name = c.name
		params.Size = strconv.Itoa(int(c.size))
		params.Description = c.description
		params.File = runtime.NamedReader("chunk", io.NopCloser(bytes.NewReader(p)))
		createResp, err := c.client.ContentLibraryImage.CreateContentLibraryImage(params)
		if err != nil {
			switch sdkErr := err.(type) {
			case *image.CreateContentLibraryImageBadRequest:
				return 0, fmt.Errorf("%s,props:\r\n%v", *sdkErr.Payload.Message, sdkErr.Payload.Props)
			case *image.CreateContentLibraryImageInternalServerError:
				return 0, fmt.Errorf("%s,props:\r\n%v", *sdkErr.Payload.Message, sdkErr.Payload.Props)
			case *image.CreateContentLibraryImageNotFound:
				return 0, fmt.Errorf("%s,props:\r\n%v", *sdkErr.Payload.Message, sdkErr.Payload.Props)
			}
			return 0, err
		}
		c.uploadTaskId = *createResp.Payload[0].ID
	} else {
		params.UploadTaskID = &c.uploadTaskId
		params.File = runtime.NamedReader("chunk", io.NopCloser(bytes.NewReader(p)))
		_, err := c.client.ContentLibraryImage.CreateContentLibraryImage(params)
		if err != nil {
			switch sdkErr := err.(type) {
			case *image.CreateContentLibraryImageBadRequest:
				return 0, fmt.Errorf("%s,props:\r\n%v", *sdkErr.Payload.Message, sdkErr.Payload.Props)
			case *image.CreateContentLibraryImageInternalServerError:
				return 0, fmt.Errorf("%s,props:\r\n%v", *sdkErr.Payload.Message, sdkErr.Payload.Props)
			case *image.CreateContentLibraryImageNotFound:
				return 0, fmt.Errorf("%s,props:\r\n%v", *sdkErr.Payload.Message, sdkErr.Payload.Props)
			}
			return 0, err
		}
	}
	c.pos += len(p)
	return len(p), nil
}

func importIso(
	server string, isoPath string,
	name string, description string, clusterId string,
	chunkSize int,
) error {

	client, err := apiclient.NewWithUserConfig(apiclient.ClientConfig{
		Host:     server,
		BasePath: "/v2/api",
		Schemes:  []string{"http"},
	}, apiclient.UserConfig{
		Name:     "username",
		Password: "password",
		Source:   models.UserSourceLOCAL,
	})

	if err != nil {
		return err
	}

	file, err := os.Open(isoPath)
	if err != nil {
		return err
	}
	fileStat, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fileStat.Size()
	_writer := CloudTowerWriter{
		pos:         0,
		clusterId:   clusterId,
		server:      server,
		size:        fileSize,
		client:      *client,
		description: description,
		name:        name,
	}

	writer := bufio.NewWriterSize(&_writer, chunkSize)
	_, err = io.Copy(writer, file)
	if err != nil {
		return errors.Wrap(err, "failed to copy IO")
	}
	err = writer.Flush()
	if err != nil {
		return errors.Wrap(err, "failed to flush writer")
	}
	return nil
}

func main() {
	importIso("tower_server_address", "D:\\SMTX_VMTOOLS_2_12_1_20210521153518.iso", "name", "description", "clusterId", 8388608) // 8388608 是目前tower上传代码的切片大小
}
```
