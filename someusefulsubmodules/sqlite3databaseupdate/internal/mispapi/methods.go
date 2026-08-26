package mispapi

import (
	"context"
	"fmt"
	"time"
)

// GetEvents поиск события по его идентификатору
func (mr *ModuleRequest) GetEvent(ctx context.Context, timeout time.Duration, eventId string) (int, []byte, error) {
	if err := ctx.Err(); err != nil {
		return 0, []byte{}, err
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	client, err := NewClientMISP(mr.authKey, mr.host, mr.port, false)
	if err != nil {
		return 0, []byte{}, err
	}

	statusCode, resBodyByte, err := client.Post(ctx, fmt.Sprintf("/events/view/%s", eventId), []byte{})
	if err != nil {
		return statusCode, resBodyByte, err
	}

	return statusCode, resBodyByte, nil
}
