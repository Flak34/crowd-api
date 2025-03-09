package crowdapiv1

import (
	desc "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	"github.com/pkg/errors"
)

func validateCreateProjectRequest(req *desc.CreateProjectRequest) error {
	if req == nil {
		return errors.New("request must not be nil")
	}
	if req.TaskConfig == "" || req.Description == "" || req.Name == "" {
		return errors.New("TaskConfig, Description and Name must not be empty")
	}
	if req.CreatorId == 0 {
		return errors.New("CreatorId must be specified")
	}
	if req.TargetOverlap == 0 || req.TasksPerUser == 0 {
		return errors.New("TargetOverlap and TasksPerUser must be specified")
	}
	return nil
}
