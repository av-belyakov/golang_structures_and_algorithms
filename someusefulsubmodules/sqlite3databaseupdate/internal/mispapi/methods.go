package mispapi

import (
	"context"
	"fmt"
)

// GetEvents поиск события по его идентификатору
func (mr *ModuleRequest) GetEvent(ctx context.Context, eventId string) (int, []byte, error) {
	if err := ctx.Err(); err != nil {
		return 0, []byte{}, err
	}

	fmt.Println("method 'GetEvent' is start")

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
