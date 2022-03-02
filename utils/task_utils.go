package utils

import (
	"fmt"
	"time"

	"github.com/openlyinc/pointy"
	"github.com/thoas/go-funk"

	apiclient "github.com/Sczlog/cloudtower-go-sdk/client"
	"github.com/Sczlog/cloudtower-go-sdk/client/task"
	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// < 0 mean not finished, 0 mean succeed, 1 mean failed
func checkTaskFinished(task *models.Task) int8 {
	switch *task.Status {
	case models.TaskStatusSUCCESSED:
		return 0
	case models.TaskStatusFAILED:
		return 1
	default:
		return -1
	}
}

func WaitTask(client *apiclient.Cloudtower, id string) error {
	start := time.Now()
	taskParams := task.NewGetTasksParams()
	taskParams.RequestBody = &models.GetTasksRequestBody{
		Where: &models.TaskWhereInput{
			ID: pointy.String(id),
		},
	}
	for {
		tasks, err := client.Task.GetTasks(taskParams)
		if err != nil {
			return err
		}
		if len(tasks.GetPayload()) <= 0 {
			return fmt.Errorf("task %s not found", id)
		} else if checkTaskFinished(tasks.Payload[0]) >= 0 {
			if tasks.Payload[0].Status == models.TaskStatusFAILED.Pointer() {
				return fmt.Errorf("task %s failed", id)
			}
			return nil
		} else {
			time.Sleep(5 * time.Second)
			if time.Since(start) > 5*time.Minute {
				return fmt.Errorf("task %s timeout", id)
			}
		}
	}

}

func WaitTasks(client *apiclient.Cloudtower, ids []string) error {
	start := time.Now()
	errorIds := ""
	queryIds := ids
	for {
		taskParams := task.NewGetTasksParams()
		taskParams.RequestBody = &models.GetTasksRequestBody{
			Where: &models.TaskWhereInput{
				IDIn: queryIds,
			},
		}
		tasks, err := client.Task.GetTasks(taskParams)
		if err != nil {
			return err
		}

		queryIds = funk.Map(funk.Filter(tasks.GetPayload(), func(t *models.Task) bool {
			switch checkTaskFinished(t) {
			case 1:
				if len(errorIds) > 0 {
					errorIds += ", "
					errorIds += *t.ID
				} else {
					errorIds = *t.ID
				}
				fallthrough
			case 0:
				return false
			case -1:
				fallthrough
			default:
				return true
			}
		}), func(t *models.Task) string {
			return *t.ID
		}).([]string)

		if len(queryIds) == 0 {
			if len(errorIds) > 0 {
				return fmt.Errorf("tasks [%s] failed", errorIds)
			}
			return nil
		} else {
			time.Sleep(5 * time.Second)
			if time.Since(start) > 5*time.Minute {
				return fmt.Errorf("query task timeout")
			}
		}
	}
}
