package crowdapiv1

import crowd_api_v1 "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"

type Implementation struct {
	crowd_api_v1.UnimplementedCrowdAPIV1Server
}

func NewCrowdAPIV1() *Implementation {
	return &Implementation{}
}
