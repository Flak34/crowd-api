package crowdapiv1

import (
	"encoding/json"
	crowdapiv1 "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	model "github.com/Flak34/crowd-api/internal/project/model"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"time"
)

type ProjectDTO struct {
	ID                 int
	CreatorID          int
	Description        string
	Name               string
	Instruction        string
	Config             ProjectConfigDTO
	TargetOverlap      int
	TasksPerUser       int
	AnnotatorTimeLimit time.Duration
	CreatedAt          time.Time
}

type ProjectConfigDTO struct {
	InputData        []ProjectInputDataDTO  `json:"inputData"`
	OutputData       []ProjectOutputDataDTO `json:"outputData"`
	SerializedLayout string                 `json:"serializedLayout"`
}

type ProjectInputDataDTO struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type ProjectOutputDataDTO struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	WithAggregation bool   `json:"withAggregation"`
}

func protoToProjectDTO(req *crowdapiv1.CreateProjectRequest) (ProjectDTO, error) {
	err := validateProjectReq(req)
	if err != nil {
		return ProjectDTO{}, err
	}
	var config ProjectConfigDTO
	err = json.Unmarshal([]byte(req.TaskConfig), &config)
	if err != nil {
		return ProjectDTO{}, errors.Errorf("unmarshal config: %s", err.Error())
	}
	return ProjectDTO{
		CreatorID:          int(req.CreatorId),
		Description:        req.Description,
		Name:               req.Name,
		Instruction:        req.Instruction,
		Config:             config,
		TargetOverlap:      int(req.TargetOverlap),
		TasksPerUser:       int(req.TasksPerUser),
		AnnotatorTimeLimit: req.AnnotatorTimeLimit.AsDuration(),
	}, nil
}

func modelToProjectConfigDTO(config model.Config) ProjectConfigDTO {
	return ProjectConfigDTO{
		InputData: lo.Map(config.InputData, func(data model.InputData, _ int) ProjectInputDataDTO {
			return ProjectInputDataDTO{
				Type: string(data.Type),
				Name: data.Name,
			}
		}),
		OutputData: lo.Map(config.OutputData, func(data model.OutputData, _ int) ProjectOutputDataDTO {
			return ProjectOutputDataDTO{
				Type:            string(data.Type),
				Name:            data.Name,
				WithAggregation: data.WithAggregation,
			}
		}),
		SerializedLayout: config.Layout,
	}
}

func validateProjectReq(req *crowdapiv1.CreateProjectRequest) error {
	if req == nil {
		return errors.New("request must not be nil")
	}
	if req.TaskConfig == "" || req.Description == "" || req.Name == "" {
		return errors.New("Config, Description and Name must not be empty")
	}
	if req.CreatorId == 0 {
		return errors.New("CreatorId must be specified")
	}
	if req.TargetOverlap == 0 || req.TasksPerUser == 0 {
		return errors.New("TargetOverlap and TasksPerUser must be specified")
	}
	return nil
}

func projectDTOToModel(dto ProjectDTO) model.Project {
	return model.Project{
		ID:          dto.ID,
		CreatorID:   dto.CreatorID,
		Description: dto.Description,
		Name:        dto.Name,
		Instruction: dto.Instruction,
		Config: model.Config{
			InputData: lo.Map(dto.Config.InputData,
				func(data ProjectInputDataDTO, _ int) model.InputData {
					return model.InputData{
						Type: model.DataType(data.Type),
						Name: data.Name,
					}
				}),
			OutputData: lo.Map(dto.Config.OutputData,
				func(data ProjectOutputDataDTO, _ int) model.OutputData {
					return model.OutputData{
						Type:            model.DataType(data.Type),
						Name:            data.Name,
						WithAggregation: data.WithAggregation,
					}
				}),
			Layout: dto.Config.SerializedLayout,
		},
		TargetOverlap:      dto.TargetOverlap,
		TasksPerUser:       dto.TasksPerUser,
		CreatedAt:          dto.CreatedAt,
		AnnotatorTimeLimit: dto.AnnotatorTimeLimit,
	}
}
