package utils

import (
	"context"
	"fmt"
	"time"

	apiclient "github.com/smartxworks/cloudtower-go-sdk/v2/client"
	"github.com/smartxworks/cloudtower-go-sdk/v2/client/task"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	"github.com/thoas/go-funk"
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

func WaitTask(ctx context.Context, client *apiclient.Cloudtower, id *string, interval time.Duration) error {
	if interval < 1*time.Second {
		interval = 1 * time.Second
	}
	if id == nil {
		return nil
	}
	taskParams := task.NewGetTasksParams()
	taskParams.RequestBody = &models.GetTasksRequestBody{
		Where: &models.TaskWhereInput{
			ID: id,
		},
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			tasks, err := client.Task.GetTasks(taskParams)
			if err != nil {
				return err
			}
			if len(tasks.GetPayload()) != 0 && checkTaskFinished(tasks.GetPayload()[0]) >= 0 {
				if *tasks.Payload[0].Status == models.TaskStatusFAILED {
					return fmt.Errorf("task %s failed: %s", *id, *tasks.Payload[0].ErrorMessage)
				}
				return nil
			}
			time.Sleep(interval)
		}
	}
}

func WaitTasks(ctx context.Context, client *apiclient.Cloudtower, ids []string, interval time.Duration) error {
	if interval < 1*time.Second {
		interval = 1 * time.Second
	}
	errorMsgs := ""
	queryIds := ids
	doneMap := make(map[string]bool)
	for _, id := range ids {
		doneMap[id] = false
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			taskParams := task.NewGetTasksParams()
			taskParams.RequestBody = &models.GetTasksRequestBody{
				Where: &models.TaskWhereInput{
					IDIn: queryIds,
				},
			}
			tasksRes, err := client.Task.GetTasks(taskParams)
			if err != nil {
				return err
			}
			tasks := tasksRes.GetPayload()
			for _, t := range tasks {
				switch checkTaskFinished(t) {
				case 1:
					errorMsgs += fmt.Sprintf("\n\t%s: %s", *t.ID, *t.ErrorMessage)
					fallthrough
				case 0:
					doneMap[*t.ID] = true
				}
			}

			queryIds = funk.Filter(ids, func(id string) bool {
				return !doneMap[id]
			}).([]string)

			if len(queryIds) == 0 {
				if len(errorMsgs) > 0 {
					return fmt.Errorf("tasks failed:%s", errorMsgs)
				}
				return nil
			}
			time.Sleep(interval)
		}
	}
}
