package crowdapiv1

import (
	"context"
	"fmt"
	crowd_api_v1 "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
)

func (*Implementation) Ping(ctx context.Context, request *crowd_api_v1.PingRequest) (*crowd_api_v1.PingResponse, error) {
	fmt.Println("got request")
	return &crowd_api_v1.PingResponse{}, nil
}
